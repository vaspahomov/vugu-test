package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Bpi struct {
	Time struct {
		Updated string `json:"updated"`
	} `json:"time"`
	BPI map[string]struct {
		Code      string  `json:"code"`
		Symbol    string  `json:"symbol"`
		RateFloat float64 `json:"rate_float"`
	} `json:"bpi"`
}

// Client test
type Client struct {
}

func (c *Client) Request() (*Bpi, error) {
	res, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Printf("Error fetch()ing: %v", err)
		return nil, nil
	}
	defer res.Body.Close()
	newb := &Bpi{}
	err = json.NewDecoder(res.Body).Decode(newb)
	if err != nil {
		log.Printf("Error JSON decoding: %v", err)
		return nil, nil
	}
	return newb, nil
}
