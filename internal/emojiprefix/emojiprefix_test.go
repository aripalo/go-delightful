package emojiprefix

import (
	"fmt"
	"testing"

	"github.com/enescakir/emoji"
	"github.com/stretchr/testify/assert"
)

func TestDeterminePad(t *testing.T) {
	tests := []struct {
		name     string
		input    emoji.Emoji
		expected int
	}{
		{
			name:     "pile of poop",
			input:    emoji.PileOfPoo,
			expected: 0,
		},
		{
			name:     "airplane departure",
			input:    emoji.AirplaneDeparture,
			expected: 0,
		},
		{
			name:     "checkbox with check",
			input:    emoji.CheckBoxWithCheck,
			expected: 1,
		},
		{
			name:     "double rune: gear",
			input:    emoji.Gear,
			expected: 1,
		},
		{
			name:     "man with red hair",
			input:    emoji.Emoji(emoji.Man.Tone(emoji.Default)),
			expected: 0,
		},
		{
			name:     "woman with white hair",
			input:    emoji.Emoji(emoji.WomanGesturingNo.Tone(emoji.Default)),
			expected: 0,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := determinePad(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		input    emoji.Emoji
		expected string
	}{
		{
			name:     "pile of poop",
			input:    emoji.PileOfPoo,
			expected: "💩",
		},
		{
			name:     "airplane departure",
			input:    emoji.AirplaneDeparture,
			expected: "🛫",
		},
		{
			name:     "checkbox with check",
			input:    emoji.CheckBoxWithCheck,
			expected: "☑️ ",
		},
		{
			name:     "double rune: gear",
			input:    emoji.Gear,
			expected: "⚙️ ",
		},
		{
			name:     "man with red hair",
			input:    emoji.Emoji(emoji.Man.Tone(emoji.Default)),
			expected: "👨",
		},
		{
			name:     "woman with white hair",
			input:    emoji.Emoji(emoji.WomanGesturingNo.Tone(emoji.Default)),
			expected: "🙅‍♀️",
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := padRight(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
