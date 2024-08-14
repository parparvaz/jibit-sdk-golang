package jibit

import (
	"context"
	"net/http"
)

type CorporationInquiryService struct {
	c    *Client
	code string
}

func (j *CorporationInquiryService) Code(code string) *CorporationInquiryService {
	j.code = code

	return j
}

func (j *CorporationInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *CorporationInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/corporation/identity",
		secType:  secTypeAccessToken,
	}

	r.setParam("code", j.code)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(CorporationInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CorporationInquiry struct{}

type MilitaryServiceInquiryService struct {
	c            *Client
	nationalCode string
}

func (j *MilitaryServiceInquiryService) NationalCode(nationalCode string) *MilitaryServiceInquiryService {
	j.nationalCode = nationalCode

	return j
}

func (j *MilitaryServiceInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *MilitaryServiceInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/social/msq",
		secType:  secTypeAccessToken,
	}

	r.setParam("nationalCode", j.nationalCode)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MilitaryServiceInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MilitaryServiceInquiry struct{}

type CardInquiryService struct {
	c      *Client
	number string
}

func (j *CardInquiryService) Number(number string) *CardInquiryService {
	j.number = number

	return j
}

func (j *CardInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *CardInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/cards",
		secType:  secTypeAccessToken,
	}

	r.setParam("number", j.number)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CardInquiry{}, err
	}
	res = new(CardInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &CardInquiry{}, err
	}
	return res, nil
}

type CardInquiry struct{}

type IbanInquiryService struct {
	c     *Client
	value string
}

func (j *IbanInquiryService) Value(value string) *IbanInquiryService {
	j.value = value

	return j
}

func (j *IbanInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *IbanInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/ibans",
		secType:  secTypeAccessToken,
	}

	r.setParam("value", j.value)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(IbanInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type IbanInquiry struct{}

type LegalInquiryService struct {
	c            *Client
	nationalCode string
}

func (j *LegalInquiryService) NationalCode(nationalCode string) *LegalInquiryService {
	j.nationalCode = nationalCode

	return j
}

func (j *LegalInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *LegalInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/identity/legal",
		secType:  secTypeAccessToken,
	}

	r.setParam("nationalCode", j.nationalCode)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(LegalInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type LegalInquiry struct{}

type PostalInquiryService struct {
	c    *Client
	code string
}

func (j *PostalInquiryService) Code(code string) *PostalInquiryService {
	j.code = code

	return j
}

func (j *PostalInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *PostalInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/postal",
		secType:  secTypeAccessToken,
	}

	r.setParam("code", j.code)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(PostalInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type PostalInquiry struct{}

type PostalWithWorldGeodeticSystemInquiryService struct {
	c    *Client
	code string
}

func (j *PostalWithWorldGeodeticSystemInquiryService) Code(code string) *PostalWithWorldGeodeticSystemInquiryService {
	j.code = code

	return j
}

func (j *PostalWithWorldGeodeticSystemInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *PostalWithWorldGeodeticSystemInquiry, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/postal/wgs",
		secType:  secTypeAccessToken,
	}

	r.setParam("code", j.code)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(PostalWithWorldGeodeticSystemInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type PostalWithWorldGeodeticSystemInquiry struct{}
