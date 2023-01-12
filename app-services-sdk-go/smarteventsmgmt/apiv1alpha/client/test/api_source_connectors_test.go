/*
Red Hat Openshift SmartEvents Fleet Manager V2

Testing SourceConnectorsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package smarteventsmgmtclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_smarteventsmgmtclient_SourceConnectorsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SourceConnectorsApiService SourceConnectorsAPICreateSourceConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bridgeId string

        resp, httpRes, err := apiClient.SourceConnectorsApi.SourceConnectorsAPICreateSourceConnector(context.Background(), bridgeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SourceConnectorsApiService SourceConnectorsAPIDeleteSourceConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bridgeId string
        var sourceId string

        resp, httpRes, err := apiClient.SourceConnectorsApi.SourceConnectorsAPIDeleteSourceConnector(context.Background(), bridgeId, sourceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SourceConnectorsApiService SourceConnectorsAPIGetSourceConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bridgeId string
        var sourceId string

        resp, httpRes, err := apiClient.SourceConnectorsApi.SourceConnectorsAPIGetSourceConnector(context.Background(), bridgeId, sourceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SourceConnectorsApiService SourceConnectorsAPIGetSourceConnectors", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bridgeId string

        resp, httpRes, err := apiClient.SourceConnectorsApi.SourceConnectorsAPIGetSourceConnectors(context.Background(), bridgeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SourceConnectorsApiService SourceConnectorsAPIUpdateSourceConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bridgeId string
        var sourceId string

        resp, httpRes, err := apiClient.SourceConnectorsApi.SourceConnectorsAPIUpdateSourceConnector(context.Background(), bridgeId, sourceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
