package main

import (
	"github.com/dharanad/journal/internal/base"
	"github.com/dharanad/journal/internal/base/server"
	"log"
)

func main() {
	logger, err := base.NewZapLogger()
	defer logger.Sync()
	if err != nil {
		log.Fatal("Error setting up zap logger", err)
	}
	r := server.NewServer()
	if err := r.Run(":80"); err != nil {
		panic("Error starting web server. Reason :" + err.Error())
	}
}
