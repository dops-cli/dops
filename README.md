<!--suppress HtmlDeprecatedAttribute -->
<a href="https://github.com/dops-cli/dops/">
<img src="https://raw.githubusercontent.com/dops-cli/assets/master/DOPS-Icon-Text-Horizontal.svg" alt="Dops">
</a>

<h1 align="center">:computer: DOPS - CLI DevOps Toolkit</h1>

<p align="center">

<a href="https://github.com/dops-cli/dops/actions?query=workflow%3A%22Go+Test%22">
<img src="https://img.shields.io/github/workflow/status/dops-cli/dops/Go%20Test?style=flat-square" alt="Go Test">
</a>

<a href="https://github.com/dops-cli/dops/issues">
<img src="https://img.shields.io/github/issues/dops-cli/dops.svg?style=flat-square" alt="Issues">
</a>

<a href="https://github.com/dops-cli/dops/stargazers">
<img src="https://img.shields.io/github/stars/dops-cli/dops.svg?style=flat-square" alt="Stars">
</a>

<a href="https://github.com/dops-cli/dops/fork">
<img src="https://img.shields.io/github/forks/dops-cli/dops.svg?style=flat-square" alt="Forks">
</a>

<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

----

<p align="center">
<strong><a href="https://github.com/dops-cli/dops#installation">Installation</a></strong>
|
<strong><a href="https://github.com/dops-cli/dops#usage">Usage</a></strong>
|
<strong><a href="https://github.com/dops-cli/dops/blob/master/CONTRIBUTING.md">Contributing</a></strong>
|
<strong><a href="https://github.com/dops-cli/dops/issues/new?assignees=MarvinJWendt&labels=bug&template=report-a-bug.md&title=">Report Bug</a></strong>
|
<strong><a href="https://github.com/dops-cli/dops/issues/new?assignees=MarvinJWendt&labels=enhancement&template=request-a-feature.md&title=">Submit Feature Idea</a></strong>
</p>

----

## Installation

> To install dops just run the following commands

### Multiplatform using go

> If you have already installed go, you can install `dops` as follows:

```console
go get github.com/dops-cli/dops
```

To update `dops` using go, you have to run:

```console
go get -u github.com/dops-cli/dops
```

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

---

> [dops-cli.com](https://dops-cli.com) &nbsp;&middot;&nbsp;
> GitHub [@dops-cli](https://github.com/dops-cli)
