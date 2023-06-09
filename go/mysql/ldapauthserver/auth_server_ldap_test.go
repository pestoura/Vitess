/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ldapauthserver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	ldap "gopkg.in/ldap.v2"
)

type MockLdapClient struct{}

func (mlc *MockLdapClient) Connect(network string, config *ServerConfig) error { return nil }
func (mlc *MockLdapClient) Close()                                             {}
func (mlc *MockLdapClient) Bind(username, password string) error {
	if username != "testuser" || password != "testpass" {
		return fmt.Errorf("invalid credentials: %s, %s", username, password)
	}
	return nil
}
func (mlc *MockLdapClient) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return &ldap.SearchResult{}, nil
}

func TestValidateClearText(t *testing.T) {
	asl := &AuthServerLdap{
		Client:         &MockLdapClient{},
		User:           "testuser",
		Password:       "testpass",
		UserDnPattern:  "%s",
		RefreshSeconds: 1,
	}
	_, err := asl.validate("testuser", "testpass")
	require.NoError(t, err, "AuthServerLdap failed to validate valid credentials. Got: %v", err)

	_, err = asl.validate("invaliduser", "invalidpass")
	require.Error(t, err, "AuthServerLdap validated invalid credentials.")

}
