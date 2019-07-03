package routes

import (
	"io"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "hello")
}