package conversion

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

func convetASTToIntModule(defs []Definition) (ret []intModule) {
	converter := intModConverter{}
	for _, def := range defs {
		switch def := def.(type) {
		case *InitBlock:
			converter.convertInitBlock(def)
		default:
			panic("Not implemented")
		}
	}
	return
}

// ========================================

type varEnv struct {
	upper *varEnv
	// Variable name to NuSMV level variable name
	mapping map[string]string
}

func newVarEnv() (ret *varEnv) {
	ret = new(varEnv)
	ret.mapping = make(map[string]string)
	return
}

func newVarEnvFromUpper(upper *varEnv) (ret *varEnv) {
	ret = newVarEnv()
	ret.upper = upper
	return
}

func (env *varEnv) add(name, nusmvName string) {
	env.mapping[name] = nusmvName
}

func (env *varEnv) lookup(name string) string {
	if nusmvName, found := env.mapping[name]; found {
		return nusmvName
	}
	if env.upper != nil {
		return env.upper.lookup(name)
	} else {
		return ""
	}
}

// ========================================

type intModConverter struct {
	env *varEnv
}

func newIntModConverter() (converter *intModConverter) {
	converter = new(intModConverter)
	converter.env = newVarEnv()
	return
}

func (x *intModConverter) convertInitBlock(def *InitBlock) {
}

// ========================================

// intModule represents intermediate module between Sandal and NuSMV.
type intModule interface {
	intmodule()
}

type intProcModule struct {
	Name      string
	Args      []string
	Vars      []intVar
	InitState intState
	Trans     map[intState][]intTransition
	Defaults  map[string]string
	Defs      []intAssign
}

func (x intProcModule) intmodule() {}

type intState string

type intVar struct {
	Name string
	Type string
}

type intTransition struct {
	Condition string
	Actions   map[intState][]intAssign
}

type intAssign struct {
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

type assignCond struct {
	state     string
	condition string
	value     string
}

// AssignCond holds assignment condition to the variables.
type AssignCond struct {
	cond map[string][]struct {
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

func extractStates(module intModule) (states []string) {
	states_map := make(map[intState]bool)
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

func extractAssignCondition(state intState, trans intTransition, assignCond map[string][]string) {
}
