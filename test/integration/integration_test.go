package integration

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	server "github.com/asxcandrew/secret-server"
	mock "github.com/asxcandrew/secret-server/test/mock"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateSecret_Positive(t *testing.T) {
	logger := log.NewNopLogger()
	a := assert.New(t)
	s := server.NewSecretService(mock.NewStorageMock())
	routes := mux.NewRouter()

	routes.PathPrefix("/secret").Handler(server.MakeSecretHandler(s, logger))

	data := url.Values{}
	data.Set("secret", "foo")
	data.Set("expireAfterViews", "2")
	data.Set("expireAfter", "20")

	req, _ := http.NewRequest("POST", "/secret", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	rr := httptest.NewRecorder()

	routes.ServeHTTP(rr, req)

	a.NotEqual(http.StatusOK, rr.Code)
	a.Equal(http.StatusOK, rr.Code)
}
