package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	RequestType int    `json:"request_type"`
	Equation    string `json:"equation"`
}

func NewRequest(requestType int) Request {
	return Request{RequestType: requestType}
}

type Response struct {
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

var currentStatus string
var currentResult float64

func AgentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		var resp Response
		resp.Status = currentStatus
		resp.Result = 0

		if currentStatus == "done" {
			resp.Result = currentResult
		}

		r.Header.Set("Content-Type", "application/json")
		_, err := fmt.Fprint(w, resp)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		currentStatus = "counting"
	}
}

func main() {
	currentStatus = "free"
	http.Handle("/", http.HandlerFunc(AgentHandler))
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		panic(err)
	}

}
