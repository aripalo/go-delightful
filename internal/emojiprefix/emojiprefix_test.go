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
			name:     "single rune: pile of poop",
			input:    emoji.PileOfPoo,
			expected: 3,
		},
		{
			name:     "single rune: airplane departure",
			input:    emoji.AirplaneDeparture,
			expected: 3,
		},
		{
			name:     "double rune: checkbox with check",
			input:    emoji.CheckBoxWithCheck,
			expected: 2,
		},
		{
			name:     "double rune: gear",
			input:    emoji.Gear,
			expected: 2,
		},
		{
			name:     "triple rune: man with red hair",
			input:    emoji.Emoji(emoji.ManWithRedHair.String()),
			expected: 1,
		},
		{
			name:     "triple rune: woman with white hair",
			input:    emoji.Emoji(emoji.WomanWithWhiteHair.String()),
			expected: 1,
		},
		{
			name:     "quadruple rune: woman gesturing ok",
			input:    emoji.Emoji(emoji.WomanGesturingOk.String()),
			expected: 0,
		},
		{
			name:     "quadruple rune: rainbow flag",
			input:    emoji.Emoji(emoji.RainbowFlag.String()),
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
