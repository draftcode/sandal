package conversion

import (
	"bytes"
	"strings"
	"text/template"
)

type tmplModule struct {
	Name    string
	Args    []string
	Vars    []tmplVar
	Assigns []tmplAssign
	Defs    []tmplAssign
}

type tmplVar struct {
	Name string
	Type string
}

type tmplAssign struct {
	LHS string
	RHS string
}

const moduleTemplate = `
MODULE {{.Name}}({{args .Args}}){{if .Vars}}
  VAR{{range .Vars}}
    {{.Name}} : {{.Type}};{{end}}{{end}}{{if .Assigns}}
  ASSIGN{{range .Assigns}}
    {{.LHS}} :={{rhsFormat .RHS}};{{end}}{{end}}{{if .Defs}}
  DEFINE{{range .Defs}}
    {{.LHS}} :={{rhsFormat .RHS}};{{end}}{{end}}
`

func rhsFormat(rhs string) string {
	if strings.ContainsRune(rhs, '\n') {
		return "\n      " + strings.Join(strings.Split(rhs, "\n"), "\n      ")
	} else {
		return " " + rhs
	}
}

var funcMap template.FuncMap = template.FuncMap{
	"args":      argJoin,
	"rhsFormat": rhsFormat,
}

func instantiateTemplate(module tmplModule) string {
	tmpl, err := template.New("NuSMVModule").Funcs(funcMap).Parse(moduleTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, module)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
