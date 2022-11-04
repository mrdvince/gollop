package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mrdvince/gollop/gophercises/cyoa"
)

func main() {

	storyfile := flag.String("story-file", "gopher.json", "The json file containing the CYOA story file")
	flag.Parse()
	f, err := os.Open(*storyfile)
	if err != nil {
		panic(err)
	}

	var story cyoa.Story
	decode := json.NewDecoder(f)
	if err := decode.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
