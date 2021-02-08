package infrastructure

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/atbys/koremiyo/interfaces/scraper"
)

type Scraper struct {
	baseURL   string
	targetURL string
	Doc       *goquery.Document
}

func NewScraper() *Scraper {
	scraper := &Scraper{
		baseURL: "https://filmarks.com",
	}
	return scraper
}

func (scraper Scraper) GetPage(url string) (scraper.Document, error) {
	result, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	doc := new(ScrapedSelection)
	doc.Selection = result.Selection
	return doc, err
}

type ScrapedSelection struct {
	Selection *goquery.Selection
}

func (ssel ScrapedSelection) Find(sel string) scraper.Selection {
	resultSel := ssel.Selection.Find(sel)
	returnSel := new(ScrapedSelection)
	returnSel.Selection = resultSel
	//*goquery.Find()は*goquery.Selectionを返すため，このメソッドの返り値の型と合わない
	//scraper.Selectionに合わせたScrapedSelectionの型にして返している
	return returnSel
}

func (ssel ScrapedSelection) Text() string {
	return ssel.Selection.Text()
}
