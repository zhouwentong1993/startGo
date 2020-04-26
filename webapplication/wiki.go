package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	p := &Page{Title: "test", Body: []byte("hello world中文")}
	err := p.save()
	if err != nil {
		return
	} else {
		page, err := loadPage("test")
		if err != nil {
			return
		} else {
			println(page.Body)
		}
	}

	//http.HandleFunc("/", handle)
	http.HandleFunc("/view/", viewPageHandler)
	http.HandleFunc("/edit/", editPageHandler)
	http.HandleFunc("/save/", savePageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filePath := page.Title + ".txt"
	return ioutil.WriteFile(filePath, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filePath := title + ".txt"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	} else {
		return &Page{Body: file, Title: title}, nil
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, I love %s!", r.URL.Path[1:])
}

func viewPageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	page, err := loadPage(path)
	if err != nil {
		http.Redirect(w, r, "/edit/"+path, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func renderTemplate(w http.ResponseWriter, path string, page *Page) {
	files, err := template.ParseFiles(path + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		err = files.Execute(w, page)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editPageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/edit/"):]
	page, _ := loadPage(path)
	renderTemplate(w, "edit", page)
}

func savePageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	newPage := &Page{Title: path, Body: []byte(body)}
	err := newPage.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+path, http.StatusFound)
}
