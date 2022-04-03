---
id: advanced
title: Advanced installation
sidebar_label: Advanced installation
slug: /installation/advanced
---

This type of installation is recommanded in the following cases:

- You have Go installed on your computer
- You are interested in contributing
- You want to read the sources
- You want to compile it by yourself
- You want to control every single step of the installation process (and maybe adapt it for your own usage)


## Clone the repository

```bash
git clone git@github.com:epsxy/aurora.git
cd aurora
```

After, you can choose your option below:

1. Building and running on the fly using go commands
2. Installing aurora in Go special `bin` directory, under `$GOPATH/bin`
3. Build your own binary from sources and copy it to one your system bin directory (`/usr/bin` or `/usr/local/bin`).

## Build & run on-the-fly

Build with `go build`, run with `go run`. Makefile rules are already defined in `aurora/Makefile`:

```bash
# Go run
make run

# Build and run the generated binary
make build
./bin/aurora
```

## Install in `$GOPATH/bin`

Make sure your `$GOPATH` and `$PATH` env variables are set up. You can type the following command in your terminal, or add them in your bash configuration file (`.bashrc`, `.zshrc`, etc) if you want this configuration to persist after you close your terminal:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Build and install the project binary. This will add the generated executable in `$GOPATH/bin`:
```bash
go install
```

You should be ready to go now, the following command should work in your terminal:
```bash
aurora
```

### Troubleshooting

- Verify your environment variables have been correctly set up: `echo $GOPATH`, `echo $PATH`.
- Open and close a new terminal to apply the changes you made to your shell configuration file. Or source it using `source ~/.bashrc`.
- Verify the binary has been correctly generated in your `GOPATH/bin` directory: `ls $GOPATH/bin`.

## System-wide install from sources

The goal is to generate a binary and install it in a system directory, for example `/usr/local/bin/` for OSX users.

```bash
go build -o ./bin/aurora main.go
sudo cp ./bin/aurora /usr/local/bin/aurora
```