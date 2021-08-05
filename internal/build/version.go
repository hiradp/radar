package build

import (
	"fmt"
)

// Version is dynamically set by the toolchain or overridden by the build system.
var Version = "local"

// Date is dynamically set at build time.
var Date = "now" // YYYY-MM-DD

type Info struct{}

func (i Info) String() string {
	return fmt.Sprintf("radar %s (%s)", Version, Date)
}
