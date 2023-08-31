package main

import (
	"exam-pgapi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8070....")

	log.Fatal(http.ListenAndServe(":8070", r))
}