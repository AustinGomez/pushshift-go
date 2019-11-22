package comment

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const commentSearchBaseURL = "https://api.pushshift.io/reddit/comment/search/"

// Comment is the struct representing a submission.
// If it's unclear what the type is from the docs, type will be
// interface{}
type Comment struct {
	IsRedditMediaDomain          bool
	AllAwardings                 []interface{} `json:"all_awardings"`
	ApprovedAtUTC                int           `json:"approved_at_utc"`
	AssociatedAward              interface{}   `json:"associated_award"`
	Author                       string        `json:"author"`
	AuthorFlairBackgroundColor   interface{}   `json:"author_flair_background_color"`
	AuthorFlairRichtext          interface{}   `json:"author_flair_richtext"`
	AuthorFlairTextColor         interface{}   `json:"author_flair_text_color"`
	AuthorFlairType              string        `json:"author_flair_type"`
	AuthorFullname               string        `json:"author_fullname"`
	AuthorPatreonFlair           bool          `json:"author_patreon_flair"`
	Awarders                     []interface{} `json:"awarders"`
	BannedAtUTC                  int           `json:"banned_at_utc"`
	Body                         string        `json:"body"`
	CanModPost                   bool          `json:"can_mod_post"`
	Collapsed                    bool          `json:"collapsed"`
	CollapsedBecauseCrowdControl bool          `json:"collapsed_because_crowd_control"`
	CollapsedReason              interface{}   `json:"collapsed_reason"`
	CreatedUTC                   int           `json:"created_utc"`
	Distinguished                interface{}   `json:"distinguished"`
	Edited                       bool          `json:"edited"`
	Gildings                     interface{}   `json:"gildings"`
	ID                           string        `json:"id"`
	IsSubmitter                  bool          `json:"is_submitter"`
	LinkID                       string        `json:"link_id"`
	Locked                       bool          `json:"locked"`
	NoFollow                     bool          `json:"no_follow"`
	ParentID                     string        `json:"parent_id"`
	Permalink                    string        `json:"permalink"`
	RetrievedOn                  int           `json:"retrieved_on"`
	Score                        int64         `json:"score"`
	StewardReports               []interface{} `json:"steward_reports"`
	Sticked                      bool          `json:"stickied"`
	Subreddit                    string        `json:"subreddit"`
	SubredditID                  string        `json:"subreddit_id"`
	TotalAwardsReceived          int           `json:"total_awards_received"`
}

// SearchParams holds query params for a Comment search.
type SearchParams struct {
	Over18        bool `url:"over_18,omitempty"`
	ReplyDelay    int  `url:"reply_delay,omitempty"`
	NestLevel     int  `url:"nest_level,omitempty"`
	SubReplyDelay int  `url:"sub_reply_delay,omitempty"`
	UTCHourOfWeek int  `url:"utc_hour_of_week,omitempty"`
	LinkID        int  `url:"link_id,omitempty"`
	ParentID      int  `url:"parent_id,omitempty"`

	// Common Filters. Not sure how to abstract these out while keeping the API nice.
	Sort          string `url:"sort,omitempty"`
	SortType      string `url:"sort_type,omitempty"`
	After         int    `url:"after,omitempty"`
	Before        int    `url:"before,omitempty"`
	AfterID       int    `url:"after_id,omitempty"`
	BeforeID      int    `url:"before_id,omitempty"`
	CreatedUTC    int    `url:"created_utc,omitempty"`
	Score         int    `url:"score,omitempty"`
	Gilded        int    `url:"gilded,omitempty"`
	Edited        bool   `url:"edited,omitempty"`
	Author        string `url:"author,omitempty"`
	Subreddit     string `url:"subreddit,omitempty"`
	Distinguished string `url:"distinguished,omitempty"`
	RetrievedOn   int    `url:"retrieved_on,omitempty"`
	LastUpdated   int    `url:"last_updated,omitempty"`
	Q             string `url:"q,omitempty"`
	ID            int    `url:"id,omitempty"`
	Metadata      bool   `url:"metadata,omitempty"`
	Pretty        bool   `url:"pretty,omit"`
	Limit         int    `url:"limit,omitempty"`
}

// List is a struct representing a list of Comments returned from the PushShift API.
type List struct {
	Data []*Comment `json:"data"`
}

// Client is used to make the HTTP requests for Comments.
type Client struct {
	Backend *http.Client
}

// Search returns a List from the given params.
func Search(params SearchParams) List {
	return getClient().Search(params)
}

// Search returns a List from the given params.
func (c Client) Search(params SearchParams) List {
	v, err := query.Values(params)
	if err != nil {
		log.Panic(err)
	}

	response, fetchErr := c.Backend.Get(commentSearchBaseURL + "?" + v.Encode())
	if fetchErr != nil {
		log.Panic(fetchErr)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Panic(readErr)
	}

	var data List
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}
	return data

}

func getClient() Client {
	return Client{Backend: &http.Client{}}
}
