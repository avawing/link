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

	dfs(doc, " ")
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	//dfs(doc, "")
	// unneeded dfs now
	return nil, nil
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ret = append(ret, linkNodes(child)...)
	}

	// what is basecase?
	return ret
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	// tree traversal -> for every first child where c is not nil, we go to the next sibling
	// cool way to traverse the dom -> not as verbose as other things
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, padding+"  ")
	}
}
