package eout

import (
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

var LocusColor *color.Color = color.New(color.FgWhite, color.Bold)
var WarningColor *color.Color = color.New(color.FgMagenta, color.Bold)
var ErrorColor *color.Color = color.New(color.FgRed, color.Bold)
var FatalColor *color.Color = color.New(color.FgWhite, color.Bold, color.BgRed)

func AlwaysColor() {
	color.NoColor = false
}

func AutoColor() {
	color.NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stderr.Fd()) && !isatty.IsCygwinTerminal(os.Stderr.Fd()))
}

func NeverColor() {
	color.NoColor = true
}
