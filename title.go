package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Title prints a heading message in blue color.
// No-op if Silent Mode active.
func (m *Message) Title(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Title.Print(formatted)
}

// Titleln prints a heading message line in blue color.
// No-op if Silent Mode active.
func (m *Message) Titleln(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Title.Println(formatted)
}
