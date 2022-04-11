package colors

import "github.com/gookit/color"

var (
	Banner  = color.New(color.FgMagenta, color.OpBold)
	Debug   = color.New(color.FgDarkGray)
	Info    = color.New(color.FgGray, color.OpBold)
	Title   = color.New(color.FgLightBlue, color.OpBold)
	Warning = color.New(color.Yellow)
	Failure = color.New(color.FgRed, color.OpBold)
	Success = color.New(color.FgGreen, color.OpBold)
	Prompt  = color.New(color.FgCyan, color.OpBold)
)
