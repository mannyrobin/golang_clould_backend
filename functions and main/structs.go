package main
// Data that is presented

type PresentedData struct {
	Project string 		`json:"project"`
	Owner string 		`json:"owner"`
	Committer string 	`json:"committer"`
	Commits int			`json:"commits"`
	Language []string	`json:"language"`
}

// Data from repo front
type RepoData struct{
	Message string 		`json:"message"`
	Project string 		`json:"name"`
	Owner Owner 		`json:"owner"`
	Contributors string `json:"contributors_url"`
	Languages string	`json:"languages_url"`
}

// Used to retrieve login string
type Owner struct{
	Name string 		`json:"login"`
}

type Contributer struct{
	Name string			`json:"login"`
	Contributes int		`json:"contributions"`
}
