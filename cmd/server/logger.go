package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// buildLogger builds an application-level logger.
func buildLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// Set logger level.
	if os.Getenv("GOENV") == "development" {
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}
