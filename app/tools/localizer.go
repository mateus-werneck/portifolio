package tools

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	pt "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

var (
	Bundle     *i18n.Bundle
	Translator ut.Translator
)

func init() {
	Bundle = NewLanguageBundle()
	SetPtBrTransaltor()
}

func SetEnTransalator() {
	pt := pt.New()
	en := en.New()

	uni := ut.New(en, en, pt)
	trans, _ := uni.GetTranslator("en_US")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	Translator = trans
}

func SetPtBrTransaltor() {
	pt := pt.New()
	en := en.New()

	uni := ut.New(pt, pt, en)
	trans, _ := uni.GetTranslator("pt_BR")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR.RegisterDefaultTranslations(v, trans)
	}

	Translator = trans
}

func NewLanguageBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.BrazilianPortuguese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	bundle.MustLoadMessageFile("translations/en.toml")
	bundle.MustLoadMessageFile("translations/pt-BR.toml")

	return bundle
}
