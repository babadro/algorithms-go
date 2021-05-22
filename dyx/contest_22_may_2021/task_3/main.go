package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

type Feed struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	OfferID   string `json:"offer_id"`
	MarketSKU int    `json:"market_sku"`
	Price     int    `json:"price"`
}

func main() {
	offers := parseOffers()

	sort.Slice(offers, func(i, j int) bool {
		price1, price2 := offers[i].Price, offers[j].Price
		if price1 != price2 {
			return price1 < price2
		}

		return offers[i].OfferID < offers[j].OfferID
	})

	resFeed := Feed{Offers: offers}

	resBytes, err := json.Marshal(resFeed)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile("output.txt", resBytes, 0644); err != nil {
		panic(err)
	}
}

func parseOffers() (arr []Offer) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	flag := false
	for scanner.Scan() {
		b := scanner.Bytes()

		if !flag {
			flag = true
		} else {
			var feed Feed
			if err = json.Unmarshal(b, &feed); err != nil {
				panic(err)
			}

			arr = append(arr, feed.Offers...)
		}

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}

	return arr
}
