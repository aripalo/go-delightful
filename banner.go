package delightful

import (
	"fmt"

	"github.com/aripalo/go-delightful/colors"
	"github.com/aripalo/go-delightful/ruler"
)

type BannerOptions struct {
	Command string
	Version string
	URL     string
	Extra   string
}

func (m *Message) Banner(b BannerOptions) {
	m.HR("=")
	fmt.Fprintln(m.target, colors.Banner(m.appName))

	if m.verboseMode {
		if b.Version != "" {
			fmt.Fprintln(m.target, colors.Debug(fmt.Sprintf("version: %s", b.Version)))
		}
		if b.URL != "" {
			fmt.Fprintln(m.target, colors.Debug(fmt.Sprintf("version: %s", b.URL)))
		}
		if b.Command != "" {
			fmt.Fprintln(m.target, colors.Debug(fmt.Sprintf("command: %s", b.Command)))
		}
		if b.Extra != "" {
			fmt.Fprintln(m.target, colors.Debug(b.Extra))
		}
	}

	m.HR("-")
}

func (m *Message) HR(char string) {
	hr := ruler.Generate(char)
	fmt.Fprintln(m.target, colors.Debug(hr))
}
