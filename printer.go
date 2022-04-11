package delightful

import (
	"github.com/aripalo/go-delightful/colors"
	"github.com/aripalo/go-delightful/internal/emojiprefix"
	"github.com/enescakir/emoji"
)

func (m *Message) Title(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Title.Println(formatted)
}

func (m *Message) Debug(prefix emoji.Emoji, rendered string) {
	if m.verboseMode {
		formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
		colors.Debug.Println(formatted)
	}
}

func (m *Message) Info(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Info.Println(formatted)
}

func (m *Message) Prompt(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Prompt.Println(formatted)
}

func (m *Message) Success(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Success.Println(formatted)
}

func (m *Message) Warning(prefix emoji.Emoji, rendered string) {
	if m.silentMode {
		return
	}
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Warning.Println(formatted)
}

func (m *Message) Failure(prefix emoji.Emoji, rendered string) {
	formatted := emojiprefix.Format(m.emojiMode, prefix, rendered)
	colors.Failure.Println(formatted)
}
