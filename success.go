package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Success prints a message in green.
// No-op if Silent Mode active.
func (m *Message) Success(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Success.Print(formatted)
}

// Successln prints a message line in green.
// No-op if Silent Mode active.
func (m *Message) Successln(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Success.Println(formatted)
}
