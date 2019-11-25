package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

var defaultStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
var selectedStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

func createScreen() (*tcell.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("Failed to create the screen -- %s", err)
	}
	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("Failed to initialize the screen -- %s", err)
	}

	screen.SetStyle(defaultStyle)
	return &screen, nil
}

func drawString(screen *tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		width := runewidth.RuneWidth(c)
		if width == 0 {
			comb = []rune{c}
			c = ' '
			width = 1
		}

		(*screen).SetContent(x, y, c, comb, style)
		x += width
	}
}

func drawDivider(screen *tcell.Screen, y int, style tcell.Style) {
	w, _ := (*screen).Size()

	for i := 0; i < w; i++ {
		(*screen).SetContent(i, y, tcell.RuneHLine, nil, style)
	}
}

func clearBottom(screen *tcell.Screen, height int, style tcell.Style) {
	w, h := (*screen).Size()
	for delta := 1; delta <= height; delta++ {
		for i := 0; i < w; i++ {
			(*screen).SetContent(i, h-delta, ' ', nil, style)
		}
	}
}

func redrawUI(screen *tcell.Screen, story *Story, selectedOptionIndex int) {
	const optionsCount = 5

	currentArc := story.CurrentArc
	currentPage := selectedOptionIndex / optionsCount
	pageFirstItem := optionsCount * currentPage
	maxItem := len(currentArc.Options) - 1
	if pageFirstItem+optionsCount-1 < maxItem {
		maxItem = pageFirstItem + optionsCount - 1
	}

	(*screen).Clear()

	title := fmt.Sprintf("\"%s\" by %s", story.Name, story.Author)
	drawString(screen, 0, 0, defaultStyle, title)

	drawDivider(screen, 1, defaultStyle)

	for i, line := range currentArc.TextWrap {
		drawString(screen, 0, i+2, defaultStyle, line)
	}

	_, h := (*screen).Size()
	drawDivider(screen, h-(optionsCount+1), defaultStyle)

	clearBottom(screen, optionsCount, defaultStyle)

	line := h - optionsCount
	for i := pageFirstItem; i <= maxItem; i++ {
		style := defaultStyle
		if i == selectedOptionIndex {
			style = selectedStyle
		}

		txt := fmt.Sprintf("[%d]. %s", i+1, currentArc.OptionNames[i])
		drawString(screen, 0, line, style, txt)
		line++
	}

	(*screen).Show()
}
