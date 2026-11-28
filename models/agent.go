package models

import (
	"gorm.io/gorm"
)

type Person struct {
	ID 				uint
	Willingness 	uint
	Susceptibility	uint
}
