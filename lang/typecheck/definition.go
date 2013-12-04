package typecheck

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

// ========================================
// typeCheckDefinition

func typeCheckDefinitions(defs []Definition, env *typeEnv) error {
	// Put all definitions to the env first. Module and toplevel definition
	// has a scope that can see all names within the block.
	for _, def := range defs {
		switch def := def.(type) {
		case DataDefinition:
			namedType := NamedType{Name: def.Name}
			for _, elem := range def.Elems {
				env.add(elem, namedType)
			}
		case ModuleDefinition:
			params := make([]Type, len(def.Parameters))
			for _, p := range def.Parameters {
				params = append(params, p.Type)
			}
			env.add(def.Name, CallableType{Parameters: params})
		case ConstantDefinition:
			env.add(def.Name, def.Type)
		case ProcDefinition:
			params := make([]Type, len(def.Parameters))
			for _, p := range def.Parameters {
				params = append(params, p.Type)
			}
			env.add(def.Name, CallableType{Parameters: params})
		case InitBlock:
			// Do nothing
		default:
			panic("Unknown definition type")
		}
	}

	for _, def := range defs {
		if err := typeCheckDefinition(def, env); err != nil {
			return err
		}
	}
	return nil
}

func typeCheckDefinition(x Definition, env *typeEnv) error {
	switch x := x.(type) {
	case DataDefinition:
		return typeCheckDataDefinition(x, env)
	case ModuleDefinition:
		return typeCheckModuleDefinition(x, env)
	case ConstantDefinition:
		return typeCheckConstantDefinition(x, env)
	case ProcDefinition:
		return typeCheckProcDefinition(x, env)
	case InitBlock:
		return typeCheckInitBlock(x, env)
	}
	panic("Unknown Definition")
}

func typeCheckDataDefinition(def DataDefinition, env *typeEnv) error {
	return nil
}

func typeCheckModuleDefinition(def ModuleDefinition, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	for _, def := range def.Definitions {
		if err := typeCheckDefinition(def, env); err != nil {
			return err
		}
	}
	return nil
}

func typeCheckConstantDefinition(def ConstantDefinition, env *typeEnv) error {
	if err := typeCheckExpression(def.Expr, env); err != nil {
		return err
	}
	actual := typeOfExpression(def.Expr, env)
	if !actual.Equal(def.Type) {
		return fmt.Errorf("Expect %+#v to have type %+#v but has %+#v",
			def.Expr, def.Type, actual)
	}
	return nil
}

func typeCheckProcDefinition(def ProcDefinition, env *typeEnv) error {
	procEnv := newTypeEnvFromUpper(env)
	for _, stmt := range def.Statements {
		if err := typeCheckStatement(stmt, procEnv); err != nil {
			return err
		}
		switch s := stmt.(type) {
		case ConstantDefinition:
			env.add(s.Name, s.Type)
		case VarDeclStatement:
			env.add(s.Name, s.Type)
		}
	}
	return nil
}

func typeCheckInitBlock(b InitBlock, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	names := make(map[string]bool)
	for _, initVar := range b.Vars {
		if _, defined := names[initVar.VarName()]; defined {
			return fmt.Errorf("Varname %s is duplicated", initVar.VarName())
		}
		names[initVar.VarName()] = true

		switch initVar := initVar.(type) {
		case ChannelVar:
			env.add(initVar.Name, initVar.Type)
		case InstanceVar:
			calleeType := env.lookup(initVar.ProcDefName)
			if calleeType == nil {
				return fmt.Errorf("%q should be a callable type", initVar.ProcDefName)
			}
			env.add(initVar.Name, calleeType)
		default:
			panic("Unknown initvar type")
		}
	}

	for _, initVar := range b.Vars {
		switch initVar := initVar.(type) {
		case ChannelVar:
			switch initVar.Type.(type) {
			case HandshakeChannelType, BufferedChannelType:
				// OK
			default:
				return fmt.Errorf("%s should be a channel", initVar.Name)
			}
		case InstanceVar:
			calleeType := env.lookup(initVar.ProcDefName)
			if t, isCallableType := calleeType.(CallableType); isCallableType {
				if len(t.Parameters) != len(initVar.Args) {
					return fmt.Errorf("Argument count mismatch")
				}
				for i := 0; i < len(t.Parameters); i++ {
					if err := typeCheckExpression(initVar.Args[i], env); err != nil {
						return err
					}
					argType := typeOfExpression(initVar.Args[i], env)
					if !argType.Equal(t.Parameters[i]) {
						return fmt.Errorf("Argument type mismatch")
					}
				}
			} else {
				return fmt.Errorf("%q should be a callable type", initVar.ProcDefName)
			}
		default:
			panic("Unknown initvar type")
		}
	}
	return nil
}
