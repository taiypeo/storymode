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
			"Option C":             "opt_c",
			"Finish your journey!": endArcName,
		},
	}
	arc3 := &Arc{
		Name: "opt_c",
		Text: "You chose option C!",
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
			"opt_c":      arc3,
		},
		CurrentArc: arc1,
	}

	fmt.Printf("checkStory() output -- %s\n", story.checkStory())
	fmt.Printf("currentArc name -- %s\n", story.CurrentArc.Name)
	fmt.Printf("changeArc() output -- %s\n", story.changeArc("opt_b"))
	fmt.Printf("currentArc name -- %s\n", story.CurrentArc.Name)
}
