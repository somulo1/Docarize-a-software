package Web

import (
	"fmt"
	"net/http"
	"text/template"

	"web/Lib"
)

// func FormHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "static/submitForm.html")
// }

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	var bnStyle, inputStr string
	fmt.Println(r.Method)
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("static/submitForm.html"))
		tmpl.Execute(w, "")
	} else if r.Method == http.MethodPost {

		bnStyle = r.FormValue("style")
		inputStr = r.FormValue("inputStr")

		// fmt.Printf("%q\n", inputStr)

		output := Lib.AsciiArt(inputStr, bnStyle+".txt")

		tmpl := template.Must(template.ParseFiles("static/submitForm.html"))
		tmpl.Execute(w, struct{ AsciiArt string }{AsciiArt: output})
	} else {
		http.Error(w, "Invalid Request Method!", http.StatusMethodNotAllowed)
	}
}
