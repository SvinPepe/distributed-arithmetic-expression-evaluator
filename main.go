package main

import (
	"awesomeProject2/db"
	. "awesomeProject2/orchestrator"
	"fmt"
	"net/http"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		select {
		case <-ticker.C:
			queue, err := db.GetQueue()
			if err != nil {
				panic(err)
			}

			if len(queue) > 0 {

				resp, err := SendGetRequestToDemon()
				if err != nil {
					panic(err)
				}
				if resp.Status == "done" {
					err := db.SetResult(queue[0].Expression, resp.Result)
					if err != nil {
						panic(err)
					}
				}
				if resp.Status == "free" {
					req := NewRequest(1)
					req.Equation = queue[0].Expression
					resp, err := SendPostRequestToDemon(req)

					if resp == nil {
						http.Error(w, err.Error(), 500)
					}
				}

				fmt.Fprintf(w, resp.Status)

			} else {
				fmt.Fprint(w, "empty queue")
			}
		default:
			fmt.Fprint(w, "burgers")
		}
	})
	http.Handle("/hello", OrchestratorMiddleware(http.HandlerFunc(OrchestratorHandler)))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
