package delightful

import (
	"fmt"

	"github.com/aripalo/go-delightful/colors"
)

func (m *Message) formatWithEmoji(emoji string, rendered string) string {
	// empty emoji means we'll skip it
	if emoji == "" {
		return rendered
	}

	// format the string with emoji prefix if emoji mode is on
	if m.emojiMode {
		return fmt.Sprintf("%s %s", emoji, rendered)
	}

	// otherwise just return the string
	return rendered
}

func (m *Message) Title(emoji string, rendered string) {
	colors.Title.Println(m.formatWithEmoji(emoji, rendered))
}

func (m *Message) Debug(emoji string, rendered string) {
	if m.verboseMode {
		colors.Debug.Println(m.formatWithEmoji(emoji, rendered))
	}
}

func (m *Message) Info(emoji string, rendered string) {
	colors.Info.Println(m.formatWithEmoji(emoji, rendered))
}

func (m *Message) Prompt(emoji string, rendered string) {
	colors.Prompt.Println(m.formatWithEmoji(emoji, rendered))
}

func (m *Message) Success(emoji string, rendered string) {
	colors.Success.Println(m.formatWithEmoji(emoji, rendered))
}

func (m *Message) Warning(emoji string, rendered string) {
	colors.Warning.Println(m.formatWithEmoji(emoji, rendered))
}

func (m *Message) Failure(emoji string, rendered string) {
	colors.Failure.Println(m.formatWithEmoji(emoji, rendered))
}
