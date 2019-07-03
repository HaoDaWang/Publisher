package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func SendTpl(res http.ResponseWriter, filename string, data ...interface{}){
	var tplData interface{}

	if len(data) == 0 {
		tplData = ""
	} else {
		tplData = data[0]
	}

	var tplPath = fmt.Sprintf("templates/%s.html", filename)

	t := template.Must(template.ParseFiles(tplPath))

	t.ExecuteTemplate(res, filename, tplData)
}