package foundation

import (
	"github.com/fatih/color"
)

func PrintError(msg string) {
	color.Red("Error: " + msg)
}

func PrintSuccess(msg string) {
	color.Green(msg)
}
