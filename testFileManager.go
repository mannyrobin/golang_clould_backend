package main

import (
	"strings"
	"io/ioutil"
)

// Only used to save []byte data to file for testing
func SaveBodyTestData(url string, body []byte) {
	if strings.Contains(url, "git/git") {
		if strings.Contains(url, "contributors") {
			fileErr := ioutil.WriteFile("contributorsBodyTest", body, 0644)
			if fileErr != nil {
				panic(fileErr)
			}
		} else if strings.Contains(url, "language") {
			fileErr := ioutil.WriteFile("languagesBodyTest", body, 0644)
			if fileErr != nil {
				panic(fileErr)
			}
		} else {
			fileErr := ioutil.WriteFile("reposTest", body, 0644)
			if fileErr != nil {
				panic(fileErr)
			}
		}
	}else{
		println("Please use /git/git!")
	}

}


func LoadTestBodyData(filename string)[]byte{
	body, err := ioutil.ReadFile("./testFiles/" + filename)
	if err != nil{
		println("Failed to find test data")
		panic(err)
	}
	return body
}
