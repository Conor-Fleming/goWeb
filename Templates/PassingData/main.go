package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dataTemp.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "dataTemp.gohtml", "Conor Fleming passed in HERE")
	if err != nil {
		log.Fatalln(err)
	}
}
