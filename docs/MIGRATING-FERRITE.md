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
of the `AsXXX()` functions in Dodeca's `config` package. Note that in some cases
Ferrite offers more purpose-built solutions than Dodeca, such as
[`ferrite.NetworkPort()`], [`ferrite.File()`] and
[`ferrite.KubernetesService()`].

| Dodeca function                | Direct Ferrite equivalent | Most relevant example | Notes                       |
| ------------------------------ | ------------------------- | --------------------- | --------------------------- |
| [`AsBool()`]                   | [`ferrite.Bool()`]        | [Bool (Required)]     |
| [`AsBoolDefault()`]            | [`ferrite.Bool()`]        | [Bool (Default)]      |
| [`AsBoolF()`]                  | [`ferrite.Bool()`]        | [Bool (Default)]      |
| [`AsBoolT()`]                  | [`ferrite.Bool()`]        | [Bool (Default)]      |
| [`AsBytes()`]                  | pending [#22]             | [File (Required)]     | see also [`ferrite.File()`] |
| [`AsBytesDefault()`]           | pending [#22]             | [File (Default)]      | see also [`ferrite.File()`] |
| [`AsDuration()`]               | [`ferrite.Duration()`]    | [Duration (Required)] |
| [`AsDurationBetween()`]        | [`ferrite.Duration()`]    | [Duration (Limits)]   |
| [`AsDurationDefault()`]        | [`ferrite.Duration()`]    | [Duration (Default)]  |
| [`AsDurationDefaultBetween()`] | [`ferrite.Duration()`]    | [Duration (Limits)]   |
| [`AsFloat32()`]                | [`ferrite.Float()`]       | [Float (Required)]    |
| [`AsFloat32Between()`]         | [`ferrite.Float()`]       | [Float (Limits)]      |
| [`AsFloat32Default()`]         | [`ferrite.Float()`]       | [Float (Default)]     |
| [`AsFloat32DefaultBetween()`]  | [`ferrite.Float()`]       | [Float (Limits)]      |
| [`AsFloat64()`]                | [`ferrite.Float()`]       | [Float (Required)]    |
| [`AsFloat64Between()`]         | [`ferrite.Float()`]       | [Float (Limits)]      |
| [`AsFloat64Default()`]         | [`ferrite.Float()`]       | [Float (Default)]     |
| [`AsFloat64DefaultBetween()`]  | [`ferrite.Float()`]       | [Float (Limits)]      |
| [`AsInt()`]                    | [`ferrite.Signed()`]      | [Signed (Required)]   |
| [`AsInt16()`]                  | [`ferrite.Signed()`]      | [Signed (Required)]   |
| [`AsInt16Between()`]           | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt16Default()`]           | [`ferrite.Signed()`]      | [Signed (Default)]    |
| [`AsInt16DefaultBetween()`]    | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt32()`]                  | [`ferrite.Signed()`]      | [Signed (Required)]   |
| [`AsInt32Between()`]           | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt32Default()`]           | [`ferrite.Signed()`]      | [Signed (Default)]    |
| [`AsInt32DefaultBetween()`]    | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt64()`]                  | [`ferrite.Signed()`]      | [Signed (Required)]   |
| [`AsInt64Between()`]           | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt64Default()`]           | [`ferrite.Signed()`]      | [Signed (Default)]    |
| [`AsInt64DefaultBetween()`]    | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt8()`]                   | [`ferrite.Signed()`]      | [Signed (Required)]   |
| [`AsInt8Between()`]            | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsInt8Default()`]            | [`ferrite.Signed()`]      | [Signed (Default)]    |
| [`AsInt8DefaultBetween()`]     | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsIntBetween()`]             | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsIntDefault()`]             | [`ferrite.Signed()`]      | [Signed (Default)]    |
| [`AsIntDefaultBetween()`]      | [`ferrite.Signed()`]      | [Signed (Limits)]     |
| [`AsString()`]                 | [`ferrite.String()`]      | [String (Required)]   | see also [`ferrite.File()`] |
| [`AsStringDefault()`]          | [`ferrite.String()`]      | [String (Default)]    | see also [`ferrite.File()`] |
| [`AsURL()`]                    | pending [#27]             |
| [`AsURLDefault()`]             | pending [#27]             |
| [`AsUint()`]                   | [`ferrite.Unsigned()`]    | [Unsigned (Required)] |
| [`AsUint16()`]                 | [`ferrite.Unsigned()`]    | [Unsigned (Required)] |
| [`AsUint16Between()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint16Default()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Default)]  |
| [`AsUint16DefaultBetween()`]   | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint32()`]                 | [`ferrite.Unsigned()`]    | [Unsigned (Required)] |
| [`AsUint32Between()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint32Default()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Default)]  |
| [`AsUint32DefaultBetween()`]   | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint64()`]                 | [`ferrite.Unsigned()`]    | [Unsigned (Required)] |
| [`AsUint64Between()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint64Default()`]          | [`ferrite.Unsigned()`]    | [Unsigned (Default)]  |
| [`AsUint64DefaultBetween()`]   | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint8()`]                  | [`ferrite.Unsigned()`]    | [Unsigned (Required)] |
| [`AsUint8Between()`]           | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUint8Default()`]           | [`ferrite.Unsigned()`]    | [Unsigned (Default)]  |
| [`AsUint8DefaultBetween()`]    | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUintBetween()`]            | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |
| [`AsUintDefault()`]            | [`ferrite.Unsigned()`]    | [Unsigned (Default)]  |
| [`AsUintDefaultBetween()`]     | [`ferrite.Unsigned()`]    | [Unsigned (Limits)]   |

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
[`ferrite.file()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#File
[`ferrite.float()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Float
[`ferrite.kubernetesservice()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#KubernetesService
[`ferrite.networkport()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#NetworkPort
[`ferrite.signed()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Signed
[`ferrite.string()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#String
[`ferrite.unsigned()`]: https://pkg.go.dev/github.com/dogmatiq/ferrite#Unsigned

<!-- ferrite examples -->

[bool (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Bool-Default
[bool (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Bool-Required
[duration (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Duration-Default
[duration (limits)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Duration-Limits
[duration (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Duration-Required
[file (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-File-Default
[file (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-File-Required
[float (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Float-Default
[float (limits)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Float-Limits
[float (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Float-Required
[signed (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Signed-Default
[signed (limits)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Signed-Limits
[signed (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Signed-Required
[string (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-String-Default
[string (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-String-Required
[unsigned (default)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Unsigned-Default
[unsigned (limits)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Unsigned-Limits
[unsigned (required)]: https://pkg.go.dev/github.com/dogmatiq/ferrite#example-Unsigned-Required

<!-- ferrite issues -->

[#22]: https://github.com/dogmatiq/ferrite/issues/22
[#27]: https://github.com/dogmatiq/ferrite/issues/27
