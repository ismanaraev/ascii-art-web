package usecase

import (
	"log"
	"os"
	"testing"
)

type asciiFontTestCase struct {
	fontname string
	file     string
	charset  string
	lines    int
	height   int
	gap      int
}

const fontnum = 4

var asciiFontTestCases = []asciiFontTestCase{
	{
		fontname: "standard",
		file:     "../../ascii/standard.txt",
		charset:  " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		lines:    855,
		height:   8,
		gap:      1,
	},
	{
		fontname: "shadow",
		file:     "../../ascii/shadow.txt",
		charset:  " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		lines:    855,
		height:   8,
		gap:      1,
	},
	{
		fontname: "thinkertoy",
		file:     "../../ascii/thinkertoy.txt",
		charset:  " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		lines:    855,
		height:   8,
		gap:      1,
	},
	{
		fontname: "doom",
		file:     "../../ascii/doom.txt",
		charset:  " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		lines:    855,
		height:   8,
		gap:      1,
	},
}

var testgenerator = NewGenerator()

func TestNewAsciiFont(t *testing.T) {
	for _, testcase := range asciiFontTestCases {
		font, err := NewAsciiFont(testcase.fontname, testcase.file, testcase.charset, testcase.lines, testcase.height, testcase.gap)
		if err != nil {
			t.Fail()
			log.Fatal(err)
		}
		testcase.charset = testcase.charset + "\n"
		for _, ch := range testcase.charset {
			if lines, ok := font.Charmap[string(ch)]; ok {
				if len(lines) != testcase.height {
					t.Fail()
					log.Print("height and len(lines) doesn't match")
				}
			} else {
				t.Fail()
				log.Print("invalid charmap")
			}
		}
		testgenerator.AddFont(*font)
	}
}

type testUserInput struct {
	input string
	font  string
	width int
	file  string
}

var testUserInputCases = []testUserInput{
	{
		input: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		font:  "doom",
		width: 280,
		file:  "ascii_for_tests/allchars_doom.txt",
	},
	{
		input: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		font:  "standard",
		width: 280,
		file:  "ascii_for_tests/allchars_standard.txt",
	},
	{
		input: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		font:  "thinkertoy",
		width: 280,
		file:  "ascii_for_tests/allchars_thinkertoy.txt",
	},
	{
		input: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		font:  "shadow",
		width: 280,
		file:  "ascii_for_tests/allchars_shadow.txt",
	},
	{
		input: "\n\n\n\n",
		font:  "doom",
		width: 280,
		file:  "ascii_for_tests/gaps.txt",
	},
	{
		input: "\n\n\n\n",
		font:  "standard",
		width: 280,
		file:  "ascii_for_tests/gaps.txt",
	},
	{
		input: "\n\n\n\n",
		font:  "shadow",
		width: 280,
		file:  "ascii_for_tests/gaps.txt",
	},
	{
		input: "\n\n\n\n",
		font:  "thinkertoy",
		width: 280,
		file:  "ascii_for_tests/gaps.txt",
	},
}

func TestAsciiGenerator(t *testing.T) {
	fontslice := testgenerator.GetFontList()
	if len(fontslice) != fontnum {
		t.Fail()
		log.Print("not enough fonts")
		return
	}
	for _, testcase := range testUserInputCases {
		err := testgenerator.CheckUserInput(testcase.input, testcase.font, testcase.width)
		if err != nil {
			t.Fail()
			log.Print(err)
			break
		}
		res := testgenerator.Generate(testcase.input, testcase.font, testcase.width)
		expected, err := os.ReadFile(testcase.file)
		if err != nil {
			t.Fail()
			log.Print(err)
			break
		}
		if res != string(expected) {
			t.Fail()
			log.Printf("string and expected string doesn't match:\n got:\n%s\n expected:\n%s\n", res, string(expected))
			break
		}
	}
}
