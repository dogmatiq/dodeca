# Dodeca

[![Build Status](http://img.shields.io/travis/com/dogmatiq/dodeca/master.svg)](https://travis-ci.com/dogmatiq/dodeca)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/dodeca/master.svg)](https://codecov.io/github/dogmatiq/dodeca)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/dodeca.svg?label=semver)](https://semver.org)
[![GoDoc](https://godoc.org/github.com/dogmatiq/dodeca?status.svg)](https://godoc.org/github.com/dogmatiq/dodeca)
[![Go Report Card](https://goreportcard.com/badge/github.com/dogmatiq/dodeca)](https://goreportcard.com/report/github.com/dogmatiq/dodeca)

Dodeca provides utilities for develoeprs of [12-Factor](http://12factor.net) applications.

## Logging

    <TODO>

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

For any given environment variable `K`, the environment variable `K_VALSRC`
indicates how the content of `K` should be interpreted.

If `K_VALSRC` is:

- empty, undefined or the value `string:plain`, then the content of `K` is treated as a standard environment variable
- the value `string:hex`, then the content of `K` is treated as a binary value, encoded as a hexadecimal string
- the value `string:base64`, then the content of `K` is treated as a binary value, encoded as a standard base-64 string
- the value `file`, then the content of `K` is treated as a path to a file containing the value

#### Consuming configuration

There are two primary ways to consume configuration. The preferred way is via
the `config.Bucket` type, which provides methods for obtaining a `config.Value`,
which in turn has methods for representing that value as a `string`, `[]byte`,
`io.ReadCloser`, or as a path to a real file on disk.

The other way is via the `config.GetEnv()` function, which is a drop-in
replacement for `os.Getenv()`. However, it should be noted that when there is a
problem loading a configuration value, such as when a non-existent file is
specified this function simply returns an empty string.
