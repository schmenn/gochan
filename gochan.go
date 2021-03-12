package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//
const (
	BoardsEndpoint string = "https://a.4cdn.org/boards.json"
	ThreadEndpoint string = "https://a.4cdn.org/board/thread/id.json"
	ThreadListEndpoint string = "https://a.4cdn.org/board/threads.json"
	CatalogEndpoint string = "https://a.4cdn.org/board/catalog.json"
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

// UpdateBoards Updates the Client's board cache
func (c *Client) UpdateBoards() error {
	req, err := http.NewRequest("GET", BoardsEndpoint, nil)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var b BoardStructure
	err = json.Unmarshal(body, &b)
	if err != nil {
		return err
	}
	c.Cache.Boards = b
	return nil
}
