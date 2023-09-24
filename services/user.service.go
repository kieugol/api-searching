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
	Api api.HttClient
}

func NewUserService() {

}

func (userSv *UserService) HandleDetail(req request.DetailRequest) (*models.User, int) {
	// Load config
	cfg := config.GetConfig()
	apiUserDetail := fmt.Sprintf(cfg.GetString("api_test.user_detail"), req.ID)
	apiAccountList := fmt.Sprintf(cfg.GetString("api_test.account_list"), req.ID)
	timeout := 5 // second

	// Prepare api
	var userReps, accReps []byte
	var userSttCode, accSttCode int
	getDetailUser := func(w *sync.WaitGroup) {
		userReps, userSttCode = userSv.Api.Get(apiUserDetail, timeout, nil)
		w.Done()
	}
	getAccounts := func(w *sync.WaitGroup) {
		accReps, accSttCode = userSv.Api.Get(apiAccountList, timeout, nil)
		w.Done()
	}
	// Call api
	var wg sync.WaitGroup
	wg.Add(2)
	go getDetailUser(&wg)
	go getAccounts(&wg)
	wg.Wait()

	// Return invalid case
	if userSttCode != http.StatusOK {
		return nil, userSttCode
	}
	if accSttCode != http.StatusOK {
		return nil, accSttCode
	}

	// Parse data
	var user *models.User
	var accounts []*models.Account
	util.ParseJSON(userReps, &user, "User")
	util.ParseJSON(accReps, &accounts, "Account")
	if user.IsEmpty() || len(accounts) == 0 || accounts[0].IsEmpty() {
		return nil, http.StatusNotFound
	}

	user.Accounts = accounts

	return user, http.StatusOK
}
