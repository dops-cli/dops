// Package color is a modification of https://github.com/fatih/color/blob/master/color.go
package color

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"

	"github.com/dops-cli/dops/global/options"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// Output defines the standard output of the print functions. By default
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()

	// Error defines a color supporting writer for os.Stderr.
	Error = colorable.NewColorableStderr()

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New returns a newly created color object.
func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if NoColor {
		return
	}

	_, err := fmt.Fprintf(Output, "%s[%dm", escape, Reset)
	if err != nil {
		log.Fatal(err)
	}
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	_, err := fmt.Fprintf(Output, c.format())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

func (c *Color) setWriter(w io.Writer) *Color {
	if c.isNoColorSet() {
		return c
	}

	_, err := fmt.Fprintf(w, c.format())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *Color) unsetWriter(w io.Writer) {
	if c.isNoColorSet() {
		return
	}

	if NoColor {
		return
	}

	_, err := fmt.Fprintf(w, "%s[%dm", escape, Reset)
	if err != nil {
		log.Fatal(err)
	}
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprint(w, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintf(w, format, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintln(w, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintln(Output, a...)
}

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color) Sprintf(format string, a ...interface{}) string {
	if options.OutputRaw {
		return fmt.Sprintf(format, a...)
	}
	return c.wrap(fmt.Sprintf(format, a...))
}

// FprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		_, err := c.Fprint(w, a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// PrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		_, err := c.Print(a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// FprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		_, err := c.Fprintf(w, format, a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// PrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		_, err := c.Printf(format, a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// FprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		_, err := c.Fprintln(w, a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// PrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		_, err := c.Println(a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

// Equals returns a boolean value indicating whether two colors are equal.
func (c *Color) Equals(c2 *Color) bool {
	if len(c.params) != len(c2.params) {
		return false
	}

	for _, attr := range c.params {
		if !c2.attrExists(attr) {
			return false
		}
	}

	return true
}

func (c *Color) attrExists(a Attribute) bool {
	for _, attr := range c.params {
		if attr == a {
			return true
		}
	}

	return false
}

func boolPtr(v bool) *bool {
	return &v
}

func getCachedColor(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		c = New(p)
		colorsCache[p] = c
	}

	return c
}

func colorPrint(format string, p Attribute, a ...interface{}) {

	if options.OutputRaw {
		return
	}

	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	if len(a) == 0 {
		_, err := c.Print(format)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := c.Printf(format, a...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func colorString(format string, p Attribute, a ...interface{}) string {

	if options.OutputRaw {
		return format
	}

	c := getCachedColor(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

// Blackf is a convenient helper function to print with black foreground. A
// newline is appended to format by default.
func Blackf(format string, a ...interface{}) { colorPrint(format, FgBlack, a...) }

// Redf is a convenient helper function to print with red foreground. A
// newline is appended to format by default.
func Redf(format string, a ...interface{}) { colorPrint(format, FgRed, a...) }

// Greenf is a convenient helper function to print with green foreground. A
// newline is appended to format by default.
func Greenf(format string, a ...interface{}) { colorPrint(format, FgGreen, a...) }

// Yellowf is a convenient helper function to print with yellow foreground.
// A newline is appended to format by default.
func Yellowf(format string, a ...interface{}) { colorPrint(format, FgYellow, a...) }

// Bluef is a convenient helper function to print with blue foreground. A
// newline is appended to format by default.
func Bluef(format string, a ...interface{}) { colorPrint(format, FgBlue, a...) }

// Magentaf is a convenient helper function to print with magenta foreground.
// A newline is appended to format by default.
func Magentaf(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }

// Cyanf is a convenient helper function to print with cyan foreground. A
// newline is appended to format by default.
func Cyanf(format string, a ...interface{}) { colorPrint(format, FgCyan, a...) }

// Whitef is a convenient helper function to print with white foreground. A
// newline is appended to format by default.
func Whitef(format string, a ...interface{}) { colorPrint(format, FgWhite, a...) }

// SBlack is a convenient helper function to return a string with black
// foreground.
func SBlack(format string, a ...interface{}) string { return colorString(format, FgBlack, a...) }

// SRed is a convenient helper function to return a string with red
// foreground.
func SRed(format string, a ...interface{}) string { return colorString(format, FgRed, a...) }

// SGreen is a convenient helper function to return a string with green
// foreground.
func SGreen(format string, a ...interface{}) string { return colorString(format, FgGreen, a...) }

// SYellow is a convenient helper function to return a string with yellow
// foreground.
func SYellow(format string, a ...interface{}) string { return colorString(format, FgYellow, a...) }

// SBlue is a convenient helper function to return a string with blue
// foreground.
func SBlue(format string, a ...interface{}) string { return colorString(format, FgBlue, a...) }

// SMagenta is a convenient helper function to return a string with magenta
// foreground.
func SMagenta(format string, a ...interface{}) string {
	return colorString(format, FgMagenta, a...)
}

// SCyan is a convenient helper function to return a string with cyan
// foreground.
func SCyan(format string, a ...interface{}) string { return colorString(format, FgCyan, a...) }

// SWhite is a convenient helper function to return a string with white
// foreground.
func SWhite(format string, a ...interface{}) string { return colorString(format, FgWhite, a...) }

// HiBlackf is a convenient helper function to print with hi-intensity black foreground. A
// newline is appended to format by default.
func HiBlackf(format string, a ...interface{}) { colorPrint(format, FgHiBlack, a...) }

// HiRedf is a convenient helper function to print with hi-intensity red foreground. A
// newline is appended to format by default.
func HiRedf(format string, a ...interface{}) { colorPrint(format, FgHiRed, a...) }

// HiGreenf is a convenient helper function to print with hi-intensity green foreground. A
// newline is appended to format by default.
func HiGreenf(format string, a ...interface{}) { colorPrint(format, FgHiGreen, a...) }

// HiYellowf is a convenient helper function to print with hi-intensity yellow foreground.
// A newline is appended to format by default.
func HiYellowf(format string, a ...interface{}) { colorPrint(format, FgHiYellow, a...) }

// HiBluef is a convenient helper function to print with hi-intensity blue foreground. A
// newline is appended to format by default.
func HiBluef(format string, a ...interface{}) { colorPrint(format, FgHiBlue, a...) }

// HiMagentaf is a convenient helper function to print with hi-intensity magenta foreground.
// A newline is appended to format by default.
func HiMagentaf(format string, a ...interface{}) { colorPrint(format, FgHiMagenta, a...) }

// HiCyanf is a convenient helper function to print with hi-intensity cyan foreground. A
// newline is appended to format by default.
func HiCyanf(format string, a ...interface{}) { colorPrint(format, FgHiCyan, a...) }

// HiWhitef is a convenient helper function to print with hi-intensity white foreground. A
// newline is appended to format by default.
func HiWhitef(format string, a ...interface{}) { colorPrint(format, FgHiWhite, a...) }

// SHiBlack is a convenient helper function to return a string with hi-intensity black
// foreground.
func SHiBlack(format string, a ...interface{}) string {
	return colorString(format, FgHiBlack, a...)
}

// SHiRed is a convenient helper function to return a string with hi-intensity red
// foreground.
func SHiRed(format string, a ...interface{}) string { return colorString(format, FgHiRed, a...) }

// SHiGreen is a convenient helper function to return a string with hi-intensity green
// foreground.
func SHiGreen(format string, a ...interface{}) string {
	return colorString(format, FgHiGreen, a...)
}

// SHiYellow is a convenient helper function to return a string with hi-intensity yellow
// foreground.
func SHiYellow(format string, a ...interface{}) string {
	return colorString(format, FgHiYellow, a...)
}

// SHiBlue is a convenient helper function to return a string with hi-intensity blue
// foreground.
func SHiBlue(format string, a ...interface{}) string { return colorString(format, FgHiBlue, a...) }

// SHiMagenta is a convenient helper function to return a string with hi-intensity magenta
// foreground.
func SHiMagenta(format string, a ...interface{}) string {
	return colorString(format, FgHiMagenta, a...)
}

// SHiCyan is a convenient helper function to return a string with hi-intensity cyan
// foreground.
func SHiCyan(format string, a ...interface{}) string { return colorString(format, FgHiCyan, a...) }

// SHiWhite is a convenient helper function to return a string with hi-intensity white
// foreground.
func SHiWhite(format string, a ...interface{}) string {
	return colorString(format, FgHiWhite, a...)
}
