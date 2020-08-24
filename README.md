<!--suppress HtmlDeprecatedAttribute -->
<a href="https://github.com/dops-cli/dops/">
<img src="https://raw.githubusercontent.com/dops-cli/assets/master/DOPS-Icon-Text-Horizontal.svg" alt="Dops">
</a>

<h1 align="center">:computer: DOPS - CLI DevOps Toolkit</h1>

<p align="center">

<a href="https://github.com/dops-cli/dops/releases">
<img src="https://img.shields.io/github/downloads/dops-cli/dops/total.svg?style=flat-square" alt="Downloads">
</a>

<a href="https://github.com/dops-cli/dops/actions?query=workflow%3A%22Go+Test%22">
<img src="https://img.shields.io/github/workflow/status/dops-cli/dops/Go%20Test?style=flat-square" alt="Go Test">
</a>

<a href="https://github.com/dops-cli/dops/stargazers">
<img src="https://img.shields.io/github/stars/dops-cli/dops.svg?style=flat-square" alt="Stars">
</a>

<a href="https://github.com/dops-cli/dops/fork">
<img src="https://img.shields.io/github/forks/dops-cli/dops.svg?style=flat-square" alt="Forks">
</a>

<a href="https://github.com/dops-cli/dops/issues">
<img src="https://img.shields.io/github/issues/dops-cli/dops.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a href="https://github.com/dops-cli/dops/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
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
<strong><a href="https://github.com/dops-cli/dops/issues/new?assignees=MarvinJWendt&labels=bug&template=report-a-bug.md&title=">Report a Bug</a></strong>
|
<strong><a href="https://github.com/dops-cli/dops/issues/new?assignees=MarvinJWendt&labels=enhancement&template=request-a-feature.md&title=">Submit Feature Idea</a></strong>
</p>

----

## What is `dops` ❓

Dops is a commandline tool, which consists of many different modules. The goal of Dops is to simplify as much DevOps work as possible. Dops runs on the most common operating systems, so that you have the same toolkit available across multiple platforms. Dops is designed to help eliminate annoying scripts and repetitive processes.  
Dops is structured in a way that it is very easy to add a new module. So everyone can write their own module in no time to make their own work easier. This also helps others who have a similar problem, who can then improve the module even further with their own ideas. Which makes adding a module a win-win situation!

## Features 🔥

Dops offers a wide range of different modules (currently there are `6`<!-- feature-count --> modules).  
The individual modules are listed in the [`MODULES.md`](https://github.com/dops-cli/dops/blob/master/MODULES.md) file with description and usage.

## Installation 💿

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

## Usage 📝

To use `dops` all you have to do is to enter `dops` into a terminal. This will automatically display the built-in help of `dops`.  
If you want to learn more about a particular module, you can do so with this command: `dops [module] -h`.

## Versioning

We are in the first major version of `dops`, with every new module, and other significant changes, the `MINOR (v1.↑.0)` number will be increased. For everything else (fixes, docs, refactor, etc.), the `PATCH (v1.X.↑)` number will be increased.

## Attribution

| Package | Original               | Last Original Commit                                                                                                       | License |
|---------|------------------------|----------------------------------------------------------------------------------------------------------------------------|---------|
| color   | github.com/fatih/color | [daf2830f2741ebb735b21709a520c5f37d642d85](https://github.com/fatih/color/commit/daf2830f2741ebb735b21709a520c5f37d642d85) | MIT     |
| cli     | github.com/urfave/cli  | [d2d2098085cee084bc50ae293acf8568cfb348e6](https://github.com/urfave/cli/commit/d2d2098085cee084bc50ae293acf8568cfb348e6)  | MIT     |

---

> [dops-cli.com](https://dops-cli.com) &nbsp;&middot;&nbsp;
> GitHub [@dops-cli](https://github.com/dops-cli)
