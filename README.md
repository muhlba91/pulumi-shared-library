# Pulumi Shared Library

[![](https://img.shields.io/github/license/muhlba91/pulumi-shared-library?style=for-the-badge)](LICENSE)
[![](https://img.shields.io/github/actions/workflow/status/muhlba91/pulumi-shared-library/verify.yml?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/actions/workflows/verify.yml)
[![](https://img.shields.io/coverallsCoverage/github/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/)
[![](https://api.scorecard.dev/projects/github.com/muhlba91/pulumi-shared-library/badge?style=for-the-badge)](https://scorecard.dev/viewer/?uri=github.com/muhlba91/pulumi-shared-library)
[![](https://img.shields.io/github/release-date/muhlba91/pulumi-shared-library?style=for-the-badge)](https://github.com/muhlba91/pulumi-shared-library/releases)

This repository contains shared Pulumi components and utilities used across personal Pulumi infrastructure-as-code (IaC) projects.

## Features

- Reusable Pulumi components and patterns
- Helper utilities for common cloud tasks

## Requirements

- Go 1.18+ (module-aware)

## Installation (Go)

Fetch the module into your project:

```bash
go get github.com/muhlba91/pulumi-shared-library
```

Import packages as needed:

```go
import "github.com/muhlba91/pulumi-shared-library/pkg/lib/..."
import "github.com/muhlba91/pulumi-shared-library/pkg/model/..."
import "github.com/muhlba91/pulumi-shared-library/pkg/util/..."
```

## Testing & CI

- Unit tests and integration checks are included where applicable.
- CI runs validation and linters; see .github/workflows for details.
