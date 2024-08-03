---
layout: "spotinst"
page_title: "Spotinst: credentials_gcp"
subcategory: "Accounts"
description: |-
  Provides a Spotinst credential GCP resource.
---

# spotinst\_credentials\_gcp

Provides a Spotinst credential GCP resource.

## Example Usage

```hcl
# set credential GCP
resource "spotinst_credentials_gcp" "cred_gcp" {
  account_id                  = "act-123456"
  type                        = "service_account"
  project_id                  = "demo-labs"
  private_key_id              = "1234567890"
  private_key                 = "-----BEGIN PRIVATE KEY-----abcd1234-----END PRIVATE KEY-----"
  client_email                = "demo-role-act-123456@demo-labs.iam.gserviceaccount.com"
  client_id                   = "1234567890"
  auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
  token_uri                   = "https://oauth2.googleapis.com/token"
  auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
  client_x509_cert_url        = "https://www.googleapis.com/robot/v1/metadata/x509/demo-role-act-123456%40demo-labs.iam.gserviceaccount.com"

  lifecycle {
    ignore_changes = [
      private_key,
      account_id
    ]
  }}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account associated with your token.
* `type` - (Required) Valid values - service_account.
* `project_id` - (Required) Name of project in GCP.
* `private_key_id` - (Required) Private key ID of JSON key created during prerequisites stage.
* `private_key` - (Required) Private key of JSON key created during prerequisites stage.
* `client_email` - (Required) Email associated with service account.
* `client_id` - (Required) Client ID of service account.
* `auth_uri` - (Required, Default: https://accounts.google.com/o/oauth2/auth) The ID of the account associated with your token.
* `token_uri` - (Required, Default: https://oauth2.googleapis.com/token) The ID of the account associated with your token.
* `auth_provider_x509_cert_url` - (Required, Default: https://www.googleapis.com/oauth2/v1/certs) The ID of the account associated with your token.
* `client_x509_cert_url` - (Required) Should be in following format - "https://www.googleapis.com/robot/v1/metadata/x509/".
