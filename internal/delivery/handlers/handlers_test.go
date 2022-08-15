package handlers

import (
	"ascii-art-web/internal/entity"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockG struct {
	fonts []string
}

func (g mockG) GetFontList() []string {
	return g.fonts
}

func (g mockG) CheckUserInput(input, font string, width int) error {
	if input == "" {
		return errors.New("empty str")
	}
	if font != "doom" {
		return errors.New("invalid font")
	}
	if width <= 0 {
		return errors.New("invalid width")
	}
	return nil
}

func (g mockG) Generate(input, font string, width int) string {
	return "Hello"
}

func (g mockG) AddFont(entity.AsciiArtFont) {
}

type mockGenerator interface {
	AddFont(entity.AsciiArtFont)
	GetFontList() []string
	CheckUserInput(input, font string, width int) error
	Generate(input, font string, width int) string
}

type getHandlerTestCase struct {
	method   string
	address  string
	expected int
}

var handlerTestCases = []getHandlerTestCase{
	{
		method:   "GET",
		address:  "/",
		expected: 200,
	},
	{
		method:   "ABOBA",
		address:  "/",
		expected: 405,
	},
	{
		method:   "GET",
		address:  "/f",
		expected: 404,
	},
}

func TestAsciiArtMainPage(t *testing.T) {
	for _, testcase := range handlerTestCases {
		m := new(mockG)
		var g mockGenerator = m
		h := NewHandler(g)
		rr := httptest.NewRecorder()
		request := httptest.NewRequest(testcase.method, testcase.address, nil)
		handler := http.HandlerFunc(h.AsciiArtMainPage)
		handler.ServeHTTP(rr, request)
		if rr.Code != testcase.expected {
			t.Fail()
		}
	}
}

type postAPITestCase struct {
	method   string
	address  string
	body     string
	expected int
}

var postAPITestCases = []postAPITestCase{
	{
		method:   "POST",
		address:  "/api",
		body:     "input=test&font=doom&width=500",
		expected: 200,
	},
	{
		method:   "POST",
		address:  "/ap",
		body:     "input=test&font=doom&width=500",
		expected: 404,
	},
	{
		method:   "POST",
		address:  "/api",
		body:     "input=test&font=dm&width=500",
		expected: 400,
	},
	{
		method:   "BREW",
		address:  "/api",
		body:     "input=test&font=doom&width=500",
		expected: 405,
	},
	{
		method:   "POST",
		address:  "/api",
		body:     "input=test&font=doom&width=aboba",
		expected: 500,
	},
}

func TestAsciiAPI(t *testing.T) {
	for _, testcase := range postAPITestCases {
		m := new(mockG)
		var g mockGenerator = m
		h := NewHandler(g)
		rr := httptest.NewRecorder()
		reader := strings.NewReader(testcase.body)
		request := httptest.NewRequest(testcase.method, testcase.address, reader)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		handler := http.HandlerFunc(h.AsciiAPI)
		handler.ServeHTTP(rr, request)
		if rr.Code != testcase.expected {
			t.Fail()
		}
	}
}

var postPageTestCases = []postAPITestCase{
	{
		method:   "POST",
		address:  "/ascii-art",
		body:     "input=test&font=doom&width=500",
		expected: 200,
	},
	{
		method:   "POST",
		address:  "/ap",
		body:     "input=test&font=doom&width=500",
		expected: 404,
	},
	{
		method:   "POST",
		address:  "/ascii-art",
		body:     "input=test&font=dm&width=500",
		expected: 400,
	},
	{
		method:   "BREW",
		address:  "/ascii-art",
		body:     "input=test&font=doom&width=500",
		expected: 405,
	},
	{
		method:   "POST",
		address:  "/ascii-art",
		body:     "input=test&font=doom&width=aboba",
		expected: 500,
	},
}

func TestAsciiReadyPage(t *testing.T) {
	for _, testcase := range postAPITestCases {
		m := new(mockG)
		var g mockGenerator = m
		h := NewHandler(g)
		rr := httptest.NewRecorder()
		reader := strings.NewReader(testcase.body)
		request := httptest.NewRequest(testcase.method, testcase.address, reader)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		handler := http.HandlerFunc(h.AsciiAPI)
		handler.ServeHTTP(rr, request)
		if rr.Code != testcase.expected {
			t.Fail()
		}
	}
}
