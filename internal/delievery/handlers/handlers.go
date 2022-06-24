package handlers

import (
	herrors "ascii-art-web/internal/delievery/errors"
	"ascii-art-web/internal/entity"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//We need Generator in handlers to generate ascii-art inside them
type Handler struct {
	generator entity.AsciiArtGenerator
}

//This struct is needed to be passed to template
type Result struct {
	Fonts []string
	Art   string
}

//Initializtion of handler
func NewHandler(generator entity.AsciiArtGenerator) *Handler {
	return &Handler{generator: generator}
}

//This handler is working with main page at /
func (H *Handler) AsciiArtMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("template/*.html")
	if err != nil {
		herrors.Error(http.StatusInternalServerError, err, w)
		return
	}
	if r.Method != "GET" {
		herrors.Error(http.StatusMethodNotAllowed, errors.New("invalid method"), w)
		return
	}
	if r.RequestURI != "/" {
		herrors.Error(http.StatusNotFound, errors.New("not found"), w)
		return
	}
	fontlist := H.generator.GetFontList()
	var output = Result{Fonts: fontlist, Art: ""}
	tmpl.ExecuteTemplate(w, "index.html", &output)
}

//This is the api handler, it returns ascii-art in plaintext based on POST-request
func (H *Handler) AsciiAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		herrors.Error(http.StatusMethodNotAllowed, errors.New("invalid method"), w)
		return
	}
	err := r.ParseForm()
	if err != nil {
		herrors.Error(http.StatusBadRequest, err, w)
		return
	}
	input := r.FormValue("input")
	font := r.FormValue("font")
	wd := r.FormValue("width")
	width, err := strconv.Atoi(wd)
	if err != nil {
		herrors.Error(http.StatusInternalServerError, err, w)
	}
	err = H.generator.CheckUserInput(input, font, width)
	if err != nil {
		log.Print(err)
		herrors.Error(http.StatusBadRequest, err, w)
		return
	}
	res := H.generator.Generate(input, font, width)
	w.Write([]byte(res))
}

//This handler returns the page with ascii-art, when you press send on main page, if you don't have javascipt enables, you will template from this handler
func (H *Handler) AsciiArtReadyPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		herrors.Error(http.StatusMethodNotAllowed, errors.New("invalid method"), w)
		return
	}
	tmpl, err := template.ParseGlob("template/*.html")
	if err != nil {
		herrors.Error(http.StatusInternalServerError, err, w)
		return
	}
	err = r.ParseForm()
	if err != nil {
		herrors.Error(http.StatusBadRequest, err, w)
		return
	}
	input := r.FormValue("input")
	font := r.FormValue("font")
	wd := r.FormValue("width")
	width, err := strconv.Atoi(wd)
	if err != nil {
		herrors.Error(http.StatusInternalServerError, err, w)
		return
	}
	err = H.generator.CheckUserInput(input, font, width)
	if err != nil {
		log.Print(err)
		herrors.Error(http.StatusBadRequest, err, w)
		return
	}
	res := H.generator.Generate(input, font, width)
	fontlist := H.generator.GetFontList()
	var output = Result{Fonts: fontlist, Art: res}
	tmpl.ExecuteTemplate(w, "index.html", &output)
}
