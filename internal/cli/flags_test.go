package cli

import (
	"flag"
	"os"
	"testing"
)

func resetFlags() {
	//We need to reinitialize the command-line flags for each test case
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func testParseFlagsWithRequiredFlags(t *testing.T) {
	//reset flags before test run
	resetFlags()

}

func testParseFlagsWithOptionalFlags(t *testing.T) {
	//reset flags before test run
	resetFlags()

}

func testParseFlagsWithOptionalConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}

func testParseFlagsWithEmptyConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}

func testParseFlagsWithNonExistentConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}
