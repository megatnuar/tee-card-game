package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// This is main model
type model struct {
	card []Card
	deck []Deck
}

type Card struct {
	suit
}

// Initialising model
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (m model) View() string {

}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
