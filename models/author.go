package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string
	Age  int
}
