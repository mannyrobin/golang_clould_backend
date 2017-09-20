
// Author: 		Aksel Hjerpbakk
// Studentnr:	992

// SOURCES:
// https://golang.org/pkg/net/http/
// for most of the connection and JSON use:
// https://blog.alexellis.io/golang-json-api-client/

package main

import (
	"net/http"
	"strings"
	"time"
	"encoding/json"
	"errors"
)
// TODO: make sure repo exist!

func getGitRepoUrl(rawUrl *http.Request) ([]string, error){

	gitRepo := strings.Split(rawUrl.URL.Path, "/")

	// Check if url is valid | [4] = owner in URL | [5] = repository in URL
	if len(gitRepo) >= 6 && gitRepo[3] == "github.com" {
		return gitRepo, nil
	}
	return gitRepo, errors.New("Invalid URL!")
}



func handlerGitURL(w http.ResponseWriter, r *http.Request){
	http.Header.Add(w.Header(), "content-type", "application/json")
	rawUrl, error := getGitRepoUrl(r)
	if error != nil{
		panic(error)
	}


	apiUrl := "https://api." + rawUrl[3] + "/repos/" + rawUrl[4] + "/" + rawUrl[5]

	myClient := http.Client{
		Timeout: time.Second * 2,
	}

	repoBody, repBError := getBody(apiUrl, myClient)
	if repBError != nil{
		panic(repBError)
	}

	repoData := RepoData{}
	jsonError := json.Unmarshal(repoBody, &repoData)
	if jsonError != nil{
		panic(jsonError)
	}

	contribBody, contStatus := getBody(repoData.Contributors, myClient)
	if contStatus != nil{
		panic(contStatus)
	}

	langBody, langStatus := getBody(repoData.Languages, myClient)
	if langStatus != nil{
		panic(langStatus)
	}

	Best, bestError := getContributor(contribBody)
	if bestError != nil {
		panic(bestError)
	}

	Language, langError := getLanguages(langBody)
	if langError != nil{
		panic(langError)
	}


	presentedData := PresentedData{}

	presentedData.Project		= rawUrl[3] + "/"  + rawUrl[4] + "/" + rawUrl[5] + "/"
	presentedData.Owner			= repoData.Owner.Name
	presentedData.Committer		= Best.Name
	presentedData.Commits		= Best.Contributes
	presentedData.Language 		= Language


	json.Marshal(&presentedData)
	json.NewEncoder(w).Encode(presentedData)

}


func main(){
	http.HandleFunc("/projectinfo/v1/", handlerGitURL)
	http.ListenAndServe("localhost:8080", nil)
}