/*
Kafka Management API

Testing EnterpriseDataplaneClustersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kafkamgmtclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_kafkamgmtclient_EnterpriseDataplaneClustersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnterpriseDataplaneClustersApiService DeleteEnterpriseClusterById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnterpriseDataplaneClustersApi.DeleteEnterpriseClusterById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnterpriseDataplaneClustersApiService GetEnterpriseClusterById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnterpriseDataplaneClustersApi.GetEnterpriseClusterById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnterpriseDataplaneClustersApiService GetEnterpriseClusterWithAddonParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnterpriseDataplaneClustersApi.GetEnterpriseClusterWithAddonParameters(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnterpriseDataplaneClustersApiService GetEnterpriseOsdClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EnterpriseDataplaneClustersApi.GetEnterpriseOsdClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnterpriseDataplaneClustersApiService RegisterEnterpriseOsdCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
