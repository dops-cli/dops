package cli

// ActiveGlobalFlags contains all global flags.
// If a global flag is not in this slice, it won't be activated.
var ActiveGlobalFlags []GlobalFlag

// ActiveModules contains all available modules.
// If a module is not in this slice, it won't be activated.
// Except for the module `modules`, which is registered in the main package.
var ActiveModules []Module

// Module is the interface of each module available in dops.
// Each module must return at least one command.
type Module interface {
	GetModuleCommands() []*Command
}

// GlobalFlag is the interface of each global flag in dops.
// Each flag module must return at least one flag.
type GlobalFlag interface {
	GetFlags() []Flag
}
