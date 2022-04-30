// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"fybrik.io/fybrik/pkg/connectors/datacatalog/clients"
	"fybrik.io/fybrik/pkg/environment"
)

const (
	envOPAServerURL             = "OPA_SERVER_URL"
	envCatalogConnectorURL      = "CATALOG_CONNECTOR_URL"
	envCatalogProviderName      = "CATALOG_PROVIDER_NAME"
	envConnectionTimeout        = "CONNECTION_TIMEOUT"
	envDefaultConnectionTimeout = 10
	commandPort                 = 8080
)

// NewRouter returns a new router.
func NewRouter(controller *ConnectorController) *gin.Engine {
	router := gin.Default()
	router.POST("/getPoliciesDecisions", controller.GetPoliciesDecisions)
	return router
}

// RootCmd defines the root cli command
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "opa-connector",
		Short: "Kubernetes based policy manager connector for Fybrik",
	}
	cmd.AddCommand(RunCmd())
	return cmd
}

// RunCmd defines the command for running the connector
func RunCmd() *cobra.Command {
	ip := ""
	port := commandPort
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run opa connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			gin.SetMode(gin.ReleaseMode)
			// Parse environment variables
			opaServerURL, err := environment.MustGetEnv(envOPAServerURL)
			if err != nil {
				return errors.Wrap(err, "failed to retrieve URL for communicating with OPA server")
			}

			if !strings.HasPrefix(opaServerURL, "https://") && !strings.HasPrefix(opaServerURL, "http://") {
				return errors.New("server URL for OPA server must have http or https schema")
			}

			catalogConnectorAddress, err := environment.MustGetEnv(envCatalogConnectorURL)
			if err != nil {
				return errors.Wrap(err, "failed to retrieve URL for communicating with data catalog")
			}

			catalogProviderName, err := environment.MustGetEnv(envCatalogProviderName)
			if err != nil {
				return errors.Wrap(err, "failed to get catalog provider name from the environment")
			}

			timeout := environment.GetEnvAsInt(envConnectionTimeout, envDefaultConnectionTimeout)
			connectionTimeout := time.Duration(timeout) * time.Second

			// Create data catalog client
			catalogClient, err := clients.NewDataCatalog(catalogProviderName, catalogConnectorAddress, connectionTimeout)
			if err != nil {
				return errors.Wrap(err, "failed to create a data catalog client")
			}

			// Create and start connector
			controller := NewConnectorController(opaServerURL, catalogClient)
			router := NewRouter(controller)
			router.Use(gin.Logger())

			bindAddress := fmt.Sprintf("%s:%d", ip, port)
			return router.Run(bindAddress)
		},
	}
	cmd.Flags().StringVar(&ip, "ip", ip, "IP address")
	cmd.Flags().IntVar(&port, "port", port, "Listening port")
	return cmd
}

func main() {
	// Run the cli
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
