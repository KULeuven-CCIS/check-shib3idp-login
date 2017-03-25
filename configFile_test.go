package main

import "testing"

func TestRetrieveValues(t *testing.T) {

	username := "some_user"
	password := "some_password"
	unsolicitedLoginBaseurl := "https://idp.example.com/idp/profile/SAML2/Unsolicited/SSO"
	providerId := "https://sp.example.com"
	shire := "https://sp.example.com/Shibboleth.sso/SAML2/POST"

	file := "testfile_config.yaml"
	config, err := retrieveValues(file)
	if err != nil {
		t.Error("Can not parse configuration.")
	}
	if config.Username != username {
		t.Error("Unexepected username value")
	}
	if config.Password != password {
		t.Error("Unexepected password value")
	}
	if config.LoginBaseURL != unsolicitedLoginBaseurl {
		t.Error("Unexepected unsolicited_login_baseurl value")
	}
	if config.ProviderId != providerId {
		t.Error("Unexepected provider_id value")
	}
	if config.Shire != shire {
		t.Error("Unexepected shire value")
	}
}
