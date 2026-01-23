package main

import (
    "go.uber.org/zap"
)

func main() {
    // production logger: zap.NewProduction()
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("starting up", zap.String("component", "zap-example"))
    logger.Debug("debug message", zap.Int("attempt", 2))
    logger.Error("example error", zap.String("err", "simulated"))
}
