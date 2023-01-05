# bp2022_netlab

## Contributing

We use Visual Studio Code with the [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) installed.

For the most convenient development, turn on "Format On Save" in VS Code settings (`@lang:go @id:editor.formatOnSave`) and set the "Default Formatter" (`@lang:go @id:editor.defaultFormatter`) to `golang.go`.

Additionally you can use the provided **pre-commit hook** to run tests and format check on commit. Activate it by running the following command inside the repository:

```bash
git config core.hooksPath .githooks
```

Please make sure to use `gofmt` and `go test` (or use pre-commit hook instead) as tests will fail if your commits do not comply with the formatting style.
