<a name="unreleased"></a>
## [Unreleased]

### Docs

### Feat
- add predefined regexes to `extract text` module


<a name="v1.23.0"></a>
## [v1.23.0] - 2020-09-14
### Chore
- go mod tidy
- clean up svg files
- **deps:** go mod tidy
- **deps:** update ping
- **deps:** go mod tidy
- **deps:** go mod tidy
- **deps:** go mod tidy
- **deps:** update golang version to `1.15`

### Ci
- disable gosum.yml
- remove gosum and integrate into dops ci
- add code scanning
- add ossar analysis
- run `go run .` instead of `dops`
- add term env
- **dops:** remove `example_casts` directory after example generation

### Docs
- update styling of sidebar
- enhance submodule documentation generator

### Feat
- add `echo` module
- add `execute` category
- add module `open`
- more examples for `random-generator`
- add example to ping command
- add another example to `random-generator`
- redesign whole documentation and add example svg generation
- change `extract-text` module to `extract`
- add escape keybind to quit dops
- add update command to `update` module

### Fix
- add exit handling to interactive mode
- remove URL from example title
- remove `ping` example, as it's not working on linux without permissions
- fix CI flag

### Refactor
- bump version to "v1.23.0"
- better flag detection in code snippets syntax highlighting
- make `ping` example use sudo
- better input handling for `extract text`
- change module structure of `extract text`
- change warning of `ci` module
- rename module variable

### Style
- go fmt
- go fmt

### BREAKING CHANGE

change cli usage of `dops extract-text [...]` to `dops extract text [...]`


<a name="v1.22.7"></a>
## [v1.22.7] - 2020-09-07
### Docs

### Refactor
- bump version to "v1.22.7"


<a name="v1.22.6"></a>
## [v1.22.6] - 2020-09-07
### Docs

### Refactor
- bump version to "v1.22.6"


<a name="v1.22.5"></a>
## [v1.22.5] - 2020-09-07
### Docs

### Refactor
- bump version to "v1.22.5"


<a name="v1.22.4"></a>
## [v1.22.4] - 2020-09-07
### Chore
- **deps:** go mod tidy

### Docs

### Refactor
- bump version to "v1.22.4"


<a name="v1.22.3"></a>
## [v1.22.3] - 2020-09-07
### Docs

### Refactor
- bump version to "v1.22.3"


<a name="v1.22.2"></a>
## [v1.22.2] - 2020-09-06
### Docs

### Refactor
- bump version to "v1.22.2"


<a name="v1.22.1"></a>
## [v1.22.1] - 2020-09-06
### Ci
- fix docup
- trigger changelog generation
- trigger changelog generation
- only run docup on pushes to master
- fix docup
- switch to own version of releasetag
- **changelog:** add changelog ci system
- **changelog:** remove action and generate docs with docup
- **changelog:** update action
- **changelog:** integrate changelog generation into docup
- **docup:** don't display `autodoc` commits
- **docup:** fix style of releases
- **docup:** fix if expression

### Docs
- **changelog:** update changelog
- **changelog:** update changelog
- **changelog:** add changelog

### Refactor
- bump version to "v1.22.1"


<a name="v1.22.0"></a>
## [v1.22.0] - 2020-09-06
### Chore
- **deps:** add go-prompt
- **deps:** go mod tidy
- **deps:** go mod tidy

### Ci
- run release test on pull requests

### Docs
- **reamde:** fix license of mpb

### Feat
- replace interactive cli with a better approach
- better progressbars
- don't display progressbars in ci or raw mode
- add global `ci` flag

### Refactor
- bump version to "v1.22.0"
- add some comments
- add list style to progress descriptions in ci/raw mode


<a name="v1.21.1"></a>
## [v1.21.1] - 2020-09-05
### Ci
- add goreleaser tests
- **goreleaser:** change configuration

### Docs

### Fix
- change maximum value of random generated integers to int32

### Refactor
- bump version to "v1.21.1"


<a name="v1.21.0"></a>
## [v1.21.0] - 2020-09-05
### Chore
- **deps:** go mod tidy

### Docs
- add better usage instructions
- restyle installation instructions
- add tabs to installation instructions
- add faq
- installation tutorial
- **gh-pages:** setup docs for github pages
- **gh-pages:** fix edit document link
- **gh-pages:** add install instructions
- **gh-pages:** update sidebar
- **gh-pages:** add sidebar entry for quick_start.md
- **modules:** update docs
- **readme:** add link to golang.org

### Feat
- add progressbar support
- add progress messages to `ci` module
- add success message format
- add docs generation

### Fix
- remove legacy tests from progressbar
- set success color to green

### Refactor
- bump version to "v1.21.0"


