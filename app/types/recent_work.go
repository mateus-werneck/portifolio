package types

import "github.com/nicksnyder/go-i18n/v2/i18n"

type RecentWork struct {
	Element     string
	Image       string
	Description string
	Website     string
	Opacity     string
}

var works = map[string]RecentWork{
	"celcoin": {
		Element:     "celcoin",
		Image:       "celcoin.svg",
		Description: "Infratech financeira para potencializar negócios",
		Website:     "https://www.celcoin.com.br",
		Opacity:     "opacity-100",
	},
	"symplicity": {
		Element:     "symplicity",
		Image:       "symplicity.webp",
		Description: "Streamline system-wide opportunities and increase student engagement",
		Website:     "https://www.symplicity.com",
		Opacity:     "opacity-100",
	},
}

func RecentWorks() map[string]RecentWork {
	return works
}

func FindWork(name string) RecentWork {
	return works[name]
}

func (w *RecentWork) Desc(localizer *i18n.Localizer) string {
	messageId := "RecentJobs." + w.Element
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageId})
}
