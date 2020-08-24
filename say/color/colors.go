package color

var (
	// Section is used for sections
	Section = New(FgHiYellow, Underline).Sprintf

	// Primary is used for primary texts
	Primary = HiCyanString

	// Secondary is used for secondary texts
	Secondary = HiGreenString

	// Flag is used for global and local flags
	Flag = GreenString

	// Separator is used for separators
	Separator = RedString

	// R resets the color
	R = New(Reset).Sprintf("")
)
