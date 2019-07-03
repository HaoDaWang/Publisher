package main

import (
	"net/http"
	"publisher/routes"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.Index)

	server := Server{
		Host:"0.0.0.0",
		Port: "3000",
		Handler:mux,
	}

	server.listen()
}