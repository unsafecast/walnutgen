package main

import (
	"fmt"

	"github.com/unsafecast/walnutgen/walnut"
)

type variables struct {
	Title string
}

func main() {
	code, _ := walnut.GenerateFile("example/index.html", &walnut.Config{
		Headers: nil,
		Footers: nil,
		Variables: variables{
			Title: "Hello",
		},
	})
	fmt.Println(code)
}
