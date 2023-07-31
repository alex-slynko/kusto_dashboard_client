package main

import (
	"context"
	"fmt"
	"os"

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
	// read first argument as dashboard id
	if len(os.Args) < 2 {
		result, err := client.GetDashboards(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Println("Dashboards:")
		fmt.Printf("%#v", result)
		os.Exit(0)
	}

	id := os.Args[1]
	// if there is --raw flag, print raw json
	if len(os.Args) > 2 && os.Args[2] == "--raw" {
		dashboard, err := client.GetDashboardRaw(context.Background(), id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Dashboard: " + id)
		fmt.Println(dashboard)
		os.Exit(0)
	}

	dashboard, err := client.GetDashboard(context.Background(), id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Dashboard: " + id)
	fmt.Printf("%#v", dashboard)

}
