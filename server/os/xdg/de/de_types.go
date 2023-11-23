package de

import (
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type updateCmd struct{}

func (c *updateCmd) OnRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "update"),
		zap.String("Reason", reason),
		zap.String("Action", "do nothing"))
}

type subscribeCmd struct {
	ch     chan<- types.DesktopEntries
	result api.BoolResult
}

func (c *subscribeCmd) OnRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "subscribe"),
		zap.String("Reason", reason),
		zap.String("Action", "subscription not activated, return false"))
	c.result.SetResult(false)
}
