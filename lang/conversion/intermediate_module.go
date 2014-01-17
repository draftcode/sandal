package conversion

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
		Trans     []intTransition
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
		FromState  intState
		NextState  intState
		Condition  string
		Actions    []intAssign
	}

	intAssign struct {
		LHS string
		RHS string
	}
)

// ========================================

type varEnv struct {
	upper   *varEnv
	mapping map[string]intInternalObj
}

func newVarEnv() (ret *varEnv) {
	ret = new(varEnv)
	ret.mapping = make(map[string]intInternalObj)
	return
}

func newVarEnvFromUpper(upper *varEnv) (ret *varEnv) {
	ret = newVarEnv()
	ret.upper = upper
	return
}

func (env *varEnv) add(name string, intVar intInternalObj) {
	env.mapping[name] = intVar
}

func (env *varEnv) lookup(name string) intInternalObj {
	if intVar, found := env.mapping[name]; found {
		return intVar
	}
	if env.upper != nil {
		return env.upper.lookup(name)
	} else {
		return nil
	}
}
