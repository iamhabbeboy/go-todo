package database

import (
	"github.com/iamhabbeboy/todoapp/models"
)

func SeedTodo() {

	db := Init()

	db.DropTable(&models.Todo{})

	db.AutoMigrate(&models.Todo{})

	todo := models.Todo{
		Title:     "Creating todo app with golang",
		Completed: false,
	}

	db.Create(&todo)
}
