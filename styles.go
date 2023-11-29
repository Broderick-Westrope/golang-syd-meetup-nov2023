package main

import "github.com/charmbracelet/lipgloss"

var focussedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AFBEE1"))

var unfocussedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

var completedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#15DA84"))

var appStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(1, 1, 0)
