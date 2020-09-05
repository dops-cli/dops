package options

var (
	// Debug returns true if dops is run in debugging mode.
	Debug bool

	// Raw is true if dops was started with the global raw flag.
	// If Raw is true, dops outputs an unformatted text.
	Raw bool

	// CI is true if dops was stated with the global ci flag.
	CI bool
)
