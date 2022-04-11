package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Warning prints a message in yellow.
// No-op if Silent Mode active.
func (m *Message) Warning(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Warning.Print(formatted)
}

// Warningln prints a message line in yellow.
// No-op if Silent Mode active.
func (m *Message) Warningln(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Warning.Println(formatted)
}
