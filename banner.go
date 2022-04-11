package delightful

import (
	"fmt"

	"github.com/aripalo/go-delightful/colors"
	"github.com/aripalo/go-delightful/ruler"
)

type BannerOptions struct {
	Command string
	Version string
	Website string
	Extra   string
}

func (m *Message) Banner(b BannerOptions) {
	if m.silentMode {
		return
	}

	m.printHr("=")
	colors.Banner.Println(m.appName)

	if m.verboseMode {
		if b.Version != "" {
			colors.Debug.Println(fmt.Sprintf("version: %s", colors.Info.Render(b.Version)))
		}
		if b.Website != "" {
			colors.Debug.Println(fmt.Sprintf("website: %s", colors.Info.Render(b.Website)))
		}
		if b.Command != "" {
			colors.Debug.Println(fmt.Sprintf("command: %s", colors.Info.Render(b.Command)))
		}
		if b.Extra != "" {
			colors.Info.Println(fmt.Sprintf("\n%s", b.Extra))
		}
	}

	m.printHr("-")
}

func (m *Message) printHr(char string) {
	if m.verboseMode {
		hr := ruler.Generate(char)
		colors.Debug.Println(hr)
	}
}

// HorizontalRuler prints equals (=) characters as wide as the terminal.
func (m *Message) HorizontalRuler() {
	m.printHr("=")
}
