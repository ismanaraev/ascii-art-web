package usecase

import (
	"ascii-art-web/internal/entity"
	"errors"
	"os"
	"strings"
)

//This struct is needed to form a font fields from input file, input file is read on outer layers
type asciifont struct {
	fontname string
	file     string
	charset  string
	charmap  map[string][]string
	lines    int
	height   int
	gap      int
}

//This function forms a new ascii-font based on input
func NewAsciiFont(fontname string, file string, charset string, lines, height, gap int) (*entity.AsciiArtFont, error) {
	tmp := &asciifont{fontname: fontname, file: file, charset: charset, lines: lines, height: height, gap: gap}
	charslice, err := tmp.readFile(tmp.file)
	if err != nil {
		return nil, err
	}
	tmp.createFontCharmap(charslice)
	res := &entity.AsciiArtFont{FontName: tmp.fontname, Charset: tmp.charset, Charmap: tmp.charmap, Height: tmp.height}
	return res, nil
}

func (a *asciifont) readFile(file string) ([]string, error) {
	res, err := os.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	splitchars := strings.Split(string(res), "\n")
	if len(splitchars)-1 != a.lines {
		return []string{}, errors.New("not enough lines")
	}
	return splitchars, nil
}

//This function creates a map from contents of fontfile split by newline, map is used for printing the symbols and checking the input string
func (a *asciifont) createFontCharmap(allchars []string) {
	tmpmap := make(map[string][]string)
	var ct int
	for i := (a.height + a.gap); i < len(allchars); i = i + (a.height + a.gap) {
		tmpmap[string(a.charset[ct])] = allchars[i-a.height : i]
		ct++
	}
	tmpmap["\n"] = []string{"", "", "", "", "", "", "", ""}
	a.charmap = tmpmap
}
