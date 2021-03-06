package main

import (
	"fmt"
	//"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	testJSON = `{"action":"test","number":123,"pull_request":{"html_url":"https://github.com/pr/url","head":{"repo":{"id":123,"name":"repo","full_name":"josemrobles/repo","owner":{"login":"josemrobles","id":123,"avatar_url":"https://avatars.githubusercontent.com/","name":"josemrobles","type":"User"},"html_url":"https://github.com/josemrobles/repo"}},"base":{"user":{"login":"creativedrive","id":123,"avatar_url":"https://avatars.githubusercontent.com","name":"josemrobles","type":"Organization"},"repo":{"id":123,"name":"repo","full_name":"josemrobles/repo","owner":{"login":"josemrobles","id":123,"avatar_url":"https: //avatars.githubusercontent.com","name":"josemrobles","type":"Organization"}}},"user":{"login":"josemrobles","id":123,"avatar_url":"https: //avatars.githubusercontent.com","name":"josemrobles","type":"Organization"}}}`
)

func getTestURL(action func(http.ResponseWriter, *http.Request)) string {
	server := httptest.NewServer(http.HandlerFunc(action))
	return fmt.Sprintf("%s/", server.URL)
}

func TestPing(t *testing.T) {
	reader := strings.NewReader(testJSON)
	request, err := http.NewRequest("POST", getTestURL(ping), reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 201 {
		t.Fatal("Expected 201 status code, received: ", res.StatusCode)
	}
}

func TestProcessPullRequest(t *testing.T) {
	reader := strings.NewReader(testJSON)
	request, err := http.NewRequest("POST", getTestURL(processPullRequest), reader)
	request.Header.Set("Token", "9543195005")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 201 {
		t.Fatal("Expected 201 status code, received: ", res.StatusCode)
	}
}
