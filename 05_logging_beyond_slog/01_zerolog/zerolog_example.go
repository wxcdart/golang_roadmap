package main

import (
    "os"
    "time"

    "errors"
    "github.com/rs/zerolog"
)

func main() {
    // Console output for demo
    out := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
    logger := zerolog.New(out).With().Timestamp().Logger()

    logger.Info().Str("component", "zerolog-example").Msg("starting up")
    logger.Debug().Int("attempt", 3).Msg("debug info")
    logger.Error().Err(errors.New("simulated"))
    logger.Error().Msg("an example error")
}
