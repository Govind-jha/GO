package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GithubApiEndpoints struct {
	CurrentUserURL                   string `json:"current_user_url"`
	CurrentUserAuthorizationsHTMLURL string `json:"current_user_authorizations_html_url"`
	AuthorizationsURL                string `json:"authorizations_url"`
	CodeSearchURL                    string `json:"code_search_url"`
	CommitSearchURL                  string `json:"commit_search_url"`
	EmailsURL                        string `json:"emails_url"`
	EmojisURL                        string `json:"emojis_url"`
	EventsURL                        string `json:"events_url"`
	FeedsURL                         string `json:"feeds_url"`
	FollowersURL                     string `json:"followers_url"`
	FollowingURL                     string `json:"following_url"`
	GistsURL                         string `json:"gists_url"`
	HubURL                           string `json:"hub_url"`
	IssueSearchURL                   string `json:"issue_search_url"`
	IssuesURL                        string `json:"issues_url"`
	KeysURL                          string `json:"keys_url"`
	NotificationsURL                 string `json:"notifications_url"`
	OrganizationRepositoriesURL      string `json:"organization_repositories_url"`
	OrganizationURL                  string `json:"organization_url"`
	PublicGistsURL                   string `json:"public_gists_url"`
	RateLimitURL                     string `json:"rate_limit_url"`
	RepositoryURL                    string `json:"repository_url"`
	RepositorySearchURL              string `json:"repository_search_url"`
	CurrentUserRepositoriesURL       string `json:"current_user_repositories_url"`
	StarredURL                       string `json:"starred_url"`
	StarredGistsURL                  string `json:"starred_gists_url"`
	TeamURL                          string `json:"team_url"`
	UserURL                          string `json:"user_url"`
	UserOrganizationsURL             string `json:"user_organizations_url"`
	UserRepositoriesURL              string `json:"user_repositories_url"`
	UserSearchURL                    string `json:"user_search_url"`
}

func main() {
	url := "https://api.github.com"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	client := &http.Client{}
	// A Client is an HTTP client

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var endpoints GithubApiEndpoints

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		log.Println(err)
	}

	fmt.Println("Emojis url - ", endpoints.EmojisURL);

}
