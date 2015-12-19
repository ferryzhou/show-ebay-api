package main

import (
	"flag"
	"fmt"
	"github.com/heatxsink/go-ebay"
	"os"
)

var (
	ebay_application_id = flag.String("appid", "", "ebay application id")
	keywords            = flag.String("keywords", "", "keywords")
)

func main() {
	flag.Parse()
	if *keywords == "" {
		panic("no keywords input, please specify with -keywords")
	}
	if *ebay_application_id == "" {
		panic("no app id input, please specify with -appid")
	}
	e := ebay.New(*ebay_application_id)
	response, err := e.FindItemsByKeywords(ebay.GLOBAL_ID_EBAY_US, *keywords, 10)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	for _, i := range response.Items {
		fmt.Println("Title: ", i.Title)
		fmt.Println("\tListing Url:     ", i.ListingUrl)
		fmt.Println("\tBin Price:       ", i.BinPrice)
		fmt.Println("\tCurrent Price:   ", i.CurrentPrice)
		fmt.Println("\tShipping Price:  ", i.ShippingPrice)
		fmt.Println()
	}
}
