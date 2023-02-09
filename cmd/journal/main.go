package main

import (
	"github.com/dharanad/journal/internal/base"
	"log"
)

func main() {
	logger, err := base.NewZapLogger()
	defer logger.Sync()
	if err != nil {
		log.Fatal("Error setting up zap logger", err)
	}
	server := base.NewJournalRouter(logger)
	r := server.SetupRouter()
	server.RegisterRoutes(r)
	if err := r.Run(":80"); err != nil {
		panic("Error starting web server. Reason :" + err.Error())
	}
}
