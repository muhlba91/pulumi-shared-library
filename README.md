# Pulumi Shared Library

[![](https://img.shields.io/github/license/muhlba91/pulumi-shared-library?style=for-the-badge)](LICENSE)
[![](https://img.shields.io/github/actions/workflow/status/muhlba91/pulumi-shared-library/verify.yml?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/actions/workflows/verify.yml)
[![](https://img.shields.io/coverallsCoverage/github/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/)
[![](https://api.scorecard.dev/projects/github.com/muhlba91/pulumi-shared-library/badge?style=for-the-badge)](https://scorecard.dev/viewer/?uri=github.com/muhlba91/pulumi-shared-library)
[![](https://img.shields.io/github/release-date/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/releases)

> [!IMPORTANT]
> This is a personal library. It is tailored to specific infrastructure patterns and may change without notice.

This repository contains a Go library of Pulumi resources and helpers for AWS, Google Cloud, Scaleway, Hetzner, GitHub, GitLab, Vault, Kubernetes, NetBird and PostgreSQL, with secure defaults such as forced S3 encryption.

## Installation

```bash
go get github.com/muhlba91/pulumi-shared-library
```

## Usage

Resources live in `pkg/lib`, helpers in `pkg/util` and shared types in `pkg/model`. Most resources follow the `Create(ctx, opts)` pattern.

```go
import (
  "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/s3/bucket"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    _, err := bucket.Create(ctx, &bucket.CreateOptions{
      Name: "my-secure-bucket",
      Labels: map[string]string{
        "Environment": "Production",
      },
    })
    return err
  })
}
```

## Development

Tool versions (Go, golangci-lint) are managed with [mise](https://mise.jdx.dev): run `mise trust` and `mise install` once, then use `make check`.

- `make help`: list the targets.
- `make check`: run `make lint` and `make test`.
- `make fix`: format the code with `golangci-lint fmt`.
- `make lint`: run `golangci-lint` and `go vet`.
- `make test`: run the tests with coverage.
- `make coverage`: open the HTML coverage report.
