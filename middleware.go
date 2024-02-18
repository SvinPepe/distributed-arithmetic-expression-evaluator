package main

import (
	"context"
	"fmt"
	"net/http"
)

func OrchestratorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		if r.Header.Get("X-AddExpression") != "" {
			ctx = context.WithValue(r.Context(), "addExpression", r.Header.Get("X-AddExpression"))
			defer fmt.Fprintf(w, "Expression: %s", r.Header.Get("X-AddExpression"))
		}
		if r.Header.Get("X-GetExpressions") != "" {
			ctx = context.WithValue(r.Context(), "getExpression", r.Header.Get("X-GetExpressions"))
			defer fmt.Fprintf(w, "need: %s", r.Header.Get("X-GetExpressions"))
		}
		if r.Header.Get("X-GetResult") != "" {
			ctx = context.WithValue(r.Context(), "getResult", r.Header.Get("X-GetResult"))
			defer fmt.Fprintf(w, "id: %s", r.Header.Get("X-GetResult"))
		}
		if r.Header.Get("X-GetOperations") != "" {
			ctx = context.WithValue(r.Context(), "getOperations", r.Header.Get("X-GetOperations"))
			defer fmt.Fprintf(w, " opers: %s ", r.Header.Get("X-GetOperations"))
		}

		if ctx != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
