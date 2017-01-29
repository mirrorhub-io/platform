package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mirrorhub-io/platform/controllers"
	"github.com/mirrorhub-io/platform/models"
)

func main() {
	log.Info("Starting server.")
	controllers.StartServer()
	defer models.Connection().Close()
}
