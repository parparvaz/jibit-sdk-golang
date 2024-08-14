package jibit

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	baseAPIMainURL = "https://napi.jibit.ir/ide"
)

const (
	CacheAccessToken  string = "access_token"
	CacheRefreshToken string = "refresh_token"
)

var UseTestnet = false

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Jibit/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Jibit-golang ", log.LstdFlags),
		cache:      newCache(),
	}
}

func NewProxyClient(apiKey, secretKey, proxyUrl string) *Client {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		log.Println(err)

		return nil
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseAPIMainURL,
		UserAgent: "Jibit/golang",
		HTTPClient: &http.Client{
			Transport: tr,
		},
		Logger: log.New(os.Stderr, "Jibit-golang ", log.LstdFlags),
		cache:  newCache(),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	ApiKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
	cache      *cache
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	body := &bytes.Buffer{}

	if r.secType == secTypeNone {
		r.setJsonParams(params{
			"apiKey":    c.ApiKey,
			"secretKey": c.SecretKey,
		})
	} else if r.secType == secTypeRefreshToken {
		accessToken, _ := c.cache.get(CacheAccessToken)
		refreshToken, _ := c.cache.get(CacheRefreshToken)
		r.setJsonParams(params{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	} else if r.secType == secTypeAccessToken {
		_, ok := c.cache.get(CacheAccessToken)
		if !ok {
			c.getAuth()
		}

		accessToken, _ := c.cache.get(CacheAccessToken)
		header.Add("Authorization", "Bearer "+accessToken)
	}

	queryString := r.query.Encode()
	bodyString := r.form.Encode()
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	if r.json != nil {
		header.Set("Content-Type", "application/json")
		body = bytes.NewBuffer(r.json)
	}

	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body

	return nil
}

func (c *Client) getAuth() {
	newClient := NewClient(c.ApiKey, c.SecretKey)

	res, err := newClient.NewGenerateTokenService().Do(context.Background())
	if err != nil {
		return
	}

	c.cache.set(CacheAccessToken, res.AccessToken)
	c.cache.set(CacheRefreshToken, res.RefreshToken)
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}

	req = req.WithContext(ctx)
	req.Header = r.header

	c.debug("request: %#v", req)

	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}

	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {

		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}

func (c *Client) SetCache(key, value string) {
	c.cache.set(key, value)
}

func (c *Client) NewCardToAccountNumberService() *CardToAccountNumberService {
	return &CardToAccountNumberService{c: c}
}
func (c *Client) NewAccountNumberToIbanService() *AccountNumberToIbanService {
	return &AccountNumberToIbanService{c: c}
}
func (c *Client) NewCardToIbanService() *CardToIbanService {
	return &CardToIbanService{c: c}
}
func (c *Client) NewHealthCheckCardToIbanService() *HealthCheckCardToIbanService {
	return &HealthCheckCardToIbanService{c: c}
}
func (c *Client) NewCorporationInquiryService() *CorporationInquiryService {
	return &CorporationInquiryService{c: c}
}
func (c *Client) NewMilitaryServiceInquiryService() *MilitaryServiceInquiryService {
	return &MilitaryServiceInquiryService{c: c}
}
func (c *Client) NewCardInquiryService() *CardInquiryService {
	return &CardInquiryService{c: c}
}
func (c *Client) NewIbanInquiryService() *IbanInquiryService {
	return &IbanInquiryService{c: c}
}
func (c *Client) NewLegalInquiryService() *LegalInquiryService {
	return &LegalInquiryService{c: c}
}
func (c *Client) NewPostalInquiryService() *PostalInquiryService {
	return &PostalInquiryService{c: c}
}
func (c *Client) NewPostalWithWorldGeodeticSystemInquiryService() *PostalWithWorldGeodeticSystemInquiryService {
	return &PostalWithWorldGeodeticSystemInquiryService{c: c}
}
func (c *Client) NewMatchCardNumberWithNationalCodeService() *MatchCardNumberWithNationalCodeService {
	return &MatchCardNumberWithNationalCodeService{c: c}
}
func (c *Client) NewMatchAccountNumberWithNationalCodeService() *MatchAccountNumberWithNationalCodeService {
	return &MatchAccountNumberWithNationalCodeService{c: c}
}
func (c *Client) NewMatchIbanWithNationalCodeService() *MatchIbanWithNationalCodeService {
	return &MatchIbanWithNationalCodeService{c: c}
}
func (c *Client) NewMatchCardNumberWithNameService() *MatchCardNumberWithNameService {
	return &MatchCardNumberWithNameService{c: c}
}
func (c *Client) NewMatchAccountNumberWithNameService() *MatchAccountNumberWithNameService {
	return &MatchAccountNumberWithNameService{c: c}
}
func (c *Client) NewMatchIbanWithNameService() *MatchIbanWithNameService {
	return &MatchIbanWithNameService{c: c}
}
func (c *Client) NewMatchNationalCodeWithMobileNumberService() *MatchNationalCodeWithMobileNumberService {
	return &MatchNationalCodeWithMobileNumberService{c: c}
}
func (c *Client) NewMatchUniversalIDNationalsWithPassportNumberService() *MatchUniversalIDNationalsWithPassportNumberService {
	return &MatchUniversalIDNationalsWithPassportNumberService{c: c}
}
func (c *Client) NewPercentageSimilarityWithIdentityInformationService() *PercentageSimilarityWithIdentityInformationService {
	return &PercentageSimilarityWithIdentityInformationService{c: c}
}
func (c *Client) NewGenerateTokenService() *GenerateTokenService {
	return &GenerateTokenService{c: c}
}
func (c *Client) NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{c: c}
}
func (c *Client) NewDailyReportService() *DailyReportService {
	return &DailyReportService{c: c}
}

const (
	CentralBankOfTheIslamicRepublicOfIran string = "MARKAZI"
	BankOfIndustryMine                    string = "SANAT_VA_MADAN"
	BankMellat                            string = "MELLAT"
	RefahKBank                            string = "REFAH"
	BankMaskan                            string = "MASKAN"
	BankSepah                             string = "SEPAH"
	BankKeshavarziIran                    string = "KESHAVARZI"
	BankMelliIran                         string = "MELLI"
	TejaratBank                           string = "TEJARAT"
	BankSaderatIran                       string = "SADERAT"
	ExportDevelopmentBankOfIran           string = "TOSEAH_SADERAT"
	PostBankIran                          string = "POST"
	ToseeTaavonBank                       string = "TOSEAH_TAAVON"
	KarafarinBank                         string = "KARAFARIN"
	ParsianBank                           string = "PARSIAN"
	EghtesadNovinBank                     string = "EGHTESAD_NOVIN"
	SamanBank                             string = "SAMAN"
	PasargadBank                          string = "PASARGAD"
	SarmayehBank                          string = "SARMAYEH"
	SinaBank                              string = "SINA"
	GharzolhasaneMehrIranBank             string = "MEHR_IRAN"
	ShahrBank                             string = "SHAHR"
	AyandehBank                           string = "AYANDEH"
	TourismBank                           string = "GARDESHGARI"
	DayBank                               string = "DAY"
	IranZaminBank                         string = "IRANZAMIN"
	ResalatGharzolhasaneBank              string = "RESALAT"
	MelalCreditInstitution                string = "MELAL"
	MiddleEastBank                        string = "KHAVARMIANEH"
	NoorCreditInstitution                 string = "NOOR"
	IranVenezuelaBiNationalBank           string = "IRAN_VENEZUELA"
	UNKNOWN                               string = "UNKNOWN"
)
