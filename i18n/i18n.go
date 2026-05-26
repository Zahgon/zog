package i18n

import (
	"github.com/Oudwins/zog/zconst"
)

const (
	// Default lang key used to get the language from the ParseContext
	LangKey = "lang"
)

// Takes a map[langKey]conf.LangMap
// usage is i18n.SetLanguagesErrsMap(map[string]zconst.LangMap{
// "es": es.Map, "en": en.Map, "ja": ja.Map,
// }, "en", i18n.WithLangKey("langKey"))
// schema.Parse(data, &dest, z.WithCtxValue("langKey", "es"))
func SetLanguagesErrsMap(m map[string]zconst.LangMap, defaultLang string, opts ...setLanguageOption) {
	_ = "STUB: not implemented"
	return
}

// use default lang if failed to get correct language map

// Override the default lang key used to get the language from the ParseContext
func WithLangKey(key string) setLanguageOption {
	_ = "STUB: not implemented"
	return *new(setLanguageOption)
}

// Please use the helper function this type may very well change in the future but the helper function's API will stay the same
type setLanguageOption = func(langKey *string)

// Proxy the type for easy use
type LangMap = zconst.LangMap
