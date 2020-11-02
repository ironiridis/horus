package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/renderInvoice", renderInvoice)

}

func renderInvoice(w http.ResponseWriter, r *http.Request) {
	var err error
	tmpl, err := template.ParseFiles("invoice.html")
	must("parse invoice template", err)

	fp, err := os.Open(r.FormValue("fn"))
	must("open invoice json", err)
	defer fp.Close()

	dec := json.NewDecoder(fp)
	var inv Invoice
	err = dec.Decode(&inv)
	must("decode json into invoice variable", err)

	err = tmpl.Execute(w, &inv)
	must("execute template against invoice variable", err)

}
