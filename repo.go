package main

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitRepo() {
	db = InitDb()
}

func CloseRepo() {
	CloseDb(db)
}

func RepoListTodo() (Todos, error) {
	var todos Todos
	if result := db.Find(&todos); result.Error != nil {
		return Todos{}, fmt.Errorf("Fetch data failed")
	}
	return todos, nil
}

func RepoCreateTodo(t Todo) (Todo, error) {
	if result := db.Create(&t); result.Error != nil {
		return t, fmt.Errorf("Create failed")
	}
	return t, nil
}

func RepoDestroyTodo(id int) error {
	var todo Todo
	if result := db.Where("id = ?", strconv.Itoa(id)).First(&todo); result.Error != nil {
		return fmt.Errorf("id = %d not found", id)
	}
	if result := db.Delete(&todo); result.Error != nil {
		return fmt.Errorf("Delete failed: id = %d", id)
	}
	return nil
}

func RepoFindTodo(id int) (Todo, error) {
	var todo Todo
	if result := db.Where("id = ?", strconv.Itoa(id)).First(&todo); result.Error != nil {
		return todo, fmt.Errorf("id = %d not found", id)
	}
	return todo, nil
}
