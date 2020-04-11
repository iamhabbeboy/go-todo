package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamhabbeboy/todoapp/controllers"
	"github.com/iamhabbeboy/todoapp/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database.SeedTodo()

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.TodoList)
	r.HandleFunc("/add", controllers.TodoAdd)
	http.Handle("/", r)

	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	port, err := os.LookupEnv("PORT")
	if !err {
		log.Fatal("Port not found in env file")
	}
	fmt.Println("Server running at " + port)
	http.ListenAndServe(":"+port, nil)
}
