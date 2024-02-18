package main

import (
	"awesomeProject2/orchestrator"
	"net/http"
)

func main() {
	http.Handle("/hello", OrchestratorMiddleware(http.HandlerFunc(orchestrator.OrchestratorHandler)))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
