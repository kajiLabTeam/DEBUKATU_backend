package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserId  int64
	Name    string
	Deleted bool
}
