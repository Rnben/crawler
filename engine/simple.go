package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
