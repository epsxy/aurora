---
id: releaser
title: Release tool
sidebar_label: Release tool
slug: /features/release
---

Aurora allows you to create automated releases.

## Command

```bash
aurora release
```

The `aurora release` command accept 3 subcommands

```
Create a new release

Usage:
  aurora release [command]

Available Commands:
  fix         Fix release
  major       Major release
  minor       Minor release
```

## Description

Aurora helps you creating release automatically. The only prerequisite is that your git repository **SHOULD** have a `VERSION` file located at its root, containing a version under the form `vX.X.X`. Where `X` represent an integer number greater or equal to 0. For now, we don't support extend types, such as `v1.2.3-beta`, etc.

```
# Supported
v0.0.1
v0.2.1
v11.29.33

# Not Supported 
v1.2.3-beta
1.2.3
v.1.2.3
...
```

Basically, when running a `release` process with aurora, the CLI will:
- Read the current version from the `VERSION` file in your repo
- Increment your version to the next one dependening on the command (see next section for the algorithm)
- Save the new version file
- Commit this update in a release commit with the following commit message: `chore: release ${VERSION}` (with, `${VERSION}` the new computed version number, eg `v1.2.3`)
- Tag this release commit with the new computed version number

You will just have to push the last commit and the last tag with `git push && git push --tags`.

## Behavior

The 3 aurora subcommands are `fix`, `minor`, `major` and references each number of the version:

```
v1.2.3
  \ \ \---- fix
   \ \----- minor
    \------ major
```

Each call to either `fix`, `minor` or `major` will increment the version number to the next suitable version. For example:

```
v1.2.3 ----> aurora release fix   ----> v1.2.4
v1.2.3 ----> aurora release minor ----> v1.3.0
v1.2.3 ----> aurora release major ----> v2.0.0
```
