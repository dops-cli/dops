# dops
A module based DevOps CLI-Toolkit written in golang.

![Go Test](https://github.com/dops-cli/dops/workflows/Go%20Test/badge.svg)

## Installation

> To install dops just run the following commands

### Linux

To install `dops` you must run the following command in a terminal:

```console
curl -s https://dops-cli.com/get/linux | sudo bash
```

### Windows

To install `dops` you must run the following command in a PowerShell terminal with **administrative privileges**:

```powershell
iwr -useb dops-cli.com/get/windows | iex
```

### Other / Manual install

If you are using a different operating system you can still use `dops`.  
Go to the latest release: https://github.com/dops-cli/dops/releases/latest and download the version you need.

## Usage

To use `dops` all you have to do is to enter `dops` into a terminal. This will automatically display the built-in help of `dops`.  
If you want to learn more about a particular module, you can do so with this command: `dops [module] -h`.

## Versioning

We are in the first major version of `dops`, with every new module, and other significant changes, the `MINOR (v1.↑.0)` number will be increased. For everything else (fixes, docs, refactor, etc.), the `PATCH (v1.X.↑)` number will be increased.
