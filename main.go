package main

// Use SURF_DEBUG_HEADERS=1 environment variable to print debug headers.

import (
	"fmt"
	"os"
)

/* Nagios exit status */
const OK = 0
const WARNING = 1
const CRITICAL = 2
const UNKNOWN = 3

/* Application defaults */
const author = "Claudio Ramirez <pub.claudio@gmail.com>"
const repo = "https://github.com/nxadm/check-shib3idp-login"
const warning = 5  // timeout seconds
const critical = 20 // timeout seconds
const version = "0.1.0"

func main() {

	/* Command line interface */
	defaults := Defaults{
		Author:   author,
		Repo:     repo,
		Warning:  warning,
		Critical: critical,
		Version:  version,
	}
	params := getParams(defaults)

	/* Configuration file */
	config, err := retrieveValues(params.ConfigFile)

	if err != nil {
		fmt.Printf("[UNKNOWN] Error reading the configuration file: %v\n", err)
		os.Exit(UNKNOWN)
	}

	/* Login */
	status, answerTime, msg := login(config, params, defaults)

	/* Exit status */
	switch status {
	case 0:
		fmt.Printf("[OK] Threshold (w:%d,c:%d), login performed in %f seconds: "+msg+".\n",
			params.Warning, params.Critical, answerTime)
		os.Exit(OK)
	case 1:
		fmt.Printf("[WARNING] Threshold (w:%d,c:%d), login performed in %f seconds: "+msg+".\n",
			params.Warning, params.Critical, answerTime)
		os.Exit(WARNING)
	case 2:
		fmt.Printf("[CRITICAL] Threshold (w:%d,c:%d), login performed in %f seconds: "+msg+".\n",
			params.Warning, params.Critical, answerTime)
		os.Exit(CRITICAL)
	default:
		fmt.Println("[UNKNOWN] Error while executing the login")
		os.Exit(UNKNOWN)

	}
}
