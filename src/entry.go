package main

import (
	"github.com/keithbro-imx/goplayground/mint"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	mint.Hello()
}
