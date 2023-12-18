package main

import (
	"time"

	"go.uber.org/zap"
)

const url = "http://google.com"

func main() {
	logger, _ := zap.NewProduction()
	// Flushes the buffers, if any
	defer logger.Sync()

	sugar := logger.Sugar()

	// In contexts where performance is nice, but no critical, use SugarredLogger
	// Which support loose (unstructured) logging
	var theSugarredLogger = func() {
		sugar.Infow("Failed to fetch URL",
			// Structured context as loosely typed key-value pairs
			"url", url,
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("Failed to fetch URL: %s", url)
	}

	// In contexts where performance and type safety are critical, use the Logger.
	// Which is even faster than the SuggaredLogger and allocates far less, but
	// it only supports structured loggin.
	var theStructuredLogger = func() {
		logger.Info("Failed to fetch URL",
			// Structured context as strongly typed field values.
			zap.String("url", url),
			zap.Int("attempt", 3),
			zap.Duration("backoff", 5*time.Second),
		)
	}

	theSugarredLogger()
	theStructuredLogger()
}
