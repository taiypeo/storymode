package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	story, err := loadStory("examples/test_story.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Loading error: %s\n", err)
		os.Exit(1)
	}

	screen, err := createScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "UI error: %s\n", err)
		os.Exit(1)
	}

	redrawUI(screen, story, 1)
	time.Sleep(time.Second * 5)
	(*screen).Fini()
}
