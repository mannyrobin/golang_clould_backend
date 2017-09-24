package main

import (
	"net/http"
	"time"
	"encoding/json"
)

func HandlerGitURL(w http.ResponseWriter, r *http.Request){
	http.Header.Add(w.Header(), "content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	rawUrl, repoError := GetGitRepoUrl(r)
	if repoError != nil{
		GetHTTP500(w,"repository url")
	}


	apiUrl := "https://api." + rawUrl[3] + "/repos/" + rawUrl[4] + "/" + rawUrl[5]

	myClient := http.Client{
		Timeout: time.Second * 2,
	}

	repoBody, repBError := GetBody(apiUrl, myClient)
	if repBError != nil{
		GetHTTP403(w,"repository")
	}

	repoData := RepoData{}
	jsonError := json.Unmarshal(repoBody, &repoData)
	if jsonError != nil{
		GetHTTP500(w,"repository")
	}

	if repoData.Message != "Not Found" {
		contribBody, contStatus := GetBody(repoData.Contributors, myClient)


		if contStatus != nil {
			GetHTTP403(w,"contributor")
		}

		langBody, langStatus := GetBody(repoData.Languages, myClient)
		if langStatus != nil {
			GetHTTP403(w,"language")
		}

		Best, bestError := GetContributor(contribBody)
		if bestError != nil {
			GetHTTP500(w,"contributor")
		}

		Language, langError := GetLanguages(langBody)
		if langError != nil {
			GetHTTP500(w,"language")
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

func HandlerWrongURL(w http.ResponseWriter, r *http.Request){
	// TODO: Wrong code atm
	http.Error(w, "please use /projectinfo/v1/github.com/*owner*/*repo*/", 404)
}
