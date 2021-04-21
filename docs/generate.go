package main

import (
	"github.com/bukowa/go-cobra-example/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, ".")
	if err != nil {
		log.Fatal(err)
	}
}
