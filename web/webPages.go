package Web

import (
	"net/http"
	"strings"
	"text/template"

	"web/lib"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	var bnStyle, inputStr string
	var tmpl *template.Template

	// Label any requests other than 'GET' and 'POST' requests as 'invalid requests'
	if !(r.Method == http.MethodGet || r.Method == http.MethodPost) {
		w.WriteHeader(http.StatusMethodNotAllowed)

		tmpl = template.Must(template.ParseFiles("static/errorPrinter.html"))
		tmpl.Execute(w, struct {
			Code  int
			Issue string
		}{Issue: "Invalid Request Method!", Code: http.StatusMethodNotAllowed})
		return

		// Serve form at initial visit of site
	} else if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			tmpl = template.Must(template.ParseFiles("static/errorPrinter.html"))
			tmpl.Execute(w, struct {
				Issue string
				Code  int
			}{Issue: "404: Page not found", Code: http.StatusNotFound})
			return
		}
		tmpl = template.Must(template.ParseFiles("static/placeHolder.html"))
		tmpl.Execute(w, nil)

		// Serve form and ascii-art/error after form submission
	} else if r.Method == http.MethodPost {

		// Extract banner style selected and text inputed in form
		bnStyle = r.FormValue("style")
		inputStr = r.FormValue("inputStr")

		// Run AsciiArt function with banner style selected and input string
		output, err := lib.AsciiArt(inputStr, bnStyle+".txt")

		// Should there occur an error, serve errorPrinter.html with the nature of error
		if err != "" {
			tmpl = template.Must(template.ParseFiles("static/errorPrinter.html"))

			if strings.Contains(err, "PRINTABLE ASCII") {
				w.WriteHeader(http.StatusBadRequest)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusBadRequest})
			}
			if strings.Contains(err, "Error reading") {
				w.WriteHeader(http.StatusNotFound)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusNotFound})
			}
			if strings.Contains(err, "modified") {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusInternalServerError})
			}

			// If no error print ascii-art below form on submitForm.html
		} else {
			// Safely load html template from submitForm.html
			tmpl = template.Must(template.ParseFiles("static/submitForm.html"))
			tmpl.Execute(w, struct{ AsciiArt, Input string }{AsciiArt: output, Input: inputStr})
		}
	}
}
