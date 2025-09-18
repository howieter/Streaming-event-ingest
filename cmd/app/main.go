package main

import (
	"fmt"

	"github.com/kproject/Streaming-event-ingest/internal/usecase"
)

func main() {
	fmt.Println("Creating server with port 8080")
	err := httpServer.CreateHTTPServer()

	if err != nil {
		fmt.Println(err)
	}
}
