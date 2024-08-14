package jibit

import (
	"context"
	"net/http"
)

type PercentageSimilarityWithIdentityInformationService struct {
	c            *Client
	nationalCode string
	birthDate    string
	firstName    string
	lastName     string
	fullName     string
	fatherName   string
}

func (j *PercentageSimilarityWithIdentityInformationService) NationalCode(nationalCode string) *PercentageSimilarityWithIdentityInformationService {
	j.nationalCode = nationalCode

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) BirthDate(birthDate string) *PercentageSimilarityWithIdentityInformationService {
	j.birthDate = birthDate

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) FirstName(firstName string) *PercentageSimilarityWithIdentityInformationService {
	j.firstName = firstName

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) LastName(lastName string) *PercentageSimilarityWithIdentityInformationService {
	j.lastName = lastName

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) FullName(fullName string) *PercentageSimilarityWithIdentityInformationService {
	j.fullName = fullName

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) FatherName(fatherName string) *PercentageSimilarityWithIdentityInformationService {
	j.fatherName = fatherName

	return j
}

func (j *PercentageSimilarityWithIdentityInformationService) Do(ctx context.Context, opts ...RequestOption) (res *PercentageSimilarityWithIdentityInformation, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/services/identity/similarity",
		secType:  secTypeAccessToken,
	}

	r.setParam("nationalCode", j.nationalCode)
	r.setParam("birthDate", j.birthDate)
	r.setParam("firstName", j.firstName)
	r.setParam("lastName", j.lastName)
	r.setParam("fullName", j.fullName)
	r.setParam("fatherName", j.fatherName)

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(PercentageSimilarityWithIdentityInformation)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type PercentageSimilarityWithIdentityInformation struct{}
