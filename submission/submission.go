package submission

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const submissionSearchBaseURL = "https://api.pushshift.io/reddit/submission/search/"

// Submission is the struct representing a submission.
type Submission struct {
	IsRedditMediaDomain   bool   `json:"is_reddit_media_domain"`
	WhitelistStatus       string `json:"whitelist_status"`
	ParentWhitelistStatus string `json:"parent_whitelist_status"`
	NoFollow              bool   `json:"no_follow"`
	SendReplies           bool   `json:"send_replies"`
	LinkFlairCSSClass     string `json:"link_flair_css"`
	LinkFlairText         string `json:"link_flair_text"`
	NumCrossposts         int    `json:"num_crossposts"`
	Over18                bool   `json:"over_18"`
	Locked                bool   `json:"locked"`
	Spoiler               bool   `json:"spoiler"`
	IsVideo               bool   `json:"is_video"`
	IsSelf                bool   `json:"is_self"`
	IsOriginalContent     bool   `json:"is_original_content"`
	IsCrosspostable       bool   `json:"is_crosspostable"`
	CanGuild              bool   `json:"can_guild"`
	Title                 string `json:"title"`
	Selftext              string `json:"selftext"`
	URL                   string `json:"URL"`
	Domain                string `json:"domain"`
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
	Limit             int    `url:"limit,omitempty"`
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
	fmt.Printf("Params in Search %v", params)
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
