package main

// Use SURF_DEBUG_HEADERS=1 environment variable to print debug headers.

import (
	"fmt"
	"os"
)

/* Nagios exit status */
const (
	OK = iota
	WARNING
	CRITICAL
	UNKNOWN
)

var rmap = map[int]string{
	OK:       "OK",
	WARNING:  "WARNING",
	CRITICAL: "CRITICAL",
	UNKNOWN:  "UNKNOWN",
}

/* Application defaults */
type Defaults struct {
	Author   string
	Critical int
	Repo     string
	Version  string
	Warning  int
}

const author = "Claudio Ramirez <pub.claudio@gmail.com>"
const repo = "https://github.com/nxadm/check-shib3idp-login"
const warning = 5   // timeout seconds
const critical = 20 // timeout seconds
const version = "v0.2.0"

var defaults = Defaults{
	Author:   author,
	Repo:     repo,
	Warning:  warning,
	Critical: critical,
	Version:  version,
}

func main() {

	/* Command line interface */
	params := getParams(defaults)

	/* Configuration file */
	config, err := retrieveValues(params.ConfigFile)

	if err != nil {
		fmt.Printf("[UNKNOWN] Error reading the configuration file: %v\n", err)
		os.Exit(UNKNOWN)
	}

	/* Login */
	result := login(config, params, defaults)

	/* Exit status */
	fmt.Printf("[%s] Threshold (w:%d,c:%d), transaction performed in %f seconds: %s\n",
		rmap[result.Code], params.Warning, params.Critical, result.Elapsed, result.Msg)
	os.Exit(result.Code)
}
