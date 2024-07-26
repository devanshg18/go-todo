package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devanshg18/go-todo/servers/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting this server on port number 8080")
	log.Fatal(http.ListenAndServe(":8080  ", r))
}
