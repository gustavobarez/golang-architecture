package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	godotenv.Load()

	r := httprouter.New()

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
