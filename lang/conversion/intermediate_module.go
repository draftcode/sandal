package conversion

import (
	. "github.com/draftcode/sandal/lang/data"
)

type intModule interface {
	intmodule()
}

type (
	intMainModule struct {
		Vars    []intVar
		Assigns []intAssign
		Defs    []intAssign
	}

	intHandshakeChannel struct {
		Name      string
		ValueType []string
	}

	intBufferedChannel struct {
		Name      string
		Length    int
		ValueType []string
	}

	intProcModule struct {
		Name      string
		Args      []string
		Vars      []intVar
		InitState intState
		Trans     map[intState][]intTransition
		Defaults  map[string]string
		Defs      []intAssign
	}
)

func (x intMainModule) intmodule()       {}
func (x intHandshakeChannel) intmodule() {}
func (x intBufferedChannel) intmodule()  {}
func (x intProcModule) intmodule()       {}

type (
	intState string

	intVar struct {
		Name string
		Type string
	}

	intTransition struct {
		Condition string
		Actions   map[intState][]intAssign
	}

	intAssign struct {
		LHS string
		RHS string
	}
)

type (
	intInternalVal interface {
		intinternalval()
	}

	intInternalChannelVal struct {
		Name       string
		ModuleName string
	}

	intInternalProcVal struct {
		Name       string
		ModuleName string
		Def        Definition
		Args       []string
		Pid        int
	}

	intInternalPrimitiveVar struct {
		Type Type
	}

	intInternalProcDef struct {
		Def Definition
	}

	intInternalConstant struct {
		Type Type
		Expr Expression
	}
)

func (x intInternalChannelVal) intinternalval()   {}
func (x intInternalProcDef) intinternalval()      {}
func (x intInternalProcVal) intinternalval()      {}
func (x intInternalPrimitiveVar) intinternalval() {}
func (x intInternalConstant) intinternalval()     {}

// ========================================

type varEnv struct {
	upper   *varEnv
	mapping map[string]intInternalVal
}

func newVarEnv() (ret *varEnv) {
	ret = new(varEnv)
	ret.mapping = make(map[string]intInternalVal)
	return
}

func newVarEnvFromUpper(upper *varEnv) (ret *varEnv) {
	ret = newVarEnv()
	ret.upper = upper
	return
}

func (env *varEnv) add(name string, intVar intInternalVal) {
	env.mapping[name] = intVar
}

func (env *varEnv) lookup(name string) intInternalVal {
	if intVar, found := env.mapping[name]; found {
		return intVar
	}
	if env.upper != nil {
		return env.upper.lookup(name)
	} else {
		return nil
	}
}

// ========================================

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
