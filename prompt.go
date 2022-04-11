package delightful

import (
	"github.com/aripalo/go-delightful/internal/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

// Prompt prints a "prompt query/question" in cyan.
// Meant to be used for asking user input, just before reading stdin.
// Printed always, even if Silent Mode active.
//
// NOTE: This method only prints the "prompt query/question" and does not
// actually read the stdin: You need to implement that yourself!
func (m *Message) Prompt(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Prompt.Print(formatted)
}

// Promptln prints a "prompt query/question" line in cyan.
// Meant to be used for asking user input, just before reading stdin.
// Printed always, even if Silent Mode active.
//
// NOTE: This method only prints the "prompt query/question" and does not
// actually read the stdin: You need to implement that yourself!
func (m *Message) Promptln(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Prompt.Println(formatted)
}
