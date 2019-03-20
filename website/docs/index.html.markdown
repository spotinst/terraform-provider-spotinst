---
layout: "spotinst"
page_title: "Provider: Spotinst"
sidebar_current: "docs-spotinst-index"
description: |-
  The Spotinst provider is used to interact with the resources supported by Spotinst. The provider needs to be configured with the proper credentials before it can be used.
---

# Spotinst Provider

The Spotinst provider is used to interact with the
resources supported by Spotinst. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Spotinst provider
provider "spotinst" {
   token   = "${var.spotinst_token}"
   account = "${var.spotinst_account}"
}

# Create an Elastigroup
resource "spotinst_elastigroup_aws" "foo" {
   # ...
}
```

## Argument Reference

The following arguments are supported:

* `token` - (Required) A Personal API Access Token issued by Spotinst. It can be sourced from the `SPOTINST_TOKEN` environment variable.
* `account` - (Optional) A valid Spotinst account ID. It can be sourced from the `SPOTINST_ACCOUNT` environment variable.

## Credential Precedence

Credentials will be set given the following precedence:
1. credentials defined in the provider block of the template
2. credentials defined as environment variables
3. credentials defined in ~/.spotinst/credentials

Please note that if you omit the Spotinst account, resources will be created using the default account for your organization.