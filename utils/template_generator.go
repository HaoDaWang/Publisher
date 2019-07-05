package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func SendTpl(res http.ResponseWriter, filenames []string, funcs template.FuncMap, data ...interface{}){
	var tplData interface{}

	if len(data) == 0 {
		tplData = ""
	} else {
		tplData = data[0]
	}

	var tplPaths []string

	for _, name := range filenames {
		tplPaths = append(tplPaths, fmt.Sprintf("templates/%s.html", name))
	}

	//tpl, err := template.New(tplPath).Funcs(template.FuncMap{
	//	"say": func() int {
	//		fmt.Println("say")
	//		return 1
	//	},
	//}).ParseFiles(tplPaths...)

	tpl := template.Must(template.New("index").Funcs(funcs).ParseFiles(tplPaths...))

	//if err != nil {
	//	log.Fatal(err)
	//}

	tpl.ExecuteTemplate(res, "index", tplData)
}