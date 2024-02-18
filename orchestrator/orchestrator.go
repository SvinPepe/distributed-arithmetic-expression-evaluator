package orchestrator

import (
	. "awesomeProject2/db"
	"fmt"
	"net/http"
	"strconv"
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

func OrchestratorHandler(w http.ResponseWriter, r *http.Request) {

	if val := r.Context().Value("addExpression"); val != nil {
		exp := Expression{Expression: val.(string), Status: "sent", Result: 0, TimeSpent: 0}

		err := AddExpressionToDB(exp)
		if err != nil {
			panic(err)
		}
	}
	if val := r.Context().Value("getExpression"); val != nil {

		exps, err := GetExpressions()

		for _, exp := range exps {
			fmt.Fprintf(w, exp.Expression+exp.Status)
		}
		if err != nil {
			panic(err)
		}
	}
	if val := r.Context().Value("getResult"); val != nil {

		convertedValue, err := strconv.Atoi(val.(string))
		if err != nil {
			fmt.Fprintf(w, "BIU")
			panic(err)
		}
		result, err := GetResult(convertedValue)

		fmt.Fprintf(w, "result : %f ", result)
		if err != nil {
			panic(err)
		}
	}
	if val := r.Context().Value("getOperations"); val != nil {

		exps, err := GetOperations()

		for _, oper := range exps {
			fmt.Fprintf(w, "operation:%s  time: %d\n", oper.Operation, oper.Time)
		}
		if err != nil {
			panic(err)
		}
	}

}

/*
curl -X GET http://localhost:8080/hello -H "X-AddExpression: 1 + 1"
curl -X GET http://localhost:8080/hello -H "X-GetExpressions: 1"
curl -X GET http://localhost:8080/hello -H "X-GetResult: 1
curl -X GET http://localhost:8080/hello -H "X-GetOperations: 1
*/
