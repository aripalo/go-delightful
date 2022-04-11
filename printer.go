package delightful

import (
	"fmt"
	"unicode/utf8"

	"github.com/aripalo/go-delightful/colors"
	"github.com/enescakir/emoji"
)

func emojiPad(em emoji.Emoji) string {
	runeCount := utf8.RuneCountInString(string(em))
	padding := 3
	if runeCount%2 != 0 {
		padding = 1
	}
	return fmt.Sprintf("%-*s", padding, string(em))
}

func (m *Message) formatWithEmoji(em emoji.Emoji, rendered string) string {
	// empty emoji means we'll skip it
	if string(em) == "" {
		return rendered
	}

	// format the string with emoji prefix if emoji mode is on
	if m.emojiMode {
		return fmt.Sprintf("%s %s", emojiPad(em), rendered)
	}

	// otherwise just return the string
	return rendered
}

func (m *Message) Title(em emoji.Emoji, rendered string) {
	if !m.silentMode {
		colors.Title.Println(m.formatWithEmoji(em, rendered))
	}
}

func (m *Message) Debug(em emoji.Emoji, rendered string) {
	if m.verboseMode {
		colors.Debug.Println(m.formatWithEmoji(em, rendered))
	}
}

func (m *Message) Info(em emoji.Emoji, rendered string) {
	if !m.silentMode {
		colors.Info.Println(m.formatWithEmoji(em, rendered))
	}
}

func (m *Message) Prompt(em emoji.Emoji, rendered string) {
	colors.Prompt.Println(m.formatWithEmoji(em, rendered))
}

func (m *Message) Success(em emoji.Emoji, rendered string) {
	if !m.silentMode {
		colors.Success.Println(m.formatWithEmoji(em, rendered))
	}
}

func (m *Message) Warning(em emoji.Emoji, rendered string) {
	if !m.silentMode {
		colors.Warning.Println(m.formatWithEmoji(em, rendered))
	}
}

func (m *Message) Failure(em emoji.Emoji, rendered string) {
	colors.Failure.Println(m.formatWithEmoji(em, rendered))
}
