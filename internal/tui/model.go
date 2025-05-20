package tui

import (
	"algo/internal/usecase"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	platforms = []string{"baekjoon", "programmers"}

	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	activeItemStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("42"))
	selectedItemStyle = activeItemStyle

	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))

	focusedPromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	focusedInputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredPromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	blurredInputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

const (
	selectPlatform = iota
	inputID
	inputPath
	confirm
)

type Status struct {
	step      int
	platform  string
	id        string
	path      string
	cursor    int
	err       error
	message   string
	quitting  bool
	idInput   textinput.Model
	pathInput textinput.Model
}

func New() Status {
	idInput := textinput.New()
	idInput.Placeholder = "Problem ID"
	idInput.CharLimit = 32
	idInput.Width = 30
	idInput.Prompt = fmt.Sprintf("  %-14s:   ", "Problem ID")
	idInput.PromptStyle = focusedPromptStyle
	idInput.TextStyle = focusedInputStyle

	pathInput := textinput.New()
	pathInput.Placeholder = "File Path (e.g., ./problems/problem.json)"
	pathInput.CharLimit = 128
	pathInput.Width = 50
	pathInput.Prompt = fmt.Sprintf("  %-14s:   ", "Save Path")
	pathInput.PromptStyle = blurredPromptStyle
	pathInput.TextStyle = blurredInputStyle

	return Status{
		step:      selectPlatform,
		cursor:    0,
		message:   "Select a platform using ↑/↓ and Enter",
		idInput:   idInput,
		pathInput: pathInput,
	}
}

func (s Status) Init() tea.Cmd {
	return textinput.Blink
}

func (s Status) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			s.quitting = true
			return s, tea.Quit

		case tea.KeyEnter:
			switch s.step {
			case selectPlatform:
				s.platform = platforms[s.cursor]
				s.step = inputID
				s.message = "Enter problem ID"
				s.idInput.Focus()
				return s, textinput.Blink

			case inputID:
				if s.idInput.Value() == "" {
					s.err = fmt.Errorf("problem ID is required")
					return s, nil
				}
				s.id = s.idInput.Value()
				s.step = inputPath
				s.message = "Enter save path"
				s.idInput.Blur()
				s.pathInput.Focus()
				return s, textinput.Blink

			case inputPath:
				if s.pathInput.Value() == "" {
					s.err = fmt.Errorf("save path is required")
					return s, nil
				}
				s.path = s.pathInput.Value()
				s.step = confirm
				s.message = fmt.Sprintf("Add problem %s-%s to %s? (Enter to confirm)", s.platform, s.id, s.path)
				s.pathInput.Blur()
				return s, nil

			case confirm:
				s.message = fmt.Sprintf("Adding problem: %s - %s to %s...", s.platform, s.id, s.path)
				s.err = nil

				err := usecase.Add(s.platform, s.id, s.path)
				if err != nil {
					s.err = err
					s.message = ""
				} else {
					s.message = fmt.Sprintf("Successfully added: %s - %s to %s. Press Esc to quit.", s.platform, s.id, s.path)
				}
				return s, nil
			}

		case tea.KeyUp, tea.KeyDown:
			if s.step == selectPlatform {
				if msg.Type == tea.KeyUp {
					s.cursor--
					if s.cursor < 0 {
						s.cursor = len(platforms) - 1
					}
				} else {
					s.cursor++
					if s.cursor >= len(platforms) {
						s.cursor = 0
					}
				}
			}

		default:
		}
	}

	if s.step == inputID {
		s.idInput, cmd = s.idInput.Update(msg)
		return s, cmd
	} else if s.step == inputPath {
		s.pathInput, cmd = s.pathInput.Update(msg)
		return s, cmd
	}

	return s, nil
}

func (s Status) View() string {
	if s.quitting {
		return "Bye!\n"
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("Algo (github.com/himitery/algo)"))
	b.WriteString("\n\n")

	switch s.step {
	case selectPlatform:
		b.WriteString(titleStyle.Render("Select a platform:"))
		b.WriteString("\n\n")

		for i, platform := range platforms {
			if i == s.cursor {
				b.WriteString(activeItemStyle.Render("▶ " + platform))
			} else {
				b.WriteString(itemStyle.Render(platform))
			}
			b.WriteString("\n")
		}

	case inputID:
		b.WriteString(titleStyle.Render("Enter problem details:"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  %-14s: %s\n\n", "Platform", selectedItemStyle.Render(s.platform)))
		b.WriteString(s.idInput.View())

	case inputPath:
		b.WriteString(titleStyle.Render("Enter problem details:"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  %-14s: %s\n", "Platform", selectedItemStyle.Render(s.platform)))
		b.WriteString(fmt.Sprintf("  %-14s: %s\n\n", "Problem ID", selectedItemStyle.Render(s.id)))
		b.WriteString(s.pathInput.View())

	case confirm:
		b.WriteString(titleStyle.Render("Confirm details"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  %-14s: %s\n", "Platform", selectedItemStyle.Render(s.platform)))
		b.WriteString(fmt.Sprintf("  %-14s: %s\n", "Problem ID", selectedItemStyle.Render(s.id)))
		b.WriteString(fmt.Sprintf("  %-14s: %s\n\n", "Save Path", selectedItemStyle.Render(s.path)))
		b.WriteString(activeItemStyle.Render("▶ Press Enter to confirm or Esc to cancel"))
	}

	if s.message != "" {
		b.WriteString("\n\n" + statusStyle.Render(s.message))
	}
	if s.err != nil {
		b.WriteString("\n" + errorStyle.Render(fmt.Sprintf("Error: %v", s.err)))
	}

	var helpText string
	switch s.step {
	case selectPlatform:
		helpText = "↑/↓: navigate • enter: select • esc: quit"
	case inputID, inputPath:
		helpText = "enter: confirm • esc: quit"
	case confirm:
		helpText = "enter: confirm • esc: quit"
	}
	b.WriteString("\n\n" + helpStyle.Render(helpText))

	return b.String()
}
