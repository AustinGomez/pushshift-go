package submission

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const submissionSearchBaseURL = "https://api.pushshift.io/reddit/submission/search/"

// Submission is the struct representing a submission.
type Submission struct {
	IsRedditMediaDomain   bool    `json:"is_reddit_media_domain"`
	WhitelistStatus       string  `json:"whitelist_status"`
	ParentWhitelistStatus string  `json:"parent_whitelist_status"`
	NoFollow              bool    `json:"no_follow"`
	SendReplies           bool    `json:"send_replies"`
	LinkFlairCSSClass     string  `json:"link_flair_css"`
	LinkFlairText         string  `json:"link_flair_text"`
	NumCrossposts         int     `json:"num_crossposts"`
	Over18                bool    `json:"over_18"`
	Locked                bool    `json:"locked"`
	Spoiler               bool    `json:"spoiler"`
	IsVideo               bool    `json:"is_video"`
	IsSelf                bool    `json:"is_self"`
	IsOriginalContent     bool    `json:"is_original_content"`
	IsCrosspostable       bool    `json:"is_crosspostable"`
	CanGuild              bool    `json:"can_guild"`
	Title                 string  `json:"title"`
	Selftext              string  `json:"selftext"`
	URL                   string  `json:"URL"`
	Domain                string  `json:"domain"`
	Media                 Media   `json:"media"`
	Preview               Preview `json:"preview"`
	MediaOnly             bool    `json:"media_only"`
	NumComments           int     `json:"num_comments"`
	Pinned                bool    `json:"pinned"`
	PostHint              string  `json:"post_hint"`
	Pwls                  int     `json:"pwls"`
	RemovedBy             string  `json:"removed_by"`
	Score                 int     `json:"score"`
}

// SearchParams holds query params for a submission search.
type SearchParams struct {
	Over18            bool   `url:"over_18,omitempty"`
	Locked            bool   `url:"locked,omitempty"`
	Spoiler           bool   `url:"spoiler,omitempty"`
	IsVideo           bool   `url:"is_video,omitempty"`
	IsSelf            bool   `url:"is_self,omitempty"`
	IsOriginalContent bool   `url:"is_original_content,omitempty"`
	IsCrosspostable   bool   `url:"is_crosspostable,omitempty"`
	CanGuild          bool   `url:"can_guild,omitempty"`
	Title             string `url:"title,omitempty"`
	Selftext          string `url:"selftext,omitempty"`
	URL               string `url:"url,omitempty"`
	Domain            string `url:"domain,omitempty"`

	// Common filters. Not sure how to abstract these out while keeping the API nice.
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

// List is a struct representing a list of Submissions returned from the PushShift API.
type List struct {
	Data []*Submission `json:"data"`
}

// Client is used to make the HTTP requests for Submissions.
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

	response, fetchErr := c.Backend.Get(submissionSearchBaseURL + "?" + v.Encode())
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

// Media is a struct within Submisison that holds information about embedded media.
type Media struct {
	Oembed Oembed `json:"oembed"`
	Type   string `json:"type"`
}

// Oembed is a struct within Media that holds information about embedded media.
type Oembed struct {
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	Height          int    `json:"height"`
	HTML            string `json:"html"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	Version         string `json:"version"`
	Width           int    `json:"width"`
}

// Preview is a struct within Submission that holds image preview information
type Preview struct {
	Enabled bool
	Images  []Image
}

// Image holds metadata about images
type Image struct {
	ID          string
	Resolutions []Resolution
	Variants    interface{}
}

// Resolution holds metadata about image resolutions.
type Resolution struct {
	Height int
	URL    string
	Width  int
}
