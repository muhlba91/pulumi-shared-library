# Pulumi Shared Library

[![](https://img.shields.io/github/license/muhlba91/pulumi-shared-library?style=for-the-badge)](LICENSE)
[![](https://img.shields.io/github/actions/workflow/status/muhlba91/pulumi-shared-library/verify.yml?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/actions/workflows/verify.yml)
[![](https://img.shields.io/coverallsCoverage/github/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/)
[![](https://api.scorecard.dev/projects/github.com/muhlba91/pulumi-shared-library/badge?style=for-the-badge)](https://scorecard.dev/viewer/?uri=github.com/muhlba91/pulumi-shared-library)
[![](https://img.shields.io/github/release-date/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/releases)

> [!IMPORTANT]
> This is a personal library. It is tailored to specific infrastructure patterns and may change without notice.

This repository contains shared Pulumi components and utilities used across personal infrastructure-as-code (IaC) projects. It provides high-level abstractions for common cloud resources to ensure consistency and security by default.

## Features

- **Multi-Cloud Support**: Helpers for AWS, Google Cloud (GCP), Scaleway, and Hetzner.
- **Service Integrations**: Utilities for GitHub, GitLab, Vault, and Kubernetes.
- **Opinionated Defaults**: Security-focused resource configurations (e.g., S3 with forced encryption).
- **Utility Helpers**: Common tasks like string manipulation, file hashing, and Pulumi type conversions.

## Usage

This library is built with Go. You can include it in your Pulumi project by fetching the module:

```bash
go get github.com/muhlba91/pulumi-shared-library
```

### Example: Creating an S3 Bucket

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

## Structure

- [`pkg/lib/`](pkg/lib/): Domain-specific resource abstractions (AWS, GCP, etc.).
- [`pkg/util/`](pkg/util/): General purpose utilities and Pulumi-specific helpers.
- [`pkg/model/`](pkg/model/): Shared data structures and configuration models.

## Development

### Requirements

- Go 1.22+
- `golangci-lint`

### Makefile Commands

To simplify common development tasks, use the following `make` commands:

- `make fix`: Automatically fix linting issues using `golangci-lint fmt`.
- `make lint`: Run code quality checks with `golangci-lint` and `go vet`.
- `make test`: Run all tests with coverage reporting.
- `make coverage`: Generate and open a local HTML coverage report.
