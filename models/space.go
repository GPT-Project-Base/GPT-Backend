package models

import (
	"gorm.io/gorm"
)

type Environment struct {
	ID 				uint
	Connecting		[]uint
}
