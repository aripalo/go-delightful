package delightful

import (
	"fmt"

	"github.com/aripalo/go-delightful/colors"
)

// BannerOptions enables passing optional/additional information into the banner
// output. Values given here are only printed if Verbose Mode is active.
type BannerOptions struct {
	Command string
	Version string
	Website string
	Extra   string
}

// Banner prints information about the current application/command.
// Application name shown in magenta color, other info in gray/white.
// No-op if Silent Mode active.
// BannerOptions printed only if Verbose Mode active.
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
