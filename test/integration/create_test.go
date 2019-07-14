package integration

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/asxcandrew/secret-server/test/mock"

	server "github.com/asxcandrew/secret-server"
	"github.com/asxcandrew/secret-server/storage"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateSecret_Positive(t *testing.T) {
	logger := log.NewNopLogger()
	a := assert.New(t)

	srMock := &mock.SecretRepositoryMock{}
	srMock.On("Create").Return(nil)

	s := server.NewSecretService(storage.Storage{
		Secret: srMock,
	})
	routes := mux.NewRouter()

	routes.PathPrefix("/secret").Handler(server.MakeSecretHandler(s, logger))

	data := url.Values{}
	data.Set("secret", "foo")
	data.Set("expireAfterViews", "2")
	data.Set("expireAfter", "20")

	req, _ := http.NewRequest("POST", "/secret", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Accept", "application/json")

	rr := httptest.NewRecorder()

	routes.ServeHTTP(rr, req)

	a.NotEqual(http.StatusOK, rr.Code)
	a.Equal(http.StatusOK, rr.Code)
}
