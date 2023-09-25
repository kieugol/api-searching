package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/coding-challenge/api-searching/config"
	"github.com/coding-challenge/api-searching/controllers"
	"github.com/coding-challenge/api-searching/helpers/respond"
	"github.com/coding-challenge/api-searching/helpers/util"
	"github.com/coding-challenge/api-searching/models"
	request "github.com/coding-challenge/api-searching/request/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

// Init system
func init() {
	gin.SetMode(gin.TestMode)
	config.Init("development", "/go/src/github.com/coding-challenge/api-searching/config")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Mock function
func (m *UserServiceMock) HandleDetail(req request.DetailRequest) (*models.User, int) {
	agrs := m.Called(req)

	var data *models.User
	byteD, _ := json.Marshal(agrs.Get(0))
	util.ParseJSON(byteD, &data, "User")

	return data, agrs.Int(1)
}

// Prepare params
func (m *UserServiceMock) initParams(id int) request.DetailRequest {
	req := request.DetailRequest{
		ID: id,
	}
	return req
}

func Test_Case_UserController_1_Success(t *testing.T) {
	// Mock data
	userSrvMock := new(UserServiceMock)
	req := userSrvMock.initParams(1)
	pathF1, _ := filepath.Abs("../mock_data/user_service/success_data_user.json")
	pathF2, _ := filepath.Abs("../mock_data/user_service/success_data_account.json")
	userDataMock := string(util.ReadFile(pathF1))
	dataAccountMock := string(util.ReadFile(pathF2))

	var userData *models.User
	var accountData []*models.Account
	util.ParseJSON([]byte(userDataMock), &userData, "User")
	util.ParseJSON([]byte(dataAccountMock), &accountData, "Account")
	if userData != nil && accountData != nil {
		userData.Accounts = accountData
	}
	userSrvMock.On("HandleDetail", req).Return(userData, http.StatusOK)

	dataExpected, _ := json.Marshal(respond.Success(userData, "Success"))

	// Execute test Controller
	userCtrlTest := controllers.NewUserController(userSrvMock)
	r := SetUpRouter()
	r.GET("/v1/users/:id", userCtrlTest.Detail)
	reqTest, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d", req.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqTest)

	// Assert test result
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(dataExpected), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_Case_UserController_2_Failed_NotFoundData(t *testing.T) {
	// Mock data
	userSrvMock := new(UserServiceMock)
	req := userSrvMock.initParams(2)
	userSrvMock.On("HandleDetail", req).Return(nil, http.StatusNotFound)
	dataExpected, _ := json.Marshal(respond.NotFound())

	// Execute test Controller
	userCtrlTest := controllers.NewUserController(userSrvMock)
	r := SetUpRouter()
	r.GET("/v1/users/:id", userCtrlTest.Detail)
	reqTest, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d", req.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqTest)

	// Assert test result
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(dataExpected), string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_Case_UserController_3_Failed_InternalServerError(t *testing.T) {
	// Mock data
	userSrvMock := new(UserServiceMock)
	req := userSrvMock.initParams(3)
	userSrvMock.On("HandleDetail", req).Return(nil, http.StatusInternalServerError)
	dataExpected, _ := json.Marshal(respond.InternalServerError())

	// Execute test Controller
	userCtrlTest := controllers.NewUserController(userSrvMock)
	r := SetUpRouter()
	r.GET("/v1/users/:id", userCtrlTest.Detail)
	reqTest, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d", req.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqTest)

	// Assert test result
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(dataExpected), string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_Case_UserController_3_Failed_BadRequest(t *testing.T) {
	// Mock data
	userSrvMock := new(UserServiceMock)
	req := userSrvMock.initParams(3)
	userSrvMock.On("HandleDetail", req).Return(nil, http.StatusNotFound)
	dataExpected, _ := json.Marshal(respond.MissingParams())

	// Execute test Controller
	userCtrlTest := controllers.NewUserController(userSrvMock)
	r := SetUpRouter()
	r.GET("/v1/users/:id", userCtrlTest.Detail)
	reqTest, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%s", "krol"), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqTest)

	// Assert test result
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(dataExpected), string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
