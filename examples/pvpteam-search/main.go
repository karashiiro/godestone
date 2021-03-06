package main

import (
	"log"
	"os"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func main() {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	opts := godestone.PVPTeamOptions{
		Name: os.Args[1],
	}

	for pvpTeam := range s.SearchPVPTeams(opts) {
		if pvpTeam.Error != nil {
			log.Fatalln(pvpTeam.Error)
		}

		log.Println(pvpTeam)
		log.Println(pvpTeam.CrestLayers)
	}
}
