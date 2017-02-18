package main

import (
	geo "golangexp/geometry"
	"golangexp/github"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
</table>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//var report = template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(issueList))

func main() {
	// first example: Marshalling json
	whl := geo.Wheel{geo.Circle{geo.Point{2, 4}, 5}, 20}
	whlJson, err := json.MarshalIndent(whl, "", "	")
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %s", err)
	}
	fmt.Printf("%s\n\n", whlJson)

	// Second example: github search issue
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	// Third Example
	results, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, results); err != nil {
		log.Fatal(err)
	}

}
