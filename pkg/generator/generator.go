package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

type Survey struct {
	SourceFile  string
	Name        string      `yaml:"name"`
	PackageName string      `yaml:"packageName"`
	Questions   []*Question `yaml:"questions"`
	AllRequired bool        `yaml:"allRequired"`
}

type Question struct {
	Name         string `yaml:"name"`
	InputType    string `yaml:"type"`
	QuestionType string
	AnswerType   string   `yaml:"answer"`
	Help         string   `yaml:"help"`
	Required     bool     `yaml:"required"`
	Prompt       string   `yaml:"prompt"`
	Default      string   `yaml:"default"`
	Options      []string `yaml:"options"`
}

var isSurveyTmpl = regexp.MustCompile(".*_survey.ya?ml$")

type Generator struct {
	tpl *template.Template
}

func New() *Generator {
	tpl := template.Must(template.New("survey.tmpl").
		Funcs(sprig.GenericFuncMap()).
		Parse(SurveyTmpl))

	return &Generator{
		tpl,
	}
}

func (g *Generator) Generate(root string) error {
	total := 0

	err := filepath.Walk(root, func(srcFile string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !isSurveyTmpl.MatchString(info.Name()) {
			return nil
		}

		total += 1
		filename := strings.Replace(path.Base(srcFile), "_survey.yaml", "", 1)
		dirname := path.Dir(srcFile)

		fmt.Printf("Found survey %s... \n", srcFile)
		s, err := g.readSurvey(srcFile)
		if err != nil {
			return err
		}

		p := g.getPackageDir(dirname, s)

		return g.writeSurvey(p, filename, s)
	})
	if err != nil {
		return err
	}
	if total == 0 {
		return fmt.Errorf("no survey found on '%s' to generate code", root)
	}
	return nil
}

func (g *Generator) getPackageDir(dir string, s *Survey) string {
	if s.PackageName == "" {
		s.PackageName = path.Base(dir)
		return dir
	}
	pdir := path.Join(dir, s.PackageName)
	os.MkdirAll(pdir, 0755)

	return pdir
}

func (g *Generator) writeSurvey(dst, filename string, survey *Survey) error {
	var buf bytes.Buffer
	if err := g.tpl.Execute(&buf, survey); err != nil {
		return err
	}

	// Pass gofmt to output
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	filename = path.Join(dst, fmt.Sprintf("%s_survey.go", path.Clean(filename)))
	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = out.Write(p)
	if err != nil {
		return err
	}

	return err
}

func (g *Generator) readSurvey(srcPath string) (*Survey, error) {
	var surveyData *Survey

	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &surveyData)
	if err != nil {
		log.Fatal(err)
	}
	surveyData.SourceFile = srcPath

	for _, q := range surveyData.Questions {
		switch q.InputType {
		case "confirm":
			q.QuestionType = "bool"
		case "select":
			q.QuestionType = "string"
		case "input":
			q.QuestionType = "string"
		case "password":
			q.QuestionType = "string"
		case "multiline":
			q.QuestionType = "string"
		case "multiselect":
			q.QuestionType = "[]string"
		default:
			q.QuestionType = "string"
		}
	}

	err = validateSurvey(surveyData)
	if err != nil {
		return nil, err
	}
	return surveyData, nil
}

func validateSurvey(data *Survey) error {
	// TODO: validate survey content
	return nil
}

/*
type QuestionType string

const (
	Multiline = QuestionType("multiline")
	Confirm   = QuestionType("confirm")
	Select    = QuestionType("select")
	Input     = QuestionType("input")
)

func mainOld() {
	var answers struct {
		Q1     bool
		Q2     bool
		Q3     int
		Option string
	}
	qs, err := GenerateQuestions("test.yaml")

	err = survey.Ask(qs, &answers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answers)
}

func GenerateQuestions(ymlFile string) ([]*survey.Question, error) {
	var yamlData Yaml
	var qs []*survey.Question

	file, err := ioutil.ReadFile(ymlFile)
	if err != nil {
		return nil, fmt.Errorf("error reading questions file: %v", err)
	}

	err = yaml.Unmarshal(file, &yamlData)
	if err != nil {
		return nil, fmt.Errorf("error parsing questions: %v", err)
	}

	for _, question := range yamlData.Questions {
		var prompt survey.Prompt
		name := question.Name

		if name == "" {
			return nil, fmt.Errorf("all questions should have valid name")
		}

		switch question.Type {
		case Input:
			prompt = &survey.Input{
				Message: question.Prompt,
			}
		case Confirm:
			prompt = &survey.Confirm{
				Message: question.Prompt,
			}
		case Select:
			prompt = &survey.Select{
				Message: question.Prompt,
				Options: question.Options,
			}
		case Multiline:
			prompt = &survey.Multiline{
				Message: question.Prompt,
			}

		default:
			return nil, fmt.Errorf("unknown type: '%s' prompt: '%s", question.Type, question.Prompt)
		}

		surveyQuestion := &survey.Question{
			Name:   name,
			Prompt: prompt,
		}

		if question.Required || yamlData.AllRequired {
			surveyQuestion.Validate = survey.Required
		}

		qs = append(qs, surveyQuestion)

	}

	return qs, nil
}
*/
