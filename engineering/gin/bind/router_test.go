package bind

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func Test_setupRouter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func Test_setupFormData(t *testing.T) {
	router := setupFormData()

	data := url.Values{
		"id":    {"1024"},
		"name":  {"Terry"},
		"attrs": {`{height: 200, weight: 100, "hobby": ["coding", "reading"]}`},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `id = 1024, name = Terry, attrs = {height: 200, weight: 100, "hobby": ["coding", "reading"]}`, w.Body.String())
}

func Test_setupBindJson(t *testing.T) {
	router := setupBindJson()

	w := httptest.NewRecorder()
	data := `{"user": "admin", "password": "123456"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"you are logged in"}`, w.Body.String())
}

func Test_setupBindXml(t *testing.T) {
	router := setupBindXml()

	w := httptest.NewRecorder()
	data := `
<?xml version="1.0" encoding="UTF-8"?>
<root>
	<user>admin</user>
	<password>123456</password>
</root>
`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"you are logged in"}`, w.Body.String())
}

func Test_setupBindForm(t *testing.T) {
	router := setupBindForm()

	//data := url.Values{
	//	"user": {"admin"},
	//	"password": {"123456"},
	//}
	//data.Encode() = user=admin&password=123456

	w := httptest.NewRecorder()
	data := "user=admin&password=123456"
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"you are logged in"}`, w.Body.String())
}

func MockRequestFile(filePath, fieldName string) (multipart.File, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}
	_ = writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	formFile, _, err := req.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	return formFile, nil
}
