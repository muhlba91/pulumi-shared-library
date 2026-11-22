# Pulumi Shared Library

Go library of Pulumi resources and helpers, imported by other repositories. Public API: everything under `pkg/`.

## Layout

- `pkg/lib/<provider>/...`: resource creation, one `Create` per resource.
- `pkg/model/`: shared data structs.
- `pkg/util/`: helpers.
- `test/mocks/`: Pulumi mocks for tests.

## Commands

- Run `make check` before finishing: `golangci-lint` (`.golangci.yml`), `go vet`, tests.
- Run `make fix` to format.
- Tools come from `mise.toml`; run `mise install` once.

## Conventions

- Resources: `Create(ctx, opts)` or `Create(ctx, name, opts)`; context first, exported options struct last.
- Make optional option fields pointers; read them with `defaults.GetOrDefault`.
- Document every exported function: one summary line, then one `name: description.` line per parameter.
- Wrap errors with `%w` and context; write lowercase messages without a trailing period.
- Declare sentinel errors in `error.go`, named `ErrX`.
- Test in `package x_test` with `testify`; mirror the source file as `<file>_test.go`.
- Add no comment the code already says.
- Suppress lint only as `//nolint:<linter> // <reason>`, lowercase reason, no period.

## Compatibility

- Keep exported names, signatures and defaults stable; the repository is consumed as a Go module.
- Keep resource names and URNs unchanged; a rename replaces the resource in consumers.
- Add options as optional fields.

## Commits

- Conventional Commits, lowercase, imperative, at most 100 characters.
- Allowed scopes: `ci`, `deps`, `release`.
- Add no attribution lines.

## Boundaries

- Allowed: edit code, tests and docs; run lint, vet and tests.
- Ask first: change exported API, resource names or defaults; edit `.github/workflows/release.yml` or the release-please files.
- Owner only: commit, push, tag, merge.
