package builders

import (
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type HomePageIntro struct {
	Title       string
	SubTitle    string
	SubTitleTwo string
}

type HomePageSummary struct {
	Greeting    string
	GreetingTwo string
	Paragraph   string
}

type HomePageButtons struct {
	ContactMe  string
	DownloadCv string
}

type UserLanguage struct {
	ChangeLanguage string
	LanguageName   string
	LanguageFlag   string
}

type TechLead struct {
	Title         string
	Description   string
	FirstSection  string
	SecondSection string
	SkillOne      string
	SkillTwo      string
	SkillThree    string
	SkillFour     string
	SkillFive     string
}

type HomePageData struct {
	Title            string
	LanguageSettings UserLanguage
	Intro            HomePageIntro
	Buttons          HomePageButtons
	Summary          HomePageSummary
	TechLead         TechLead
	RecentWork       map[string]types.RecentWork
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
	data := HomePageData{
		Intro: HomePageIntro{},
		LanguageSettings: UserLanguage{
			ChangeLanguage: "en",
			LanguageName:   "Inglês",
			LanguageFlag:   "/static/images/us.svg",
		},
		Summary:  HomePageSummary{},
		TechLead: TechLead{},
		Buttons:  HomePageButtons{},
	}

	data.Title = b.Title

	data.Intro.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.Title"})
	data.Intro.SubTitle = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.SubTitle"})
	data.Intro.SubTitleTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.SubTitleTwo"})

	if b.Language == "en" {
		data.LanguageSettings.ChangeLanguage = "pt-br"
		data.LanguageSettings.LanguageName = "Português"
		data.LanguageSettings.LanguageFlag = "/static/images/br.svg"
	}

	data.Summary.Greeting = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.Greeting"})
	data.Summary.GreetingTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.GreetingTwo"})
	data.Summary.Paragraph = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.Paragraph"})

	data.TechLead.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.Title"})
	data.TechLead.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.Description"})
	data.TechLead.FirstSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.FirstSection"})
	data.TechLead.SecondSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SecondSection"})
	data.TechLead.SkillOne = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillOne"})
	data.TechLead.SkillTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillTwo"})
	data.TechLead.SkillThree = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillThree"})
	data.TechLead.SkillFour = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillFour"})
	data.TechLead.SkillFive = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillFive"})

	data.RecentWork = types.RecentWorks()

	data.Buttons.ContactMe = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.DownloadCv"})
	data.Buttons.DownloadCv = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.ContactMe"})

	return data
}
