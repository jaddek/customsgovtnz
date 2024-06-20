package rate

import "net/http"

const (
	NZCUSTOMSGOV_URL        = "https://www.customs.govt.nz/"
	NZCUSTOMSGOV_RATES_PATH = "/api/datafiles/current-exchange"
)

type INZCustomsGovHttpClient interface {
	GetRates() (*http.Response, error)
}

type NZCustomsGovHttpClient struct {
	HOST string
}

func MakeNZCustomsGovHttpClient(host string) *NZCustomsGovHttpClient {
	return &NZCustomsGovHttpClient{HOST: host}
}

func (c *NZCustomsGovHttpClient) GetRates() (*http.Response, error) {
	return http.DefaultClient.Get(c.HOST + NZCUSTOMSGOV_RATES_PATH)
}
