package models

import (
	"gorm.io/gorm"
)

type Relation struct {
	ID 				uint
	Observer 		uint
	Observee		uint
	Familiarity		uint
	Judgment		uint
}
