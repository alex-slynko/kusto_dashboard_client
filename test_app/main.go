package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alex-slynko/kusto_dashboard_client"
)

func main() {
	// Create an Azure Identity using default credentials
	// https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-defaultazurecredential
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	client := kusto_dashboard_client.NewDashboardClient(credential)

	result, err := client.GetDashboards(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", result)
}
