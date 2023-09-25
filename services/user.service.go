package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/coding-challenge/api-searching/config"
	"github.com/coding-challenge/api-searching/helpers/api"
	"github.com/coding-challenge/api-searching/helpers/util"
	"github.com/coding-challenge/api-searching/models"
	request "github.com/coding-challenge/api-searching/request/user"
)

type UserService struct {
	Api api.IHttpClient
}

func NewUserService(client api.IHttpClient) *UserService {
	return &UserService{
		Api: client,
	}
}

func (userSrv *UserService) HandleDetail(req request.DetailRequest) (*models.User, int) {
	// Load config
	cfg := config.GetConfig()
	apiUserDetail := fmt.Sprintf(cfg.GetString("api_test.user_detail"), req.ID)
	apiAccountList := fmt.Sprintf(cfg.GetString("api_test.account_list"), req.ID)
	timeout := 5 // second

	// Prepare api
	var userReps, accReps string
	var userSttCode, accSttCode int
	getDetailUser := func(w *sync.WaitGroup) {
		userReps, userSttCode = userSrv.Api.SendGet(api.Params{
			URL:     apiUserDetail,
			Timeout: timeout,
			Header:  nil,
		})
		w.Done()
	}
	getAccounts := func(w *sync.WaitGroup) {
		accReps, accSttCode = userSrv.Api.SendGet(api.Params{
			URL:     apiAccountList,
			Timeout: timeout,
			Header:  nil,
		})
		w.Done()
	}
	// Call api
	var wg sync.WaitGroup
	wg.Add(2)
	go getDetailUser(&wg)
	go getAccounts(&wg)
	wg.Wait()

	if userSttCode != http.StatusOK {
		return nil, userSttCode
	}
	if accSttCode != http.StatusOK {
		return nil, accSttCode
	}

	// Parse data
	var user *models.User
	var accounts []*models.Account
	util.ParseJSON([]byte(userReps), &user, "User")
	util.ParseJSON([]byte(accReps), &accounts, "Account")

	if user == nil || accounts == nil || user.IsEmpty() || len(accounts) == 0 {
		return nil, http.StatusNotFound
	}

	user.Accounts = accounts

	return user, http.StatusOK
}