<a name="v1.20.1"></a>
## [v1.20.1] - 2020-08-29
### Docs

### Test
- remove legacy tests from cli


<a name="v1.20.0"></a>
## [v1.20.0] - 2020-08-29
### Chore
- **deps:** go mod tidy
- **deps:** bump gopkg.in/yaml.v2 from 2.2.2 to 2.3.0

### Ci
- **dependabot:** add dependabot config

### Docs

### Feat
- add submodule support and `random-generator` module


<a name="v1.19.0"></a>
## [v1.19.0] - 2020-08-27
### Ci
- **github-linguist:** don't detect irrelevant files

### Feat
- add ping module

### Perf
- some performance improvements

### Refactor
- comment code
- comment forked `cli` code


<a name="v1.18.2"></a>
## [v1.18.2] - 2020-08-26
### Chore
- **deps:** go mod tidy

### Docs

### Fix
- fix tests

### Refactor
- bump version to "v1.18.2"


<a name="v1.18.1"></a>
## [v1.18.1] - 2020-08-26
### Fix
- fix interactive cli not working when `OptionFlag` is missing

### Refactor
- clean up code


<a name="v1.18.0"></a>
## [v1.18.0] - 2020-08-25
### Feat
- module `rename-files` now appends backups

### Refactor
- bump version to "v1.18.0"
- add unique methods for slices


<a name="v1.17.1"></a>
## [v1.17.1] - 2020-08-25
### Docs

### Fix
- module `rename-files` now removes duplicates when hashing
- module `rename-files` now removes duplicates when hashing


<a name="v1.17.0"></a>
## [v1.17.0] - 2020-08-25
### Chore
- **deps:** go mod tidy
- **deps:** go mod tidy

### Ci
- run tests from all packages

### Docs
- **readme:** add attribution table
- **readme:** change platform badge
- **readme:** add platform badge

### Feat
- complete `rename-file` module
- add module `renamefiles`
- show options for `OptionFlag` in command help
- add `OptionFlag` and dropdown to interactive cli

### Refactor
- bump version to "v1.17.0"
- clean up code
- remove todos
- include github.com/urfave/cli/ in source


<a name="v1.16.1"></a>
## [v1.16.1] - 2020-08-24
### Chore
- **deps:** go mod tidy
- **deps:** go mod tidy

### Ci
- trigger `go.sum` regeneration
- add cross-platform testing
- add own token to `tidy` pull request
- remove unused `go dep`

### Docs
- add downloads badge

### Feat
- bigger inputs for interactive mode

### Refactor
- bump version to "v1.16.1"
- remove dot alias imports
- switch to `log.Fatal` for `color` error handling
- clean up code, add comments and configure styling
- clean up code


<a name="v1.16.0"></a>
## [v1.16.0] - 2020-08-23
### Chore
- **deps:** update go deps

### Feat
- accept url as input option
- enable autocomplete support

### Refactor
- ignore error if `clear` command is not found


<a name="v1.15.1"></a>
## [v1.15.1] - 2020-08-20
### Docs

### Fix
- remove "Message" after boolean flag in interactive cli

### Refactor
- bump version to "v1.15.1"


<a name="v1.15.0"></a>
## [v1.15.0] - 2020-08-20
### Chore
- **deps:** go mod tidy

### Ci
- go tests on all branches

### Docs

### Feat
- add append option to output
- add `FileOrStdout` to utils
- add `FileOrStdin` to utils
- add scrollbar to interactive module list

### Perf
- improve performance

### Refactor
- bump version to "v1.15.0"


<a name="v1.14.0"></a>
## [v1.14.0] - 2020-08-17
### Docs

### Feat
- add categories to interactive cli

### Style
- sort imports


<a name="v1.13.1"></a>
## [v1.13.1] - 2020-08-17
### Refactor
- change markdown output


<a name="v1.13.0"></a>
## [v1.13.0] - 2020-08-17
### Chore
- **deps:** go mod tidy
- **deps:** update
- **deps:** go mod tidy

### Docs

### Feat
- show description in interactive cli mode
- show flag usage in interactive cli
- add interactive cli and demo module

### Fix
- don't cut the output in interactive cli mode

### Refactor
- bump version to "v1.13.0"
- clean code
- remove unused `ShowEmpty`
- rename `screens` package to `interactive`


<a name="v1.12.0"></a>
## [v1.12.0] - 2020-08-15
### Feat
- support updating of dops on windows


<a name="v1.11.1"></a>
## [v1.11.1] - 2020-08-14
### Ci
- update release-action to v1.0.2

### Docs

### Fix
- remove debugging output

### Refactor
- bump version to "v1.11.1"


<a name="v1.11.0"></a>
## [v1.11.0] - 2020-08-14
### Feat
- add stdin to extract-text


