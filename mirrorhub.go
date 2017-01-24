package main

import (
    "./models"
    "./controllers"
    log "github.com/Sirupsen/logrus"
)

func main() {
  models.Connection()
  log.Info("Connection should be open.")
  controllers.StartServer()
  defer models.Connection().Close()
}
