package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Info prints an informational message in white color.
// No-op if Silent Mode active.
func (m *Message) Info(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Info.Print(formatted)
}

// Infoln prints an informational message line in white color.
// No-op if Silent Mode active.
func (m *Message) Infoln(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Info.Println(formatted)
}
