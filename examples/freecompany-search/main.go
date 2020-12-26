package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/godestone"
	"github.com/karashiiro/godestone/search"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := search.FreeCompanyOptions{}
	if len(os.Args) > 1 {
		opts = search.FreeCompanyOptions{
			Name: os.Args[1],
		}
	}

	for fc := range s.SearchFreeCompanies(opts) {
		if fc.Error != nil {
			log.Fatalln(fc.Error)
		}

		fcJSON, err := json.MarshalIndent(fc, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(fcJSON))
	}
}
