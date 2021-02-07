// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"time"

	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	app "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	"github.com/ibm/the-mesh-for-data/manager/controllers/app/modules"
	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	dc "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	"github.com/ibm/the-mesh-for-data/pkg/serde"
)

// GetConnectionDetails calls the data catalog service
func GetConnectionDetails(req *modules.DataInfo, input *app.M4DApplication) error {
	// Set up a connection to the data catalog interface server.
	conn, err := grpc.Dial(utils.GetDataCatalogServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := dc.NewDataCatalogServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := c.GetDatasetInfo(ctx, &dc.CatalogDatasetRequest{
		AppId:     utils.CreateAppIdentifier(input),
		DatasetId: req.Context.DataSetID,
	})
	if err != nil {
		return err
	}

	details := response.GetDetails()

	protocol, err := utils.GetProtocol(details)
	if err != nil {
		return err
	}
	format, err := utils.GetDataFormat(details)
	if err != nil {
		return err
	}

	connection, err := serde.ToRawExtension(details.DataStore)
	if err != nil {
		return err
	}

	req.DataDetails = &modules.DataDetails{
		Name: details.Name,
		Interface: app.InterfaceDetails{
			Protocol:   protocol,
			DataFormat: format,
		},
		Geography:  details.Geo,
		Connection: *connection,
	}

	return nil
}

// GetCredentials calls the credentials manager service
// TODO: Choose appropriate catalog connector based on the datacatalog service indicated as part of datasetID
func GetCredentials(req *modules.DataInfo, input *app.M4DApplication) error {
	// Set up a connection to the data catalog interface server.
	conn, err := grpc.Dial(utils.GetCredentialsManagerServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := dc.NewDataCredentialServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	dataCredentials, err := c.GetCredentialsInfo(ctx, &dc.DatasetCredentialsRequest{
		DatasetId: req.Context.DataSetID,
		AppId:     utils.CreateAppIdentifier(input)})
	if err != nil {
		return err
	}
	req.Credentials = dataCredentials

	return nil
}

// RegisterAsset registers a new asset in the specified catalog
// Input arguments:
// - catalogID: the destination catalog identifier
// - info: connection and credential details
// Returns:
// - an error if happened
// - the new asset identifier
func (r *M4DApplicationReconciler) RegisterAsset(catalogID string, info *app.DatasetDetails) (string, error) {
	// Set up a connection to the data catalog interface server.
	conn, err := grpc.Dial(utils.GetDataCatalogServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := dc.NewDataCatalogServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var datasetDetails *dc.DatasetDetails
	if err := serde.FromRawExtention(info.Details, datasetDetails); err != nil {
		return "", err
	}
	var creds *dc.Credentials
	if creds, err = r.SecretToCredentials(info.SecretRef); err != nil {
		return "", err
	}

	response, err := c.RegisterDatasetInfo(ctx, &dc.RegisterAssetRequest{
		Creds:                creds,
		DatasetDetails:       datasetDetails,
		DestinationCatalogId: catalogID,
	})
	if err != nil {
		return "", err
	}
	return response.GetAssetId(), nil
}

// SecretToCredentials fetches a secret and constructs Credentials structure
func (r *M4DApplicationReconciler) SecretToCredentials(secretName string) (*dc.Credentials, error) {
	// fetch a secret
	secret := &corev1.Secret{}
	if err := r.Client.Get(context.Background(), types.NamespacedName{Name: secretName, Namespace: utils.GetSystemNamespace()}, secret); err != nil {
		return nil, err
	}
	creds := &dc.Credentials{}
	for key, val := range secret.Data {
		switch key {
		case "accessKeyID":
			creds.AccessKey = string(val)
		case "accessKey":
			creds.AccessKey = string(val)
		case "secretAccessKey":
			creds.SecretKey = string(val)
		case "secretKey":
			creds.SecretKey = string(val)
		case "apiKey":
			creds.ApiKey = string(val)
		case "resourceInstanceId":
			creds.ResourceInstanceId = string(val)
		}
	}
	return creds, nil
}
