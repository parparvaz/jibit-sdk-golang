package jibit

import (
	"context"
	"net/http"
)

type GenerateTokenService struct {
	c *Client
}

func (j *GenerateTokenService) Do(ctx context.Context, opts ...RequestOption) (res *GenerateToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v1/tokens/generate",
		secType:  secTypeNone,
	}
	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GenerateToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GenerateToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenService struct {
	c *Client
}

func (j *RefreshTokenService) Do(ctx context.Context, opts ...RequestOption) (res *RefreshToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v1/tokens/refresh",
		secType:  secTypeRefreshToken,
	}

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(RefreshToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RefreshToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type BalancesService struct {
	c            *Client
	accessToken  string
	refreshToken string
}

func (j *BalancesService) Do(ctx context.Context, opts ...RequestOption) (res *Balances, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v1/balances",
		secType:  secTypeAccessToken,
	}

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(Balances)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type Balances struct{}

type DailyReportService struct {
	c            *Client
	yearMonthDay string
}

func (j *DailyReportService) YearMonthDay(yearMonthDay string) *DailyReportService {
	j.yearMonthDay = yearMonthDay

	return j
}

func (j *DailyReportService) Do(ctx context.Context, opts ...RequestOption) (res *DailyReport, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/reports/daily",
		secType:  secTypeAccessToken,
	}

	r.setParam("yearMonthDay", j.yearMonthDay)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DailyReport)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DailyReport struct{}
