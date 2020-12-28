package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildCWLSSearchCollector builds the collector used for processing the page.
func BuildCWLSSearchCollector(
	meta *models.Meta,
	lastURI string,
	searchSelectors *selectors.SearchSelectors,
	output chan *models.CWLSSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(41),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	cwlsSearchSelectors := searchSelectors.CWLS
	entrySelectors := cwlsSearchSelectors.Entry

	revisitedOnce := false
	c.OnHTML(cwlsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := cwlsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCWLS := models.CWLSSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
			}

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextCWLS.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextCWLS
		})

		if nextURI != "javascript:void(0);" {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.CWLSSearchResult{
					Error: err,
				}
			}
		} else if !revisitedOnce && nextURI != "" /* "Your search yielded no results." */ {
			revisitedOnce = true
			err := container.Request.Visit(lastURI)
			if err != nil {
				output <- &models.CWLSSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
