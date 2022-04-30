// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package datauser

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:gosec
var (
	credserverurl = "http://localhost:8080/v1/creds/usercredentials"
	cred1         = "{\"SecretName\": \"notebook\",\"System\": \"Egeria\",\"Credentials\": {\"username\": \"user1\"}}"
	cred2         = "{\"SecretName\": \"notebook\",\"System\": \"OPA\", \"Credentials\": {\"username\": \"user2\"}}"
	name          = "notebook"
)

func storeCredentials(t *testing.T, cred string) {
	body := strings.NewReader(cred)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, credserverurl, body)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusCreated, "Failed to store credentials")
}

func readCredentials(t *testing.T, path string) {
	url := credserverurl + "/" + path
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Failed to read credentials")
}

func deleteCredentials(t *testing.T, path string) {
	url := credserverurl + "/" + path
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, http.NoBody)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Failed to delete credentials")
}

func TestCredentialAPIs(t *testing.T) {
	SkipOnClosedSocket("localhost:8080", t)
	storeCredentials(t, cred1)
	storeCredentials(t, cred2)
	readCredentials(t, name)
	deleteCredentials(t, name)
}
