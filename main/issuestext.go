// 尝试 text/templ 使用
package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}}issues:
{{range .Items}}---------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | Printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo, "Printf": fmt.Printf}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}