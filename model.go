package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items     []string
	cursor    int
	completed map[int]struct{}
}

func NewModel() model {
	return model{
		items: []string{
			"Headphones",
			"Lipgloss",
			"Bubble tea",
			"Gopher toy",
		},
		completed: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	output := "What should we buy at the market?\n\n"

	for i, choice := range m.items {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.completed[i]; ok {
			checked = "x"
		}

		output += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	output += "\nPress q to quit.\n"
	return output
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			if _, ok := m.completed[m.cursor]; ok {
				delete(m.completed, m.cursor)
			} else {
				m.completed[m.cursor] = struct{}{}
			}
		case "s":
			return m, m.saveToFile
		}
	case savedMsg:
		if msg.err != nil {
			fmt.Printf("Alas, there's been an error: %v\n", msg.err)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) saveToFile() tea.Msg {
	list := ""
	for _, item := range m.items {
		list += item + "\n"
	}

	err := os.WriteFile("shopping-list.txt", []byte(list), 0644)
	return savedMsg{err}
}

type savedMsg struct {
	err error
}
