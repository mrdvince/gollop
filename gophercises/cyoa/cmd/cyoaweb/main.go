package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mrdvince/gollop/gophercises/cyoa"
)

func main() {
	storyfile := flag.String("story-file", "gopher.json", "The json file containing the CYOA story file")
	port := flag.Int("port", 3000, "The server port")
	flag.Parse()
	f, err := os.Open(*storyfile)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)

	fmt.Printf("Server running on port %d\n",*port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port), h))
}
