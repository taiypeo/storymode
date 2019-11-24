package main

import (
	"errors"
	"fmt"
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
		return fmt.Errorf("%s should be present in Arcs in the Story struct", startArcName)
	}
	if _, ok := s.Arcs[endArcName]; ok {
		return fmt.Errorf("%s should NOT be present in Arcs in the Story struct", endArcName)
	}

	for name, arc := range s.Arcs {
		if arc == nil {
			return fmt.Errorf("%s has an invalid empty value in Arcs", name)
		}
		if name == "" || arc.Name == "" {
			return errors.New("Arc name cannot be empty")
		}
		if name != arc.Name {
			return fmt.Errorf("%s has an invalid key in Arcs (%s)", arc.Name, name)
		}
		if arc.Text == "" {
			return fmt.Errorf("%s's Text cannot be empty", arc.Name)
		}
		if arc.Options == nil {
			return fmt.Errorf("%s's Options cannot be nil", arc.Name)
		}

		for optionName, targetArcName := range arc.Options {
			if optionName == "" {
				return fmt.Errorf("%s's option names cannot be empty", arc.Name)
			}
			if targetArcName == "" {
				return fmt.Errorf("%s's option targets cannot be empty", arc.Name)
			}
			if _, ok := s.Arcs[targetArcName]; targetArcName != endArcName && !ok {
				return fmt.Errorf("%s's options should point to existing arcs", arc.Name)
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
		return fmt.Errorf("Cannot change to arc %s -- no such arc", arcName)
	}

	s.CurrentArc = arc
	return nil
}
