package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
    log.WithFields(log.Fields{
        "component": "logrus-example",
        "attempt":   1,
    }).Info("starting up")

    log.WithField("module", "example").Debug("debug message")
    log.WithError(nil).Error("example error")
}
