package main

import (
	"encoding/json"
	"fmt"
	"lrs/structs"
	"net/http"
)

//Hard coded data to load
var andyEvents = []structs.Event{
	structs.Event{
		Course: "Matlab", Section: 1, Completed: true,
	},
	structs.Event{
		Course: "Matlab", Section: 2, Completed: true,
	},
	structs.Event{
		Course: "Simulink", Section: 1, Completed: true,
	},
}
var andy = structs.Message{Name: "Andy", Events: andyEvents}

func main() {

	//Start listening
	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)

}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		send(w, r)
	default:
		fmt.Println("Wrong method in request!")
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	//Method: get by name
	//its a get request, it requires a name, and returns all events associated wiht that name
	andyMarshalled, _ := json.Marshal(andy)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(andyMarshalled)

}
