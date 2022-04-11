package delightful

import (
	"github.com/aripalo/go-delightful/colors"
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
	colors.Title.Println(formatted)
}

// Debug prints a message in gray for debugging/testing purposes.
// No-op, unless Verbose Mode active.
func (m *Message) Debug(prefix emoji.Emoji, rendered string) {
	if m.verboseMode {
		formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
		colors.Debug.Println(formatted)
	}
}

// Info prints an informational message in white color.
// No-op if Silent Mode active.
func (m *Message) Info(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Info.Println(formatted)
}

// Prompt prints a "prompt query/question" in cyan.
// Meant to be used for asking user input, just before reading stdin.
// Printed always, even if Silent Mode active.
//
// NOTE: This method only prints the "prompt query/question" and does not
// actually read the stdin: You need to implement that yourself!
func (m *Message) Prompt(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Prompt.Println(formatted)
}

// Success prints a message in green.
// No-op if Silent Mode active.
func (m *Message) Success(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Success.Println(formatted)
}

// Warning prints a message in yellow.
// No-op if Silent Mode active.
func (m *Message) Warning(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Warning.Println(formatted)
}

// Failure prints a message in red.
// Usefule for notifying user about errors.
// Printed always, even if Silent Mode active.
func (m *Message) Failure(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Failure.Println(formatted)
}
