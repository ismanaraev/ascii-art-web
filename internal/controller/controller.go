package controller

import (
	"ascii-art-web/internal/entity"
	"ascii-art-web/internal/usecase"
	"encoding/json"
	"os"
	"strconv"
)

//This structs is used only for reading json config file
type fontList struct {
	Fontname string `json:"fontname"`
	File     string `json:"file"`
	Linenum  string `json:"linenum"`
	Charset  string `json:"charset"`
	Height   string `json:"height"`
	Gaps     string `json:"gaps"`
}

type config struct {
	Port  string     `json:"port"`
	Fonts []fontList `json:"fonts"`
}

//This function reads config file, sets the port as environment variable, initializes generator and fills it with fonts
func ReadConfigFile(file string) (entity.AsciiArtGenerator, error) {
	cfg, err := readJSON(file)
	if err != nil {
		return nil, err
	}
	port := os.Getenv("ASCII_WEB_PORT")
	if port == "" {
		err = os.Setenv("ASCII_WEB_PORT", cfg.Port)
		if err != nil {
			return nil, err
		}
	}
	var generator entity.AsciiArtGenerator = usecase.NewGenerator()
	for _, font := range cfg.Fonts {
		lines, err := strconv.Atoi(font.Linenum)
		if err != nil {
			return nil, err
		}
		height, err := strconv.Atoi(font.Height)
		if err != nil {
			return nil, err
		}
		gap, err := strconv.Atoi(font.Gaps)
		if err != nil {
			return nil, err
		}
		asciiFont, err := usecase.NewAsciiFont(font.Fontname, font.File, font.Charset, lines, height, gap)
		if err != nil {
			return nil, err
		}
		generator.AddFont(*asciiFont)
	}
	return generator, nil
}

//This function reads JSON config file
func readJSON(file string) (*config, error) {
	configFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
