package googlebusinessprofile

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Name              string `json:"name"`
	AccountName       string `json:"accountName"`
	Type              string `json:"type"`
	VerificationState string `json:"verificationState"`
	VettedState       string `json:"vettedState"`
}

func (service *Service) Accounts() (*[]Account, *errortools.Error) {
	accountsReponse := AccountsResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           "https://mybusinessaccountmanagement.googleapis.com/v1/accounts",
		ResponseModel: &accountsReponse,
	}

	_, _, e := service.googleService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &accountsReponse.Accounts, nil
}

func (service *Service) Account(accountName string) (*Account, *errortools.Error) {
	var account Account

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           fmt.Sprintf("https://mybusinessaccountmanagement.googleapis.com/v1/%s", accountName),
		ResponseModel: &account,
	}

	_, _, e := service.googleService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &account, nil
}
