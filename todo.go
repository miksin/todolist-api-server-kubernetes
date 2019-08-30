package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name     string
	Complete bool
	Due      time.Time
}

type Todos []Todo
