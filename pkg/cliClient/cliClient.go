package cliClient

import (
	"github.com/datreeio/datree/pkg/httpClient"
	"github.com/datreeio/datree/pkg/restyClient"
	"github.com/go-resty/resty/v2"
)

type HTTPClient interface {
	Request(method string, resourceURI string, body interface{}, headers map[string]string) (httpClient.Response, error)
}
type NetworkValidator interface {
	IdentifyNetworkError(errStr string) error
	IsLocalMode() bool
}

type CliClient struct {
	baseUrl          string
	httpClient       HTTPClient
	timeoutClient    HTTPClient
	httpErrors       []string
	networkValidator NetworkValidator
	flagsHeaders     map[string]string
	restyClient      *resty.Client
}

func NewCliClient(url string, networkValidator NetworkValidator) *CliClient {
	httpClient := httpClient.NewClient(url, nil)
	return &CliClient{
		baseUrl:          url,
		httpClient:       httpClient,
		timeoutClient:    nil,
		httpErrors:       []string{},
		networkValidator: networkValidator,
		flagsHeaders:     make(map[string]string),
		restyClient:      restyClient.New(url),
	}
}
