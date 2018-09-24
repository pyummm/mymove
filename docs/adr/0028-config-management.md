# Config Management

* Status: accepted
* Deciders: @dynamike, @pjdufour-truss
* Date: 2018-09-21

## Context and Problem Statement

To make MyMove a canonical [12-factor application](https://12factor.net/), we must do a better job of managing [config](https://12factor.net/config).  Our webserver uses flags and environment variables, but our use of config throughout the application is not managed in a cohesive way.  The use of a more robust config framework with standard patterns will enable the seamless integration of new options and application environmental contexts, e.g., a local docker server, end-to-end test context, mock per-branch environments, etc.

## Decision Drivers

* Maintained (new commits less than 6 months ago)
* Support environment variables, command line flags, and config files.
* Supports integer, duration, and time variables.
* Enables using JSON as config values
* Mark variables as required and implement sanity checks

## Considered Options

* Built-in flag package
* Cobra/Viper/pflag
* github.com/namsral/flag
* github.com/jessevdk/go-flags

## Decision Outcome

Chosen option: "Cobra/Viper/pflag".  This option has the most community support and will give us continued flexibility as the code base grows over time.

## Pros and Cons of the Options

### Built-in flag package

Go ships with a built-in [flag](https://godoc.org/flag) package that provides support for command line flags.

* Good, no additional dependencies.
* Good, maintained but shouldn't receive any improvements either.
* Good, supports bool, (u)int, (u)int64, (u)float64, time.Duration, and string.
* Bad, no support for JSON variables.
* Bad, cannot mark variables as required (only provide defaults)
* Bad, invalid flag values cause panic (making custom sanity checks impossible)

### Cobra/Viper/pflag

[Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), and [pflag](https://github.com/spf13/pflag) are 3 packages that are used together to enable 12-factor applications in Go.  Cobra is used by some of the biggest and most important Go projects, including `Kubernetes`, `Hugo`, `Moby/Docker`, `Delve`, and `GopherJS`.

* Good, cobra/viper/pflag each have over 50 contributors and are actively maintained.
* Good, cobra, viper, and pflag are "owned" by a [Steve Francia](https://github.com/spf13/), a Google employee, and the creator of Hugo.
* Good, supports aliases to enable non-breaking improvements.
* Good, supports bool, int, int64, float64, duration, string, map[string]string, []string, map[string][]string, and time.Time.
* Bad, no support for JSON variables.
* Good, can unmarshal flag values into structs.
* Good, can mark flag as required using Cobra (can also do defaults).
* Good, doesn't panic on bad values and can retrieve errors from pflag if needed.
* Good, supports json, toml, yaml, properties, and hcl config file formats.

### github.com/namsral/flag

[flag](github.com/namsral/flag) is a drop-in replacement for Go's flag package that adds support for environment variables.  Currently used by our webserver.

* Bad, not maintained (the last code update was December 28, 2016).
* Good, supports bool, (u)int, (u)int64, (u)float64, time.Duration, and string (drop in replacement for built-in flag package)
* Bad, no support for JSON variables.
* Good, supports environment variables.
* Bad, only supports `name=value` and `name value` config file formats
* Bad, cannot mark variables as required.
* Bad, invalid values cause panic (making custom sanity checks impossible)

### github.com/jessevdk/go-flags

[go-flags](https://github.com/jessevdk/go-flags) enhances the functionality of the builtin `flag` package with support for many useful features.  You define your config as a single struct using fields and [struct tags](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e).  Currently used by [truss-aws-tools](https://github.com/trussworks/truss-aws-tools).

* Borderline, last updated on March 31, 2018.
* Good, supports a variety of integer, float, string, and maps, including `[]*string{}`
* Bad, no support for JSON variables.
* Good, supports environment variables
* Bad, must unmarshal values into a single config struct.  Creates some baseline structure and increases readability, but reduces flexibility for responsively handling multiple contexts.
* Good, can mark variables as required
* Bad, hard to make custom sanity checks, since the config is parsed into a struct all at once.
