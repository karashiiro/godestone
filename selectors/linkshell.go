package selectors

import (
	"encoding/json"

	"github.com/xivapi/godestone/v2/internal/pack/css"
)

// LinkshellBasicSelectors contains the CSS selectors for the basic information on the linkshell page.
type LinkshellBasicSelectors struct {
	Name SelectorInfo `json:"NAME"`
}

// LinkshellMemberSelectors contains the CSS selectors for the member list on the linkshell page.
type LinkshellMemberSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root              SelectorInfo `json:"ROOT"`
		Avatar            SelectorInfo `json:"AVATAR"`
		ID                SelectorInfo `json:"ID"`
		Name              SelectorInfo `json:"NAME"`
		Rank              SelectorInfo `json:"RANK"`
		RankIcon          SelectorInfo `json:"RANK_ICON"`
		LinkshellRank     SelectorInfo `json:"LINKSHELL_RANK"`
		LinkshellRankIcon SelectorInfo `json:"LINKSHELL_RANK_ICON"`
		Server            SelectorInfo `json:"SERVER"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
}

// LinkshellSelectors contains the CSS selectors for the linkshell page.
type LinkshellSelectors struct {
	Basic   *LinkshellBasicSelectors
	Members *LinkshellMemberSelectors
}

// LoadLinkshellSelectors loads the CSS selectors for the linkshell page.
func LoadLinkshellSelectors() *LinkshellSelectors {
	basicBytes, _ := css.Asset("linkshell/ls.json")
	basicSelectors := LinkshellBasicSelectors{}
	json.Unmarshal(basicBytes, &basicSelectors)

	membersBytes, _ := css.Asset("linkshell/members.json")
	membersSelectors := LinkshellMemberSelectors{}
	json.Unmarshal(membersBytes, &membersSelectors)

	return &LinkshellSelectors{
		Basic:   &basicSelectors,
		Members: &membersSelectors,
	}
}
