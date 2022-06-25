package usecase

import (
	"ascii-art-web/internal/entity"
	"errors"
	"strings"
)

//This struct is needed as struct that implements Methods of AsciiArtGenerator
type generator struct {
	Fonts map[string]entity.AsciiArtFont
}

//This function is needed for struct initialization
func NewGenerator() *generator {
	return &generator{}
}

//This function returns list of available fonts for handlers to draw that list at template
func (g generator) GetFontList() []string {
	var res []string
	for font := range g.Fonts {
		res = append(res, font)
	}
	return res
}

//This function adds new fonts to generator
func (g *generator) AddFont(font entity.AsciiArtFont) {
	if g.Fonts == nil {
		g.Fonts = make(map[string]entity.AsciiArtFont)
		g.Fonts[font.FontName] = font
		return
	}
	g.Fonts[font.FontName] = font
}

//This function validates user input
func (g generator) CheckUserInput(input, font string, width int) error {
	if _, ok := g.Fonts[font]; !ok {
		return errors.New("invalid font")
	}
	for _, ch := range input {
		if _, ok := g.Fonts[font].Charmap[string(ch)]; !ok {
			return errors.New("invalid input")
		}
	}
	if width <= 0 {
		return errors.New("invalid width")
	}
	return nil
}

//This is the core function, it hyphenates input using hyphenate, then forms a string with resulting ascii-art from input
func (g generator) Generate(input, font string, width int) string {
	sb := strings.Builder{}
	lines := g.hyphenate(input, font, width)
	for _, line := range lines {
		if line == "" {
			sb.WriteByte(10) //10 is \n
			continue
		}
		for i := 0; i < g.Fonts[font].Height; i++ {
			for _, ch := range line {
				sb.WriteString(g.Fonts[font].Charmap[string(ch)][i])
			}
			sb.WriteByte(10)
		}
	}
	res := sb.String()
	return res
}

//This function hyphenates string if it doesn't fit the screen with set width (width in symbols, as in terminal)
func (g *generator) hyphenate(input, font string, width int) []string {
	var wordlen int
	var word string
	var newinputlines []string
	s := strings.Split(input, "\n")
	tw := width
	charmap := g.Fonts[font].Charmap
	for _, item := range s {
		if item == "" {
			newinputlines = append(newinputlines, "")
		}
		for _, ch := range item {
			if l := wordlen + len(charmap[string(ch)][0]); l <= tw {
				wordlen += len(charmap[string(ch)][0])
				word += string(ch)
			} else {
				newinputlines = append(newinputlines, word)
				wordlen = len(charmap[string(ch)][0])
				word = string(ch)
			}
		}
		wordlen = 0
		if word != "" {
			newinputlines = append(newinputlines, word)
			word = ""
		}
	}
	return newinputlines
}
