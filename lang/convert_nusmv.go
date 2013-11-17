package lang

type converter struct {
	number int
}

type DefinitionClosure struct {
	def Definition
	env *DefEnv
}

type DefEnv struct {
	upper *DefEnv
	scope map[string]DefinitionClosure
}

func NewDefEnvWithUpper(env *DefEnv) *DefEnv {
	return &DefEnv{upper: env, scope: make(map[string]DefinitionClosure)}
}

func ConvertNuSMV(defs []Definition) string {
	conv := new(converter)
	return conv.ConvertNuSMV(defs)
}

func (conv *converter) ConvertNuSMV(defs []Definition) (converted string) {
	env := NewDefEnvWithUpper(nil)

	// Collect definitions exposed in toplevel.
	for _, def := range defs {
		switch d := def.(type) {
		case *DataDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *ModuleDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *ConstantDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *ProcDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *InitBlock:
			// Do nothing
		}
	}

	// Convert initblock.
	for _, def := range defs {
		if d, isInitBlock := def.(*InitBlock); isInitBlock {
			converted += conv.ConvertInitBlock(d, env)
		}
	}
	return
}

func (conv *converter) ConvertInitBlock(def *InitBlock, env *DefEnv) (converted string) {
	env = NewDefEnvWithUpper(env)
	return
}
