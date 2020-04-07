package main

import (
	"fmt"
	"net/http"
	"service/calculate"
	"service/utils"
)

func main() {
	fmt.Println("Service started!")

	http.HandleFunc("/", router)
	http.ListenAndServe(":8100", nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	default:
		utils.Write(w, 405, "Method not supported!")
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	body := utils.LrsGet(r)
	report := calculate.Calculate(body)
	utils.Write(w, 200, report)
}
