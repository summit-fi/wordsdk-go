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
bin/wordsdk export --api-key <API_KEY> [--dynamic-key <KEY>]
```

Options:

- `--api-key` (required) – API key used to authenticate.
- `--dynamic-key` – export dynamic translations for the provided key instead of static ones.
- `--environment` – target API environment (`production` or `stage`, default `production`).
- `--output`, `-o` – destination directory for generated files (default `./exported`).

These options can also be supplied via environment variables. If both a flag and an environment variable are provided, the flag takes precedence.

| Flag          | Environment variable   |
|---------------|------------------------|
| `--api-key`   | `WORDSDK_API_KEY`      |
| `--dynamic-key` | `WORDSDK_DYNAMIC_KEY` |
| `--environment` | `WORDSDK_ENVIRONMENT` |
| `--output`    | `WORDSDK_OUTPUT`       |

Instead of building you can run the CLI directly:

```bash
go run ./cli export --api-key <API_KEY> [--dynamic-key <KEY>]
```
