package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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
	http.HandleFunc("/view/", makeHandler(viewPageHandler))
	http.HandleFunc("/edit/", makeHandler(editPageHandler))
	http.HandleFunc("/save/", makeHandler(savePageHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

func renderTemplate(w http.ResponseWriter, path string, page *Page) {
	err := templates.ExecuteTemplate(w, path+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewPageHandler(w http.ResponseWriter, r *http.Request, path string) {
	page, err1 := loadPage(path)
	if err1 != nil {
		http.Redirect(w, r, "/edit/"+path, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editPageHandler(w http.ResponseWriter, r *http.Request, path string) {
	page, err := loadPage(path)
	if err != nil {
		page = &Page{
			Title: path,
		}
	}
	renderTemplate(w, "edit", page)
}

func savePageHandler(w http.ResponseWriter, r *http.Request, path string) {

	body := r.FormValue("body")
	newPage := &Page{Title: path, Body: []byte(body)}
	err1 := newPage.save()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+path, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		m := validPath.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(writer, request)
			return
		} else {
			fn(writer, request, m[2])
		}
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		return "", errors.New("not found")
	} else {
		return m[2], nil
	}
}
