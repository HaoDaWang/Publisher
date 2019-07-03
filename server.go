package main

import (
	"fmt"
	"net/http"
	"publisher/utils"
)

type Server struct {
	Host string
	Port string
	Handler *http.ServeMux
}

func (server *Server) listen(){
	addr := fmt.Sprintf("%s:%s", server.Host, server.Port)

	httpServer := http.Server{
		Addr: addr,
		Handler: server.Handler,
	}

	utils.Log("Listening...")

	httpServer.ListenAndServe()
}