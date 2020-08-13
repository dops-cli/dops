package color

var (
	Section   = New(FgHiYellow, Underline).Sprintf
	Primary   = HiCyanString
	Secondary = HiGreenString
	Flag      = GreenString
	Separator = RedString
)
