package link

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

// Link represents a link in an html document
type Link struct {
	Href string
	Text string
}

// Parse takes in an html document, returns slice of links parsed from document
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	_ = doc
	return nil, nil
}

func dfs(node *html.Node, padding string) {
	fmt.Println(padding, node.Data)
	// tree traversal -> for every first child where c is not nil, we go to the next sibling
	// cool way to traverse the dom -> not as verbose as other things
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
