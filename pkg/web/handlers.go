package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) uppyTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	buf := new(bytes.Buffer)
	tmpl := template.New("page")
	var err error
	if tmpl, err = tmpl.Parse(page); err != nil {
		fmt.Println(err)
	}

	// Execute the template set, passing in any dynamic data.
	if err := tmpl.Execute(buf, make(map[string]string)); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		fmt.Println(err)
		return
	}
}
