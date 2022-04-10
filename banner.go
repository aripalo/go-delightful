package delightful

import "fmt"

type BannerOptions struct {
	Version string
	URL     string
	Extra   string
}

func (m *Message) Banner(b BannerOptions) {
	var args []string

	if b.Version != "" {
		args = append(args, "version: ", b.Version, "\n")
	}

	fmt.Fprint(m.target, m.appName, "\n", args)
}
