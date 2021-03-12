package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Catalog post catalog
type Catalog []CatalogElement

// CatalogElement part of Catalog
type CatalogElement struct {
	Page    int64           `json:"page"`
	Threads []CatalogThread `json:"threads"`
}

// CatalogThread part of CatalogElement
type CatalogThread struct {
	Number          uint   `json:"no"`
	ResTo           uint   `json:"resto"`
	IsSticky        uint   `json:"sticky"`
	IsClosed        uint   `json:"closed"`
	Name            string `json:"name"`
	Sub             string `json:"sub,omitempty"`
	Comment         string `json:"com,omitempty"`
	Filename        string `json:"filename,omitempty"`
	EXT             string `json:"ext,omitempty"`
	ImageWidth      int64  `json:"w,omitempty"`
	ImageHeight     int64  `json:"h,omitempty"`
	ThumbnailWidth  int64  `json:"tn_w,omitempty"`
	ThumbnailHeight int64  `json:"tn_h,omitempty"`
	Tim             int64  `json:"tim,omitempty"`
	Time            int64  `json:"time"`
	MD5hash         string `json:"md5,omitempty"`
	Filesize        int64  `json:"fsize,omitempty"`
	Capcode         string `json:"capcode,omitempty"`
	SemanticURL     string `json:"semantic_url"`
	Replies         int64  `json:"replies"`
	Images          int64  `json:"images"`
	OmittedPosts    int64  `json:"omitted_posts,omitempty"`
	OmittedImages   int64  `json:"omitted_images,omitempty"`
	LastReplies     []Post `json:"last_replies"`
	LastModified    int64  `json:"last_modified"`
	Bumplimit       int64  `json:"bumplimit,omitempty"`
	Imagelimit      int64  `json:"imagelimit,omitempty"`
	Trip            string `json:"trip,omitempty"`
	Filedeleted     uint8  `json:"filedeleted,omitempty"`
}

// GetCatalog  returns a thread
func (c Client) GetCatalog(board BoardName) (Catalog, error) {
	var catalog Catalog

	url := strings.ReplaceAll(CatalogEndpoint, "board", string(board))
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return catalog, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return catalog, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return catalog, err
	}

	err = json.Unmarshal(body, &catalog)
	if err != nil {
		return catalog, err
	}
	//fmt.Println(string(body))
	return catalog, nil
}
