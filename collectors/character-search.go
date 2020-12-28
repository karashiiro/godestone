package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildCharacterSearchCollector builds the collector used for processing the page.
func BuildCharacterSearchCollector(
	meta *models.Meta,
	lastURI string,
	searchSelectors *selectors.SearchSelectors,
	output chan *models.CharacterSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(41),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	charSearchSelectors := searchSelectors.Character
	entrySelectors := charSearchSelectors.Entry

	revisitedOnce := false
	c.OnHTML(charSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := charSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCharacter := models.CharacterSearchResult{
				Avatar:   entrySelectors.Avatar.ParseThroughChildren(e)[0],
				Name:     entrySelectors.Name.ParseThroughChildren(e)[0],
				Lang:     entrySelectors.Lang.ParseThroughChildren(e)[0],
				RankIcon: entrySelectors.RankIcon.ParseThroughChildren(e)[0],
			}

			idStr := entrySelectors.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextCharacter.ID = uint32(id)
			}

			gcRank := entrySelectors.Rank.ParseThroughChildren(e)[0]
			nextCharacter.Rank = gcrank.Parse(gcRank)

			worldDC := entrySelectors.Server.ParseThroughChildren(e)
			nextCharacter.World = worldDC[0]
			nextCharacter.DC = worldDC[1]

			output <- &nextCharacter
		})

		if nextURI != "javascript:void(0);" {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.CharacterSearchResult{
					Error: err,
				}
			}
		} else if !revisitedOnce && nextURI != "" /* "Your search yielded no results." */ {
			revisitedOnce = true
			err := container.Request.Visit(lastURI)
			if err != nil {
				output <- &models.CharacterSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
