<div align="center">

# Dodeca

Go utilities for developers of [12-Factor](http://12factor.net) applications.

[![Documentation](https://img.shields.io/badge/go.dev-documentation-007d9c?&style=for-the-badge)](https://pkg.go.dev/github.com/dogmatiq/dodeca)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/dodeca.svg?&style=for-the-badge&label=semver)](https://github.com/dogmatiq/dodeca/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/dogmatiq/dodeca/ci.yml?style=for-the-badge&branch=main)](https://github.com/dogmatiq/dodeca/actions/workflows/ci.yml)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/dodeca/main.svg?style=for-the-badge)](https://codecov.io/github/dogmatiq/dodeca)

</div>

> **This project is deprecated.**
>
> - All Dogmatiq projects that use the Dodeca `logging` package are
>   transitioning to use [`slog`](https://pkg.go.dev/golang.org/x/exp/slog)
>   instead, in anticipation of its inclusion in Go's standard library.
>
> - The Dodeca `config` package has been superceded by
>   [`dogmatiq/ferrite`](https://github.com/dogmatiq/ferrite). Please see the
>   [migration guide](docs/MIGRATING-FERRITE.md) for more information.
>
> This repository will be archived once it is no longer used by other Dogmatiq
> projects.

## Logging

The `logging` package provides a very simple logging interface.

The [12 Factor](https://12factor.net/logs) methodology states that all
application logs should be written to STDOUT, and as such this is the default
behavior of the `logging` package.

Additionally, the logger discriminates between application messages and debug
messages, as per [Dave Cheney's post about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

## Configuration

The `config` package provides features to abstract the **specification** of a
configuration value, from the **consumption** of a configuration value, using
environment variables.

The [12 Factor](https://12factor.net/config) methodology states that all
application configuration should be read from environment variables. However,
it is not uncommon to have configuration requirements with complexity that
exceeds the capability of simple key/value pairs. In such situations, the
obvious solution is to use a configuration file.

The approach we've taken is to allow configuration to be **specified** as a
regular environment variable, or as an environment variable that describes the
path to a configuration file.

The developer can then chose to **consume** this configuration as a `string`,
`[]byte`, `io.ReadCloser`, or as a path to a real file on disk, regardless of
how the configuration as **specified**.

### Usage

#### Specifying configuration

For any given environment variable `K`, the environment variable `K__DATASOURCE`
indicates how the content of `K` should be interpreted.

If `K__DATASOURCE` is:

- empty, undefined or the value `string:plain`, then `K` is a regular variable
- the value `string:hex`, then `K` contains a binary value with hexadecimal encoding
- the value `string:base64`, then `K` contains a binary value with base-64 encoding
- the value `file`, then `K` contains a path to a file containing the value

#### Consuming configuration

There are three primary approaches to consuming configuration. The preferred way
is to use once of the "typed" functions, such as `AsBool()`, `AsInt()`, etc.

These accept a `config.Bucket` type. The second approach is to use the bucket
directly, which gives access to `config.Value` which in turn has methods for
representing that value as a `string`, `[]byte`, `io.ReadCloser`, or as a path
to a real file on disk.

Finally, the `config.GetEnv()` function can be used as a drop-in replacement for
`os.Getenv()`. However, it should be noted that when there is a problem loading
a configuration value, such as when a non-existent file is specified this
function simply returns an empty string.
