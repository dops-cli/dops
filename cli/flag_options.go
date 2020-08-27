package cli

import (
	"errors"
	"flag"
	"strings"
)

// Option contains the selected option and a slice of all available options
type Option struct {
	Option  string
	Options []string
}

// Set parses the value into an integer and appends it to the list of values
func (i *Option) Set(value string) error {
	if !sliceContainsString(i.Options, value) {
		return errors.New("value '" + value + "' is not in: " + strings.Join(i.Options, ", "))
	}

	i.Option = value

	return nil
}

// Serialize allows IntSlice to fulfill Serializer
// func (i *Option) Serialize() string {
// 	return i.Option
// }

// Get returns the slice of ints set by this flag
func (i *Option) Get() interface{} {
	return *i
}

// String returns a readable representation of this value (for usage defaults)
func (i *Option) String() string {
	if i == nil {
		return ""
	}
	return i.Option
}

// OptionFlag is a flag with type string
type OptionFlag struct {
	Aliases     []string
	EnvVars     []string
	Options     []string
	Name        string
	Usage       string
	FilePath    string
	DefaultText string
	Value       *Option
	Destination *string
	Required    bool
	Hidden      bool
	TakesFile   bool
	HasBeenSet  bool
}

// IsSet returns whether or not the flag has been set through env or file
func (f *OptionFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *OptionFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag
func (f *OptionFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether or not the flag is required
func (f *OptionFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *OptionFlag) TakesValue() bool {
	return true
}

// GetUsage returns the usage string for the flag
func (f *OptionFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty
// string if the flag takes no value at all.
func (f *OptionFlag) GetValue() string {
	return f.Value.String()
}

// Apply populates the flag given the flag set and environment
func (f *OptionFlag) Apply(set *flag.FlagSet) error {
	f.Value = &Option{
		Options: f.Options,
	}

	if val, ok := flagFromEnvOrFile(f.EnvVars, f.FilePath); ok {
		f.HasBeenSet = true
		err := f.Value.Set(val)
		if err != nil {
			return err
		}
		f.HasBeenSet = true
	}

	for _, name := range f.Names() {
		if f.Value == nil {
			f.Value = &Option{
				Options: f.Options,
			}
		}

		if f.Destination != nil {
			set.StringVar(f.Destination, name, f.Value.Option, f.Usage)
			continue
		}
		set.Var(f.Value, name, f.Usage)
	}

	return nil
}

// Option looks up the value of a local OptionFlag, returns "" if not found
func (c *Context) Option(name string) string {
	if fs := lookupFlagSet(name, c); fs != nil {
		return lookupOption(name, fs)
	}
	return ""
}

func lookupOption(name string, set *flag.FlagSet) string {
	f := set.Lookup(name)

	if f != nil {
		if e, ok := f.Value.(*Option); ok {
			return e.Option
		}
	}
	return ""
}
