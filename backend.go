package main

import (
	"errors"
	"os"
)

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
	if s.Name == "" {
		return errors.New("Story's name cannot be empty")
	}
	if s.Author == "" {
		return errors.New("Story's author cannot be empty")
	}
	if s.Arcs == nil {
		return errors.New("Arcs cannot be nil in the Story struct")
	}
	if s.CurrentArc == nil {
		return errors.New("CurrentArc cannot be nil in the Story struct")
	}

	if _, ok := s.Arcs[startArcName]; !ok {
		return errors.New(startArcName + " should be present in Arcs in the Story struct")
	}
	if _, ok := s.Arcs[endArcName]; ok {
		return errors.New(endArcName + " should NOT be present in Arcs in the Story struct")
	}

	for name, arc := range s.Arcs {
		if arc == nil {
			return errors.New(name + " has an invalid value in Arcs")
		}
		if arc.Name == "" {
			return errors.New("Arc name cannot be empty")
		}
		if name != arc.Name {
			return errors.New(arc.Name + " has an invalid key in Arcs")
		}
		if arc.Text == "" {
			return errors.New(arc.Name + "'s Text cannot be empty")
		}
		if arc.Options == nil {
			return errors.New(arc.Name + "'s Options cannot be nil")
		}

		for optionName, targetArcName := range arc.Options {
			if optionName == "" {
				return errors.New(arc.Name + "'s option names cannot be empty")
			}
			if targetArcName == "" {
				return errors.New(arc.Name + "'s option targets cannot be empty")
			}
			if _, ok := s.Arcs[targetArcName]; targetArcName != endArcName && !ok {
				return errors.New(arc.Name + "'s options should point to existing arcs")
			}
		}
	}

	return nil
}

func (s *Story) changeArc(arcName string) error {
	if arcName == endArcName {
		// TODO: add resource deallocation
		os.Exit(0)
	}

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
