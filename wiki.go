package main

import (
	"fmt"
	"html/template"
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
	subdirectory := "wikipages/"
	filename := title + ".txt"
	body, err := os.ReadFile(subdirectory + filename)
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

// this function allows us to create and edit pages
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	dir := "templates/"
	log.Print(dir + "edit.html")
	t, _ := template.ParseFiles(dir + "edit.html")
	err = t.Execute(w, p)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
