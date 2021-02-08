package scraper

type Scraper interface {
	GetPage(string) (Document, error)
}

type Document interface {
	Selection
}

type Selection interface {
	//Each(func(int, Selection)) Selection
	Find(string) Selection
	Text() string
}
