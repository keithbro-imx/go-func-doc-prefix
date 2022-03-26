package utils

import (
	log "github.com/sirupsen/logrus"
)

// Yo dawg this is Hello in utils.
func Hello() string {
	log.Debugln("In Hello.utils()")

	return "Hello, world."
}
