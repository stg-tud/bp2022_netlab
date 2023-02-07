# bp2022_netlab

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
