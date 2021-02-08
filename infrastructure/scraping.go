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
	doc := new(GoquerySelection)
	doc.Selection = result.Selection
	return doc, err
}

type GoquerySelection struct {
	Selection *goquery.Selection
}

func (ssel GoquerySelection) Find(sel string) scraper.Selection {
	resultSel := ssel.Selection.Find(sel)
	returnSel := new(GoquerySelection)
	returnSel.Selection = resultSel
	//*goquery.Find()は*goquery.Selectionを返すため，このメソッドの返り値の型と合わない
	//scraper.Selectionに合わせたScrapedSelectionの型にして返している
	return returnSel
}

func (gsel GoquerySelection) Text() string {
	return gsel.Selection.Text()
}

func (gsel GoquerySelection) FindAll(sel_name string) []scraper.Selection {
	var sels []scraper.Selection
	base := gsel.Selection.Find(sel_name)
	base.Each(func(i int, sel *goquery.Selection) {
		tmp := &GoquerySelection{Selection: sel}
		sels = append(sels, tmp)
	})

	return sels
}

func (gsel GoquerySelection) Attr(attrName string) (string, bool) {
	return gsel.Selection.Attr(attrName)
}
