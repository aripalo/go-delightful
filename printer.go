package delightful

import "fmt"

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
