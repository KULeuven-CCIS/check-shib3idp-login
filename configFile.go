package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

/* Interface with the yaml configuration file */
type Config struct {
	Username     string `yaml:"username,omitempty"`
	Password     string `yaml:"password,omitempty"`
	LoginBaseURL string `yaml:"unsolicited_login_baseurl,omitempty"`
	ProviderId   string `yaml:"provider_id,omitempty"`
	Shire        string `yaml:"shire,omitempty"`
}

func retrieveValues(file string) (Config, error) {
	c := Config{}

	// Read it
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	// Unmarshall it
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}

	// Sanity check
	missing := make([]string, 0, 0)
	switch {
	case c.Username == "":
		missing = append(missing, "username")
		fallthrough
	case c.Password == "":
		missing = append(missing, "password")
		fallthrough
	case c.LoginBaseURL == "":
		missing = append(missing, "unsolicited_login_baseurl")
		fallthrough
	case c.ProviderId == "":
		missing = append(missing, "provider_id")
		fallthrough
	case c.Shire == "":
		missing = append(missing, "shire")
	}
	if len(missing) > 0 {
		err = errors.New("Missing configuration entries: " + strings.Join(missing, ", ") + ".")
		return c, err
	}

	return c, nil
}

func printSampleConfig() {
	sampleConf :=
		`---
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
`
	fmt.Println(sampleConf)
}
