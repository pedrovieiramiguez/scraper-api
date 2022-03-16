package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scraper-api/scrapers"
)

// Pingodoce is a http.Handler
type Pingodoce struct {
	l *log.Logger
}

// NewPingodoce creates a pingodoce handler with the given logger
func NewPingodoce(l *log.Logger) *Pingodoce {
	return &Pingodoce{l}
}

// GetProducts scrapes the pingo doce website for a list of products
func (p *Pingodoce) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Pingodoce")

	vars := mux.Vars(r)
	product := vars["product"] //need to validate

	// create scraper
	parser := &scrapers.PingodoceParser{}
	scraper := scrapers.InitScraper(product, parser)

	products := scraper.Scrape()

	// serialize the list to JSON
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
