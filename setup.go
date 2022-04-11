package delightful

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aripalo/go-delightful/pkg/ttyaccess"
	"github.com/gookit/color"
	changecase "github.com/ku/go-change-case"
)

// Message struct holds the configurationg for messaging and you can call
// various methods on it that print messages.
type Message struct {
	appName     string
	colorMode   bool
	emojiMode   bool
	verboseMode bool
	silentMode  bool
	target      io.Writer
}

// New returns a new Output instance with default configuration
func New(appName string) Message {

	// set up initial color mode and disable color output if needed
	colorMode := allowColor(appName)
	if !colorMode {
		color.Disable()
	}

	// choose target io.Writer
	target := ttyaccess.GetWithFallback()
	color.SetOutput(target)

	// set up initial verbose mode
	verboseMode := enableVerboseMode(appName)

	// return an instance
	return Message{
		appName:     appName,
		colorMode:   colorMode,
		emojiMode:   allowEmoji(appName),
		verboseMode: verboseMode,
		silentMode:  false,
		target:      target,
	}
}

// SetColorMode controls if the messages are printed with emojis and colors.
// ColorMode is enabled (true) by default.
// May not have any effect if user has disabled color & emojis via environment variable.
// Or if user's terminal environment does not support colors.
func (m *Message) SetColorMode(colorMode bool) {
	if allowColor(m.appName) {
		m.colorMode = colorMode
		if !m.colorMode {
			color.Disable()
			m.emojiMode = false
		}
	}
}

// SetEmojiMode controls if emojis are printed with the messages:
// Can be disabled even when colors are enabled, but not enabled
// when colors are disabled.
func (m *Message) SetEmojiMode(emojiMode bool) {
	if allowEmoji(m.appName) {
		if m.colorMode {
			m.emojiMode = emojiMode
		}
	}
}

// SetVerboseMode controls additional debug messages.
// Verbose output is disabled by default unless user has set
// VERBOSE or <APP_NAME>_VERBOSE environment variables.
// Enabling verbose mode also disables silent mode.
func (m *Message) SetVerboseMode(verboseMode bool) {
	if !enableVerboseMode(m.appName) {
		m.verboseMode = verboseMode
		if m.verboseMode {
			m.silentMode = false
		}
	}
}

// SetSilentMode controls if info/warning/success
// messages are shown or not (i.e. silent mode).
// Silent mode can not be enabled if verbose mode is active.
func (m *Message) SetSilentMode(silentMode bool) {
	if !m.verboseMode {
		m.silentMode = silentMode
	}
}

// SetMessageTarget overrides the default output target (tty/stderr).
// Mainly used for testing.
func (m *Message) SetMessageTarget(target io.Writer) {
	m.target = target
	color.SetOutput(m.target)
}

// allowColor is responsible for ensuring that user configuration is respected
// by checking if user environment dictates coloured output should not be used.
func allowColor(appName string) bool {
	// Check if NO_COLOR set https://no-color.org/
	if isEnvVarSetOrTruthy("NO_COLOR") {
		return false
	}

	// Check if app-specific _NO_COLOR set https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46
	if isEnvVarSetOrTruthy(formatPrefixedEnvVar(appName, "NO_COLOR")) {
		return false
	}

	// Check if $TERM=dumb https://unix.stackexchange.com/a/43951
	if os.Getenv("TERM") == "dumb" {
		return false
	}

	// Otherwise default to colors being enabled
	return true
}

// allowEmoji enables separate control of emoji outout
func allowEmoji(appName string) bool {

	// if colors are disabled, then disable emojis too
	if !allowColor(appName) {
		return false
	}

	// Check if NO_EMOJI
	if isEnvVarSetOrTruthy("NO_EMOJI") {
		return false
	}

	// Check if app-specific _NO_EMOJI set https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46
	if isEnvVarSetOrTruthy(formatPrefixedEnvVar(appName, "NO_EMOJI")) {
		return false
	}

	// Otherwise default to emojis being enabled
	return true
}

// enableVerboseMode sets the initial verbosity setting by looking at the user
// environment for specific variables.
func enableVerboseMode(appName string) bool {
	if isEnvVarSetOrTruthy("VERBOSE") {
		return true
	}

	if isEnvVarSetOrTruthy(formatPrefixedEnvVar(appName, "VERBOSE")) {
		return true
	}

	return false
}

func isEnvVarSetOrTruthy(envVar string) bool {
	value, set := os.LookupEnv(envVar)

	// check if not set at all
	if !set {
		return false
	}

	// check if falsy value
	lowerValue := strings.ToLower(value)
	if lowerValue == "false" || lowerValue == "0" {
		return false
	}

	// is set (or has truthy value)
	return true
}

func formatPrefixedEnvVar(prefix string, envVar string) string {
	return fmt.Sprintf(
		"%s_%s",
		changecase.Constant(strings.TrimSpace(prefix)),
		changecase.Constant(strings.TrimSpace(envVar)),
	)
}
