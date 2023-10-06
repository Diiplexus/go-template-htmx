package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Year     string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Year: "1972", Director: "Francis Ford Coppola"},
				{Title: "The Shawshank Redemption", Year: "1994", Director: "Frank Darabont"},
				{Title: "Schindler's List", Year: "1993", Director: "Steven Spielberg"},
			},
		}
		template.Execute(w, films) //this is the line that writes to the browser
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		// log.Print(r.Header.Get("HX-Request"))//this is the line that writes to the terminal that the request received is an HX request
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		year := r.PostFormValue("year")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Year: year, Director: director})
	}

	http.HandleFunc("/add-film/", h2)
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
