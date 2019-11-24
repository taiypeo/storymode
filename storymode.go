package main

import (
	"fmt"
	"os"
)

func main() {
	story, err := loadStory("examples/test_story.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(story)
}
