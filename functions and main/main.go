
// Author: 		Aksel Hjerpbakk
// Studentnr:	992

// SOURCES:
// https://golang.org/pkg/net/http/
// for most of the connection and JSON use:
// https://blog.alexellis.io/golang-json-api-client/

package main

import (
	"net/http"
	"time"
	"encoding/json"
)


func handlerGitURL(w http.ResponseWriter, r *http.Request){
	http.Header.Add(w.Header(), "content-type", "application/json")
	rawUrl, repoError := getGitRepoUrl(r)
	if repoError != nil{
		getHTTP500(w,"repository url")
	}


	apiUrl := "https://api." + rawUrl[3] + "/repos/" + rawUrl[4] + "/" + rawUrl[5]

	myClient := http.Client{
		Timeout: time.Second * 2,
	}

	repoBody, repBError := getBody(apiUrl, myClient)
	if repBError != nil{
		getHTTP403(w,"repository")
	}

	repoData := RepoData{}
	jsonError := json.Unmarshal(repoBody, &repoData)
	if jsonError != nil{
		getHTTP500(w,"repository")
	}

	if repoData.Message != "Not Found" {
		contribBody, contStatus := getBody(repoData.Contributors, myClient)


		if contStatus != nil {
			getHTTP403(w,"contributor")
		}

		langBody, langStatus := getBody(repoData.Languages, myClient)
		if langStatus != nil {
			getHTTP403(w,"language")
		}

		Best, bestError := getContributor(contribBody)
		if bestError != nil {
			getHTTP500(w,"contributor")
		}

		Language, langError := getLanguages(langBody)
		if langError != nil {
			getHTTP500(w,"language")
		}

		presentedData := PresentedData{}

		presentedData.Project = rawUrl[3] + "/" + rawUrl[4] + "/" + rawUrl[5] + "/"
		presentedData.Owner = repoData.Owner.Name
		presentedData.Committer = Best.Name
		presentedData.Commits = Best.Contributes
		presentedData.Language = Language

		json.Marshal(&presentedData)
		json.NewEncoder(w).Encode(presentedData)
	} else{
		http.Error(w, "Repository not found", 404)
	}

}


func main(){
	http.HandleFunc("/projectinfo/v1/", handlerGitURL)
	http.ListenAndServe("localhost:8080", nil)
}