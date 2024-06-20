package rate

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expected = `<?xml version="1.0"?>
<exchangeRateList>
    <exchangeRate>
        <countryName>Australia</countryName>
        <currencyCode>AUD</currencyCode>
        <dateNow>2024-06-23</dateNow>
        <rateNow>0.91</rateNow>
        <dateFuture>2024-07-07</dateFuture>
        <rateFuture>0.91</rateFuture>
        <currencyName>Australian Dollar</currencyName>
    </exchangeRate>
</exchangeRateList>
`
	expectedJson = `{"exchangeRate":[{"country":"Australia","currencyCode":"AUD","dateNow":"2024-06-23","rateNow":"0.91","dateFuture":"2024-07-07","rateFuture":"0.91","currencyName":"Australian Dollar"}]}
`
)

func TestGetNzCustomsRates(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	httpClient := MakeCustomsGovtNZHttpClient(svr.URL)

	content := string(GetCustomsGovtNZRates(httpClient))

	assert.Equal(t, content, expected)
}

func TestMakeExchangeRateList(t *testing.T) {
	envelope := MakeExchangeRateList([]byte(expected))

	currency, _ := envelope.GetRateValueByCurrency("AUD")
	assert.NotEmpty(t, currency)

	currencies := envelope.GetCurrencies()
	assert.NotEmpty(t, currencies)
	assert.Equal(t, 1, len(currencies))
}

func TestEnvelopeGetEnvelopeAsJson(t *testing.T) {
	envelope := MakeExchangeRateList([]byte(expected))

	assert.JSONEq(t, expectedJson, string(envelope.GetEnvelopeAsJson()))
}
