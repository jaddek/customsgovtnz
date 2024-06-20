package rate

import "net/http"

const (
	CUSTOMSGOVTNZ_URL        = "https://www.customs.govt.nz/"
	CUSTOMSGOVTNZ_RATES_PATH = "/api/datafiles/current-exchange"
)

type ICustomsGovtNZHttpClient interface {
	GetRates() (*http.Response, error)
}

type CustomsGovtNZHttpClient struct {
	HOST string
}

func MakeCustomsGovtNZHttpClient(host string) *CustomsGovtNZHttpClient {
	return &CustomsGovtNZHttpClient{HOST: host}
}

func (c *CustomsGovtNZHttpClient) GetRates() (*http.Response, error) {
	return http.DefaultClient.Get(c.HOST + CUSTOMSGOVTNZ_RATES_PATH)
}
