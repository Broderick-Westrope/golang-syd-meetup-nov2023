package main

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit   key.Binding
	Up     key.Binding
	Down   key.Binding
	Submit key.Binding
	Save   key.Binding
	Help   key.Binding
}

func NewKeyMap() KeyMap {
	return KeyMap{
		Quit:   key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		Up:     key.NewBinding(key.WithKeys("up"), key.WithHelp("up arrow", "up")),
		Down:   key.NewBinding(key.WithKeys("down"), key.WithHelp("down arrow", "down")),
		Submit: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("enter", "toggle completed")),
		Save:   key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save to file")),
		Help:   key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Submit},
		{k.Save, k.Help, k.Quit},
	}
}
