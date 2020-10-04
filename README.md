# gommitizen

A simple cross platform git commit CLI. Inspired from [commitizen](https://github.com/commitizen/cz-cli).

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

# Quickstart

```bash
# build
make build

# run
make run
```

# Configuration file

You can set up your custom commit scopes and types in a configuration file at the route of your repository. Create a `.gommitizen.yml` file matching the following structure:

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

This file can be present, present partially, or absent. If the file or one of its option is missing, the default parameters will apply. This repository contains a `.gommitizen.yml` example file.

# Licence

The MIT Licence