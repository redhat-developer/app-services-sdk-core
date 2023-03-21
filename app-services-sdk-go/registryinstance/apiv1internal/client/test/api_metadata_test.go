/*
Apicurio Registry API [v2]

Testing MetadataApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package registryinstanceclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_registryinstanceclient_MetadataApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetadataApiService DeleteArtifactVersionMetaData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string
		var version string

		httpRes, err := apiClient.MetadataApi.DeleteArtifactVersionMetaData(context.Background(), groupId, artifactId, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService GetArtifactMetaData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string

		resp, httpRes, err := apiClient.MetadataApi.GetArtifactMetaData(context.Background(), groupId, artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService GetArtifactOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string

		resp, httpRes, err := apiClient.MetadataApi.GetArtifactOwner(context.Background(), groupId, artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService GetArtifactVersionMetaData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string
		var version string

		resp, httpRes, err := apiClient.MetadataApi.GetArtifactVersionMetaData(context.Background(), groupId, artifactId, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService GetArtifactVersionMetaDataByContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string

		resp, httpRes, err := apiClient.MetadataApi.GetArtifactVersionMetaDataByContent(context.Background(), groupId, artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService UpdateArtifactMetaData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string

		httpRes, err := apiClient.MetadataApi.UpdateArtifactMetaData(context.Background(), groupId, artifactId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService UpdateArtifactOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string

		httpRes, err := apiClient.MetadataApi.UpdateArtifactOwner(context.Background(), groupId, artifactId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataApiService UpdateArtifactVersionMetaData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var artifactId string
		var version string

		httpRes, err := apiClient.MetadataApi.UpdateArtifactVersionMetaData(context.Background(), groupId, artifactId, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
