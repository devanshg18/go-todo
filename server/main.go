package main

import (
	"fmt"
	"go-todo/server/router"
	"log"
	"net/http"
	//"go-todo/server/middleware"
)

func main() {
	r := router.Router()
	fmt.Println("Starting this server on port number 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
