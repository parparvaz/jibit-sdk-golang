package jibit

import (
	"context"
	"net/http"
)

type HealthCheckCardToIbanService struct {
	c *Client
}

func (j *HealthCheckCardToIbanService) Do(ctx context.Context, opts ...RequestOption) (res *HealthCheckCardToIban, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/availability",
		secType:  secTypeAccessToken,
	}

	r.setParam("cardToIBAN", true)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(HealthCheckCardToIban)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type HealthCheckCardToIban struct{}
