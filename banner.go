package delightful

import "fmt"

type BannerOptionalInfo struct {
	version string
	url     string
	extra   string
}

func (m *Message) Banner(b BannerOptionalInfo) {
	var args []string

	if b.version != "" {
		args = append(args, "version: ", b.version, "\n")
	}

	fmt.Fprint(m.target, m.appName, "\n", args)
}
