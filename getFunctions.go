package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetContributor ...
// @Return: The first contributor in array
func GetContributor(body []byte) (Contributer, error) {

	Contributers := []Contributer{}

	jsonError := json.Unmarshal(body, &Contributers)
	if jsonError != nil {
		return Contributers[0], jsonError
	}

	return Contributers[0], nil
}

// GetLanguages ...
// @Source: Nataniel Gåsøy, modified by author
// @Return: Array with all languages with random order
func GetLanguages(body []byte) ([]string, error) {

	// TODO: fix random order!
	var Languages []string

	LanguageMap := make(map[string]int)
	jsonError := json.Unmarshal(body, &LanguageMap)
	if jsonError != nil {
		return Languages, jsonError
	}

	for key := range LanguageMap {
		Languages = append(Languages, key)
	}

	return Languages, nil
}

// GetBody ...
// @Return: Body and error/nil
func GetBody(url string, myClient http.Client) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		body := []byte(nil)
		return body, err
	}

	// Apply header to request
	req.Header.Set("User-Agent", "projectinfo")

	// Try to execute request
	res, doError := myClient.Do(req)
	if doError != nil {
		body := []byte(nil)
		return body, doError
	}

	body, readError := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if readError != nil {
		rtrError := readError
		return body, rtrError
	}

	//SaveBodyTestData(url, body)

	return body, nil
}

// GetGitRepoURL ...
// @Return: []string of url based on request and error
func GetGitRepoURL(rawURL *http.Request) ([]string, error) {

	gitRepo := strings.Split(rawURL.URL.Path, "/")

	// Check if url is valid | [4] = owner in URL | [5] = repository in URL
	if len(gitRepo) >= 6 && gitRepo[3] == "github.com" {
		return gitRepo, nil
	}
	return gitRepo, errors.New("invalid URL")
}

// GetHTTP403 ...
// Redirect user to errorpage
func GetHTTP403(w http.ResponseWriter, failed string) {
	w.WriteHeader(http.StatusForbidden)
	http.Error(w, "Could not get "+failed+" data", 403)

}

// GetHTTP500 ...
// Redirect user to errorpage
func GetHTTP500(w http.ResponseWriter, failed string) {
	http.Error(w, "Could not process"+failed+"data", 500)
}
