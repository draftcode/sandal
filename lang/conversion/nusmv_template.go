package lang

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

// ========================================
// AbstractModule

type absNuSMVModule struct {
	Name      string
	Args      []string
	Vars      []absVar
	InitState absState
	Trans     map[absState][]absTransition
	Defaults  map[string]string
	Defs      []absAssign
}

type absState string

type absVar struct {
	Name string
	Type string
}

type absTransition struct {
	Condition string
	Actions   map[absState][]absAssign
}

type absAssign struct {
	LHS string
	RHS string
}

const caseTemplate = `case{{range .Cases}}
  {{.Condition}} : {{.Value}}{{end}}
  TRUE : {{.Default}}
esac;`

type caseTmplValue struct {
	Cases []struct {
		Condition string
		Value     string
	}
	Default string
}

// AssignCond holds assignment condition to the variables.
type AssignCond struct {
	cond map[string][]struct {
		condition string
		value     string
	}
}

func NewAssignCond() *AssignCond {
	return &AssignCond{make(map[string][]struct {
		condition string
		value     string
	})}
}

func (cond *AssignCond) Add(variable, condition, value string) {
	cond.cond[variable] = append(cond.cond[variable], struct {
		condition string
		value     string
	}{condition, value})
}

func ConvertAbstractModuleToTemplate(module absNuSMVModule) (tmpl tmplNuSMVModule, err error) {
	tmpl.Name = module.Name
	tmpl.Args = append([]string{"running_pid", "pid"}, module.Args...)
	tmpl.Vars = []tmplVar{
		{"state", "{" + strings.Join(extractStates(module), ", ") + "}"},
	}
	for _, absvar := range module.Vars {
		tmpl.Vars = append(tmpl.Vars, tmplVar{absvar.Name, absvar.Type})
	}
	assignCond := make(map[string]map[string]string)
	for state, transes := range module.Trans {
		for _, trans := range transes {
			extractAssignCondition(state, trans, assignCond)
		}
	}
	for variable, cases := range assignCond {
		var defaultValue string
		if variable == "next(state)" {
			defaultValue = "state"
		} else if defaultValue, hasValue := module.Defaults[variable]; !hasValue {
			return tmplNuSMVModule{}, fmt.Errorf("There is no default value for %s", variable)
		}
	}
	return
}

func extractStates(module absNuSMVModule) (states []string) {
	states_map := make(map[absState]bool)
	states_map[module.InitState] = true
	for s, transes := range module.Trans {
		states_map[s] = true
		for _, trans := range transes {
			for t, _ := range trans.Actions {
				states_map[t] = true
			}
		}
	}

	for state, _ := range states_map {
		states = append(states, string(state))
	}
	sort.StringSlice(states).Sort()
	return
}

func extractAssignCondition(state absState, trans absTransition, assignCond map[string][]string) {
}

// ========================================
// Template

type tmplNuSMVModule struct {
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

func args(args []string) string {
	return strings.Join(args, ", ")
}

func rhsFormat(rhs string) string {
	if strings.ContainsRune(rhs, '\n') {
		return "\n      " + strings.Join(strings.Split(rhs, "\n"), "\n      ")
	} else {
		return " " + rhs
	}
}

var funcMap template.FuncMap = template.FuncMap{
	"args":      args,
	"rhsFormat": rhsFormat,
}

func InstantiateTemplate(module tmplNuSMVModule) string {
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
