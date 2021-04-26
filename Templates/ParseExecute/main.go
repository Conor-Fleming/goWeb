package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") //parsing given file and saving to container 'tpl'
	if err != nil {
		log.Fatalln(err)
	}

	//parsing more files into container
	tpl, err = tpl.ParseFiles("anotherone.gohtml", "anotherx2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //after parsing html template execute uses writer stdout to print to terminal
	//you could also use Create and write to a file
	if err != nil {
		log.Fatalln(err)

	}

}
