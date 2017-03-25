package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
	"strconv"
)

type Defaults struct {
	Author   string
	Critical int
	Repo     string
	Version  string
	Warning  int
}

type Params struct {
	ConfigFile string
	Critical   int
	Warning    int
}

func getParams(defaults Defaults) Params {
	args := docoptArgs(defaults)

	/* Fill defaults */
	p := Params{}
	p.Critical = defaults.Critical
	p.Warning = defaults.Warning

	/* Short-cut actions */
	if args["-s"] == true {
		printSampleConfig()
		os.Exit(UNKNOWN)
	}

	/* Fill struct from cli parameters + convert from string if necessary */
	// Required
	if v, ok := args["-f"]; ok {
		p.ConfigFile = v.(string)
	} else {
		fmt.Println("A configuration file is required. Try '-h'.\n")
		os.Exit(UNKNOWN)
	}

	// Optional
	if v, ok := args["-w"]; ok {
		if v != nil {
			int, err := strconv.Atoi(v.(string))
			if err == nil {
				p.Warning = int
			} else {
				fmt.Println("Invalid threshold. Try '-h'.\n")
				os.Exit(UNKNOWN)
			}
		}
	}
	if v, ok := args["-c"]; ok {
		if v != nil {
			int, err := strconv.Atoi(v.(string))
			if err == nil {
				p.Critical = int
			} else {
				fmt.Println("Invalid threshold. Try '-h'.\n")
				os.Exit(UNKNOWN)
			}
		}
	}

	return p
}

func docoptArgs(defaults Defaults) map[string]interface{} {
	versionMsg := "check-shib3idp-login " + defaults.Version + "."
	usage := versionMsg + "\n" +
		`Nagios/Icinga check for an end-to-end Shibboleth 3 IdP login.
Code, bugs and feature requests: ` + defaults.Repo + `.
Author: ` + defaults.Author + `.
        _       _       _       _       _       _       _       _
     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  check-shib3idp-login
      -f <file>
      [-w <threshold> -c <threshold>]
  check-shib3idp-login -s
  check-shib3idp-login -h
  check-shib3idp-login --version

Options:
  -f <file>       Configuration file
  -w <threshold>  Threshold for warning state in seconds
                  [default:` + fmt.Sprintf("%d", defaults.Warning) + `]
  -c <threshold>  Threshold for critical state in seconds
                  [default:` + fmt.Sprintf("%d", defaults.Critical) + `]
  -s              Print a sample YAML configuration file to STDOUT
  -h, --help      Show this screen
  --version       Show version
`
	args, _ := docopt.Parse(usage, nil, true, versionMsg, false)
	return args
}
