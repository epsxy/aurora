---
id: quickstart
title: Quickstart
sidebar_label: Quickstart
slug: /installation/quickstart
---

This type of installation is recommanded in the following cases:

- You don't have Go installed on your computer
- You want to use _Aurora_ out of the box
- You're not interested in contributing nor cloning the repository

## How does it work?

GoBinaries.com allows you to request an on-demand compilation from the latest tagged Aurora version. The binary is compiled on the fly and installed on your computer under for example `/usr/local/bin/aurora` for OSX users.

## Installation

### Latest version

```bash
curl -sf https://gobinaries.com/epsxy/aurora | sh
```

### Specific version

To install version `vX.Y.Z`

```bash
curl -sf https://gobinaries.com/epsxy/aurora@X.Y.Z | sh
```

## After-install checks

You can type `aurora` in your terminal. The following output should show up:

```bash
aurora

A Go program to help you create conventional commits

Usage:
  aurora [command]

Available Commands:
  changelog   Run changelog generator
  commit      Run commit formatter
  help        Help about any command
  lint        Run commit linter

Flags:
  -h, --help   help for aurora

Use "aurora [command] --help" for more information about a command.
```