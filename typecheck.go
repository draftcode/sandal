package sandal

import (
	"fmt"
)

func TypeCheck(defs []Definition, env *TypeEnv) error {
	// Put all definitions to the env first. Module and toplevel definition
	// has a scope that can see all names within the block.
	for _, def := range defs {
		switch def := def.(type) {
		case *DataDefinition:
			namedType := &NamedType{Name: def.Name}
			for _, elem := range def.Elems {
				env.Add(elem, namedType)
			}
		case *ModuleDefinition:
			params := make([]Type, len(def.Parameters))
			for _, p := range def.Parameters {
				params = append(params, p.Type)
			}
			env.Add(def.Name, &CallableType{Parameters: params})
		case *ConstantDefinition:
			env.Add(def.Name, def.Type)
		case *ProcDefinition:
			params := make([]Type, len(def.Parameters))
			for _, p := range def.Parameters {
				params = append(params, p.Type)
			}
			env.Add(def.Name, &CallableType{Parameters: params})
		case *InitBlock:
			// Do nothing
		default:
			panic("Unknown definition type")
		}
	}

	for _, def := range defs {
		if err := def.typecheck(env); err != nil {
			return err
		}
	}
	return nil
}

// ========================================
// Typecheck of definitions

func (def *DataDefinition) typecheck(env *TypeEnv) error {
	return nil
}

func (def *ModuleDefinition) typecheck(env *TypeEnv) error {
	return TypeCheck(def.Definitions, NewTypeEnvFromUpper(env))
}

func (def *ConstantDefinition) typecheck(env *TypeEnv) error {
	if err := def.Expr.typecheck(env); err != nil {
		return err
	}
	actual := def.Expr.type_()
	if !actual.equal(def.Type) {
		return fmt.Errorf("Expect %+#v to have type %+#v but has %+#v",
			def.Expr, def.Type, actual)
	}
	return nil
}

func (def *ProcDefinition) typecheck(env *TypeEnv) error {
	// procEnv := NewTypeEnvFromUpper(env)
	// for _, stmt := range def.Statements {
	// 	if err := stmt.typecheck(procEnv); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (b *InitBlock) typecheck(env *TypeEnv) error {
	// TODO
	return nil
}

// ========================================
// Typecheck of definitions

func (x *IdentifierExpression) type_() Type {
	return nil
}

func (x *NumberExpression) type_() Type {
	return &NamedType{Name: "int"}
}

func (x *NotExpression) type_() Type {
	return x.SubExpr.type_()
}

func (x *UnarySubExpression) type_() Type {
	return nil
}

func (x *ParenExpression) type_() Type {
	return nil
}

func (x *BinOpExpression) type_() Type {
	return nil
}

func (x *TimeoutRecvExpression) type_() Type {
	return nil
}

func (x *TimeoutPeekExpression) type_() Type {
	return nil
}

func (x *NonblockRecvExpression) type_() Type {
	return nil
}

func (x *NonblockPeekExpression) type_() Type {
	return nil
}

func (x *ArrayExpression) type_() Type {
	return nil
}

func (x *IdentifierExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *NumberExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *NotExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *UnarySubExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *ParenExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *BinOpExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *TimeoutRecvExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *TimeoutPeekExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *NonblockRecvExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *NonblockPeekExpression) typecheck(env *TypeEnv) error {
	return nil
}

func (x *ArrayExpression) typecheck(env *TypeEnv) error {
	return nil
}
