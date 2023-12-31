package unit

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/coding-challenge/api-searching/config"
	"github.com/coding-challenge/api-searching/helpers/api"
	"github.com/coding-challenge/api-searching/helpers/util"
	"github.com/coding-challenge/api-searching/models"
	request "github.com/coding-challenge/api-searching/request/user"
	"github.com/coding-challenge/api-searching/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

// Init system
func init() {
	config.Init("development", "/go/src/github.com/coding-challenge/api-searching/config")
}

// Mock function
func (m *HttpClientMock) SendGet(params api.Params) (string, int) {
	agrs := m.Called(params)
	return agrs.String(0), agrs.Int(1)
}

// Prepare params
func (m *HttpClientMock) initParams(id int) (request.DetailRequest, api.Params, api.Params) {
	cfg := config.GetConfig()
	req := request.DetailRequest{
		ID: id,
	}
	apiUserDetailParams := api.Params{
		URL:     fmt.Sprintf(cfg.GetString("api_test.user_detail"), req.ID),
		Timeout: 5,
		Header:  nil,
	}
	apiAccountListParams := api.Params{
		URL:     fmt.Sprintf(cfg.GetString("api_test.account_list"), req.ID),
		Timeout: 5,
		Header:  nil,
	}
	return req, apiUserDetailParams, apiAccountListParams
}

func Test_Case_UserService_1_Success(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/success_data_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	userDataMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(1)
	httpClientM.On("SendGet", apiUserDetailParams).Return(userDataMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Parse data expected
	var dataExpected *models.User
	var accountData []*models.Account
	util.ParseJSON([]byte(userDataMock), &dataExpected, "User")
	util.ParseJSON([]byte(dataAccountMock), &accountData, "Account")
	if dataExpected != nil && accountData != nil {
		dataExpected.Accounts = accountData
	}
	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Assert test result
	assert.Equal(t, dataExpected.Name, dataActual.Name)
	for i, row := range dataExpected.Accounts {
		assert.Equal(t, row.ID, dataActual.Accounts[i].ID)
		assert.Equal(t, row.Name, dataActual.Accounts[i].Name)
		assert.Equal(t, row.Balance, dataActual.Accounts[i].Balance)
	}
	assert.Equal(t, http.StatusOK, sttCodeActual)
}

func Test_Case_UserService_2_Failed_Api_GetUser(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/not_found_data.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	userDataMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(2)
	httpClientM.On("SendGet", apiUserDetailParams).Return(userDataMock, http.StatusNotFound)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Data expected
	var dataExpected *models.User

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusNotFound, sttCodeActual)
}

func Test_Case_UserService_3_Success_Api_GetAccountError(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/success_data_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/not_found_data.json")
	userDataMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(3)
	httpClientM.On("SendGet", apiUserDetailParams).Return(userDataMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusNotFound)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Parse data expected
	var dataExpected *models.User
	util.ParseJSON([]byte(userDataMock), &dataExpected, "User")

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusOK, sttCodeActual)
}

func Test_Case_UserService_4_Failed_Api_GetUser500(t *testing.T) {
	// Mock data
	pathF, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	dataAccountMock := string(util.ReadFile(pathF))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(4)
	httpClientM.On("SendGet", apiUserDetailParams).Return("", http.StatusInternalServerError)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Data expected
	var dataExpected *models.User

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusInternalServerError, sttCodeActual)
}

func Test_Case_UserService_5_Failed_Api_ErrorStructDataUser(t *testing.T) {
	// Mock data
	pathF2, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	dataUserMock := `[]`
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(5)
	httpClientM.On("SendGet", apiUserDetailParams).Return(dataUserMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Data expected
	var dataExpected *models.User

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusNotFound, sttCodeActual)
}

func Test_Case_UserService_6_Success_Api_ErrorStructDataAccount(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/success_data_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/wrong_structure_data.json")
	dataUserMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(6)
	httpClientM.On("SendGet", apiUserDetailParams).Return(dataUserMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Parse data expected
	var dataExpected *models.User
	util.ParseJSON([]byte(dataUserMock), &dataExpected, "User")

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusOK, sttCodeActual)
}

func Test_Case_UserService_7_Failed_Api_ErrorDataTypeUser(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/wrong_data_type_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	dataUserMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(7)
	httpClientM.On("SendGet", apiUserDetailParams).Return(dataUserMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Data expected
	var dataExpected *models.User

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusNotFound, sttCodeActual)
}

func Test_Case_UserService_8_Success_Api_ErrorDataTypeAccount(t *testing.T) {
	// Mock data
	pathF1, _ := filepath.Abs("../mock_data/user_service/success_data_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/wrong_data_type_account.json")
	dataUserMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))
	httpClientM := new(HttpClientMock)
	req, apiUserDetailParams, apiAccountListParams := httpClientM.initParams(8)
	httpClientM.On("SendGet", apiUserDetailParams).Return(dataUserMock, http.StatusOK)
	httpClientM.On("SendGet", apiAccountListParams).Return(dataAccountMock, http.StatusOK)

	// Execute test service
	userSrvTest := services.NewUserService(context.TODO(), httpClientM)
	dataActual, sttCodeActual := userSrvTest.HandleDetail(req)

	// Data expected
	var dataExpected *models.User
	util.ParseJSON([]byte(dataUserMock), &dataExpected, "User")

	// Assert test result
	assert.Equal(t, dataExpected, dataActual)
	assert.Equal(t, http.StatusOK, sttCodeActual)
}
