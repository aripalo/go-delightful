package colors

import "github.com/gookit/color"

var (
	Banner  = color.New(color.FgMagenta, color.OpBold)
	Debug   = color.New(color.FgGray)
	Info    = color.New(color.FgWhite)
	Title   = color.New(color.FgLightBlue, color.OpBold)
	Warning = color.New(color.FgLightYellow)
	Failure = color.New(color.FgRed, color.OpBold)
	Success = color.New(color.FgGreen, color.OpBold)
	Prompt  = color.New(color.FgLightCyan, color.OpBold)
)
