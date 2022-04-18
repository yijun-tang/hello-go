package tdddemo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	t.Parallel()

	mockApiClient, mockApiServer := mockClientAndServer()
	defer mockApiServer.Close()

	t.Run("return 200 given third party API return 200", func(t *testing.T) {
		response, statusCode, _ := mockApiClient.GetMessage("api/v1/message/1")

		assert.EqualValues(t, http.StatusOK, statusCode)
		assert.EqualValues(t, "Hello world!", response)
	})

	t.Run("return 404 given third party API return 404", func(t *testing.T) {
		response, statusCode, _ := mockApiClient.GetMessage("api/v1/message/notExist")

		assert.EqualValues(t, http.StatusNotFound, statusCode)
		assert.EqualValues(t, `{"error":"record not found"}`, response)
	})
}

func mockClientAndServer() (*Client, *httptest.Server) {
	mockApiServer := mock_server.ThirdPartyAPIServerMock()
	mockClient := MockClient(mockApiServer)

	return mockClient, mockApiServer
}

// Use this function only for tests
func MockClient(testServer *httptest.Server) *Client {
	return &Client{
		apiBasePath: testServer.URL + "/",
		client:      testServer.Client(),
	}
}
