package unit

import (
	"log"
	"net/http"
	"testing"

	request "github.com/coding-challenge/api-searching/request/user"
	"github.com/coding-challenge/api-searching/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// smsServiceMock để giả định cho một MessageService
type httpClientMock struct {
	mock.Mock
}

func (m *httpClientMock) SendGet(url string, timeout int, header map[string]string) ([]byte, int) {
	return nil, http.StatusOK
}

func TestCase1(t *testing.T) {
	httpClient := new(httpClientMock)

	dataUser := []byte(`{"id":"someID","data":["str1","str2", 1337, {"my": "obj", "id": 42}]}`)
	dataAccount := []byte(`{"id":"someID","data":["str1","str2", 1337, {"my": "obj", "id": 42}]}`)
	httpClient.On("SendGet", "https://mfx-recruit-dev.herokuapp.com/users/1").
		Return(dataUser, http.StatusOK)
	httpClient.On("SendGet", "https://mfx-recruit-dev.herokuapp.com/users/1/accounts").
		Return(dataAccount, http.StatusOK)

	userSvTest := &services.UserService{
		Api: httpClient,
	}
	req := request.DetailRequest{
		ID: 1,
	}
	data, code := userSvTest.HandleDetail(req)
	log.Println("data:", data)
	log.Println("stt_code:", code)

	assert.Equal(t, 200, code)

	httpClient.AssertExpectations(t)
}
