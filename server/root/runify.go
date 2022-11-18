package root

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/logger"
	dsX11 "github.com/ReanGD/runify/server/os/x11"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/provider"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/ReanGD/runify/server/x11"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

var logModule = zap.String("module", "runify")

type Runify struct {
	appID    uuid.UUID
	cfg      *config.Config
	logger   *logger.Logger
	rpc      *rpc.Rpc
	x11      *x11.X11
	ds       *dsX11.X11
	provider *provider.Provider

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
	r.rpc = rpc.New()
	r.x11 = x11.New()
	r.ds = dsX11.New()
	r.provider = provider.New()
	r.runifyLogger = nil

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

	for _, it := range []struct {
		moduleName string
		initCh     <-chan error
	}{
		{rpc.ModuleName, r.rpc.OnInit(r.cfg, r.provider, rootLogger)},
		{x11.ModuleName, r.x11.OnInit(r.cfg, r.rpc, rootLogger)},
		{dsX11.ModuleName, r.ds.OnInit(r.cfg, rootLogger)},
		{provider.ModuleName, r.provider.OnInit(r.cfg, r.x11, rootLogger)},
	} {
		err := <-it.initCh
		if err != nil {
			r.runifyLogger.Error("OnInit finished with error",
				zap.String("module", it.moduleName),
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

	cfgCh := r.cfg.OnStart(ctx, wg, r.logger.GetRoot())
	rpcCh := r.rpc.OnStart(ctx, wg)
	x11Ch := r.x11.OnStart(ctx, wg)
	dsCh := r.ds.OnStart(ctx, wg)
	providerCh := r.provider.OnStart(ctx, wg)

	r.runifyLogger.Info("Start")

	var err error
	var name string
	select {
	case err = <-cfgCh:
		name = config.ModuleName
	case err = <-rpcCh:
		name = rpc.ModuleName
	case err = <-x11Ch:
		name = x11.ModuleName
	case err = <-dsCh:
		name = dsX11.ModuleName
	case err = <-providerCh:
		name = provider.ModuleName
	}

	r.runifyLogger.Error("Module finished with error. Start to cancel the work of the other modules",
		zap.String("module", name),
		zap.Error(err))

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
