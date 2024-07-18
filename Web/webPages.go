package Web

import (
	"net/http"
	"text/template"

	"web/Lib"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	var bnStyle, inputStr string
	var tmpl *template.Template
	if r.Method == http.MethodGet {
		tmpl = template.Must(template.ParseFiles("static/submitForm.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		bnStyle = r.FormValue("style")
		inputStr = r.FormValue("inputStr")

		//fmt.Printf("%q\n", inputStr)

		output, err := Lib.AsciiArt(inputStr, bnStyle+".txt")

		if err != "" {
			tmpl = template.Must(template.ParseFiles("static/errorPrinter.html"))
			tmpl.Execute(w, struct{ Issue string }{Issue: err})
		} else {
			//Safely load html template from submitForm.html
			tmpl = template.Must(template.ParseFiles("static/submitForm.html"))
			tmpl.Execute(w, struct{ AsciiArt string }{AsciiArt: output})
		}

	} else {
		http.Error(w, "Invalid Request Method!", http.StatusMethodNotAllowed)
	}
}
