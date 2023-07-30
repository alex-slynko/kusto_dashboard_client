module github.com/alex-slynko/kusto-dashboard-api

go 1.20

require github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.3.0

require github.com/Azure/azure-sdk-for-go/sdk/azcore v1.8.0-beta.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
)

// load azure-kusto-go from other directory
replace github.com/Azure/azure-sdk-for-go/sdk/azidentity => ../../../other/azure-sdk-for-go/sdk/azidentity
