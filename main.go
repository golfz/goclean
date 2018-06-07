package main

import (
	"fmt"
	"log"
	"net/http"

	"goclean/router"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("====================================================")
	log.Println("Example GoClean's service is starting")

	router := router.Router()
	log.Print("The service is ready to listen and serve on port : 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}
