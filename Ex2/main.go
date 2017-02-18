package main

import (
	"fmt"
	"golangexp/github"
	"log"
	"os"
)

// example of parameter in the command line:
// go run main.go repo:golang/go is:open json decoder
func main() {

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
}
