package builders

import (
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type HomePageData struct {
	Title          string
	ContactButton  string
	ChangeLanguage string
	LanguageName   string
	LanguageFlag   string
	RecentWork     map[string]types.RecentWork
}

type HomePageBuilder struct {
	Title     string
	Language  string
	Localizer *i18n.Localizer
}

func NewHomePage() PageBuilder {
	return &HomePageBuilder{}
}

func (b *HomePageBuilder) SetTitle(title string) PageBuilder {
	b.Title = title
	return b
}

func (b *HomePageBuilder) SetLanguage(language string) PageBuilder {
	b.Language = language
	return b
}

func (b *HomePageBuilder) SetLocalizer(localizer *i18n.Localizer) PageBuilder {
	b.Localizer = localizer
	return b
}

func (b *HomePageBuilder) Build() interface{} {
	var data HomePageData

	data.Title = b.Title

	data.ContactButton = b.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "ContactButton",
	})

	data.ChangeLanguage = "en"
	data.LanguageName = "Inglês"
	data.LanguageFlag = "/static/images/us.svg"

	if b.Language == "en" {
		data.ChangeLanguage = "pt-br"
		data.LanguageName = "Português"
		data.LanguageFlag = "/static/images/br.svg"
	}

	data.RecentWork = types.RecentWorks()

	return data
}
