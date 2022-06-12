package main

import (
	"fmt"
	"link"
	"strings"
)

var html = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page"><span>A link to another page</span></a>
  <a href="/second-page">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(html)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
