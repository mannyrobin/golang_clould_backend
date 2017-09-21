package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
	"errors"
)

// @Return:
func getContributor(body []byte) (Contributer, error) {

	Contributers := []Contributer{}

	jsonError := json.Unmarshal(body, &Contributers)
	if jsonError != nil{
		return Contributers[0], jsonError
	}

	return Contributers[0], nil
}

// Source: Nataniel Gåsøy, modified by author
// @Return: Array with all languages
func getLanguages(body []byte) ([]string, error) {

	var Languages []string
	LanguageMap := make(map [string] int)
	jsonError := json.Unmarshal(body, &LanguageMap)
	if jsonError != nil{
		return Languages, jsonError
	}
	for key:=range LanguageMap{
		Languages = append(Languages, key)
	}

	return Languages, nil
}

// @Return: Body and error/nil
func getBody(url string, myClient http.Client) ([]byte, error){

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		body := []byte(nil)
		return body, err
	}

	//TODO: make username changeable
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
	return body, nil
}

func getGitRepoUrl(rawUrl *http.Request) ([]string, error){

	gitRepo := strings.Split(rawUrl.URL.Path, "/")

	// Check if url is valid | [4] = owner in URL | [5] = repository in URL
	if len(gitRepo) >= 6 && gitRepo[3] == "github.com" {
		return gitRepo, nil
	}
	return gitRepo, errors.New("Invalid URL!")
}

func getHTTP403(w http.ResponseWriter, failed string){
	http.Error(w, "Could not get" + failed + "data" , 403)
}

func getHTTP500(w http.ResponseWriter, failed string){
	http.Error(w, "Could not process" + failed + "data" , 500)
}