# Contributing

If you submit a pull request, please keep the following guidelines in mind:

1. Code should be `go fmt` compliant.
2. Types, structs and funcs should be documented.
3. Tests pass.

## Getting set up

Assuming your `$GOPATH` is set up according to your desires, run:

```sh
go get -d github.com/spotinst/spotinst-sdk-go/
```

## Running tests

When working on code in this repository, tests can be run via:

```sh
make test
```

## Changelog

You can see all release changes in the `CHANGELOG.md` file at the root of the
repository. The release notes added to this file will contain service client
updates, and major SDK changes.