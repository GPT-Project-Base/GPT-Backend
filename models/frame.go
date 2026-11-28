package models

import (
	"gorm.io/gorm"
)

type Frame struct {
	ID 				uint
	PersonID		uint
	NodeID			uint
	Frame			uint
	InterID			uint
}
