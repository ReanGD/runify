package root

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/provider"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

var logModule = zap.String("module", "runify")

type Runify struct {
	appID    uuid.UUID
	cfg      *config.Config
	logger   *logger.Logger
	provider *provider.Provider
	rpc      *rpc.Rpc

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
	r.provider = provider.New()
	r.rpc = rpc.New()
	r.runifyLogger = nil

	return nil
}

func (r *Runify) checkOnInit(name string, ch <-chan error) error {
	err := <-ch
	if err != nil {
		r.runifyLogger.Error("OnInit finished with error",
			zap.String("module", name),
			zap.Error(err))
		return err
	}
	return nil
}

func (r *Runify) init(cfgFile string) bool {
	if len(cfgFile) == 0 {
		cfgFile = filepath.Join(paths.GetAppConfig(), "runify.cfg")
	}
	r.cfg.OnInit(cfgFile)

	var err error
	if r.logger, err = logger.New(r.cfg, r.appID.String()); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load logger: %s", err)
		return false
	}

	rootLogger := r.logger.GetRoot()
	r.runifyLogger = rootLogger.With(logModule)
	r.cfg.AddVersionToLog(rootLogger)

	providerCh := r.provider.OnInit(r.cfg, rootLogger)
	rpcCh := r.rpc.OnInit(r.cfg, rootLogger, r.provider)

	if err = r.checkOnInit(provider.ModuleName, providerCh); err != nil {
		return false
	}

	if err = r.checkOnInit(rpc.ModuleName, rpcCh); err != nil {
		return false
	}

	r.runifyLogger.Info("Init")
	return true
}

func (r *Runify) start() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	cfgCh := r.cfg.OnStart(ctx, wg, r.logger.GetRoot())
	providerCh := r.provider.OnStart(ctx, wg)
	rpcCh := r.rpc.OnStart(ctx, wg)

	r.runifyLogger.Info("Start")

	var err error
	var name string
	select {
	case err = <-cfgCh:
		name = config.ModuleName
	case err = <-providerCh:
		name = provider.ModuleName
	case err = <-rpcCh:
		name = rpc.ModuleName
	}

	r.runifyLogger.Error("Module finished with error. Start to cancel the work of the other modules",
		zap.String("module", name),
		zap.Error(err))

	cancel()
	wg.Wait()
}

func (r *Runify) Run(cfgFile string, buildCfg *config.BuildCfg) {
	if err := r.create(buildCfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if !r.init(cfgFile) {
		if r.runifyLogger != nil {
			_ = r.runifyLogger.Sync()
		}
		return
	}

	r.start()
}