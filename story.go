package main

import (
	"errors"
	"fmt"
	"github.com/mattn/go-runewidth"
	"strings"
)

const startArcName = "~~start~~"
const endArcName = "~~end~~"

type Arc struct {
	Name, Text          string
	TextSplit, TextWrap []string
	Options             map[string]string // option name -> arc name
	OptionNames         []string
	// Sound, Music (*sound struct)
}

func (a *Arc) calculateTextSplit() {
	a.TextSplit = strings.Split(a.Text, " ")
}

// Solves the word wrap problem for the text (https://xxyxyz.org/line-breaking/)
func (a *Arc) recalculateTextWrap(width int) {
	const MaxInt = int(^uint(0) >> 1) // all 1 except for the first bit

	offsets := make([]int, len(a.TextSplit)+1)
	offsets[0] = 0
	for i, word := range a.TextSplit {
		offsets[i+1] = offsets[i] + runewidth.StringWidth(word)
	}

	minima := make([]int, len(a.TextSplit)+1)
	minima[0] = 0
	for i := 1; i < len(minima); i++ {
		minima[i] = MaxInt
	}

	breaks := make([]int, len(a.TextSplit)+1)
	for i := 1; i < len(breaks); i++ {
		breaks[i] = 0
	}

	for i := 0; i < len(a.TextSplit); i++ {
		j := i + 1
		for j <= len(a.TextSplit) {
			w := offsets[j] - offsets[i] + j - i - 1
			if w > width {
				break
			}

			cost := minima[i] + (width-w)*(width-w)
			if cost < minima[j] {
				minima[j] = cost
				breaks[j] = i
			}

			j++
		}
	}

	j := len(a.TextSplit)
	for j > 0 {
		i := breaks[j]
		a.TextWrap = append(a.TextWrap, strings.Join(a.TextSplit[i:j], " "))
		j = i
	}

	// Reverse a.TextWrap
	for i := len(a.TextWrap)/2 - 1; i >= 0; i-- {
		opposite := len(a.TextWrap) - 1 - i
		a.TextWrap[i], a.TextWrap[opposite] = a.TextWrap[opposite], a.TextWrap[i]
	}
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
		if arc.OptionNames == nil {
			return fmt.Errorf("%s's OptionNames cannot be nil", arc.Name)
		}
		if arc.TextSplit == nil {
			return fmt.Errorf("%s's TextSplit cannot be nil", arc.Name)
		}
		if arc.TextWrap == nil {
			return fmt.Errorf("%s's TextWrap cannot be nil", arc.Name)
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

func (s *Story) changeArc(arcName string) (bool, error) {
	if arcName == endArcName {
		return true, nil
	}

	arc, ok := s.Arcs[arcName]
	if !ok {
		return false, fmt.Errorf("Cannot change to arc %s -- no such arc", arcName)
	}

	s.CurrentArc = arc
	return false, nil
}
