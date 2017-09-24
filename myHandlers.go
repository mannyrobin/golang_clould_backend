package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HandlerGitURL ...
// Functionality: controller
func HandlerGitURL(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	rawURL, repoError := GetGitRepoURL(r)
	if repoError != nil {
		GetHTTP403(w, "repository url")
		return
	}

	apiURL := "https://api." + rawURL[3] + "/repos/" + rawURL[4] + "/" + rawURL[5]

	myClient := http.Client{
		Timeout: time.Second * 2,
	}

	repoBody, repBError := GetBody(apiURL, myClient)
	if repBError != nil {
		GetHTTP403(w, "repository")
		return
	}

	repoData := RepoData{}
	jsonError := json.Unmarshal(repoBody, &repoData)
	if jsonError != nil {
		GetHTTP403(w, "repository")
		return
	}

	if repoData.Message != "Not Found" {
		contribBody, contStatus := GetBody(repoData.Contributors, myClient)
		if contStatus != nil {
			GetHTTP403(w, "contributor")
			return
		}

		langBody, langStatus := GetBody(repoData.Languages, myClient)
		if langStatus != nil {
			GetHTTP403(w, "language")
			return
		}

		Best, bestError := GetContributor(contribBody)
		if bestError != nil {
			GetHTTP403(w, "contributor")
			return
		}

		Language, langError := GetLanguages(langBody)
		if langError != nil {
			GetHTTP403(w, "language")
			return
		}

		presentedData := PresentedData{}

		presentedData.Project = rawURL[3] + "/" + rawURL[4] + "/" + rawURL[5] + "/"
		presentedData.Owner = repoData.Owner.Name
		presentedData.Committer = Best.Name
		presentedData.Commits = Best.Contributes
		presentedData.Language = Language

		json.Marshal(&presentedData)
		json.NewEncoder(w).Encode(presentedData)
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Repository not found", 404)
	}

}

// HandlerWrongURL ...
func HandlerWrongURL(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "please use /projectinfo/v1/github.com/*owner*/*repo*/", 404)
}
