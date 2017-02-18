package github

import (
	"net/url"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"
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
