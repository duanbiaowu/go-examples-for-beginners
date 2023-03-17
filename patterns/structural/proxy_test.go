package structural

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Proxy(t *testing.T) {
	nginx := NewNginxServer(1)

	getUrl := "/app/list"
	createUrl := "/app/create"

	code, body := nginx.handleRequest(getUrl, http.MethodGet)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, http.StatusText(http.StatusOK), body)

	code, body = nginx.handleRequest(createUrl, http.MethodPost)
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, http.StatusText(http.StatusCreated), body)

	code, body = nginx.handleRequest(getUrl, http.MethodGet)
	assert.Equal(t, http.StatusForbidden, code)
	assert.Equal(t, http.StatusText(http.StatusForbidden), body)

	code, body = nginx.handleRequest(createUrl, http.MethodPost)
	assert.Equal(t, http.StatusForbidden, code)
	assert.Equal(t, http.StatusText(http.StatusForbidden), body)
}
