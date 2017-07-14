package github

import (
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp,err := 

}
