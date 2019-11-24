package main

import "fmt"

func main() {
	story, err := loadStory("examples/test_story.json")
	fmt.Println(story)
	fmt.Println(err)
}
