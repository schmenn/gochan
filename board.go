package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type falseBool uint8

// BoardName Board Name
type BoardName string

// 4chan boards
const (
	Board3DCG                  BoardName = "3"
	BoardAnime                 BoardName = "a"
	BoardAdultCartoons         BoardName = "aco"
	BoardAdvice                BoardName = "adv"
	BoardAnimals               BoardName = "an"
	BoardAlternativeSports     BoardName = "asp"
	BoardRandom                BoardName = "b"
	BoardInternationalRandom   BoardName = "bant"
	BoardBusiness              BoardName = "biz"
	BoardAnimeCute             BoardName = "c"
	BoardCosplay               BoardName = "cgl"
	BoardCooking               BoardName = "ck"
	BoardCuteMale              BoardName = "cm"
	BoardComics                BoardName = "co"
	BoardHentaiAlternative     BoardName = "d"
	BoardDIY                   BoardName = "diy"
	BoardEcchi                 BoardName = "e"
	BoardFlash                 BoardName = "f"
	BoardFashion               BoardName = "fa"
	BoardFitness               BoardName = "fit"
	BoardTechnology            BoardName = "g"
	BoardGraphicDesign         BoardName = "gd"
	BoardAdultGIF              BoardName = "gif"
	BoardHentai                BoardName = "h"
	BoardHardcore              BoardName = "hc"
	BoardHistory               BoardName = "his"
	BoardHandsomeMen           BoardName = "hm"
	BoardHighRes               BoardName = "hr"
	BoardOekaki                BoardName = "i"
	BoardArtwork               BoardName = "ic"
	BoardInternational         BoardName = "int"
	BoardOtaku                 BoardName = "jp"
	BoardWeapons               BoardName = "k"
	BoardLGBT                  BoardName = "lgbt"
	BoardLiterature            BoardName = "lit"
	BoardMecha                 BoardName = "m"
	BoardPony                  BoardName = "mpl"
	BoardMusic                 BoardName = "mu"
	BoardTransportation        BoardName = "n"
	BoardNews                  BoardName = "news"
	BoardAuto                  BoardName = "o"
	BoardOutdoors              BoardName = "out"
	BoardPhotography           BoardName = "p"
	BoardOrigami               BoardName = "po"
	BoardPoliticallyIncorrect  BoardName = "pol"
	BoardProfessionalWrestling BoardName = "pw"
	BoardQandA                 BoardName = "qa"
	BoardQuests                BoardName = "qst"
	BoardAdultRequests         BoardName = "r"
	BoardROBOT9001             BoardName = "r9k"
	BoardSexyBeautifulWomen    BoardName = "s"
	BoardShit4chanSays         BoardName = "s4s"
	BoardScienceAndMath        BoardName = "sci"
	BoardCamsAndMeetups        BoardName = "soc"
	BoardSports                BoardName = "sp"
	BoardTorrents              BoardName = "t"
	BoardTraditionalGames      BoardName = "tg"
	BoardToys                  BoardName = "toy"
	BoardOffTopic              BoardName = "trash"
	BoardTravel                BoardName = "trv"
	BoardTelevisionAndFilm     BoardName = "tv"
	BoardYuri                  BoardName = "u"
	BoardVideoGames            BoardName = "v"
	BoardVideoGameGenerals     BoardName = "vg"
	BoardVeryImportantPosts    BoardName = "vip"
	BoardVideoGamesMultiplayer BoardName = "vm"
	BoardVideoGamesMobile      BoardName = "vmg"
	BoardPokémon               BoardName = "vp"
	BoardRetroGames            BoardName = "vr"
	BoardVideoGamesRPG         BoardName = "vrpg"
	BoardVideoGamesStrategy    BoardName = "vst"
	BoardVirtualYouTubers      BoardName = "vt"
	BoardAnimeWallpapers       BoardName = "w"
	BoardWallpapers            BoardName = "wg"
	BoardWorksafeGIF           BoardName = "wsg"
	BoardWorksafeRequests      BoardName = "wsr"
	BoardParanormal            BoardName = "x"
	BoardExtremeSports         BoardName = "xs"
	BoardYaoi                  BoardName = "y"
)

// Board 4chan board structure
type Board struct {
	// Tag tag the board appears as, e.g. "b"
	Tag BoardName `json:"board"`
	// Title the full title of the Board
	Title string `json:"title"`
	// IsWorkSafe whether the board is Safe-For-Work
	IsWorksafe falseBool `json:"ws_boards"`
	// ThreadsPerPage how many threads are shown per board page
	ThreadsPerPage uint8 `json:"per_page"`
	// Pages number of pages
	Pages uint `json:"pages"`
	// MaxFileSize Maximum File Size in kiloBytes
	MaxFileSize uint `json:"max_filesize"`
	// MaxWebmFileSize Maximum File Size for .webm files in kiloBytes
	MaxWebmFileSize uint `json:"max_webm_filesize"`
	// MaxCommentChars comment character limit
	MaxCommentChars uint16 `json:"max_comment_chars"`
	// MaxWebmDuration Maximum duration for .webm files
	MaxWebmDuration uint `json:"max_webm_duration"`
	// BumpLimit after this limit is reached bumps no longer bump the thread
	BumpLimit uint `json:"bump_limit"`
	// ImageLimit how many image comments can be posted
	ImageLimit uint `json:"image_limit"`
	// Cooldowns Board Cooldowns
	Cooldowns BoardCooldowns `json:"cooldowns"`
	// MetaDescription description of the board
	MetaDescription string `json:"meta_description"`
	// Spoilers if Spoilers are enabled
	Spoilers       falseBool `json:"spoilers,omitempty"`
	CustomSpoilers uint8     `json:"custom_spoilers,omitempty"`
	// IsArchived whether archives are enabled for the board
	IsArchived falseBool `json:"is_archived,omitempty"`
	// TrollFlags whether troll flags are enabled on the board
	TrollFlags     falseBool `json:"troll_flags,omitempty"`
	CountryFlags   falseBool `json:"country_flags,omitempty"`
	UserIDs        falseBool `json:"user_ids,omitempty"`
	Oekaki         falseBool `json:"oekaki,omitempty"`
	SjisTags       falseBool `json:"sjis_tags,omitempty"`
	CodeTags       falseBool `json:"code_tags,omitempty"`
	MathTags       falseBool `json:"math_tags,omitempty"`
	TextOnly       falseBool `json:"text_only,omitempty"`
	ForcedAnon     falseBool `json:"forced_anon,omitempty"`
	WebmAudio      falseBool `json:"webm_audio,omitempty"`
	RequireSubject falseBool `json:"require_subject,omitempty"`
	MinImageWidth  uint16    `json:"min_image_width,omitempty"`
	MinImageHeight uint16    `json:"min_image_height,omitempty"`
}

// BoardCooldowns Board Cooldowns
type BoardCooldowns struct {
	Threads uint `json:"threads"`
	Replies uint `json:"replies"`
	Images  uint `json:"images"`
}

// BoardStructure structure of the boards.json response
type BoardStructure struct {
	Boards     []Board           `json:"boards"`
	TrollFlags map[string]string `json:"troll_flags,omitempty"`
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

func (f falseBool) Bool() bool {
	return f == 1
}
