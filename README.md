# Dodeca

[![Build Status](https://github.com/dogmatiq/dodeca/workflows/CI/badge.svg)](https://github.com/dogmatiq/dodeca/actions?workflow=CI)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/dodeca/master.svg)](https://codecov.io/github/dogmatiq/dodeca)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/dodeca.svg?label=semver)](https://semver.org)
[![Documentation](https://img.shields.io/badge/go.dev-reference-007d9c)](https://pkg.go.dev/github.com/dogmatiq/dodeca)
[![Go Report Card](https://goreportcard.com/badge/github.com/dogmatiq/dodeca)](https://goreportcard.com/report/github.com/dogmatiq/dodeca)

Dodeca provides utilities for developers of [12-Factor](http://12factor.net) applications.

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

There are two primary ways to consume configuration. The preferred way is via
the `config.Bucket` type, which provides methods for obtaining a `config.Value`,
which in turn has methods for representing that value as a `string`, `[]byte`,
`io.ReadCloser`, or as a path to a real file on disk.

The other way is via the `config.GetEnv()` function, which is a drop-in
replacement for `os.Getenv()`. However, it should be noted that when there is a
problem loading a configuration value, such as when a non-existent file is
specified this function simply returns an empty string.
