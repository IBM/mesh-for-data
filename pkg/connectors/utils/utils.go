// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"net/http"

	"github.com/rs/zerolog"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"

	fybrikTLS "fybrik.io/fybrik/pkg/tls"
)

// GetHTTPClient returns an object of type *http.Client.
func GetHTTPClient(log *zerolog.Logger, scheme *kruntime.Scheme) *http.Client {
	client, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	if err != nil {
		log.Error().Err(err)
		return nil
	}
	config, err := fybrikTLS.GetClientTLSConfig(log, client)
	if err != nil {
		log.Error().Err(err)
		return nil
	}
	if config != nil {
		transport := &http.Transport{TLSClientConfig: config}
		return &http.Client{Transport: transport}
	}
	return http.DefaultClient
}
