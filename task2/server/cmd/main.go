package main

import (
	"task2/server/request"
	"task2/server/server"
)

func main() {
	server := server.CreateServer(":8080")
	requests := []request.Request{
		{Handler: request.PrintAPI, Path: "/version"},
		{Handler: request.Decode, Path: "/decode"},
		{Handler: request.HardOperation, Path: "/hard-op"},
	}
	for _, val := range requests {
		server.AddRequest(val)
	}
	server.Start()
}
