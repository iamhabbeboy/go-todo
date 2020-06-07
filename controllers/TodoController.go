package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/iamhabbeboy/todoapp/database"
	"github.com/iamhabbeboy/todoapp/models"
)

func TodoAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		TodoForm(w, r)
	case "POST":
		ProcessTodoForm(w, r)
	}
}

func TodoForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func ProcessTodoForm(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	if title == "" {
		http.Redirect(w, r, "/", 302)
		return
	}
	db := database.Init()

	todo := models.Todo{
		Title:     title,
		Completed: false,
	}
	db.Create(&todo)

	http.Redirect(w, r, "/", 302)
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/list.html")
	if err != nil {
		log.Fatal("Template not found ")
	}

	db := database.Init()
	var todos []models.Todo

	query := db.Find(&todos)

	defer query.Close()

	view.Execute(w, todos)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Redirect(w, r, "/", 302)
		return
	}
	db := database.Init()
	var todos []models.Todo
	db.Where("id = ?", id).Delete(&todos)
	http.Redirect(w, r, "/", 302)
}
