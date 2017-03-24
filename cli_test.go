package main

import (
	"os"
	"testing"
)

var defaultsTest = Defaults{
	Author:   author,
	Warning:  warningThreshold,
	Critical: criticalThreshold,
	Version:  version,
}

func TestGetParams(t *testing.T) {

	// Define the cli combinations to test
	okCliTests := make(map[string][]string)
	okCliTests["okMinimalCli"] = []string{"cmd", "-f", "some_file"}
	okCliTests["okMaximalCli"] =
		[]string{"cmd", "-f", "some_file", "-w", "1", "-c", "1"}
	for _, cli := range okCliTests {
		os.Args = cli
		getParams(defaultsTest)
	}
}
