package main

import (
	"fmt"
	"net/http"

	"github.com/awebisam/go-web/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about/", handlers.About)

	fmt.Printf("Starting server in port %s", portNumber)

	err := http.ListenAndServe(
		portNumber,
		nil,
	)
	if err != nil {
		return
	}
}
