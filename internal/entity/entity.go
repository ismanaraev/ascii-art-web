package entity

//main entity struct for AsciiArt
type AsciiArtFont struct {
	FontName string
	Charset  string
	Charmap  map[string][]string
	Height   int
}

//This interface defines behavior of core function of the project. Implemetation are in usecases
type AsciiArtGenerator interface {
	AddFont(AsciiArtFont)
	GetFontList() []string
	CheckUserInput(input, font string, width int) error
	Generate(input, font string, width int) string
}
