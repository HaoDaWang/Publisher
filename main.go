package main

import (
	"net/http"
	"publisher/routes"
)

func main(){
	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/task/add", routes.AddTask)
	mux.HandleFunc("/task/all", routes.GetTasks)
	mux.HandleFunc("/task/remove", routes.RemoveTask)
	mux.HandleFunc("/branches", routes.GetBranches)

	server := Server{
		Host: "0.0.0.0",
		Port: "3000",
		Handler: mux,
	}

	server.listen()
}