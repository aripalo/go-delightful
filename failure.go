package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Failure prints a message in red.
// Usefule for notifying user about errors.
// Printed always, even if Silent Mode active.
func (m *Message) Failure(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Failure.Print(formatted)
}

// Failureln prints a message line in red.
// Usefule for notifying user about errors.
// Printed always, even if Silent Mode active.
func (m *Message) Failureln(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Failure.Println(formatted)
}
