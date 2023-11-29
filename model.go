package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items     []string
	cursor    int
	completed map[int]struct{}
	keys      KeyMap
	help      help.Model
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
		keys:      NewKeyMap(),
		help:      help.New(),
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
			choice = focussedStyle.Render(choice)
		} else {
			choice = unfocussedStyle.Render(choice)
		}

		checked := " "
		if _, ok := m.completed[i]; ok {
			checked = completedStyle.Render("x")
		}

		output += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	output = appStyle.Render(output)
	output += "\n" + m.help.View(m.keys)
	return output
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Submit):
			if _, ok := m.completed[m.cursor]; ok {
				delete(m.completed, m.cursor)
			} else {
				m.completed[m.cursor] = struct{}{}
			}
		case key.Matches(msg, m.keys.Save):
			return m, m.saveToFile
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
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
