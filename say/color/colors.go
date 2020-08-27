package color

var (
	// Section is used for sections
	Section = New(FgHiYellow, Underline).Sprintf

	// Primary is used for primary texts
	Primary = SHiCyan

	// Secondary is used for secondary texts
	Secondary = SHiGreen

	// Flag is used for global and local flags
	Flag = SGreen

	// Separator is used for separators
	Separator = SRed

	// R resets the color
	R = New(Reset).Sprintf("")
)
