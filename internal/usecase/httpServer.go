package httpServer

import (
	"net/http"

	"github.com/kproject/Streaming-event-ingest/internal/domain"
)

type HTTPServer struct {
	cache map[int32]domain.Event
}

func CreateHTTPServer() error {
	err := http.ListenAndServe(":8080", nil)
	return err
}
