package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Service started!")

	//Start listening
	http.HandleFunc("/", router)
	http.ListenAndServe(":8100", nil)
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

	resp, err := http.Get("http://localhost:8000/" + r.URL.Path[1:])
	if err != nil {
		fmt.Println("Error making request!")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body!")
	}

	fmt.Println(string(body))

	//marshalled, _ := json.Marshal(andy)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)

	return

}
