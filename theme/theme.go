package theme

import "github.com/pterm/pterm"

var (
	Primary   = pterm.LightCyan
	Secondary = pterm.LightMagenta
	Section   = pterm.NewStyle(pterm.FgYellow, pterm.Bold).Sprintln
	Flag      = pterm.NewStyle(pterm.FgGreen).Sprint
	Separator = pterm.NewStyle(pterm.FgGray).Sprint
)
