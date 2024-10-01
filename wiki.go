package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	dir := "templates/"
	t, _ := template.ParseFiles(dir + tmpl + ".html")
	err := t.Execute(w, p)
	if err != nil {
		log.Fatalln(err)
	}
}

// this function loads urls prefixed with the /view/ pattern
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// this function allows us to create and edit pages
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
