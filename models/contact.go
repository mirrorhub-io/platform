package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model

	Name    string
	EMail   string
	Mirrors []Mirror
}
