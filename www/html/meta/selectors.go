package meta

const (
	TitleSelector       = "html head title"
	DescriptionSelector = "meta[name=\"description\" i][content]"
	IconSelector        = "link[rel=\"icon\" i], link[rel=\"shortcut icon\" i]"
	KeywordsSelector    = "meta[name=\"keywords\" i][content]"
	RobotsSelector      = "meta[name=\"robots\" i][content]"
	RatingSelector      = "meta[name=\"rating\" i][content]"
	CanonicalSelector   = "link[rel=\"canonical\" i][href]"
	AlternateSelector   = "link[rel=\"alternate\" i]"
)
