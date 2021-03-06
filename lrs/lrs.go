package main

import (
	"encoding/json"
	"fmt"
	"lrs/structs"
	"net/http"
)

//Hard coded data
////////////////////////////////////////////////////////////////////////
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

var georgeEvents = []structs.Event{
	structs.Event{
		Course: "Matlab", Section: 1, Completed: true,
	},
	structs.Event{
		Course: "Matlab", Section: 2, Completed: true,
	},
	structs.Event{
		Course: "Matlab", Section: 3, Completed: true,
	},
	structs.Event{
		Course: "Symbolic Math", Section: 1, Completed: true,
	},
}
var george = structs.Message{Name: "George", Events: georgeEvents}

////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println("lrs started!")

	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)

}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		send(w, r)
	default:
		marshalled, _ := json.Marshal("Method not supported!")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(405)
		w.Write(marshalled)
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path[1:] {
	case "andy":
		marshalled, _ := json.Marshal(andy)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(marshalled)
	case "george":
		marshalled, _ := json.Marshal(george)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(marshalled)
	default:
		marshalled, _ := json.Marshal("Name doesn't exist!")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write(marshalled)
	}
	return
}
