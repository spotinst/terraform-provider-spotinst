# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 1.4.x
-	[Go](https://golang.org/doc/install) 1.21 (to build the provider plugin)

## Building the Provider

Clone the repository:

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd "$_"
$ git clone https://github.com/spotinst/terraform-provider-spotinst.git
```

Choose your build method:

#### 1. `make build` and install it globally

If you don't mind installing the development version of the provider globally,
you can use `make build` in the provider directory which will build and link the
binary into your `$GOPATH/bin` directory.

```sh
$ cd $GOPATH/src/github.com/spotinst/terraform-provider-spotinst
$ make build
```

#### 2. `go build` and install it local to your changes

If you would rather install the provider locally and not impact the stable
version you already have installed, you can use the `~/.terraformrc` file to tell
Terraform where your provider is. You do this by building the provider using Go.

```sh
$ cd $GOPATH/src/github.com/spotinst/terraform-provider-spotinst
$ go build -o terraform-provider-spotinst
```

Then, update your `~/.terraformrc` file to point at the location you've built it.

```hcl
providers {
  spotinst = "${GOPATH}/src/github.com/spotinst/terraform-provider-spotinst/terraform-provider-spotinst"
}
```

A caveat with this approach is that you will need to run `terraform init`
whenever the provider is rebuilt. You'll also need to remember to comment
it/remove it when it's not in use to avoid tripping yourself up.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine (version 1.15+ is *required*). You'll also need to
correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well
as adding `$GOPATH/bin` to your `$PATH`.

See above for which option suits your workflow for building the provider.

## Testing the Provider

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Dependencies

Terraform providers use [Go modules](https://github.com/golang/go/wiki/Modules)
to manage the dependencies. To add or update a dependency, you would run the
following (`v1.2.3` of `foo` is a new package we want to add):

```
$ go get foo@v1.2.3
$ go mod tidy
```

Stepping through the above commands:

- `go get foo@v1.2.3` fetches version `v1.2.3` from the source (if needed) and
adds it to the `go.mod` file for use.
- `go mod tidy` cleans up any dangling dependencies or references that aren't
defined in your module file.

If you wish to remove a dependency, you can remove the reference from `go.mod`
and use the same commands above but omit the initial `go get`.
