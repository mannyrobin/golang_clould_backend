package main

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
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
		// Check if req has errors
		panic(err)
	}

	//TODO: make username changeable
	// Apply header to request
	req.Header.Set("User-Agent", "projectinfo")

	// Try to execute request
	res, doError := myClient.Do(req)
	if doError != nil {
		panic(err)
	}

	body, readError := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if readError != nil {
		rtrError := readError
		return body, rtrError
	}
	return body, nil
}