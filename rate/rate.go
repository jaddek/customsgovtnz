package rate

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
)

type ExchangeRateList struct {
	ExchangeRate []Rate `json:"exchangeRate"  xml:"exchangeRate"`
}

type Rate struct {
	CountryName  string `json:"country" xml:"countryName"`
	CurrencyCode string `json:"currencyCode" xml:"currencyCode"`
	DateNow      string `json:"dateNow" xml:"dateNow"`
	RateNow      string `json:"rateNow" xml:"rateNow"`
	DateFuture   string `json:"dateFuture" xml:"dateFuture"`
	RateFuture   string `json:"rateFuture" xml:"rateFuture"`
	CurrencyName string `json:"currencyName" xml:"currencyName"`
}

func (exchangeRateList *ExchangeRateList) GetCurrencies() []string {
	rates := exchangeRateList.GetRates()

	currencies := make([]string, 0)
	for _, rate := range rates {
		currencies = append(currencies, rate.CurrencyCode)
	}

	return currencies
}

func (exchangeRateList *ExchangeRateList) GetCurrenciesAsJson() []byte {
	currencies := exchangeRateList.GetCurrencies()

	return asJson(currencies)
}

func (envelope *ExchangeRateList) GetRatesAsJson() []byte {
	rates, err := json.Marshal(envelope.GetRates())

	if nil != err {
		log.Fatal("Error marshalling to JSON", err)
	}

	return rates
}

func (exchangeRateList *ExchangeRateList) GetRateObjectByCurrency(currency string) (*Rate, error) {
	rates := exchangeRateList.GetRates()

	for _, rate := range rates {
		if rate.CurrencyCode == currency {
			return &rate, nil
		}
	}

	return nil, fmt.Errorf("currency %s not found", currency)
}

func (exchangeRateList *ExchangeRateList) GetRateByCurrencyAsJson(currency string) []byte {
	rate, err := exchangeRateList.GetRateObjectByCurrency(currency)

	if err != nil {
		log.Panic(err)
	}

	return asJson(rate)
}

func (exchangeRateList *ExchangeRateList) GetRateValueByCurrency(currency string) (string, error) {
	rates := exchangeRateList.GetRates()

	for _, rate := range rates {
		if rate.CurrencyCode == currency {
			return string(rate.RateNow), nil
		}
	}

	return "", fmt.Errorf("currency %s not found", currency)
}

func (exchangeRateList *ExchangeRateList) GetRates() []Rate {
	return exchangeRateList.ExchangeRate
}

func (envelope *ExchangeRateList) GetEnvelopeAsJson() []byte {
	return asJson(envelope)
}

func asJson(o interface{}) []byte {
	result, err := json.Marshal(o)
	if nil != err {
		log.Fatal("Error marshalling to JSON", err)
	}

	return result
}

func GetNzCustomsRates(httpClient INZCustomsGovHttpClient) []byte {
	resp, err := httpClient.GetRates()

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func MakeExchangeRateList(body []byte) *ExchangeRateList {
	exchangeRateList := &ExchangeRateList{}
	err := xml.Unmarshal(body, exchangeRateList)

	if nil != err {
		log.Fatal("Error unmarshalling from XML", err)
	}

	return exchangeRateList
}
