package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/asxcandrew/secret-server/test/mock"

	server "github.com/asxcandrew/secret-server"
	"github.com/asxcandrew/secret-server/storage"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetSecret_Positive(t *testing.T) {
	logger := log.NewNopLogger()
	a := assert.New(t)

	srMock := &mock.SecretRepositoryMock{}
	srMock.On("Get").Return(nil)

	s := server.NewSecretService(storage.Storage{
		Secret: srMock,
	})
	routes := mux.NewRouter()

	routes.PathPrefix("/secret").Handler(server.MakeSecretHandler(s, logger))

	req, _ := http.NewRequest("GET", "/secret/6ef10f70-c75f-4cc5-b8db-b111ea2dec56", strings.NewReader(""))

	req.Header.Add("Accept", "application/json")

	rr := httptest.NewRecorder()

	routes.ServeHTTP(rr, req)

	a.NotEqual(http.StatusOK, rr.Code)
	a.Equal(http.StatusOK, rr.Code)
}