<a name="v1.10.1"></a>
## [v1.10.1] - 2020-08-14
### Feat
- add better descriptions to the global flags

### Refactor
- bump version to "v1.10.1"


<a name="v1.10.0"></a>
## [v1.10.0] - 2020-08-14
### Feat
- add `r` alias to global `raw` flag

### Style
- format code


<a name="v1.9.0"></a>
## [v1.9.0] - 2020-08-14
### Docs

### Refactor
- bump version to "v1.9.0"
- replace `say` prefixes
- replace `say.Raw` with `say.Text`
- coloring terminal output


<a name="v1.8.1"></a>
## [v1.8.1] - 2020-08-13
### Docs
- **readme:** edit description of dops
- **readme:** add description of dops

### Refactor
- bump version to "v1.8.1"
- change module registration
- change to html line breaks


<a name="v1.8.0"></a>
## [v1.8.0] - 2020-08-13
### Ci
- add docs updater action

### Docs

### Feat
- update module descriptions and usages
- add modules documentation in markdown format
- add `count` flag to `modules` module

### Fix
- offset text in module description
- `markdown` flag of `modules`

### Refactor
- bump version to "v1.8.0"
- `markdown` flag of `modules`


<a name="v1.7.0"></a>
## [v1.7.0] - 2020-08-13
### Feat
- add `markdown` flag to modules

### Refactor
- bump version to "v1.7.0"


<a name="v1.6.0"></a>
## [v1.6.0] - 2020-08-13
### Docs
- **readme:** styling
- **readme:** update styling
- **readme:** add logo

### Feat
- add short option handling
- add `--describe` to `modules` module

### Fix
- raw no longer outputs color

### Refactor
- bump version to "v1.6.0"
- change from HelpName to Name
- change variable names fof help templates
- change apps usage


<a name="v1.5.0"></a>
## [v1.5.0] - 2020-08-13
### Chore
- **deps:** go mod tidy

### Docs
- **funding:** remove unused platforms

### Feat
- change color theme
- add description to module `modules`
- add color to modules help

### Refactor
- bump version to "v1.5.0"
- extract colors to `colors` file
- extract template strings
- capitalize dops as title name
- remove unused `Sprintf`
- add static color import
- move help templates to single file
- add CommandHelpTemplate


<a name="v1.4.1"></a>
## [v1.4.1] - 2020-08-13
### Chore
- **deps:** update deps

### Docs
- **readme:** add install description for go

### Refactor
- bump version to "v1.4.1"


<a name="v1.4.0"></a>
## [v1.4.0] - 2020-08-13
### Ci
- automatically release a new version, if a version change is pushed
- change version comment
- add version comment

### Feat
- bump to version "v1.4.0"
- add more color to dops help


<a name="v1.3.0"></a>
## [v1.3.0] - 2020-08-12
### Feat
- add color to app help
- add `extract-text` module
- add categories to commands

### Refactor
- change categories to constants


<a name="v1.2.0"></a>
## [v1.2.0] - 2020-08-12
### Refactor
- change version number to v1.2.0


<a name="v1.1.2"></a>
## [v1.1.2] - 2020-08-12
### Ci
- fix goreleaser


<a name="v1.1.1"></a>
## [v1.1.1] - 2020-08-12
### Ci
- fix goreleaser


<a name="v1.1.0"></a>
## [v1.1.0] - 2020-08-12
### Chore
- comment code
- fix IntelliJ settings
- **gitignore:** disable uploading of output file.

### Ci
- change goproxy in releaser tool to official one
- change name of golang action
- automate `build` and `test`
- remove changelog generation
- trying to fix changelog generation
- refactor workflow script

### Docs
- add issue templates
- add code of conduct
- add contributing file
- rename `feature request` to `module idea`
- clear feature request template
- **contributing:** add all basic steps to contribute
- **contributing:** add tutorial on how to contribute
- **license:** change copyright
- **readme:** uppercase `dops` again
- **readme:** lowercase `dops`
- **readme:** add badges
- **readme:** add header
- **readme:** change badges to flat style
- **readme:** move license badge to the end
- **readme:**  add license badge and footer
- **readme:** more styling
- **readme:** update README.md

### Feat
- add `bd` alias to `bulkdownload`
- add global `raw` flag to output unstyled text
- add `modules` module to list and search modules
- add sponsor button to GitHub repository
- add terminal color support
- show description instead of usage in `dops -h`
- auto generate changelog
- add update module

### Refactor
- change `registered` to `active`
- usages to lowercase
- change `GFlag` to `GlobalFlag`
- change variable names to shorter ones
- change `modules` package to `module`
- add own interface for global flags `GFlag`
- set version to v1.0.0

### Style
- reformat project


