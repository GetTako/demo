package app

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
)

// Todo represents a single task.
type Todo struct {
	Text    string
	Checked bool
}

// TodoBox implements contracts.Component.
type TodoBox struct {
	ctx        *tako.Context
	todos      []Todo
	inputModel textinput.Model
	cursor     int
	listMode   bool // true: navigate list, false: type input
}

func (t *TodoBox) ID() string { return "todo-box" }

func (t *TodoBox) Init() {
	var stored []Todo
	if err := t.ctx.Storage().Get("todos", &stored); err == nil && len(stored) > 0 {
		t.todos = stored
	}

	ti := textinput.New()
	ti.Placeholder = "What needs to be done?"
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(50)

	styles := textinput.DefaultDarkStyles()
	styles.Focused.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("204")).Bold(true)
	styles.Blurred.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Focused.Text = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	styles.Cursor.Color = lipgloss.Color("204") // Cursor color doesn't use lipgloss style directly in v2 virtual cursor but we set it
	ti.SetStyles(styles)

	t.inputModel = ti
}

func (t *TodoBox) save() {
	t.ctx.Storage().Set("todos", t.todos)
}

func (t *TodoBox) Render() any {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).MarginBottom(1)
	b.WriteString(titleStyle.Render("✓ Tako Todo List") + "\n")

	// Get terminal width from storage
	var termWidth int
	_ = t.ctx.Storage().Get("term_width", &termWidth)
	if termWidth <= 0 {
		termWidth = 60 // Fallback
	}

	// Calculate responsive width with some margin
	boxWidth := termWidth - 4
	if boxWidth < 20 {
		boxWidth = 20
	}

	// Render List
	for i, todo := range t.todos {
		itemStr := ""
		if todo.Checked {
			itemStr = fmt.Sprintf("%d. [x] %s", i+1, todo.Text)
		} else {
			itemStr = fmt.Sprintf("%d. [ ] %s", i+1, todo.Text)
		}

		style := lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Width(boxWidth)
		if todo.Checked {
			style = style.Strikethrough(true).Foreground(lipgloss.Color("240"))
		}

		if t.listMode && i == t.cursor {
			// Highlight cursor in list mode with a calm color (soft blue/slate)
			style = style.Background(lipgloss.Color("61")).Foreground(lipgloss.Color("255")).Bold(true)
			itemStr = "> " + itemStr
		} else {
			itemStr = "  " + itemStr
		}

		b.WriteString(style.Render(itemStr) + "\n")
	}

	b.WriteString("\n")

	// Manage Focus for TextInput
	if !t.listMode {
		t.inputModel.Focus()
		t.inputModel.Prompt = "Input> "
	} else {
		t.inputModel.Blur()
		t.inputModel.Prompt = "Input> "
	}
	t.inputModel.SetWidth(boxWidth - 10) // leave room for prompt

	b.WriteString(t.inputModel.View())

	b.WriteString("\n\n")
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	if t.listMode {
		b.WriteString(helpStyle.Render("Mode: LIST | [tab] to Input | [space] toggle | [d] delete | [up/down] move"))
	} else {
		b.WriteString(helpStyle.Render("Mode: INPUT | [tab] to List | [enter] add | [left/right] move cursor"))
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(2, 4).
		Width(60)

	return containerStyle.Render(b.String())
}

func (t *TodoBox) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(t.ID())

	// Tab toggles mode
	z.Bind("tab", func() {
		t.listMode = !t.listMode
	})

	// Capture all unnormalized keys for text input (e.g. 'A', 'a', '!', '1')
	z.OnFallback(func(key string) {
		if !t.listMode {
			// Convert string to rune(s). Usually it's a single character.
			runes := []rune(key)
			if len(runes) > 0 {
				msg := tea.KeyPressMsg{Text: key, Code: runes[0]}
				t.inputModel, _ = t.inputModel.Update(msg)
			}
		} else {
			// If in list mode, 'd' or 'D' deletes the active item
			if key == "d" || key == "D" {
				if len(t.todos) > 0 {
					t.todos = append(t.todos[:t.cursor], t.todos[t.cursor+1:]...)
					if t.cursor >= len(t.todos) && t.cursor > 0 {
						t.cursor--
					}
					t.save()
				}
			}
		}
	})

	z.Bind("space", func() {
		if !t.listMode {
			msg := tea.KeyPressMsg{Text: " ", Code: uv.KeySpace}
			t.inputModel, _ = t.inputModel.Update(msg)
		} else {
			// Toggle checked in list mode
			if len(t.todos) > 0 {
				t.todos[t.cursor].Checked = !t.todos[t.cursor].Checked
				t.save()
			}
		}
	})

	z.Bind("backspace", func() {
		if !t.listMode {
			msg := tea.KeyPressMsg{Code: uv.KeyBackspace}
			t.inputModel, _ = t.inputModel.Update(msg)
		} else {
			if len(t.todos) > 0 {
				t.todos = append(t.todos[:t.cursor], t.todos[t.cursor+1:]...)
				if t.cursor >= len(t.todos) && t.cursor > 0 {
					t.cursor--
				}
				t.save()
			}
		}
	})

	z.Bind("enter", func() {
		if !t.listMode {
			if val := t.inputModel.Value(); strings.TrimSpace(val) != "" {
				t.todos = append(t.todos, Todo{Text: val, Checked: false})
				t.inputModel.SetValue("")
				t.save()
			}
		}
	})

	z.Bind("left", func() {
		if !t.listMode {
			t.inputModel, _ = t.inputModel.Update(tea.KeyPressMsg{Code: uv.KeyLeft})
		}
	})

	z.Bind("right", func() {
		if !t.listMode {
			t.inputModel, _ = t.inputModel.Update(tea.KeyPressMsg{Code: uv.KeyRight})
		}
	})

	z.Bind("up", func() {
		if t.listMode && t.cursor > 0 {
			t.cursor--
		}
	})

	z.Bind("down", func() {
		if t.listMode && t.cursor < len(t.todos)-1 {
			t.cursor++
		}
	})
}
