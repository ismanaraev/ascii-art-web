package herrors

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Err struct {
	Error     string
	ErrorText string
}

func Error(status int, inputError error, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("template/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	log.Print(inputError)
	w.WriteHeader(status)
	fmt.Println(status, inputError.Error())
	errortemp := &Err{Error: http.StatusText(status), ErrorText: inputError.Error()}
	tmpl.Execute(w, &errortemp)
}
