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
	Greeting       string
	GreetingTwo    string
	ParagraphOne   string
	ParagraphTwo   string
	ParagraphThree string
}

type HomePageButtons struct {
	ContactMe  string
	DownloadCv string
	LearnMore  string
	Visit      string
	Proposal   string
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

type Backend struct {
	Title         string
	Description   string
	FirstSection  string
	SecondSection string
	ThirdSection  string
}

type Frontend struct {
	Title         string
	Description   string
	FirstSection  string
	SecondSection string
	ThirdSection  string
}

type RecentJobs struct {
	Title       string
	Description string
	Jobs        map[string]types.RecentWork
}

type ShowInterest struct {
	Title       string
	Description string
}

type Footer struct {
	Message   string
	CreatedBy string
	Copyright string
}

type HomePageData struct {
	Title            string
	LanguageSettings UserLanguage
	Intro            HomePageIntro
	Buttons          HomePageButtons
	Summary          HomePageSummary
	TechLead         TechLead
	Backend          Backend
	Frontend         Frontend
	RecentJobs       RecentJobs
	ShowInterest     ShowInterest
	Footer           Footer
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
			ChangeLanguage: "pt-BR",
			LanguageName:   "Inglês",
			LanguageFlag:   "/static/images/us.svg",
		},
		Summary:  HomePageSummary{},
		TechLead: TechLead{},
		Buttons:  HomePageButtons{},
	}

	if b.Language == "pt-BR" {
		data.LanguageSettings.ChangeLanguage = "en-US"
		data.LanguageSettings.LanguageName = "Português"
		data.LanguageSettings.LanguageFlag = "/static/images/br.svg"
	}

	data.Title = b.Title

	data.Intro.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.Title"})
	data.Intro.SubTitle = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.SubTitle"})
	data.Intro.SubTitleTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Intro.SubTitleTwo"})

	data.Summary.Greeting = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.Greeting"})
	data.Summary.GreetingTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.GreetingTwo"})
	data.Summary.ParagraphOne = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.ParagraphOne"})
	data.Summary.ParagraphTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.ParagraphTwo"})
	data.Summary.ParagraphThree = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Summary.ParagraphThree"})

	data.TechLead.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.Title"})
	data.TechLead.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.Desc"})
	data.TechLead.FirstSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.FirstSection"})
	data.TechLead.SecondSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SecondSection"})
	data.TechLead.SkillOne = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillOne"})
	data.TechLead.SkillTwo = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillTwo"})
	data.TechLead.SkillThree = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillThree"})
	data.TechLead.SkillFour = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillFour"})
	data.TechLead.SkillFive = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TechLead.SkillFive"})

	data.Backend.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Backend.Title"})
	data.Backend.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Backend.Desc"})
	data.Backend.FirstSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Backend.FirstSection"})
	data.Backend.SecondSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Backend.SecondSection"})
	data.Backend.ThirdSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Backend.ThirdSection"})

	data.Frontend.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Frontend.Title"})
	data.Frontend.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Frontend.Desc"})
	data.Frontend.FirstSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Frontend.FirstSection"})
	data.Frontend.SecondSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Frontend.SecondSection"})
	data.Frontend.ThirdSection = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Frontend.ThirdSection"})

	data.RecentJobs.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "RecentJobs.Title"})
	data.RecentJobs.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "RecentJobs.Desc"})
	data.RecentJobs.Jobs = types.RecentWorks()

	data.ShowInterest.Title = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ShowInterest.Title"})
	data.ShowInterest.Description = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ShowInterest.Desc"})

	data.Footer.Message = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Footer.Message"})
	data.Footer.CreatedBy = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Footer.CreatedBy"})
	data.Footer.Copyright = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Footer.Copyright"})

	data.Buttons.ContactMe = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.ContactMe"})
	data.Buttons.DownloadCv = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.DownloadCv"})
	data.Buttons.LearnMore = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.LearnMore"})
	data.Buttons.Visit = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.Visit"})
	data.Buttons.Proposal = b.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.Proposal"})

	return data
}
