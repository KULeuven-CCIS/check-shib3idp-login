package main

import (
	"gopkg.in/headzoo/surf.v1"
	"net/url"
	"regexp"
	"time"
)

func login(config Config, params Params, defaults Defaults) (int, float64, string) {

	// Prepare the request
	escapedSpEntityID := url.QueryEscape(config.ProviderId)
	escapedSpPostBindingURL := url.QueryEscape(config.Shire)
	unsolicitedUrl :=
		config.LoginBaseURL + "?providerId=" + escapedSpEntityID + "&shire=" + escapedSpPostBindingURL
	browser := surf.NewBrowser()
	browser.SetUserAgent("check-shib3idp-login/" + defaults.Version)

	// Open it
	start := time.Now()
	err := browser.Open(unsolicitedUrl)
	if err != nil {
		return CRITICAL, time.Since(start).Seconds(), err.Error()
	}

	// Submit intermediate pag
	if len(browser.Forms()) >= 1 {
		form := browser.Forms()[1]
		if form != nil && form.Submit() != nil {
			return CRITICAL, time.Since(start).Seconds(), err.Error()
		}
	} else {
		return CRITICAL, time.Since(start).Seconds(), "Login failed (unsolicited SSO error)"
	}

	// Login page
	if len(browser.Forms()) >= 1 {
		form := browser.Forms()[1]
		if form != nil {
			err = form.Set("username", config.Username)
			if err != nil {
				return CRITICAL, time.Since(start).Seconds(), err.Error()
			}
			err = form.Set("password", config.Password)
			if err != nil {
				return CRITICAL, time.Since(start).Seconds(), err.Error()
			}
			if form.Submit() != nil {
				return CRITICAL, time.Since(start).Seconds(), err.Error()
			}
		}
	} else {
		return CRITICAL, time.Since(start).Seconds(), "Login failed (user/pass form not found)"
	}
	elapsed := time.Since(start).Seconds()
	matched, _ := regexp.MatchString("\\bname=\"SAMLResponse\"", browser.Body())

	// Exit status
	var nagiosCode int
	var msg string
	switch {
	case !matched:
		nagiosCode = CRITICAL
		msg = "login failed"
	case elapsed >= float64(params.Critical):
		nagiosCode = CRITICAL
		msg = "login is too slow"
	case elapsed >= float64(params.Warning):
		msg = "login is slow"
		nagiosCode = WARNING
	default:
		msg = "everything is OK"
		nagiosCode = OK
	}

	return nagiosCode, elapsed, msg
}
