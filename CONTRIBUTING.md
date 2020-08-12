# Contributing to dops

Jeder ist herzlich eingeladen bei dops mitzumachen. Unser Ziel ist es so viele nützliche Module wie möglich in dops zu integrieren.  
Du kannst auf verschiedene Arten bei der Entwicklung von dops mithelfen. Du kannst Features vorschlagen, oder direkt selber ein feature coden.  

## Writing a new module

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
1. The file you just created should now implement the `Module` interface (`/module/module.go | Module{}`)

Here is an example of such a file:

```go
package modulename

import "github.com/urfave/cli/v2"

type Module struct{}

func (Module) GetCommands() []*cli.Command {
	return nil
}
```

#### Activating your module

To be able to use your module, you have to activate it. This is easily done by putting your module in the file `module/module.go`, in the module slice variable `ActiveModules`. (add to the end: `,yourmodule.Module{}`)