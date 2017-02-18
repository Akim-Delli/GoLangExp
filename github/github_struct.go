// Package github provide a go API for the github issue tracker.
// See https://developer.github.com/v3/search/#search-issues/
package github

import "time"

const IssueUrl = "https://api.github.com/search/issues"

type SearchResults struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Title     string    `json:"title"`
	Number    int       `json:"number"`
	User      *User     `json:user`
	CreatedAt time.Time `json:"created_at""`
	HTMLURL   string    `json:"url"`
	State     string    `json:"state"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"url"`
}
