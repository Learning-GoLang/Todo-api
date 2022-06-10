package main

import (
	"encoding/json"
	"log"       //log error etc.
	"math/rand" //just bringing rand for random number
	"net/http"  //to work with http
	"strconv"   // string converter

	"github.com/gorilla/mux"
)

//Struct is almost class
// Todo Struct (Model)
type Todo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

//Init todos var

var todos []Todo

// Get All Todos
func getTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

// Get Single Todo
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //Get params
	//with id find

	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

// Create a new Todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

//Update a Todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(todos)
}

//Delete a Todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	//Mock Data

	todos = append(todos, Todo{ID: "1", Task: "TEST1", Status: false})
	todos = append(todos, Todo{ID: "2", Task: "TEST2", Status: true})
	// Route Handlers / Endpoints

	r.HandleFunc("/api/todo", getTodos).Methods("GET")
	r.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")

	//if error it will show
	log.Fatal(http.ListenAndServe(":8000", r)) // set port 8000
}
