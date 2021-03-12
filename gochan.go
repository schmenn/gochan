package gochan

import (
	"net/http"
	"time"
)

//
const (
	BoardsEndpoint     string = "https://a.4cdn.org/boards.json"
	ThreadEndpoint     string = "https://a.4cdn.org/board/thread/id.json"
	ThreadListEndpoint string = "https://a.4cdn.org/board/threads.json"
	CatalogEndpoint    string = "https://a.4cdn.org/board/catalog.json"
)

// Client gochan client
type Client struct {
	httpClient http.Client
	Cache      ClientCache
}

// ClientCache the cache of the client
type ClientCache struct {
	Boards BoardStructure
}

// New initializes a new client
func New() Client {
	return Client{
		httpClient: http.Client{Timeout: time.Second * 5},
		Cache:      ClientCache{},
	}
}
