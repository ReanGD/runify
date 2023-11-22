package root

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/os/desktop"
	"github.com/ReanGD/runify/server/os/x11"
	"github.com/ReanGD/runify/server/os/xdg/de"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/provider"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

var logModule = zap.String("module", "runify")

type moduleFull interface {
	Create(impl api.ModuleImpl, name string, isModule bool, cfg *config.Configuration, rootLogger *zap.Logger)
	Init() <-chan error
	Start(ctx context.Context, wg *sync.WaitGroup) <-chan error

	GetName() string

	api.ModuleImpl
}

type moduleItem struct {
	item      moduleFull
	name      string
	initErrCh <-chan error
}

func newModuleItem(item moduleFull, name string) *moduleItem {
	return &moduleItem{
		item:      item,
		name:      name,
		initErrCh: nil,
	}
}

type Runify struct {
	appID    uuid.UUID
	cfg      *config.Config
	logger   *logger.Logger
	rpc      *rpc.Rpc
	ds       *x11.X11
	de       *de.XDGDesktopEntry
	desktop  *desktop.Desktop
	provider *provider.Provider
	items    []*moduleItem

	runifyLogger *zap.Logger
}

func NewRunify() *Runify {
	return &Runify{}
}

func (r *Runify) create(buildCfg *config.BuildCfg) error {
	var err error

	if r.appID, err = uuid.NewV1(); err != nil {
		return fmt.Errorf("Cannot generate UUID for server: %s", err)
	}

	if err := paths.New(); err != nil {
		return fmt.Errorf("Cannot create paths modile: %s", err)
	}

	r.cfg = config.New(buildCfg)
	r.logger = nil
	r.runifyLogger = nil

	var name string
	items := []*moduleItem{}

	r.rpc, name = rpc.New()
	items = append(items, newModuleItem(r.rpc, name))

	r.ds, name = x11.New()
	items = append(items, newModuleItem(r.ds, name))

	r.de, name = de.New()
	items = append(items, newModuleItem(r.de, name))

	r.desktop, name = desktop.New()
	items = append(items, newModuleItem(r.desktop, name))

	r.provider, name = provider.New()
	items = append(items, newModuleItem(r.provider, name))

	r.items = items

	return nil
}

func (r *Runify) init(cfgFile string, cfgSave bool) bool {
	if len(cfgFile) == 0 {
		cfgFile = filepath.Join(paths.GetAppConfig(), "runify.cfg")
	}
	r.cfg.OnInit(cfgFile)
	if cfgSave {
		if err := r.cfg.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot save logger: %s", err)
			return false
		}
	}

	var err error
	if r.logger, err = logger.New(r.cfg, r.appID.String()); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load logger: %s", err)
		return false
	}

	rootLogger := r.logger.GetRoot()
	r.runifyLogger = rootLogger.With(logModule)
	r.cfg.AddVersionToLog(rootLogger)

	configuration := r.cfg.Get()
	for _, it := range r.items {
		m := it.item
		m.Create(m, it.name, module.MODULE, configuration, rootLogger)
	}

	r.rpc.SetDeps()
	r.ds.SetDeps()
	r.de.SetDeps()
	r.desktop.SetDeps(r.ds, r.provider)
	r.provider.SetDeps(r.desktop, r.de, r.rpc)

	for _, it := range r.items {
		it.initErrCh = it.item.Init()
	}

	for _, it := range r.items {
		if err = <-it.initErrCh; err != nil {
			r.runifyLogger.Error("Module initialization finished with error",
				zap.String("module", it.item.GetName()),
				zap.Error(err),
			)

			return false
		}
	}

	r.runifyLogger.Info("Init")
	return true
}

func (r *Runify) start() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	cases := make([]reflect.SelectCase, len(r.items)+1)
	for i, it := range r.items {
		m := it.item
		errCh := reflect.ValueOf(m.Start(ctx, wg))
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: errCh}
	}

	cfgIndex := len(r.items)
	errCh := reflect.ValueOf(r.cfg.Start(ctx, wg, r.logger.GetRoot()))
	cases[cfgIndex] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: errCh}

	r.runifyLogger.Info("Start")

	chosen, recv, recvOk := reflect.Select(cases)

	moduleName := config.ModuleName
	if chosen != cfgIndex {
		moduleName = r.items[chosen].name
	}
	action := zap.String("action", "Start to cancel the work of the other modules")

	if recvOk {
		r.runifyLogger.Error("Module finished without error",
			zap.String("module", moduleName),
			action)
	} else {
		err := recv.Interface().(error)

		r.runifyLogger.Error("Module finished with error",
			zap.String("module", moduleName),
			action,
			zap.Error(err))
	}

	cancel()
	wg.Wait()
}

func (r *Runify) Run(cfgFile string, cfgSave bool, buildCfg *config.BuildCfg) {
	if err := r.create(buildCfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if !r.init(cfgFile, cfgSave) {
		if r.runifyLogger != nil {
			_ = r.runifyLogger.Sync()
		}
		return
	}

	r.start()
}
