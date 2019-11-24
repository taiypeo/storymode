package main

import "errors"

const startArcName = "~~start~~"
const endArcName = "~~end~~"

type Arc struct {
	Name, Text string
	Options    map[string]string // option name -> arc name
	// Sound, Music (*sound struct)
}

type Story struct {
	Name, Author string
	Arcs         map[string]*Arc // arc name -> arc pointer
	CurrentArc   *Arc
}

func (s *Story) checkStory() error {
	// TODO: check the story for validity

	return nil
}

func (s *Story) changeArc(arcName string) error {
	arc, ok := s.Arcs[arcName]
	if !ok {
		return errors.New("Cannot change to arc " + arcName + " -- no such arc")
	}

	s.CurrentArc = arc
	return nil
}

func loadStory(folderPath string) (*Story, error) {
	story := &Story{Arcs: make(map[string]*Arc)}

	// TODO: actually load the story from JSON

	if err := story.checkStory(); err != nil {
		return nil, err
	}

	return story, nil
}
