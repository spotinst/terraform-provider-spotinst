module github.com/terraform-providers/terraform-provider-spotinst

go 1.12

require (
	github.com/fsouza/go-dockerclient v0.0.0-20160427172547-1d4f4ae73768
	github.com/hashicorp/terraform v0.12.8
	github.com/spotinst/spotinst-sdk-go v0.0.0-20190510223236-6deb754bf781
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
