# wordsdk-go

Go SDK and command-line interface for the Word translation service.

## Installation

Install the CLI with:

```bash
go install github.com/summit-fi/wordsdk-go/cli@latest
```

This places the `wordsdk` binary in your `GOBIN` directory.

## CLI usage

Build the CLI binary:

```bash
go build -o bin/wordsdk ./cli
```

Export translations to [Fluent](https://projectfluent.org/) FTL files:

```bash
bin/wordsdk export --api-key <API_KEY>
```

Options:

- `--api-key` (required) – API key used to authenticate.
- `--environment` – target API environment (`production` or `stage`, default `production`).
- `--output`, `-o` – destination directory for generated files (default `.`).

Instead of building you can run the CLI directly:

```bash
go run ./cli export --api-key <API_KEY>
```
