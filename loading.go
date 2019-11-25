package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StoryLoader struct {
	Name   string
	Author string
	Arcs   []Arc
}

func loadFromFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func loadFromJSON(jsonByte []byte) (*StoryLoader, error) {
	var sl StoryLoader
	err := json.Unmarshal(jsonByte, &sl)
	if err != nil {
		return nil, err
	}

	return &sl, nil
}

func loadStory(filepath string) (*Story, error) {
	jsonBytes, err := loadFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the story JSON file -- %s", err)
	}

	story := &Story{Arcs: make(map[string]*Arc)}

	sl, err := loadFromJSON(jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to load the story from JSON -- %s", err)
	}

	story.Name = sl.Name
	story.Author = sl.Author
	for i := 0; i < len(sl.Arcs); i++ {
		arc := &sl.Arcs[i]
		arc.calculateTextSplit()
		arc.recalculateTextWrap(80)

		for optionName := range arc.Options {
			arc.OptionNames = append(arc.OptionNames, optionName)
		}

		story.Arcs[arc.Name] = arc
	}
	startArc, ok := story.Arcs[startArcName]
	if !ok {
		return nil, fmt.Errorf("%s missing from the JSON loaded story", startArcName)
	}
	story.CurrentArc = startArc

	if err := story.checkStory(); err != nil {
		return nil, fmt.Errorf("Invalid story -- %s", err)
	}

	return story, nil
}
