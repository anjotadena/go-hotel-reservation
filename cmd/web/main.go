package main

import (
	"fmt"
	"net/http"

	"github.com/anjotadena/projectX/pkg/handlers"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
