package kusto_dashboard_client

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const (
	// DefaultBaseURL is the default base URL for the Kusto Dashboard service.
	defaultBaseURL = "https://dashboards.kusto.windows.net/dashboards"
)

type AzureCredentials interface {
	GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error)
}

type Client struct {
	cred   AzureCredentials
	client *http.Client
}

func NewDashboardClient(cred AzureCredentials) *Client {
	return &Client{cred: cred, client: http.DefaultClient}
}

func (c *Client) GetDashboards(ctx context.Context) ([]Dashboard, error) {
	dashboards := []Dashboard{}
	req, err := http.NewRequest("GET", defaultBaseURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := runtime.UnmarshalAsJSON(resp, &dashboards); err != nil {
		return nil, err
	}
	return dashboards, nil
}

// func (c *Client) GetDashboard(ctx context.Context, id string) (*DashboardDetail, error) {
// return getDashboard(c.cred, id)
// }

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	token, err := c.cred.GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{"https://rtd-metadata.azurewebsites.net/"}})
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token.Token)
	return c.client.Do(req)
}
