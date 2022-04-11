package colors

import "github.com/gookit/color"

var (
	Banner  = color.New(color.FgMagenta, color.OpBold)
	Debug   = color.New(color.Gray)
	Info    = color.New(color.Gray, color.OpBold)
	Title   = color.New(color.FgLightBlue, color.OpBold)
	Warning = color.New(color.Yellow, color.OpBold)
	Failure = color.New(color.FgRed, color.OpBold)
	Success = color.New(color.FgGreen, color.OpBold)
	Prompt  = color.New(color.FgCyan, color.OpBold)
)
