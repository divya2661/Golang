package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

func Index (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, world")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	type Todos []Todo
	todos := Todos{
		Todo{Name: "Write Presentation"},
		Todo{Name: "Host meetup"},
	}
	if err:= json.NewEncoder(w).Encode(todos); err !=nil {
		panic(err)
	}
	
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId  := vars["todoId"]
	fmt.Fprintln(w,"Todo Show: ",todoId)
}