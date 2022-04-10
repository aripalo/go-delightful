package delightful

import (
	"fmt"
	"strings"
)

type BannerOptions struct {
	Version string
	URL     string
	Extra   string
}

func (m *Message) Banner(b BannerOptions) {
	var args []string

	args = append(args, m.appName, "\n")

	if b.Version != "" {
		args = append(args, "Version: ", b.Version, "\n")
	}

	if b.URL != "" {
		args = append(args, "URL: ", b.URL, "\n")
	}

	if b.Extra != "" {
		args = append(args, b.Extra, "\n")
	}

	fmt.Fprint(m.target, strings.Join(args, ""))
}
