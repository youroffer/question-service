package ogen

import (
	"fmt"
	"net/http"
)

func TestMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("from test md")
		next.ServeHTTP(w, r)
	})
}
