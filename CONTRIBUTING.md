# Contributing to dops

Everyone is welcome to join dops. Our goal is to integrate as many useful modules as possible into dops.  
You can help with the development of dops in different ways. You can suggest features, or directly code a feature yourself.  

## Adding a new module

Modules are what defines dops. With `dops modules --list` you can display all modules.  
If you want to write a module yourself, you can do so as follows:  

### Setup your local fork

> If you already have experience with GitHub, you can skip this chapter

1. [Fork the repository](https://github.com/dops-cli/dops/fork) into your account.
1. Clone your fork to your PC using the command `git clone https://github.com/YOUR_USERNAME/dops` (replace `dops` at the end of the URL if you have named your fork differently)

You are now ready to code!

### Setting up a new module

1. In your local fork go to the path `module/`.
1. Create a folder with the name of your module.
1. Go into that directory.
1. Now create a `.go` file with the same name as the directory.
1. The file you just created should now implement the `Module` interface (`/module/module.go | Module{}`).
1. The `GetCommands()` function should return a new slice of `cli.Command`.
1. Your `cli.Command` slice must have at least one `cli.Command`.

Here is an example of such a file:

```go
package modulename

import (

"github.com/dops-cli/dops/say"
"github.com/urfave/cli/v2"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
    		{
    			Name:        "modulename",
    			Usage:       "usage of modulename",
    			Description: "description of modulename",
    			Action: func(c *cli.Context) error {
    				// This runs when modulename is executed
                    say.Text("Hello, World!")
    				return nil
    			},
    		},
    	}
}
```

#### Activating your module

To be able to use your module, you have to activate it. This is easily done by putting your module in the file `module/module.go`, in the module slice variable `ActiveModules` (add to the end: `,yourmodule.Module{}`).

Your new module is now ready to use, but it won't do anything, yet :)

*Tip: Test if your module is active by running `go run . yourmodulename` in your project root*

### Give your module functionality

The easiest way is to have a look at already existing modules. It is quite easy to understand the code of these modules. Especially simpler modules like [`modules`](https://github.com/dops-cli/dops/blob/master/module/modules/modules.go) are a good starting point.  

You can also use the [CLI API](https://github.com/urfave/cli/blob/master/docs/v2/manual.md), which is well documented.  

### Submitting your module

After you have finished coding your module (and tested it well), you can submit it to us.

1. Create a commit. (Our commit messages follow the style of [conventionalcommits.org](http://conventionalcommits.org). If your commits are not in this style, we will squash your commits into one and give it a message that fits the style).
1. Push your commit to your fork on GitHub.
1. [Create a pull request](https://github.com/dops-cli/dops/compare).

That's it :rocket: We will review and process your pull request as soon as possible.