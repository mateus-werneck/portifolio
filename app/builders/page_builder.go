package builders

import "github.com/nicksnyder/go-i18n/v2/i18n"

type PageBuilder interface {
	SetTitle(title string) PageBuilder
	SetLanguage(language string) PageBuilder
	SetLocalizer(localizer *i18n.Localizer) PageBuilder
	Build() interface{}
}
