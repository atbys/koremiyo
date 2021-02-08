package scraper

type Scraper interface {
	GetPage(string) (Document, error)
}

type Document interface {
	Selection
}

type Selection interface {
	Find(string) Selection
	FindAll(string) []Selection
	Text() string
	Attr(string) (string, bool)
}
