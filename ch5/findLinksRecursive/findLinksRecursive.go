package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinksRecursive: %v\n", err)
		os.Exit(1)
	}
	// visit(nil, doc)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

// visit function traverses an HTML node tree founc in n and returns the result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				links = append(links, a.Val)
			}

		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)

	return links
}

/*
$ go build ../ch1/fetch
$ go build findlinks
$ ./fetch https://red-badger.com/ | ./findlinks
#
/what-we-do
/technology
/our-work
/about-us
/people
/blog
/events
/jobs
/
/what-

*/
