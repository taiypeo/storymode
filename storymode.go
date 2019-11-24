package main

import "fmt"

func main() {
	arc1 := &Arc{
		Name: startArcName,
		Text: "You embark on your journey...",
		Options: map[string]string{
			"Option B":             "opt_b",
			"Finish your journey!": endArcName,
		},
	}
	arc2 := &Arc{
		Name: "opt_b",
		Text: "You chose option B!",
		Options: map[string]string{
			"Finish your journey!": endArcName,
		},
	}

	story := Story{
		Name:   "Boring story",
		Author: "Name Surname",
		Arcs: map[string]*Arc{
			startArcName: arc1,
			"opt_b":      arc2,
		},
		CurrentArc: arc1,
	}

	fmt.Printf("checkStory() output -- %s\n", story.checkStory())
}
