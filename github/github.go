package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the github issue tracker.
func SearchIssues(terms []string) (*SearchResults, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	url := IssueUrl + "?q=" + q
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	fmt.Println(url)
	return &result, nil
}
