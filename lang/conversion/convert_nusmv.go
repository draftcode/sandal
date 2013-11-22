package lang

import (
	"github.com/draftcode/sandal/lang/data"
)

type converter struct {
	number int
}

type DefinitionClosure struct {
	def data.Definition
	env *DefEnv
}

type DefEnv struct {
	upper *DefEnv
	scope map[string]DefinitionClosure
}

func NewDefEnvWithUpper(env *DefEnv) *DefEnv {
	return &DefEnv{upper: env, scope: make(map[string]DefinitionClosure)}
}

func ConvertNuSMV(defs []data.Definition) string {
	conv := new(converter)
	return conv.ConvertNuSMV(defs)
}

func (conv *converter) ConvertNuSMV(defs []data.Definition) (converted string) {
	env := NewDefEnvWithUpper(nil)

	// Collect definitions exposed in toplevel.
	for _, def := range defs {
		switch d := def.(type) {
		case *data.DataDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *data.ModuleDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *data.ConstantDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *data.ProcDefinition:
			env.scope[d.Name] = DefinitionClosure{d, env}
		case *data.InitBlock:
			// Do nothing
		}
	}

	// Convert initblock.
	for _, def := range defs {
		if d, isInitBlock := def.(*data.InitBlock); isInitBlock {
			converted += conv.ConvertInitBlock(d, env)
		}
	}
	return
}

func (conv *converter) ConvertInitBlock(def *data.InitBlock, env *DefEnv) (converted string) {
	env = NewDefEnvWithUpper(env)
	return
}
