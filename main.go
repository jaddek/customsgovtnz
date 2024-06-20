package main

import (
	"fmt"

	"jaddek.com/jaddek/nzcustomsgov/rate"
)

func main() {
	httpClient := rate.MakeNZCustomsGovHttpClient(rate.NZCUSTOMSGOV_URL)
	body := rate.GetNzCustomsRates(httpClient)
	envelope := rate.MakeExchangeRateList(body)

	fmt.Println(string(envelope.GetEnvelopeAsJson()))
}
