---
id: conf-file
title: Configuration file
sidebar_label: Configuration file
slug: /features/configuration
---

## Description

Aurora can be configured throught a configuration file, which allows you to customize the values of some parameters for the aurora commands. For now, you can only customize the following parameters for the `aurora commit` command:

- Commit types
- Commit scopes

## Usage

Create a `.aurora.yml` at the root of your repository. This file can be present, present partially, or absent. If the file or one of its option is missing, the default parameters will apply. The following keys are supported

- `types`: list of different commit types allowed.
- `scopes`: list of different commit scopes allowed. Empty scopes is supported.

When you provide a list of scopes in the `.aurora.yml` configuration file, the input prompted by the CLI is replaced by a select input.

## Example

```yaml
types:
    - feat
    - fix
    - chore
    - refactor
scopes:
    - client
    - server
    - 
```

This file will allow you to create commits like:

```
feat(server): add UsersController
refactor: prefix all api routes by `v1`
fix(client): fix create account button callback 
```

## Todo

- [ ] Add commit line length configuration
- [ ] Add linter configuration
- [ ] Add changelog configuration