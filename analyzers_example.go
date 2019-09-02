package main

import (
	"azuretextanalysis"
	"html/template"
	"net/http"
)

var (
	api       string = ""
	resource  string = ""
	documents        = []map[string]string{
		{"id": "1", "language": "en", "text": "This is a super cool test for my super cool package."},
		{"id": "2", "language": "en", "text": "This is a stupid test for my stupid package."},
		{"id": "3", "language": "en", "text": "The DoD is a very big operation compared to Banner Health."},
		{"id": "4", "language": "en", "text": "I really like CostCo chicken nuggets and German beer."},
	}
)

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html>
	<html>
	<body>
		<center>
			<h1>Analyze</h1>
			<h2>{{.}}</h2>
</center></body></html>
`))

func main() {

	// Run Entities at `localhost:8080/entities`
	http.HandleFunc("/entities", func(w http.ResponseWriter, r *http.Request) {
		data := azuretextanalysis.Entities(api, resource, documents)

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Run Phrases at `localhost:8080/phrases`
	http.HandleFunc("/phrases", func(w http.ResponseWriter, r *http.Request) {
		data := azuretextanalysis.Phrases(api, resource, documents)

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Run Language at `localhost:8080/language`
	http.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
		data := azuretextanalysis.Language(api, resource, documents)

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Run Sentiment at `localhost:8080/sentiment`
	http.HandleFunc("/sentiment", func(w http.ResponseWriter, r *http.Request) {
		data := azuretextanalysis.Sentiment(api, resource, documents)

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start listener
	http.ListenAndServe(":8080", nil)
}
