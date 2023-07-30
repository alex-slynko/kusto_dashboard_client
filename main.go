package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Dashboard struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func main() {
	url := "https://dashboards.kusto.windows.net/dashboards"
	httpClient := http.DefaultClient
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	token, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{"https://rtd-metadata.azurewebsites.net/"}})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)
	// handle response
	// read response body
	// result := Dashboard{}
	// if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
	// 	panic(err)
	// }
	// print body
	resultB, err := ioutil.ReadAll(resp.Body)
	result := string(resultB)
	fmt.Printf("%#v", result)
}
