package locales

import (
	// Import these packages to trigger the init() function
	_ "github.com/NerdDoc/server/res/locales/ca"
	_ "github.com/NerdDoc/server/res/locales/de"
	_ "github.com/NerdDoc/server/res/locales/el"
	_ "github.com/NerdDoc/server/res/locales/en"
	_ "github.com/NerdDoc/server/res/locales/es"
	_ "github.com/NerdDoc/server/res/locales/fr"
	_ "github.com/NerdDoc/server/res/locales/it"
	_ "github.com/NerdDoc/server/res/locales/nl"
	_ "github.com/NerdDoc/server/res/locales/tr"
	_ "github.com/NerdDoc/server/res/locales/el"
	_ "github.com/NerdDoc/server/res/locales/ru"

)

// Locales is the list of locales's tags and names
// Please check if the language is supported in https://github.com/tebeka/snowball,
// if it is please add the correct language name.
var Locales = []Locale{
	{
		Tag:  "en",
		Name: "english",
	},
	{
		Tag:  "de",
		Name: "german",
	},
	{
		Tag:  "fr",
		Name: "french",
	},
	{
		Tag:  "es",
		Name: "spanish",
	},
	{
		Tag:  "ca",
		Name: "catalan",
	},
	{
		Tag:  "it",
		Name: "italian",
	},
	{
		Tag:  "tr",
		Name: "turkish",
	},
	{
		Tag:  "nl",
		Name: "dutch",
	},
	{
		Tag:  "el",
		Name: "greek",
	},
	{
		Tag:  "ru",
		Name: "russian",
	},
}

// A Locale is a registered locale in the file
type Locale struct {
	Tag  string
	Name string
}

// GetNameByTag returns the name of the given locale's tag
func GetNameByTag(tag string) string {
	for _, locale := range Locales {
		if locale.Tag != tag {
			continue
		}

		return locale.Name
	}

	return ""
}

// GetTagByName returns the tag of the given locale's name
func GetTagByName(name string) string {
	for _, locale := range Locales {
		if locale.Name != name {
			continue
		}

		return locale.Tag
	}

	return ""
}

// Exists checks if the given tag exists in the list of locales
func Exists(tag string) bool {
	for _, locale := range Locales {
		if locale.Tag == tag {
			return true
		}
	}

	return false
}
