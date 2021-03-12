package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Thread structure of the thread
type Thread struct {
	Posts []Post `json:"posts"`
}

// Timestamp int64 alias
type Timestamp int64

// Post a 4chan post
type Post struct {
	Number               uint      `json:"no"`
	ResTo                uint      `json:"resto"`
	IsSticky             uint      `json:"sticky"`
	IsClosed             uint      `json:"closed"`
	Now                  string    `json:"now"`
	Timestamp            Timestamp `json:"time"`
	Author               string    `json:"name"`
	TripCode             string    `json:"trip"`
	ID                   string    `json:"id"`
	CapCode              string    `json:"capcode"`
	Country              string    `json:"country"`
	CountryName          string    `json:"country_name "`
	Subject              string    `json:"sub"`
	Comment              string    `json:"com"`
	ImageTimestamp       Timestamp `jsson:"tim"`
	Filename             string    `json:"filename"`
	Filetype             string    `json:"ext"`
	Filesize             uint      `json:"fsize"`
	MD5hash              string    `json:"md5"`
	ImageWidth           uint      `json:"w"`
	ImageHeight          uint      `json:"h"`
	ThumbnailWidth       uint      `json:"tn_w"`
	ThumbnailHeight      uint      `json:"tn_h"`
	FileDeleted          uint8     `json:"filedeleted"`
	IsSpoilered          uint8     `json:"spoiler"`
	CustomSpoiler        uint8     `json:"custom_spoiler"`
	Replies              uint      `json:"replies"`
	Images               uint      `json:"images"`
	BumpLimitReached     uint8     `json:"bumplimit"`
	ImageLimitReached    uint8     `json:"imagelimit"`
	Tag                  string    `json:"tag"`
	SemanticURL          string    `json:"semantic_url"`
	Since4Pass           string    `json:"since4pass"`
	UniquePosters        uint      `json:"unique_ips"`
	MobileOptimizedImage uint8     `json:"m_img"`
	IsArchived           uint8     `json:"archived"`
	ArchivedAt           Timestamp `json:"archived_on"`
}

// GetThread  returns a thread
func (c Client) GetThread(board BoardName, id string) (Thread, error) {
	var thread Thread

	url := strings.ReplaceAll(strings.ReplaceAll(ThreadEndpoint, "board", string(board)), "id", id)
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return thread, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return thread, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return thread, err
	}

	err = json.Unmarshal(body, &thread)
	if err != nil {
		return thread, err
	}
	//fmt.Println(string(body))
	return thread, nil
}

// GetThreads gets a list of threads from a board
func (c Client) GetThreads(board BoardName) (ThreadList, error) {
	var tl ThreadList
	url := strings.ReplaceAll(ThreadListEndpoint, "board", string(board))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return tl, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return tl, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tl, err
	}
	err = json.Unmarshal(body, &tl)
	if err != nil {
		return tl, err
	}
	return tl, nil
}

// ThreadList list of threads
type ThreadList []ThreadListElement

// ThreadListElement Element in ThreadList
type ThreadListElement struct {
	Page    uint                      `json:"page"`
	Threads []ThreadListElementThread `json:"threads"`
}

// ThreadListElementThread Element in ThreadListElement
type ThreadListElementThread struct {
	No           uint      `json:"no"`
	LastModified Timestamp `json:"last_modified"`
	Replies      uint      `json:"replies"`
}

// ToTime turns a unix timestamp into time.Time
func (t Timestamp) ToTime() time.Time {
	return time.Unix(int64(t), 0)
}

// FullFilename returns the full file name
func (p Post) FullFilename() string {
	return p.Filename + p.Filetype
}
