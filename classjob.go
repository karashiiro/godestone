package godestone

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/v2/internal/selectors"
)

var nonDigits = regexp.MustCompile("[^\\d]")

func (s *Scraper) buildClassJobCollector(charData *Character) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	classJobSelectors := s.getProfileSelectors().ClassJob
	charData.ClassJobs = make([]*ClassJob, 0)

	cjRefs := []*selectors.OneClassJobSelectors{
		&classJobSelectors.Paladin,
		&classJobSelectors.Warrior,
		&classJobSelectors.DarkKnight,
		&classJobSelectors.Gunbreaker,
		&classJobSelectors.Monk,
		&classJobSelectors.Dragoon,
		&classJobSelectors.Ninja,
		&classJobSelectors.Samurai,
		&classJobSelectors.WhiteMage,
		&classJobSelectors.Scholar,
		&classJobSelectors.Astrologian,
		&classJobSelectors.Bard,
		&classJobSelectors.Machinist,
		&classJobSelectors.Dancer,
		&classJobSelectors.BlackMage,
		&classJobSelectors.Summoner,
		&classJobSelectors.RedMage,
		&classJobSelectors.BlueMage,
		&classJobSelectors.Carpenter,
		&classJobSelectors.Blacksmith,
		&classJobSelectors.Armorer,
		&classJobSelectors.Goldsmith,
		&classJobSelectors.Leatherworker,
		&classJobSelectors.Weaver,
		&classJobSelectors.Alchemist,
		&classJobSelectors.Culinarian,
		&classJobSelectors.Miner,
		&classJobSelectors.Botanist,
		&classJobSelectors.Fisher,
	}

	for _, ref := range cjRefs {
		curRef := ref
		curCj := ClassJob{}

		c.OnHTML(curRef.Exp.Selector, func(e *colly.HTMLElement) {
			expStrs := curRef.Exp.Parse(e)
			expStrs[0] = nonDigits.ReplaceAllString(expStrs[0], "")
			expStrs[1] = nonDigits.ReplaceAllString(expStrs[1], "")

			curExp, err := strconv.ParseUint(expStrs[0], 10, 32)
			if err == nil {
				curCj.ExpLevel = uint32(curExp)
			}

			maxExp, err := strconv.ParseUint(expStrs[1], 10, 32)
			if err == nil {
				curCj.ExpLevelMax = uint32(maxExp)
			}

			curCj.ExpLevelTogo = curCj.ExpLevelMax - curCj.ExpLevel
		})

		c.OnHTML(curRef.Level.Selector, func(e *colly.HTMLElement) {
			levelStr := curRef.Level.Parse(e)[0]
			level, err := strconv.ParseUint(levelStr, 10, 8)
			if err == nil {
				curCj.Level = uint8(level)
			}
		})

		c.OnHTML(curRef.UnlockState.Selector, func(e *colly.HTMLElement) {
			curCj.UnlockedState.Name = curRef.UnlockState.Parse(e)[0]

			// This bit is kinda wack, should fix it
			curCj.Name = e.Attr("data-tooltip")
			names := strings.Split(curCj.Name, " / ")
			names[0] = strings.TrimRight(strings.Split(names[0], "(")[0], " ")
			names[0] = strings.TrimRight(strings.Split(names[0], "[")[0], " ")

			curCj.IsSpecialized = strings.Contains(e.Attr("class"), "meister")

			jobInfo := s.dataProvider.ClassJob(names[0])
			curCj.JobID = uint8(jobInfo.ID)

			if len(names) > 1 {
				classInfo := s.dataProvider.ClassJob(names[1])
				curCj.ClassID = uint8(classInfo.ID)
			} else {
				curCj.ClassID = uint8(curCj.JobID)
			}

			cjInfo := s.dataProvider.ClassJob(curCj.UnlockedState.Name)
			curCj.UnlockedState.ID = uint8(cjInfo.ID)
		})

		charData.ClassJobs = append(charData.ClassJobs, &curCj)
	}

	cjb := &ClassJobBozja{}
	c.OnHTML(classJobSelectors.Bozja.Level.Selector, func(e *colly.HTMLElement) {
		levelStr := classJobSelectors.Bozja.Level.Parse(e)[0]
		level, err := strconv.ParseUint(levelStr, 10, 8)
		if err == nil {
			cjb.Level = uint8(level)
		}
	})

	c.OnHTML(classJobSelectors.Bozja.Mettle.Selector, func(e *colly.HTMLElement) {
		mettleStr := classJobSelectors.Bozja.Mettle.Parse(e)[0]
		mettleStr = nonDigits.ReplaceAllString(mettleStr, "")
		mettle, err := strconv.ParseUint(mettleStr, 10, 32)
		if err == nil {
			cjb.Mettle = uint32(mettle)
		}
	})

	c.OnHTML(classJobSelectors.Bozja.Name.Selector, func(e *colly.HTMLElement) {
		cjb.Name = classJobSelectors.Bozja.Name.Parse(e)[0]
	})
	charData.ClassJobBozjan = cjb

	cje := &ClassJobEureka{}
	c.OnHTML(classJobSelectors.Eureka.Level.Selector, func(e *colly.HTMLElement) {
		levelStr := classJobSelectors.Eureka.Level.Parse(e)[0]
		level, err := strconv.ParseUint(levelStr, 10, 8)
		if err == nil {
			cje.Level = uint8(level)
		}
	})

	c.OnHTML(classJobSelectors.Eureka.Exp.Selector, func(e *colly.HTMLElement) {
		expStrs := classJobSelectors.Eureka.Exp.Parse(e)
		expStrs[0] = nonDigits.ReplaceAllString(expStrs[0], "")
		expStrs[1] = nonDigits.ReplaceAllString(expStrs[1], "")

		curExp, err := strconv.ParseUint(expStrs[0], 10, 32)
		if err == nil {
			cje.ExpLevel = uint32(curExp)
		}

		maxExp, err := strconv.ParseUint(expStrs[1], 10, 32)
		if err == nil {
			cje.ExpLevelMax = uint32(maxExp)
		}

		cje.ExpLevelTogo = cje.ExpLevelMax - cje.ExpLevel
	})

	c.OnHTML(classJobSelectors.Eureka.Name.Selector, func(e *colly.HTMLElement) {
		cje.Name = classJobSelectors.Eureka.Name.Parse(e)[0]
	})
	charData.ClassJobElemental = cje

	return c
}
