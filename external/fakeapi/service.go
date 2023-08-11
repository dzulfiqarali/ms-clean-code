package fakeapi

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"github.com/ms-clean-code/configs"
	"time"
)

const defaultTimeout = 40 * time.Second

type FakeApiImpl interface {
	AddNewProductFakeApi(request RequestFakeAPI) (resp ResponseFakeAPI, err error)
}

type ClientImpl struct {
	httpClient *resty.Client
	config     *configs.Config
}

func NewClientRequest(config *configs.Config) *ClientImpl {
	return &ClientImpl{
		httpClient: resty.New().
			SetDebug(true).
			SetTimeout(defaultTimeout).
			SetBaseURL(config.FakeApi.Host).
			SetTLSClientConfig(&tls.Config{
				InsecureSkipVerify: true,
			}).
			SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)),
		config: config,
	}
}

func (c *ClientImpl) AddNewProductFakeApi(request RequestFakeAPI) (resp ResponseFakeAPI, err error) {

	res, err := c.httpClient.R().
		SetBody(request).
		SetResult(&resp).
		Post("/products")

	if err != nil {
		return
	}

	//result := res.String()
	//if err = json.Unmarshal([]byte(result), &resp); err != nil {
	//	return
	//}

	defer res.RawBody().Close()

	return
}
