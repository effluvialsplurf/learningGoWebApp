package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/* this stuff was all for the initial commit of the tutorial
* it includes writing to a text file and outputting that file on the cmd line
 */

// this is the data structure of a single page
type Page struct {
	Title string
	Body  []byte
}

// this saves pages (writes them to a text file)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// loads that page (reads the data from the text file)
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// this function loads urls prefixed with a pattern
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
