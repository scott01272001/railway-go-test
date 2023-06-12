package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name  string
	Title string
	Body  string
}
