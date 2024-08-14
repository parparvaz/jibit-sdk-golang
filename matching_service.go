package jibit

import (
	"context"
	"net/http"
)

type MatchCardNumberWithNationalCodeService struct {
	c            *Client
	cardNumber   string
	nationalCode string
	birthDate    string
}

func (j *MatchCardNumberWithNationalCodeService) CardNumber(cardNumber string) *MatchCardNumberWithNationalCodeService {
	j.cardNumber = cardNumber

	return j
}

func (j *MatchCardNumberWithNationalCodeService) NationalCode(nationalCode string) *MatchCardNumberWithNationalCodeService {
	j.nationalCode = nationalCode

	return j
}

func (j *MatchCardNumberWithNationalCodeService) BirthDate(birthDate string) *MatchCardNumberWithNationalCodeService {
	j.birthDate = birthDate

	return j
}

func (j *MatchCardNumberWithNationalCodeService) Do(ctx context.Context, opts ...RequestOption) (res *MatchCardNumberWithNationalCode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("cardNumber", j.cardNumber)
	r.setParam("nationalCode", j.nationalCode)
	r.setParam("birthDate", j.birthDate)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchCardNumberWithNationalCode)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchCardNumberWithNationalCode struct {
	Matched bool `json:"matched"`
}
type MatchAccountNumberWithNationalCodeService struct {
	c             *Client
	bank          string
	depositNumber string
	nationalCode  string
	birthDate     string
}

func (j *MatchAccountNumberWithNationalCodeService) DepositNumber(depositNumber string) *MatchAccountNumberWithNationalCodeService {
	j.depositNumber = depositNumber

	return j
}

func (j *MatchAccountNumberWithNationalCodeService) Bank(bank string) *MatchAccountNumberWithNationalCodeService {
	j.bank = bank

	return j
}

func (j *MatchAccountNumberWithNationalCodeService) NationalCode(nationalCode string) *MatchAccountNumberWithNationalCodeService {
	j.nationalCode = nationalCode

	return j
}

func (j *MatchAccountNumberWithNationalCodeService) BirthDate(birthDate string) *MatchAccountNumberWithNationalCodeService {
	j.birthDate = birthDate

	return j
}

func (j *MatchAccountNumberWithNationalCodeService) Do(ctx context.Context, opts ...RequestOption) (res *MatchAccountNumberWithNationalCode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("bank", j.bank)
	r.setParam("depositNumber", j.depositNumber)
	r.setParam("nationalCode", j.nationalCode)
	r.setParam("birthDate", j.birthDate)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchAccountNumberWithNationalCode)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchAccountNumberWithNationalCode struct{}

type MatchIbanWithNationalCodeService struct {
	c            *Client
	iban         string
	nationalCode string
	birthDate    string
}

func (j *MatchIbanWithNationalCodeService) Iban(iban string) *MatchIbanWithNationalCodeService {
	j.iban = iban

	return j
}

func (j *MatchIbanWithNationalCodeService) NationalCode(nationalCode string) *MatchIbanWithNationalCodeService {
	j.nationalCode = nationalCode

	return j
}

func (j *MatchIbanWithNationalCodeService) BirthDate(birthDate string) *MatchIbanWithNationalCodeService {
	j.birthDate = birthDate

	return j
}

func (j *MatchIbanWithNationalCodeService) Do(ctx context.Context, opts ...RequestOption) (res *MatchIbanWithNationalCode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("iban", j.iban)
	r.setParam("nationalCode", j.nationalCode)
	r.setParam("birthDate", j.birthDate)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchIbanWithNationalCode)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchIbanWithNationalCode struct{}

type MatchCardNumberWithNameService struct {
	c          *Client
	cardNumber string
	name       string
}

func (j *MatchCardNumberWithNameService) CardNumber(cardNumber string) *MatchCardNumberWithNameService {
	j.cardNumber = cardNumber

	return j
}

func (j *MatchCardNumberWithNameService) Name(name string) *MatchCardNumberWithNameService {
	j.name = name

	return j
}

func (j *MatchCardNumberWithNameService) Do(ctx context.Context, opts ...RequestOption) (res *MatchCardNumberWithName, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("cardNumber", j.cardNumber)
	r.setParam("name", j.name)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchCardNumberWithName)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchCardNumberWithName struct{}

type MatchAccountNumberWithNameService struct {
	c             *Client
	bank          string
	depositNumber string
	name          string
}

func (j *MatchAccountNumberWithNameService) DepositNumber(depositNumber string) *MatchAccountNumberWithNameService {
	j.depositNumber = depositNumber

	return j
}

func (j *MatchAccountNumberWithNameService) Bank(bank string) *MatchAccountNumberWithNameService {
	j.bank = bank

	return j
}

func (j *MatchAccountNumberWithNameService) Name(name string) *MatchAccountNumberWithNameService {
	j.name = name

	return j
}

func (j *MatchAccountNumberWithNameService) Do(ctx context.Context, opts ...RequestOption) (res *MatchAccountNumberWithName, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("bank", j.bank)
	r.setParam("depositNumber", j.depositNumber)
	r.setParam("name", j.name)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchAccountNumberWithName)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchAccountNumberWithName struct{}

type MatchIbanWithNameService struct {
	c    *Client
	iban string
	name string
}

func (j *MatchIbanWithNameService) Iban(iban string) *MatchIbanWithNameService {
	j.iban = iban

	return j
}

func (j *MatchIbanWithNameService) Name(name string) *MatchIbanWithNameService {
	j.name = name

	return j
}

func (j *MatchIbanWithNameService) Do(ctx context.Context, opts ...RequestOption) (res *MatchIbanWithName, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("iban", j.iban)
	r.setParam("name", j.name)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchIbanWithName)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchIbanWithName struct{}

type MatchNationalCodeWithMobileNumberService struct {
	c            *Client
	mobileNumber string
	nationalCode string
}

func (j *MatchNationalCodeWithMobileNumberService) MobileNumber(mobileNumber string) *MatchNationalCodeWithMobileNumberService {
	j.mobileNumber = mobileNumber

	return j
}

func (j *MatchNationalCodeWithMobileNumberService) NationalCode(nationalCode string) *MatchNationalCodeWithMobileNumberService {
	j.nationalCode = nationalCode

	return j
}

func (j *MatchNationalCodeWithMobileNumberService) Do(ctx context.Context, opts ...RequestOption) (res *MatchNationalCodeWithMobileNumber, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("mobileNumber", j.mobileNumber)
	r.setParam("nationalCode", j.nationalCode)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchNationalCodeWithMobileNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchNationalCodeWithMobileNumber struct {
	Matched bool `json:"matched"`
}

type MatchUniversalIDNationalsWithPassportNumberService struct {
	c              *Client
	fida           string
	passportNumber string
}

func (j *MatchUniversalIDNationalsWithPassportNumberService) Fida(fida string) *MatchUniversalIDNationalsWithPassportNumberService {
	j.fida = fida

	return j
}

func (j *MatchUniversalIDNationalsWithPassportNumberService) PassportNumber(passportNumber string) *MatchUniversalIDNationalsWithPassportNumberService {
	j.passportNumber = passportNumber

	return j
}

func (j *MatchUniversalIDNationalsWithPassportNumberService) Do(ctx context.Context, opts ...RequestOption) (res *MatchUniversalIDNationalsWithPassportNumber, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/matching",
		secType:  secTypeAccessToken,
	}

	r.setParam("fida", j.fida)
	r.setParam("passportNumber", j.passportNumber)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchUniversalIDNationalsWithPassportNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchUniversalIDNationalsWithPassportNumber struct{}
