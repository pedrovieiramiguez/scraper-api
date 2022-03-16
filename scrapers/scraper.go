package scrapers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"log"
	"scraper-api/data"
)

type Scraper struct {
	product       string
	parseFunction parseFunction
}

type parseFunction interface {
	parse(s *Scraper) data.Products
}

func InitScraper(product string, pf parseFunction) *Scraper {
	return &Scraper{
		product:       product,
		parseFunction: pf,
	}
}

func (s *Scraper) SetParseFunction(pf parseFunction) {
	s.parseFunction = pf
}

func (s *Scraper) SetProduct(product string) {
	s.product = product
}

func (s *Scraper) Scrape() data.Products {
	return s.parseFunction.parse(s)
}

type PingodoceParser struct{}

func (pds *PingodoceParser) parse(s *Scraper) data.Products {
	url := "https://mercadao.pt/store/pingo-doce/search?queries=" + s.product
	var products data.Products
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div._35DP2XbY0vnDR6ntQlSXMJ>pdo-product-item.P9eg53AkHYfXRP7gt5njS").Each(func(index int, item *goquery.Selection) {
				log.Println("Found an item!")
				n := item.Find("h3.pdo-heading-s.detail-title").Text()
				p := item.Find("span.pdo-inline-block.pdo-middle").Text()
				log.Printf("Name: %v\n", n)
				log.Printf("Price: %v\n", p)
				products = append(products, &data.Product{
					Name:  n,
					Price: p,
				})
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	return products
}
