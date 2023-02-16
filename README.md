# bp2022_netlab

## Installation

To install netlab, you can either use the pre-built binaries provided [here](https://github.com/stg-tud/bp2022_netlab/releases) or build one yourself. To build from source, run these commands:

```bash
# HTTPS:
git clone https://github.com/stg-tud/bp2022_netlab.git -o netlab

# SSH:
git clone git@github.com:stg-tud/bp2022_netlab.git -o netlab

cd netlab
make build
make install
```

## Usage

netlab is a simple command line application. Thus you have to run it from a command line. Either install it to a path in `$PATH` (`%Path%` for Windows respectively) or run it inside the current folder, like so:

```bash
# installed in $PATH:
netlab

# in current folder:
./netlab
```

The usage itself should be pretty self-explainatory. You can always run `netlab help` or `netlab help <command>` for help on that.

Currently, these **commands** are supported:

- `netlab version`: Shows the current version number of netlab
- `netlab test [filename]`: Tests the given TOML file for errors but does not generate any output files
- `netlab generate [filename]`: Loads the given TOML file and generates output files for all targets.

There are also some **flags** which might interest you:

- `-h --help`: shows help (same as `netlab help`)
- `-v --version`: shows the version (same as `netlab version`)
- `-d --debug`: enabled debug output/logging
- `-o --overwrite`: overwrites existing files _(`generate` only)_
- `-f <name> --folder <name>`: outputs to the specified folder name instead of the default `output` _(`generate` only)_
- `-t <target> --target <target>`: outputs for the specified target format. If specified, the targets definition from the TOML file will be ignored. Can be specified multiple times for multiple targets. Possible values are for example `core`, `coreemu-lab` and `the-one`. _(`generate` only)_

## Contributing

### VS Code

We use Visual Studio Code with the [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) installed.

For the most convenient development, turn on "Format On Save" in VS Code settings (`@lang:go @id:editor.formatOnSave`) and set the "Default Formatter" (`@lang:go @id:editor.defaultFormatter`) to `golang.go`.

### Pre-commit hook

Additionally you can use the provided **pre-commit hook** to run tests and format check on commit. Activate it by running the following command inside the repository:

```bash
git config core.hooksPath .githooks
```

Please make sure to use `gofmt` and `go test` (or use pre-commit hook instead) as tests will fail if your commits do not comply with the formatting style.

### golangci-lint

We also use the tool [golangci-lint](https://golangci-lint.run/) for some static code analysis. This is also part of our testing pipeline so please make sure to run those tests before commiting (or use pre-commit hook instead).

It can be installed as described [here](https://golangci-lint.run/usage/install/#local-installation). Please make sure it is callable by `golangci-lint` (e. g. by adding its path to the `$PATH` environment variable).

In case check errors should be intentionally be ignored please use linter directives as described [here](https://golangci-lint.run/usage/false-positives/#nolint-directive).
