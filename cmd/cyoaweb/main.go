package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/ThePSAdmin/gophercises-03-cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "The file with choose your own adventure story")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
