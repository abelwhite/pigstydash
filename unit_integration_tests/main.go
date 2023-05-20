package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abelwhite/pigsty/router"
)

func main() {
	fmt.Println("MongoDB Connect")
	r := router.Router()
	fmt.Println("Server is starting..")
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000..")
}
