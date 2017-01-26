package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mirrorhub-io/platform/controllers"
	"github.com/mirrorhub-io/platform/models"
)

func main() {
	models.Connection()
	log.Info("Connection should be open.")
	controllers.StartServer()
	defer models.Connection().Close()
}
