package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Email     string
}
