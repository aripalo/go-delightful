package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Debug prints a message in gray for debugging/testing purposes.
// No-op, unless Verbose Mode active.
func (m *Message) Debug(prefix emoji.Emoji, rendered string) {
	if m.verboseMode {
		formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
		colors.Debug.Print(formatted)
	}
}

// Debugln prints a message line in gray for debugging/testing purposes.
// No-op, unless Verbose Mode active.
func (m *Message) Debugln(prefix emoji.Emoji, rendered string) {
	if m.verboseMode {
		formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
		colors.Debug.Println(formatted)
	}
}
