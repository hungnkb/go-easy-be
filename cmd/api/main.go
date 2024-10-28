package main

import (
	"my_project/internal/server"
	"net/http"
	"os"
)

func main() {
	r := server.NewServer()
	port := ":" + string(os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
