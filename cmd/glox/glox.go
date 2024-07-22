package main

import (
	"log"
	"os"

	"github.com/samthom/glox/lib"
)

var glx lib.Glox

func main() {
	args := os.Args

	glx = lib.NewGlox()

	if len(args) > 2 {
		log.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		glx.RunFile(args[1])
	} else {
		glx.RunPrompt()
	}

}
