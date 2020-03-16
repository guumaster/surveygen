//go:generate surveygen --path demo

package main

import (
	"log"
	"os"
	"text/template"

	"github.com/AlecAivazis/survey/v2"

	"surveygen-example/demo"
)

func main() {
	s := demo.CreateMyAwesomeSurvey()

	err := survey.Ask(s.Questions, s.Answers)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("survey").Parse(`
Your answers:

  Age: {{ .Age }}
  Like: {{ if .Like }}Yes!{{ else }}meh...{{ end }}
  Level: {{ .Level }}
  Comment: {{ .Comment }}
  Countries choosen: 
    {{ range .Country -}}
    - {{ . }}
    {{- end }}
`)
	err = t.ExecuteTemplate(os.Stdout, "survey", s.Answers)
	if err != nil {
		log.Fatal(err)
	}
}
