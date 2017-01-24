package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model

	Name               string
	Storage            int64
	TrafficConsumption int64
}
