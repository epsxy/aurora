# aurora

*Current status: I'm restarting to work on this project. It has been renamed to Aurora. For now, the docs might be not up-to-date, and it seems like the install through GoBinaries doesn't work anymore. I have a few new things to do, in order to finalize this software and release the v1.0.0, e.g. improving the docs, fixing the installation issues, generalising the use of Bazel and adding new features about commit length, config files, etc. Feel free to open any issue if you're interested or if something doesn't work.*

![go report card](https://goreportcard.com/badge/github.com/epsxy/aurora)
![ci (gh-actions)](https://github.com/epsxy/aurora/actions/workflows/go.yml/badge.svg)

![logo](logo.png)

![demo](commit.gif)

A simple cross platform git commit CLI. Inspired from [commitizen](https://github.com/commitizen/cz-cli) and using the amazing [promptui](https://github.com/manifoldco/promptui) library.

This tool aims to quickly creates commits that follow the [conventional commits](https://www.conventionalcommits.org) spec. It currently supports just a simple part of the spec: generating commit messages of the form:

`<type>([optional scope])[optional !]: <commit message>`

Where,

- `type`: one element from `feat`, `fix`, `build`, `chore`, `ci`, `docs`, `style`, `refactor`, `perf`, `test`.
- `scope`: indicates where the modification took place (component affected for example). Totally optional. Can be `client`, `server`, etc.
- `!`: indicates wether the commit introduces a breaking change (if present) or not. It's optional.
- `commit message`: a single line of max 72 chars indicate the nature of the change.

Examples:

```
feat(client): add login page
feat(client)!: migrate all webservices to v2 api
ci: add an e2e tests step
```

# Todo

- [x] Create commit messages with default params
- [x] Add .env file to customize params
- [x] Release v1.0
- [x] Add installation and usage documentation
- [x] Propose to `git add -A` if running aurora without any file added
- [ ] Support more custom params (commit length, etc)
- [ ] Add tests
- [ ] Add cross platform builds to releases
- [x] Add generate changelog command
- [ ] Add commit validation command
- [x] Gobinaries installation
- [x] BUG(promptui): MacOS ring bell triggers in multi select prompt

# Quickstart

```bash
# build
make build

# run
make run
```

# Installation

## From gobinaries

You can install aurora without using or installing go by downloading it from gobinaries.com.

```
curl -sf https://gobinaries.com/epsxy/aurora | sh
```

Typing `aurora` in a terminal should work now:

```
> aurora

A Go program to help you create conventional commits

Usage:
  aurora [command]

Available Commands:
  changelog   Run changelog generator
  commit      Run commit formatter
  help        Help about any command

Flags:
  -h, --help   help for aurora

Use "aurora [command] --help" for more information about a command.
```

## From source

### Using Makefile
**Recommended**

The following commands will:
- clone the repository
- build the program with `go build`
- move the binary to `/usr/local/bin/aurora`

```
git clone git@github.com:epsxy/aurora.git
cd aurora
make install
```

To uninstall, use `make uninstall`.

You can also install/uninstall the git utilities with the following commands, which copy the tool to `usr/local/bin/git-aur`

```
make install-git
make uninstall-git
```

### Legacy (w/ go install)
**Not recommended**

You need Go installed on your computer.

- Make your your `$GOPATH` and `$PATH` env variables are set up. You can type the following command in your terminal, or add them in your bash configuration file (`.bashrc`, `.zshrc`, etc) if you want this configuration to persist after you close your terminal:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
- Build and install the project binary. This will add the generated executable in $GOPATH/bin:
```bash
go install
```

- You should be ready to go now, the following command should work in your terminal:
```bash
aurora
```

Troubleshooting:
- Verify your environment variables have been correctly set up: `echo $GOPATH`, `echo $PATH`.
- Open and close a new terminal to apply the changes you made to your shell configuration file. Or source it using `source ~/.bashrc`.
- Verify the binary has been correctly generated in your `GOPATH/bin` directory: `ls $GOPATH/bin`.

# Usage

You have to be in a git repository, otherwise aurora will return and error and close itself. You have to git add the files **before** launching aurora. Type `aurora` in a shell at your repository's path and follow the indications to proceed in formatting your commit message.

aurora will basically `git commit -m [message]` with a nicely formatted message you will create using the CLI.

# Configuration file

You can set up your custom commit scopes and types in a configuration file at the route of your repository. Create a `.aurora.yml` file matching the following structure:

```yaml
types:
    - first
    - second
    - ...
scopes:
    - first-scope
    - second-scope
    - ...
```

If you want to propose users to use an empty scope, you can add an empty option:

```yaml
scopes:
    - 
    - first
    - second
```

This file can be present, present partially, or absent. If the file or one of its option is missing, the default parameters will apply. This repository contains a `.aurora.yml` example file.

# Bazel

```
# build
bazel build //:aurora
# test all
bazel test --test_output=errors //...
# run with flags
bazel run //:aurora -- --help
```

# Licence

The MIT Licence