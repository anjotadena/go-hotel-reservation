package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Page hit...")

		next.ServeHTTP(w, r)
	})
}
