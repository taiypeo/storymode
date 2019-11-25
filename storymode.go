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
	selectedOption := 0
	go func() {
		for {
			ev := (*screen).PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(endChan)
					return
				case tcell.KeyDown:
					selectedOption++
					if selectedOption > len(story.CurrentArc.OptionNames)-1 {
						selectedOption = len(story.CurrentArc.OptionNames) - 1
					}
				case tcell.KeyUp:
					selectedOption--
					if selectedOption < 0 {
						selectedOption = 0
					}
				case tcell.KeyEnter:
					selectedOptionText := story.CurrentArc.OptionNames[selectedOption]
					targetArc := story.CurrentArc.Options[selectedOptionText]
					finished, err := story.changeArc(targetArc)
					selectedOption = 0
					if finished || err != nil {
						if err != nil {
							(*screen).Fini()
							fmt.Fprintf(os.Stderr, "Gameplay error: %s\n", err)
						}

						close(endChan)
						return
					}
				}
			case *tcell.EventResize:
				w, _ := (*screen).Size()
				story.CurrentArc.recalculateTextWrap(w)
			}

			redrawUI(screen, story, selectedOption)
		}
	}()

	<-endChan
	(*screen).Fini()
}
