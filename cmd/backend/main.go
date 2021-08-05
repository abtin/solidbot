package main

import (
	log "github.com/sirupsen/logrus"
	"solidbot/internal/backend"
)

func main() {
	be, err := backend.GetConfig("solidbot_backend")
	if err != nil {
		panic(err)
	}
	log.Info(be)
}
