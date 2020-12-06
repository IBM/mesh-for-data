package razee

import (
	"context"
	"github.com/IBM/go-sdk-core/core"
	"github.com/machinebox/graphql"
	"net/http"
)

type Response struct {
	Organizations []Organization `json:"organizations"`
}

type Organization struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cluster struct {
	ClusterId     string         `json:"clusterId"`
	Name          string         `json:"name"`
	Status        string         `json:"status"`
	ClusterGroups []ClusterGroup `json:"groups"`
}

type ClusterGroup struct {
	GroupId string `json:"uuid"`
	Name    string `json:"name"`
}

type Resource struct {
	Data string `json:"data"`
}

type RazeeClient struct {
	client *graphql.Client
}

func (r *RazeeClient) getResourceByKeys(orgId string, clusterId string, selfLink string) (string, error) {
	req := graphql.NewRequest(`
    query ($orgId: String!, $clusterId: String!, $selfLink: String!) {resourceByKeys(orgId: $orgId, clusterId: $clusterId, selfLink: $selfLink) {
	  data
	}}
	`)

	req.Var("orgId", orgId)
	req.Var("clusterId", clusterId)
	req.Var("selfLink", selfLink)

	var res struct {
		Resource Resource `json:"resourceByKeys"`
	}

	err := r.client.Run(context.Background(), req, &res)

	if err != nil {
		return "", err
	}

	return res.Resource.Data, nil
}

func (r *RazeeClient) GetOrganizations() ([]Organization, error) {
	req := graphql.NewRequest(`
    query {organizations {
  		id
  		name
	}}
	`)

	var res struct {
		Organizations []Organization `json:"organizations"`
	}
	err := r.client.Run(context.Background(), req, &res)

	if err != nil {
		return nil, err
	}
	return res.Organizations, nil
}

func (r *RazeeClient) GetClustersByOrgId(orgId string) ([]Cluster, error) {
	req := graphql.NewRequest(`
    query ($orgId: String!) {clustersByOrgId(orgId: $orgId) {
	  clusterId
	  name
	  status
	  groups {
		uuid
		name
	  }
	}}
	`)

	req.Var("orgId", orgId)
	var result struct {
		Clusters []Cluster `json:"clustersByOrgId"`
	}

	err := r.client.Run(context.Background(), req, &result)

	if err != nil {
		return nil, err
	}
	return result.Clusters, nil
}

func (r *RazeeClient) getClusterByName(orgId string, clusterName string) (*Cluster, error) {
	req := graphql.NewRequest(`
    query ($orgId: String!, $clusterName: String!) {clusterByName(orgId: $orgId, clusterName: $clusterName) {
	  clusterId
	  name
	  status
	}}
	`)

	req.Var("orgId", orgId)
	req.Var("clusterName", clusterName)
	var result struct {
		Cluster Cluster `json:"clusterByName"`
	}
	err := r.client.Run(context.Background(), req, &result)

	if err != nil {
		return nil, err
	}
	return &result.Cluster, nil
}

func (r *RazeeClient) getChannelByName(orgId string, clusterName string) (*Cluster, error) {
	req := graphql.NewRequest(`
    query ($orgId: String!, $clusterName: String!) {clusterByName(orgId: $orgId, clusterName: $clusterName) {
	  clusterId
	  name
	  status
	}}
	`)

	req.Var("orgId", orgId)
	req.Var("clusterName", clusterName)
	var result struct {
		Cluster Cluster `json:"clusterByName"`
	}
	err := r.client.Run(context.Background(), req, &result)

	if err != nil {
		return nil, err
	}
	return &result.Cluster, nil
}

func (r *RazeeClient) createGroup(orgId string, groupName string) (string, error) {
	req := graphql.NewRequest(`
    mutation($orgId: String!, $name: String!){addGroup(orgId: $orgId, name: $name) {
	  uuid
	}}
	`)

	req.Var("orgId", orgId)
	req.Var("name", groupName)
	var result struct {
		AddGroup struct {
			Uuid string `json:"uuid"`
		} `json:"addGroup"`
	}
	err := r.client.Run(context.Background(), req, &result)

	if err != nil {
		return "", err
	}
	return result.AddGroup.Uuid, nil
}

type IAMRazeeClient struct {
	apiKey string
	client *graphql.Client
}

func NewIAMRazeeClient(apiKey string) *RazeeClient {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	iamHTTPClient := http.Client{
		Transport: &IAMRoundTripper{
			authenticator: authenticator,
		}}

	client := graphql.NewClient("https://config.satellite.cloud.ibm.com/graphql", graphql.WithHTTPClient(&iamHTTPClient))

	return &RazeeClient{
		client: client,
	}
}

func NewGithubAPIKeyClient(url string, apiKey string) *RazeeClient {
	httpClient := http.Client{
		Transport: &RazeeGithubApiRoundTripper{
			apiKey: apiKey,
		}}

	client := graphql.NewClient(url, graphql.WithHTTPClient(&httpClient))

	return &RazeeClient{
		client: client,
	}
}

func NewRazeeLocalClient(url string, login string, password string) *RazeeClient {
	httpClient := http.Client{
		Transport: &RazeeLocalRoundTripper{
			url:      url,
			login:    login,
			password: password,
			token:    "",
		},
	}

	client := graphql.NewClient(url, graphql.WithHTTPClient(&httpClient))

	return &RazeeClient{
		client: client,
	}
}

type IAMHTTPClient struct {
	*http.Client
}

type IAMRoundTripper struct {
	authenticator core.Authenticator
}

func (t *IAMRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	err := t.authenticator.Authenticate(request)
	if err != nil {
		return nil, err
	}
	return http.DefaultTransport.RoundTrip(request)
}

type RazeeGithubApiRoundTripper struct {
	apiKey string
}

func (t *RazeeGithubApiRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add("x-api-key", t.apiKey)
	return http.DefaultTransport.RoundTrip(request)
}

type RazeeLocalRoundTripper struct {
	url      string
	login    string
	password string
	token    string
}

func (t *RazeeLocalRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if t.token == "" {
		req := graphql.NewRequest(`
    mutation ($login: String!, $password: String!) {
	  signIn(
		login: $login
		password: $password
	  ) {
		token
	  }
	}
	`)

		req.Var("login", t.login)
		req.Var("password", t.password)
		var result struct {
			Signin struct {
				Token string `json:"token"`
			} `json:"signIn"`
		}

		client := graphql.NewClient(t.url)

		err := client.Run(context.Background(), req, &result)

		if err != nil {
			return nil, err
		}
		t.token = result.Signin.Token
	}

	request.Header.Add("Authorization", "Bearer "+t.token)
	return http.DefaultTransport.RoundTrip(request)
}

// Authentication method used by sat-con-client library
func (t *RazeeLocalRoundTripper) Authenticate(request *http.Request) error {
	if t.token == "" {
		req := graphql.NewRequest(`
		mutation ($login: String!, $password: String!) {
		  signIn(
			login: $login
			password: $password
		  ) {
			token
		  }
		}
	`)

		req.Var("login", t.login)
		req.Var("password", t.password)
		var result struct {
			Signin struct {
				Token string `json:"token"`
			} `json:"signIn"`
		}

		client := graphql.NewClient(t.url)

		err := client.Run(context.Background(), req, &result)

		if err != nil {
			return err
		}
		t.token = result.Signin.Token
	}

	request.Header.Add("Authorization", "Bearer "+t.token)
	return nil
}