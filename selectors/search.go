package selectors

import (
	"encoding/json"

	"github.com/xivapi/godestone/v2/pack/css"
)

// CharacterSearchSelectors contains the CSS selectors for the character search interface.
type CharacterSearchSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root     SelectorInfo `json:"ROOT"`
		Avatar   SelectorInfo `json:"AVATAR"`
		ID       SelectorInfo `json:"ID"`
		Lang     SelectorInfo `json:"LANG"`
		Name     SelectorInfo `json:"NAME"`
		Rank     SelectorInfo `json:"RANK"`
		RankIcon SelectorInfo `json:"RANK_ICON"`
		Server   SelectorInfo `json:"SERVER"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// CWLSSearchSelectors contains the CSS selectors for the CWLS search interface.
type CWLSSearchSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root          SelectorInfo `json:"ROOT"`
		ID            SelectorInfo `json:"ID"`
		Name          SelectorInfo `json:"NAME"`
		DC            SelectorInfo `json:"DC"`
		ActiveMembers SelectorInfo `json:"ACTIVE_MEMBERS"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// LinkshellSearchSelectors contains the CSS selectors for the linkshell search interface.
type LinkshellSearchSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root          SelectorInfo `json:"ROOT"`
		ID            SelectorInfo `json:"ID"`
		Name          SelectorInfo `json:"NAME"`
		Server        SelectorInfo `json:"SERVER"`
		ActiveMembers SelectorInfo `json:"ACTIVE_MEMBERS"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// PVPTeamSearchSelectors contains the CSS selectors for the PVP team search interface.
type PVPTeamSearchSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root        SelectorInfo `json:"ROOT"`
		ID          SelectorInfo `json:"ID"`
		Name        SelectorInfo `json:"NAME"`
		DC          SelectorInfo `json:"DC"`
		CrestLayers struct {
			Bottom SelectorInfo `json:"BOTTOM"`
			Middle SelectorInfo `json:"MIDDLE"`
			Top    SelectorInfo `json:"TOP"`
		} `json:"CREST_LAYERS"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// FreeCompanySearchSelectors contains the CSS selectors for the Free Company search interface.
type FreeCompanySearchSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root            SelectorInfo `json:"ROOT"`
		ID              SelectorInfo `json:"ID"`
		Name            SelectorInfo `json:"NAME"`
		Server          SelectorInfo `json:"SERVER"`
		GrandCompany    SelectorInfo `json:"GRAND_COMPANY"`
		Active          SelectorInfo `json:"ACTIVE"`
		ActiveMembers   SelectorInfo `json:"ACTIVE_MEMBERS"`
		RecruitmentOpen SelectorInfo `json:"RECRUITMENT_OPEN"`
		EstateBuilt     SelectorInfo `json:"ESTATE_BUILT"`
		Formed          SelectorInfo `json:"FORMED"`
		CrestLayers     struct {
			Bottom SelectorInfo `json:"BOTTOM"`
			Middle SelectorInfo `json:"MIDDLE"`
			Top    SelectorInfo `json:"TOP"`
		} `json:"CREST_LAYERS"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// SearchSelectors contains the CSS selectors for the search interface.
type SearchSelectors struct {
	Character   *CharacterSearchSelectors
	CWLS        *CWLSSearchSelectors
	FreeCompany *FreeCompanySearchSelectors
	Linkshell   *LinkshellSearchSelectors
	PVPTeam     *PVPTeamSearchSelectors
}

// LoadSearchSelectors loads the CSS selectors for the search interface.
func LoadSearchSelectors() *SearchSelectors {
	charaBytes, _ := css.Asset("search/character.json")
	charaSearchSelectors := CharacterSearchSelectors{}
	json.Unmarshal(charaBytes, &charaSearchSelectors)

	cwlsBytes, _ := css.Asset("search/cwls.json")
	cwlsSearchSelectors := CWLSSearchSelectors{}
	json.Unmarshal(cwlsBytes, &cwlsSearchSelectors)

	fcBytes, _ := css.Asset("search/freecompany.json")
	fcSearchSelectors := FreeCompanySearchSelectors{}
	json.Unmarshal(fcBytes, &fcSearchSelectors)

	lsBytes, _ := css.Asset("search/linkshell.json")
	lsSearchSelectors := LinkshellSearchSelectors{}
	json.Unmarshal(lsBytes, &lsSearchSelectors)

	pvpTeamBytes, _ := css.Asset("search/pvpteam.json")
	pvpTeamSearchSelectors := PVPTeamSearchSelectors{}
	json.Unmarshal(pvpTeamBytes, &pvpTeamSearchSelectors)

	return &SearchSelectors{
		Character:   &charaSearchSelectors,
		CWLS:        &cwlsSearchSelectors,
		FreeCompany: &fcSearchSelectors,
		Linkshell:   &lsSearchSelectors,
		PVPTeam:     &pvpTeamSearchSelectors,
	}
}
