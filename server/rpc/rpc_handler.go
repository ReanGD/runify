package rpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type rpcHandler struct {
	uiBinaryPath     string
	unixAddr         string
	netUnixAddr      *net.UnixAddr
	showUIMultiplier *showUIMultiplier
	grpcServer       *grpc.Server
	runifyServer     *runifyServer

	moduleLogger *zap.Logger
}

func newRpcHandler() *rpcHandler {
	return &rpcHandler{
		uiBinaryPath:     "",
		unixAddr:         "",
		netUnixAddr:      nil,
		showUIMultiplier: newShowUIMultiplier(),
		grpcServer:       nil,
		runifyServer:     nil,
		moduleLogger:     nil,
	}
}

func (h *rpcHandler) onInit(
	cfg *config.Configuration, provider api.Provider, uiLogger *zap.Logger, moduleLogger *zap.Logger,
) error {
	h.moduleLogger = moduleLogger
	h.uiBinaryPath = cfg.System.UIBinaryPath

	var err error
	h.unixAddr = cfg.System.RpcAddress
	h.netUnixAddr, err = h.resolveUnixAddr(h.unixAddr)
	if err != nil {
		moduleLogger.Error("Failed resolve unit address", zap.String("address", h.unixAddr), zap.Error(err))
		return errors.New("failed resolve unit address")
	}
	h.grpcServer = grpc.NewServer()

	uiHandler := newUIHandler(uiLogger, moduleLogger)
	h.runifyServer = newRunifyServer(provider, h.showUIMultiplier, uiHandler, h.moduleLogger)

	return nil
}

func (h *rpcHandler) onStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		wg.Add(1)
		lis, err := net.ListenUnix("unix", h.netUnixAddr)
		if err != nil {
			h.moduleLogger.Error("Can't start listener for grps unix socket", zap.String("address", h.unixAddr), zap.Error(err))
			errCh <- errors.New("failed listen unix address")
			return
		}

		defer lis.Close()
		defer os.Remove(h.unixAddr)

		fi, err := os.Stat(h.unixAddr)
		if err == nil {
			os.Chmod(h.unixAddr, fi.Mode()|0o777)
		} else {
			h.moduleLogger.Info("Couldn't set permission for grps unix socket", zap.String("address", h.unixAddr), zap.Error(err))
		}

		pb.RegisterRunifyServer(h.grpcServer, h.runifyServer)
		if err = h.grpcServer.Serve(lis); err != nil {
			h.moduleLogger.Error("Grpc server finished with error", zap.String("address", h.unixAddr), zap.Error(err))
		} else {
			h.moduleLogger.Error("Grpc server finished", zap.String("address", h.unixAddr))
		}
		wg.Done()
		h.grpcServer = nil
		errCh <- err
	}()

	return errCh
}

func (h *rpcHandler) resolveUnixAddr(unixAddr string) (*net.UnixAddr, error) {
	mode, err := paths.LStatMode(unixAddr)
	if err == nil {
		if mode != paths.ModeSocket {
			return nil, errors.New("file exists but is not a socket")
		}

		if err = os.Remove(unixAddr); err != nil {
			return nil, fmt.Errorf("can't remove exists file: %s", err)
		}

		if ok, _ := paths.Exists(unixAddr); ok {
			return nil, errors.New("can't remove exists file: unknown error")
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("can't get stat for file: %s", err)
	}

	return &net.UnixAddr{Name: unixAddr, Net: "unix"}, nil
}

func (h *rpcHandler) onStop() {
	if h.grpcServer != nil {
		h.grpcServer.Stop()
	}
}

func (h *rpcHandler) onShowUI() {
	if !h.showUIMultiplier.sendToAll() {
		cmd := exec.Command(h.uiBinaryPath)
		if err := cmd.Start(); err != nil {
			h.moduleLogger.Error("Failed start runify UI process", zap.String("binary", h.uiBinaryPath), zap.Error(err))
			return
		}

		h.moduleLogger.Debug("Runify UI process started", zap.String("binary", h.uiBinaryPath))
		go cmd.Wait()
	}
}
