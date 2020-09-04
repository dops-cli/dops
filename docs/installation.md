# Quick Start - Install dops

## Windows

### Command

> [!NOTE]
> This command has to be run in an elevated powershell prompt.

```powershell
iwr -useb dops-cli.com/get/windows | iex
```

> [!WARNING]
> This command is executed with administrative rights!\
> Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

## Linux

### Command

> [!NOTE]
> This command has to be run in an elevated shell.

```bash
curl -s https://dops-cli.com/get/linux | sudo bash
```

> [!WARNING]
> This command is executed with administrative rights!\
> Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

## macOS

### Command

> [!NOTE]
> This command has to be run in an elevated shell.

```bash
curl -s https://dops-cli.com/get/linux | sudo bash
```

> [!WARNING]
> This command is executed with administrative rights!\
> Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

## Compile from source

> [!NOTE]
> To compile dops from source, you have to have [go](https://golang.org/) installed.

### Command

```bash
go get -u github.com/dops-cli/dops
```