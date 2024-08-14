package jibit

import (
	"context"
	"net/http"
)

type CardToAccountNumberService struct {
	c      *Client
	number string
}

func (j *CardToAccountNumberService) Number(number string) *CardToAccountNumberService {
	j.number = number

	return j
}

func (j *CardToAccountNumberService) Do(ctx context.Context, opts ...RequestOption) (res *CardToAccountNumber, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/cards",
		secType:  secTypeAccessToken,
	}

	r.setParam("number", j.number)
	r.setParam("deposit", true)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CardToAccountNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CardToAccountNumber struct {
	Number      string `json:"number"`
	Type        string `json:"type"`
	DepositInfo struct {
		Bank          string `json:"bank"`
		DepositNumber string `json:"depositNumber"`
	} `json:"depositInfo"`
}

type CardToIbanService struct {
	c      *Client
	number string
}

func (j *CardToIbanService) Number(number string) *CardToIbanService {
	j.number = number

	return j
}

func (j *CardToIbanService) Do(ctx context.Context, opts ...RequestOption) (res *CardToIban, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/cards",
		secType:  secTypeAccessToken,
	}

	r.setParam("number", j.number)
	r.setParam("iban", true)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(CardToIban)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CardToIban struct {
	Number   string `json:"number"`
	Type     string `json:"type"`
	IbanInfo struct {
		Bank          string `json:"bank"`
		DepositNumber string `json:"depositNumber"`
		Iban          string `json:"iban"`
		Status        string `json:"status"`
		Owners        []struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		} `json:"owners"`
	} `json:"ibanInfo"`
}

type AccountNumberToIbanService struct {
	c      *Client
	bank   string
	number string
}

func (j *AccountNumberToIbanService) Bank(bank string) *AccountNumberToIbanService {
	j.bank = bank

	return j
}

func (j *AccountNumberToIbanService) Number(number string) *AccountNumberToIbanService {
	j.number = number

	return j
}

func (j *AccountNumberToIbanService) Do(ctx context.Context, opts ...RequestOption) (res *AccountNumberToIban, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/deposits",
		secType:  secTypeAccessToken,
	}

	r.setParam("bank", j.bank)
	r.setParam("number", j.number)
	r.setParam("iban", true)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(AccountNumberToIban)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type AccountNumberToIban struct{}
