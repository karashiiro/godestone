package godestone

// SiteLang represents the scraped website language.
type SiteLang string

// Lodestone website language. Do note that all five language-versions of the website (`eu` not listed) are on
// the same physical servers in Japan. Changing the language of the website will not meaningfully improve
// response times.
const (
	JA SiteLang = "jp"
	EN SiteLang = "na"
	FR SiteLang = "fr"
	DE SiteLang = "de"
)

// SearchLang represents character language for the purposes of search.
type SearchLang uint8

// Language
const (
	NoneLang SearchLang = 1 << iota
	SearchJA
	SearchEN
	SearchDE
	SearchFR
)
