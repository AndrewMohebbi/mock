package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Write is an http write helper
func Write(w http.ResponseWriter, status int, body interface{}) {
	marshalled, _ := json.Marshal(body)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(marshalled)
}

//LrsGet queries lrs for records
func LrsGet(r *http.Request) []byte {
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

	return body
}
