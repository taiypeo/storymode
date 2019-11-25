package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
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

	w, _ := (*screen).Size()
	story.CurrentArc.recalculateTextWrap(w)

	endChan := make(chan struct{})
	go func() {
		for {
			ev := (*screen).PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(endChan)
					return
				}
			case *tcell.EventResize:
				w, _ := (*screen).Size()
				story.CurrentArc.recalculateTextWrap(w)
			}

			redrawUI(screen, story, 0)
		}
	}()

	<-endChan
	(*screen).Fini()
}
