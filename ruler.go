package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/ruler"
)

// printHr is the internal implementation of printing horizontal rulers
// with any given character.
// No-op unless Verbose Mode active.
func (m *Message) printHr(char string) {
	if m.verboseMode {
		hr := ruler.Generate(char)
		colors.Debug.Println(hr)
	}
}

// HorizontalRuler prints equals (=) characters as wide as the terminal.
func (m *Message) HorizontalRuler() {
	m.printHr("=")
}
