package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	inputFile    = ""
	outputFile   = ""
	templateFile = ""
)

type Theme struct {
	Metadata             Metadata         `yaml:"metadata"`
	ClockMode            ClockModeSection `yaml:"clock-mode"`
	DisplayPanes         Section          `yaml:"display-panes"`
	DisplayPanesActive   Section          `yaml:"display-panes-active"`
	Message              Section          `yaml:"message"`
	MessageCommand       Section          `yaml:"message-command"`
	Mode                 Section          `yaml:"mode"`
	PaneBorder           Section          `yaml:"pane-border"`
	PaneActiveBorder     Section          `yaml:"pane-active-border"`
	Status               Section          `yaml:"status"`
	StatusLeft           Section          `yaml:"status-left"`
	StatusRight          Section          `yaml:"status-right"`
	WindowStatus         Section          `yaml:"window-status"`
	WindowStatusCurrent  Section          `yaml:"window-status-current"`
	WindowStatusActivity Section          `yaml:"window-status-activity"`
}

type Metadata struct {
	Header string `yaml:"header"`
}

type Section struct {
	Colour    string `yaml:"colour"`
	Format    string `yaml:"format"`
	Interval  int    `yaml:"interval"`
	Justify   string `yaml:"justify"`
	Length    int    `yaml:"length"`
	Separator string `yaml:"separator"`
	Style     Style  `yaml:"style"`
}

type ClockModeSection struct {
	Section
	Style int `yaml:"style"`
}

type Style struct {
	Fg string `yaml:"fg"`
	Bg string `yaml:"bg"`
}

func loadThemeFile(filename string) (*Theme, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	theme := &Theme{}

	err = yaml.Unmarshal(content, theme)
	if err != nil {
		return nil, err
	}

	return theme, nil
}

func buildTemplate(filename string) (*template.Template, error) {
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	return template.New("tmuxtheme").Parse(string(templateContent))
}

func init() {
	flag.StringVar(&inputFile, "i", "", "Input file to parse.")
	flag.StringVar(&templateFile, "t", "", "Template file to use.")
	flag.StringVar(&outputFile, "o", "",
		"Output file to write to, or STDOUT if not specified.")
}

func main() {
	flag.Parse()

	// fmt.Println("inputFile:", inputFile)
	// fmt.Println("outputFile:", outputFile)
	// fmt.Println("templateFile:", templateFile)

	theme, err := loadThemeFile(inputFile)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", theme)

	tpl, err := buildTemplate(templateFile)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, theme)
	if err != nil {
		panic(err)
	}
}
