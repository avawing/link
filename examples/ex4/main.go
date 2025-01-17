package main

import (
	"fmt"
	link "github.com/avawing/link/m"
	"strings"
)

var html = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
>
`

func main() {
	r := strings.NewReader(html)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
