package conversion_deprecated

import (
	"bytes"
	"sort"
	"strings"
	"text/template"
)

type (
	tmplModule struct {
		Name     string
		Args     []string
		Vars     []tmplVar
		Trans    []string
		Assigns  []tmplAssign
		Defs     []tmplAssign
		LtlSpecs []string
		Justice  string
	}

	tmplVar struct {
		Name string
		Type string
	}

	tmplAssign struct {
		LHS string
		RHS string
	}

	tmplVars    []tmplVar
	tmplAssigns []tmplAssign
)

func (l tmplVars) Len() int              { return len(l) }
func (l tmplVars) Less(i, j int) bool    { return l[i].Name < l[j].Name }
func (l tmplVars) Swap(i, j int)         { l[i], l[j] = l[j], l[i] }
func (l tmplAssigns) Len() int           { return len(l) }
func (l tmplAssigns) Less(i, j int) bool { return l[i].LHS < l[j].LHS }
func (l tmplAssigns) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

const moduleTemplate = `
MODULE {{.Name}}({{args .Args}}){{if .Justice}}
  JUSTICE
    {{.Justice}}{{end}}{{if .Vars}}
  VAR{{range .Vars}}
    {{.Name}} : {{.Type}};{{end}}{{end}}{{if .Trans}}{{range .Trans}}
  TRANS {{.}};{{end}}{{end}}{{if .Assigns}}
  ASSIGN{{range .Assigns}}
    {{.LHS}} :={{rhsFormat .RHS}};{{end}}{{end}}{{if .Defs}}
  DEFINE{{range .Defs}}
    {{.LHS}} :={{rhsFormat .RHS}};{{end}}{{end}}{{if .LtlSpecs}}
  LTLSPEC{{range .LtlSpecs}}
    {{.}}{{end}}{{end}}
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
	sort.Sort(tmplVars(module.Vars))
	sort.Strings(module.Trans)
	sort.Sort(tmplAssigns(module.Assigns))
	sort.Sort(tmplAssigns(module.Defs))

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, module)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
