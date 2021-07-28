package main

import (
	"crud-restfull-api-1/app/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	router := routes.NewRouter()

	fmt.Println("Service started at http://localhost:5000")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err.Error())
	}

}
