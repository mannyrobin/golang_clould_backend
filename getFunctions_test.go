package main

import(
	"testing"
	"encoding/json"
	"strings"
	"net/http"
	"net/http/httptest"
)

func TestGetContributor(t *testing.T) {
	rawJson := []byte(`{"login":"gitster", "contributions":18497}`)
	contribTest := Contributer{}
	json.Unmarshal(rawJson, &contribTest)

	testBody := LoadTestBodyData("contributorsBodyTest")

	cont, err := GetContributor(testBody)
	if err != nil{
		t.Error(err)
	}

	if contribTest.Name != cont.Name {
		t.Error("Did not retrieve correct contributer")
	}

}

// @Functions: To check if all members of array exist
func TestGetLanguages(t *testing.T) {
	equalFlag := false

	wantString := strings.Split("C,Shell,Perl,Tcl,Python,C++,Makefile,"+
		 "Emacs Lisp,JavaScript,M4,Roff,Perl 6,Go,CSS,PHP,Assembly,Objective-C", ",")

	testBody := LoadTestBodyData("languagesBodyTest")

	lang, err := GetLanguages(testBody)
	if err != nil {
		t.Error(err)
	}

	wantLen := len(wantString)
	if len(lang) != wantLen {
		t.Error("lang array not same!")
	} else {
		for i := 0; i < wantLen; i++ {
			for j := 0; j < wantLen; j++ {
				if lang[i] == wantString[j] {
					equalFlag = true
					break
				}
			}
			if equalFlag != true {
				t.Error("lang array is not correct")
				return
			}
			equalFlag = false
		}
	}
}

func TestGetGitRepoUrl(t *testing.T) {
	req := CreateTestReq(t)

	splitUrl := strings.Split("/projectinfo/v1/github.com/git/git/", "/")
	rawUrl, repoError := GetGitRepoUrl(req)
	if repoError != nil{
		t.Error(repoError)
	}

	length := len(rawUrl)

	for i := 0; i < length; i++ {
		if splitUrl[i] != rawUrl[i] {
			t.Error("Did not retrieve correct string array!")
		}
	}
}

func TestHandlerGitURL(t *testing.T) {
	req := CreateTestReq(t)

	resReq := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerGitURL)

	handler.ServeHTTP(resReq, req)

	status := resReq.Code
	if status != http.StatusOK{
		t.Errorf("Asked for ok status but got %v!", status)
	}

	expected := `{"project":"github.com/git/git/","owner":"git","committer":"gitster"`
	if !strings.Contains(resReq.Body.String(), expected) {
		t.Error("Handler returned unexpected body")
	}
}

func CreateTestReq(t *testing.T)*http.Request{
	wantedUrl := "/projectinfo/v1/github.com/git/git/"

	req, err := http.NewRequest("GET", wantedUrl, nil)
	if err != nil{
		t.Fatal("Did not get request")
	}
	return req
}