<a name="v1.0.0"></a>
## v1.0.0 - 2020-08-11
### Chore
- **deps:** go mod tidy

### Ci
- change automerge label to enhancement label in go tidy action
- add gosum to tidy golang code automatically


[Unreleased]: https://github.com/dops-cli/dops/compare/v1.23.0...HEAD
[v1.23.0]: https://github.com/dops-cli/dops/compare/v1.22.7...v1.23.0
[v1.22.7]: https://github.com/dops-cli/dops/compare/v1.22.6...v1.22.7
[v1.22.6]: https://github.com/dops-cli/dops/compare/v1.22.5...v1.22.6
[v1.22.5]: https://github.com/dops-cli/dops/compare/v1.22.4...v1.22.5
[v1.22.4]: https://github.com/dops-cli/dops/compare/v1.22.3...v1.22.4
[v1.22.3]: https://github.com/dops-cli/dops/compare/v1.22.2...v1.22.3
[v1.22.2]: https://github.com/dops-cli/dops/compare/v1.22.1...v1.22.2
[v1.22.1]: https://github.com/dops-cli/dops/compare/v1.22.0...v1.22.1
[v1.22.0]: https://github.com/dops-cli/dops/compare/v1.21.1...v1.22.0
[v1.21.1]: https://github.com/dops-cli/dops/compare/v1.21.0...v1.21.1
[v1.21.0]: https://github.com/dops-cli/dops/compare/v1.20.1...v1.21.0
[v1.20.1]: https://github.com/dops-cli/dops/compare/v1.20.0...v1.20.1
[v1.20.0]: https://github.com/dops-cli/dops/compare/v1.19.0...v1.20.0
[v1.19.0]: https://github.com/dops-cli/dops/compare/v1.18.2...v1.19.0
[v1.18.2]: https://github.com/dops-cli/dops/compare/v1.18.1...v1.18.2
[v1.18.1]: https://github.com/dops-cli/dops/compare/v1.18.0...v1.18.1
[v1.18.0]: https://github.com/dops-cli/dops/compare/v1.17.1...v1.18.0
[v1.17.1]: https://github.com/dops-cli/dops/compare/v1.17.0...v1.17.1
[v1.17.0]: https://github.com/dops-cli/dops/compare/v1.16.1...v1.17.0
[v1.16.1]: https://github.com/dops-cli/dops/compare/v1.16.0...v1.16.1
[v1.16.0]: https://github.com/dops-cli/dops/compare/v1.15.1...v1.16.0
[v1.15.1]: https://github.com/dops-cli/dops/compare/v1.15.0...v1.15.1
[v1.15.0]: https://github.com/dops-cli/dops/compare/v1.14.0...v1.15.0
[v1.14.0]: https://github.com/dops-cli/dops/compare/v1.13.1...v1.14.0
[v1.13.1]: https://github.com/dops-cli/dops/compare/v1.13.0...v1.13.1
[v1.13.0]: https://github.com/dops-cli/dops/compare/v1.12.0...v1.13.0
[v1.12.0]: https://github.com/dops-cli/dops/compare/v1.11.1...v1.12.0
[v1.11.1]: https://github.com/dops-cli/dops/compare/v1.11.0...v1.11.1
[v1.11.0]: https://github.com/dops-cli/dops/compare/v1.10.1...v1.11.0
[v1.10.1]: https://github.com/dops-cli/dops/compare/v1.10.0...v1.10.1
[v1.10.0]: https://github.com/dops-cli/dops/compare/v1.9.0...v1.10.0
[v1.9.0]: https://github.com/dops-cli/dops/compare/v1.8.1...v1.9.0
[v1.8.1]: https://github.com/dops-cli/dops/compare/v1.8.0...v1.8.1
[v1.8.0]: https://github.com/dops-cli/dops/compare/v1.7.0...v1.8.0
[v1.7.0]: https://github.com/dops-cli/dops/compare/v1.6.0...v1.7.0
[v1.6.0]: https://github.com/dops-cli/dops/compare/v1.5.0...v1.6.0
[v1.5.0]: https://github.com/dops-cli/dops/compare/v1.4.1...v1.5.0
[v1.4.1]: https://github.com/dops-cli/dops/compare/v1.4.0...v1.4.1
[v1.4.0]: https://github.com/dops-cli/dops/compare/v1.3.0...v1.4.0
[v1.3.0]: https://github.com/dops-cli/dops/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/dops-cli/dops/compare/v1.1.2...v1.2.0
[v1.1.2]: https://github.com/dops-cli/dops/compare/v1.1.1...v1.1.2
[v1.1.1]: https://github.com/dops-cli/dops/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/dops-cli/dops/compare/v1.0.0...v1.1.0
