package version

import (
	"fmt"
	"runtime"
)

var (
	BuildTime    string
	CommitHash   string
	OS           = runtime.GOOS
	Architecture = runtime.GOARCH
	Major        = 1
	Minor        = 9
	Patch        = 0
)

func GetLongVersion() string {
	vLong := fmt.Sprintf("Version: %d.%d.%d\n", Major, Minor, Patch)
	vLong = vLong + fmt.Sprintf("Hash: %s\n", CommitHash)
	vLong = vLong + fmt.Sprintf("OS: %s\n", OS)
	vLong = vLong + fmt.Sprintf("Arch: %s\n", Architecture)
	vLong = vLong + fmt.Sprintf("Built: %s", BuildTime)

	return vLong
}

func GetShortVersion() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
}
