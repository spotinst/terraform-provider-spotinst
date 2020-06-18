package featureflag

import "os"

// Default features.
var (
	// Toggle the usage of merging credentials in chain provider.
	//
	// This feature allows users to configure their credentials using multiple
	// providers. For example, a token can be statically configured using a file,
	// while the account can be dynamically configured via environment variables.
	MergeCredentialsChain = New("MergeCredentialsChain", false)
)

// EnvVar is the name of the environment variable to read feature flags from.
// The value should be a comma-separated list of K=V flags, while V is optional.
const EnvVar = "SPOTINST_FEATURE_FLAGS"

func init() {
	// Set features from the environment and ignore any errors.
	Set(os.Getenv(EnvVar))
}
