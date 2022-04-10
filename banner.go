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
		args = append(args, "version: ", b.Version, "\n")
	}

	fmt.Fprint(m.target, strings.Join(args, ""))
}
