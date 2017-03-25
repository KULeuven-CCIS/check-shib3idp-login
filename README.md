# check-shib3idp-login

[![Build Status](https://travis-ci.org/nxadm/check-shib3idp-login.svg?branch=master)](https://travis-ci.org/nxadm/check-shib3idp-login)

A Nagios/Icinga plugin to check an end-to-end user/pass login to a Shibboleth Idp3 instance. The program can also be used standalone.

## Usage

See the help page:

```
$ check-shib3idp-login -h
check-shib3idp-login 0.1.0.
Nagios/Icinga check for an end-to-end Shibboleth 3 IdP login.
Code, bugs and feature requests: https://github.com/nxadm/check-shib3idp-login.
Author: Claudio Ramirez <pub.claudio@gmail.com>.
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
  		  [default:5]
  -c <threshold>  Threshold for critical state in seconds
  		  [default:20]
  -s		  Print a sample YAML configuration file to STDOUT
  -h, --help  	  Show this screen
  --version   	  Show version
```

Examples:

```
$ check-shib3idp-login -f config.yaml
[OK] Threshold (w:5,c:20), transaction performed in 1.216901 seconds: login is OK.

$ check-shib3idp-login -f config.yaml -c 1
[CRITICAL] Threshold (w:5,c:1), transaction performed in 1.169114 seconds: login is too slow.

$ check-shib3idp-login -f config.yaml -c 10 -w 1
[WARNING] Threshold (w:1,c:10), transaction performed in 1.183457 seconds: login is slow.
```

## Configuration

A configuration file is used to store the user and password. You can create a configuration file with the -s switch:

```
---
### check-shib3idp-login configuration ###
# Login information
username: "some_user"
password: "some_password"
# Base URL for unsolicited login
unsolicited_login_baseurl: "https://idp.example.com/idp/profile/SAML2/Unsolicited/SSO"
# Entity ID of a service provider to login
provider_id: "https://sp.example.com"
# ACS url of a service provider to login
shire: "https://sp.example.com/Shibboleth.sso/SAML2/POST"
```
