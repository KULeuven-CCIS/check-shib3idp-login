package main

import (
	"net/url"
	"regexp"
	"time"
	//"github.com/nxadm/surf" // Forking for upstream PR
	"gopkg.in/headzoo/surf.v1"
	"encoding/base64"
)

type Result struct {
	Code    int
	Elapsed float64
	Msg     string
}

func login(config Config, params Params, defaults Defaults) Result {
	result := Result{Code: CRITICAL}

	// Prepare the request
	escapedSpEntityID := url.QueryEscape(config.ProviderId)
	escapedSpPostBindingURL := url.QueryEscape(config.Shire)
	unsolicitedUrl :=
		config.LoginBaseURL + "?providerId=" + escapedSpEntityID + "&shire=" + escapedSpPostBindingURL
	browser := surf.NewBrowser()
	// Add timeout if upstream PR merged: https://github.com/headzoo/surf/pull/52
	//browser.SetTimeout(time.Duration(params.Critical) * time.Second)
	browser.SetUserAgent("check-shib3idp-login/" + defaults.Version)

	// Flow
	// 1. Open the unsolicited SSO page
	start := time.Now()

	//Add credentials for basic authentication,
	// ref https://wiki.shibboleth.net/confluence/display/IDP30/PasswordAuthnConfiguration#PasswordAuthnConfiguration-UserInterface
	// "The first user interface layer of the flow is actually HTTP Basic authentication;
	// if a header with credentials is supplied, the credentials are tested immediately with no prompting."
	browser.AddRequestHeader("Authorization", "Basic "+basicAuth(config.Username, config.Password))

	err := browser.Open(unsolicitedUrl)
	if err != nil {
		result.Elapsed = time.Since(start).Seconds()
		result.Msg = "Connection failed or time out reached"
		return result
	}

	// 2. Submit intermediate page when using HTML local storage
	if config.UseLocalStorage {
		if len(browser.Forms()) >= 1 {
			form := browser.Forms()[1]
			if form != nil && form.Submit() != nil {
				result.Elapsed = time.Since(start).Seconds()
				result.Msg = err.Error()
				return result
			}
		} else {
			result.Elapsed = time.Since(start).Seconds()
			result.Msg = "Login failed (unsolicited SSO error)"
			return result
		}
	}

	browser.DelRequestHeader("Authorization")


	// 3. Do something with the SAMLResponse
	result.Elapsed = time.Since(start).Seconds()
	matched, _ := regexp.MatchString("\\bname=\"SAMLResponse\"", browser.Body())

	// Exit status
	switch {
	case !matched:
		result.Code = CRITICAL
		result.Msg = "login failed / Reponse: " + browser.Body()
	case result.Elapsed >= float64(params.Critical):
		result.Code = CRITICAL
		result.Msg = "login is too slow"
	case result.Elapsed >= float64(params.Warning):
		result.Code = WARNING
		result.Msg = "login is slow"
	default:
		result.Code = OK
		result.Msg = "login is OK"
	}

	return result
}


func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
