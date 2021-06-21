package models

import "gorm.io/gorm"

type Puppy struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
	UserId int `json:"user_id" binding:"required"`
}
