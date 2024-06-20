package main

import (
	"fmt"

	"github.com/jaddek/customsgovtnz/rate"
)

func main() {
	httpClient := rate.MakeCustomsGovtNZHttpClient(rate.CUSTOMSGOVTNZ_URL)
	body := rate.GetCustomsGovtNZRates(httpClient)
	envelope := rate.MakeExchangeRateList(body)

	fmt.Println(string(envelope.GetEnvelopeAsJson()))
}
