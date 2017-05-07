package pallete

import (
	"github.com/fatih/color"
)

var (
	Bold      = color.New(color.Bold)
	Green     = color.New(color.FgGreen)
	BoldGreen = color.New(color.FgGreen, color.Bold)
	Red       = color.New(color.FgRed)
	BoldRed   = color.New(color.FgRed, color.Bold)
	Normal    = color.New(color.Reset)
)
