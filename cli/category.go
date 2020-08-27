package cli

// CommandCategories interface allows for category manipulation
type CommandCategories interface {
	// AddCommand adds a command to a category, creating a new category if necessary.
	AddCommand(category string, command *Command)

	// Categories returns a copy of the category slice
	Categories() []CommandCategory
}

type commandCategories []*commandCategory

func newCommandCategories() CommandCategories {
	ret := commandCategories([]*commandCategory{})
	return &ret
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c *commandCategories) Less(i, j int) bool {
	return lexicographicLess((*c)[i].Name(), (*c)[j].Name())
}

// Len returns the amount of categories
func (c *commandCategories) Len() int {
	return len(*c)
}

// Swap swaps the elements with indexes i and j.
func (c *commandCategories) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

// AddCommand adds a command to a category
func (c *commandCategories) AddCommand(category string, command *Command) {
	for _, commandCategory := range []*commandCategory(*c) {
		if commandCategory.name == category {
			commandCategory.commands = append(commandCategory.commands, command)
			return
		}
	}
	newVal := append(*c,
		&commandCategory{name: category, commands: []*Command{command}})
	*c = newVal
}

// Categories returns all categories
func (c *commandCategories) Categories() []CommandCategory {
	ret := make([]CommandCategory, len(*c))
	for i, cat := range *c {
		ret[i] = cat
	}
	return ret
}

// CommandCategory is a category containing commands.
type CommandCategory interface {
	// Name returns the category name string
	Name() string
	// VisibleCommands returns a slice of the Commands with Hidden=false
	VisibleCommands() []*Command
}

type commandCategory struct {
	name     string
	commands []*Command
}

// Name returns the category name string
func (c *commandCategory) Name() string {
	return c.name
}

// VisibleCommands returns a slice of the Commands with Hidden=false
func (c *commandCategory) VisibleCommands() []*Command {
	if c.commands == nil {
		c.commands = []*Command{}
	}

	var ret []*Command
	for _, command := range c.commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}
