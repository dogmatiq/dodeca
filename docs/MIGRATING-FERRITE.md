# Migrating from Dodeca's `config` package to Ferrite

## Overview

With the deprecation of `dogmatiq/dodeca`, we are encouraging users to migrate
to [`dodeca/ferrite`](https://github.com/dogmatiq/ferrite) for reading
environment variables.

Ferrite is a _declarative_ system, in contrast to Dodeca's _imperative_ style.
This means that the set of environment variables, their types and validation
requirements are declared separately to where the values are consumed.

## General usage

First, each environment variable must be declared using a "builder". This is
typically done as a package-scoped variable in the `main` package.

The example below uses a "duration builder" to declare an environment variable
that is used to configure a duration.

```go
package main

var httpTimeout = ferrite.Duration("HTTP_TIMEOUT", "the maximum amount of time to wait for an HTTP request to complete").
    Default(30 * time.Second).
    Required()
```

Next, Ferrite should be initialized at the start of the `main()` function. This
validates environment variable values. If any required environment variables are
missing, or if any environment variables have invalid values, the problems will
be printed to STDERR and the application exists with a non-zero exit code.

```go
func main() {
    ferrite.Init()
    // ... existing code ...
```

Finally, wherever the value of an environment variable is needed, call the
`Value()` method on the builder.

```go
ctx, cancel := context.WithDeadline(ctx, httpTimeout.Value())
defer cancel()
```

There are several
[examples](https://pkg.go.dev/github.com/dogmatiq/ferrite#pkg-examples) for each
builder in Ferrite's documentation.

## Migrating existing environment variables

The table below describes which Ferrite builder to use as a replacement for each
of the `AsXXX()` functions in Dodeca's `config` package.

| Dodeca                         | Ferrite                | Most relevant example |
| ------------------------------ | ---------------------- | --------------------- |
| [`AsBool()`]                   | [`ferrite.Bool()`]     | [Bool (Required)]     |
| [`AsBoolDefault()`]            | [`ferrite.Bool()`]     | [Bool (Default)]      |
| [`AsBoolF()`]                  | [`ferrite.Bool()`]     | [Bool (Default)]      |
| [`AsBoolT()`]                  | [`ferrite.Bool()`]     | [Bool (Default)]      |
| [`AsBytes()`]                  | Proposed in [#22]      |
| [`AsBytesDefault()`]           | Proposed in [#22]      |
| [`AsDuration()`]               | [`ferrite.Duration()`] | [Duration (Required)] |
| [`AsDurationBetween()`]        | Proposed in [#23]      |
| [`AsDurationDefault()`]        | [`ferrite.Duration()`] | [Duration (Default)]  |
| [`AsDurationDefaultBetween()`] | Proposed in [#23]      |
| [`AsFloat32()`]                | Proposed in [#24]      |
| [`AsFloat32Between()`]         | Proposed in [#24]      |
| [`AsFloat32Default()`]         | Proposed in [#24]      |
| [`AsFloat32DefaultBetween()`]  | Proposed in [#24]      |
| [`AsFloat64()`]                | Proposed in [#24]      |
| [`AsFloat64Between()`]         | Proposed in [#24]      |
| [`AsFloat64Default()`]         | Proposed in [#24]      |
| [`AsFloat64DefaultBetween()`]  | Proposed in [#24]      |
| [`AsInt()`]                    | [`ferrite.Signed()`]   | [Signed (Required)]   |
| [`AsInt16()`]                  | [`ferrite.Signed()`]   | [Signed (Required)]   |
| [`AsInt16Between()`]           | Proposed in [#25]      |
| [`AsInt16Default()`]           | [`ferrite.Signed()`]   | [Signed (Default)]    |
| [`AsInt16DefaultBetween()`]    | Proposed in [#25]      |
| [`AsInt32()`]                  | [`ferrite.Signed()`]   | [Signed (Required)]   |
| [`AsInt32Between()`]           | Proposed in [#25]      |
| [`AsInt32Default()`]           | [`ferrite.Signed()`]   | [Signed (Default)]    |
| [`AsInt32DefaultBetween()`]    | Proposed in [#25]      |
| [`AsInt64()`]                  | [`ferrite.Signed()`]   | [Signed (Required)]   |
| [`AsInt64Between()`]           | Proposed in [#25]      |
| [`AsInt64Default()`]           | [`ferrite.Signed()`]   | [Signed (Default)]    |
| [`AsInt64DefaultBetween()`]    | Proposed in [#25]      |
| [`AsInt8()`]                   | [`ferrite.Signed()`]   | [Signed (Required)]   |
| [`AsInt8Between()`]            | Proposed in [#25]      |
| [`AsInt8Default()`]            | [`ferrite.Signed()`]   | [Signed (Default)]    |
| [`AsInt8DefaultBetween()`]     | Proposed in [#25]      |
| [`AsIntBetween()`]             | Proposed in [#25]      |
| [`AsIntDefault()`]             | [`ferrite.Signed()`]   | [Signed (Default)]    |
| [`AsIntDefaultBetween()`]      | Proposed in [#25]      |
| [`AsString()`]                 | [`ferrite.String()`]   | [String (Required)]   |
| [`AsStringDefault()`]          | [`ferrite.String()`]   | [String (Default)]    |
| [`AsURL()`]                    | Proposed in [#27]      |
| [`AsURLDefault()`]             | Proposed in [#27]      |
| [`AsUint()`]                   | [`ferrite.Unsigned()`] | [Unsigned (Required)] |
| [`AsUint16()`]                 | [`ferrite.Unsigned()`] | [Unsigned (Required)] |
| [`AsUint16Between()`]          | Proposed in [#26]      |
| [`AsUint16Default()`]          | [`ferrite.Unsigned()`] | [Unsigned (Default)]  |
| [`AsUint16DefaultBetween()`]   | Proposed in [#26]      |
| [`AsUint32()`]                 | [`ferrite.Unsigned()`] | [Unsigned (Required)] |
| [`AsUint32Between()`]          | Proposed in [#26]      |
| [`AsUint32Default()`]          | [`ferrite.Unsigned()`] | [Unsigned (Default)]  |
| [`AsUint32DefaultBetween()`]   | Proposed in [#26]      |
| [`AsUint64()`]                 | [`ferrite.Unsigned()`] | [Unsigned (Required)] |
| [`AsUint64Between()`]          | Proposed in [#26]      |
| [`AsUint64Default()`]          | [`ferrite.Unsigned()`] | [Unsigned (Default)]  |
| [`AsUint64DefaultBetween()`]   | Proposed in [#26]      |
| [`AsUint8()`]                  | [`ferrite.Unsigned()`] | [Unsigned (Required)] |
| [`AsUint8Between()`]           | Proposed in [#26]      |
| [`AsUint8Default()`]           | [`ferrite.Unsigned()`] | [Unsigned (Default)]  |
| [`AsUint8DefaultBetween()`]    | Proposed in [#26]      |
| [`AsUintBetween()`]            | Proposed in [#26]      |
| [`AsUintDefault()`]            | [`ferrite.Unsigned()`] | [Unsigned (Default)]  |
| [`AsUintDefaultBetween()`]     | Proposed in [#26]      |

<!-- dodeca -->

[`asbool()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBool
[`asbooldefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBoolDefault
[`asboolf()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBoolF
[`asboolt()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBoolT
[`asbytes()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBytes
[`asbytesdefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsBytesDefault
[`asduration()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsDuration
[`asdurationbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsDurationBetween
[`asdurationdefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsDurationDefault
[`asdurationdefaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsDurationDefaultBetween
[`asfloat32()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat32
[`asfloat32between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat32Between
[`asfloat32default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat32Default
[`asfloat32defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat32DefaultBetween
[`asfloat64()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat64
[`asfloat64between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat64Between
[`asfloat64default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat64Default
[`asfloat64defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsFloat64DefaultBetween
[`asint()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt
[`asint16()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt16
[`asint16between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt16Between
[`asint16default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt16Default
[`asint16defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt16DefaultBetween
[`asint32()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt32
[`asint32between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt32Between
[`asint32default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt32Default
[`asint32defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt32DefaultBetween
[`asint64()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt64
[`asint64between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt64Between
[`asint64default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt64Default
[`asint64defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt64DefaultBetween
[`asint8()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt8
[`asint8between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt8Between
[`asint8default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt8Default
[`asint8defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsInt8DefaultBetween
[`asintbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsIntBetween
[`asintdefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsIntDefault
[`asintdefaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsIntDefaultBetween
[`asstring()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsString
[`asstringdefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsStringDefault
[`asurl()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsURL
[`asurldefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsURLDefault
[`asuint()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint
[`asuint16()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint16
[`asuint16between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint16Between
[`asuint16default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint16Default
[`asuint16defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint16DefaultBetween
[`asuint32()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint32
[`asuint32between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint32Between
[`asuint32default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint32Default
[`asuint32defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint32DefaultBetween
[`asuint64()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint64
[`asuint64between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint64Between
[`asuint64default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint64Default
[`asuint64defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint64DefaultBetween
[`asuint8()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint8
[`asuint8between()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint8Between
[`asuint8default()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint8Default
[`asuint8defaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUint8DefaultBetween
[`asuintbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUintBetween
[`asuintdefault()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUintDefault
[`asuintdefaultbetween()`]: https://pkg.go.dev/github.com/dogmatiq/dodeca/config#AsUintDefaultBetween

<!-- ferrite builders -->

[`ferrite.bool()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Bool
[`ferrite.duration()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Duration
[`ferrite.signed()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Signed
[`ferrite.string()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#String
[`ferrite.unsigned()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Unsigned

<!-- ferrite examples -->

[bool (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Bool-Required
[bool (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Bool-Default
[duration (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Duration-Required
[duration (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Duration-Default
[signed (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Signed-Required
[signed (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Signed-Default
[unsigned (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Unsigned-Required
[unsigned (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Unsigned-Default

<!-- ferrite issues -->

[#22]: https://github.com/dogmatiq/ferrite/issues/22
[#23]: https://github.com/dogmatiq/ferrite/issues/23
[#24]: https://github.com/dogmatiq/ferrite/issues/24
[#25]: https://github.com/dogmatiq/ferrite/issues/25
[#26]: https://github.com/dogmatiq/ferrite/issues/26
[#27]: https://github.com/dogmatiq/ferrite/issues/27
