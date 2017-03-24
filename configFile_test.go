package main

import "testing"

func TestRetrieveValues(t *testing.T) {

	username := "some_user"
	password := "some_password"
	unsollicitedLoginBaseurl := "https://idp.example.com/idp/profile/SAML2/Unsolicited/SSO"
	logoutUrl := "https://idp.example.com/idp/profile/Logout"
	serviceProviderEntityId := "https://sp.example.com"
	serviceProviderPostBindingUrl := "https://sp.example.com/Shibboleth.sso/SAML2/POST"

	file := "testfile_config.yaml"
	config, err := retrieveValues(file)
	if err != nil {
		t.Error("Can not parse configuration.")
	}
	if config.Username != username {
		t.Error("Unexepected user value")
	}
	if config.Password != password {
		t.Error("Unexepected user value")
	}
	if config.LoginBaseURL != unsollicitedLoginBaseurl {
		t.Error("Unexepected user value")
	}
	if config.LogoutURL != logoutUrl {
		t.Error("Unexepected user value")
	}
	if config.SpEntityID != serviceProviderEntityId {
		t.Error("Unexepected user value")
	}
	if config.SpPostBindingURL != serviceProviderPostBindingUrl {
		t.Error("Unexepected user value")
	}
}
