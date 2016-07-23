package main

import (
	"fmt"
)

type HttpResponseFetcher interface {
	Fetch(url string) (int, error)
}

type RealResponseFetcher struct{}

func (rsf *RealResponseFetcher) Fetch(url string) (int, error) {
	fmt.Printf("I am in real fetch %s", url)
	return 0, nil
}

func newmain() {
	rsf := &RealResponseFetcher{}
	i, err := getResponse(rsf)
	if err != nil {
		fmt.Printf("I am error, Bro!")
	}
	fmt.Printf("I am success %d", i)
}

func getResponse(fetcher HttpResponseFetcher) (int, error) {
	response, err := fetcher.Fetch("http://example.com/info")
	if err != nil {
		return 0, fmt.Errorf("I am blowing")
	}
	return response, nil
}
