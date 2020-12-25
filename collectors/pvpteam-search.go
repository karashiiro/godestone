package collectors

import (
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildPVPTeamSearchCollector builds the collector used for processing the page.
func BuildPVPTeamSearchCollector(meta *models.Meta, searchSelectors *selectors.SearchSelectors, output chan *models.PVPTeamSearchResult) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(21),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	pvpTeamSearchSelectors := searchSelectors.PVPTeam
	entrySelectors := pvpTeamSearchSelectors.Entry

	c.OnHTML(pvpTeamSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := pvpTeamSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextTeam := models.PVPTeamSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
				CrestLayers: &models.CrestLayers{
					Bottom: entrySelectors.CrestLayers.Bottom.ParseThroughChildren(e)[0],
					Middle: entrySelectors.CrestLayers.Middle.ParseThroughChildren(e)[0],
					Top:    entrySelectors.CrestLayers.Top.ParseThroughChildren(e)[0],
				},
			}

			output <- &nextTeam
		})

		if nextURI != "javascript:void(0);" && nextURI != "" /* "Your search yielded no results." */ {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.PVPTeamSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
