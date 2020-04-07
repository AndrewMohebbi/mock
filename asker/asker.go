package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var nameFlag = flag.String("name", "", "Which name do you want progress for?")

func main() {
	flag.Parse()

	//Make the request
	resp, err := http.Get("http://localhost:8100/" + *nameFlag)
	if err != nil {
		fmt.Println("Error making request!")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body!")
	}

	fmt.Println(string(body))
}
