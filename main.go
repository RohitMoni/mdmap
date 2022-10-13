package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func main() {
	fmt.Println("Started!")

	source, err := os.ReadFile("./roadmap.md")
	if err != nil {
		panic(err)
	}

	// fmt.Print(string(data))

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert([]byte(source), &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}
	metaData := meta.Get(context)
	title := metaData["title"]
	fmt.Println(title)
}
