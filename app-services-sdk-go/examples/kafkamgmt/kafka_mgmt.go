package main

import (
	"context"
	"os"

	rhoasAuth "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/auth/apiv1"
	kafkamgmt "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1"
)

func main() {
	ctx := context.Background()
	offlineToken := os.Getenv("OFFLINE_TOKEN")
	httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

	client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		HTTPClient: httpClient,
	})

	_, _, err := client.DefaultApi.GetKafkas(ctx).Execute()
	if err != nil {
		panic(err)
	}
}
