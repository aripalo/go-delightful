package colors

import "github.com/gookit/color"

var (
	Bold    = color.Bold
	Banner  = color.New(color.FgMagenta, color.OpBold)
	Debug   = color.Gray
	Info    = color.White
	Title   = color.New(color.FgLightBlue, color.OpBold)
	Warning = color.Yellow
	Failure = color.Red
	Success = color.Green
	Prompt  = color.New(color.FgCyan, color.OpBold)
)
