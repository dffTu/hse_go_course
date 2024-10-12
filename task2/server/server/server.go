package server

import (
	"fmt"
	"net/http"
	"task2/server/request"
)

type Server struct {
	address  string
	mux      *http.ServeMux
	requests []request.Request
}

func CreateServer(address string) Server {
	return Server{address: address, mux: http.NewServeMux(), requests: make([]request.Request, 0)}
}

func (server *Server) AddRequest(request request.Request) {
	server.requests = append(server.requests, request)
}

func (server *Server) Start() {
	for _, val := range server.requests {
		server.mux.HandleFunc(val.Path, val.Handler)
	}
	httpServer := http.Server{
		Addr:    server.address,
		Handler: server.mux,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
