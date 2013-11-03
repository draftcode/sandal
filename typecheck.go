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
	actual := def.Expr.type_(env)
	if !actual.equal(def.Type) {
		return fmt.Errorf("Expect %+#v to have type %+#v but has %+#v",
			def.Expr, def.Type, actual)
	}
	return nil
}

func (def *ProcDefinition) typecheck(env *TypeEnv) error {
	// TODO: Statement typecheck
	// procEnv := NewTypeEnvFromUpper(env)
	// for _, stmt := range def.Statements {
	// 	if err := stmt.typecheck(procEnv); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (b *InitBlock) typecheck(env *TypeEnv) error {
	// TODO: Statement typecheck
	return nil
}

// ========================================
// Typecheck of expression

func (x *IdentifierExpression) type_(env *TypeEnv) Type {
	return env.Lookup(x.Name)
}

func (x *NumberExpression) type_(env *TypeEnv) Type {
	return NamedType{Name: "int"}
}

func (x *NotExpression) type_(env *TypeEnv) Type {
	return x.SubExpr.type_(env)
}

func (x *UnarySubExpression) type_(env *TypeEnv) Type {
	return x.SubExpr.type_(env)
}

func (x *ParenExpression) type_(env *TypeEnv) Type {
	return x.SubExpr.type_(env)
}

var operatorResultType = map[int]Type{
	ADD:  NamedType{"int"},
	SUB:  NamedType{"int"},
	MUL:  NamedType{"int"},
	QUO:  NamedType{"int"},
	REM:  NamedType{"int"},
	AND:  NamedType{"int"},
	OR:   NamedType{"int"},
	XOR:  NamedType{"int"},
	SHL:  NamedType{"int"},
	SHR:  NamedType{"int"},
	LAND: NamedType{"bool"},
	LOR:  NamedType{"bool"},
	EQL:  NamedType{"bool"},
	LSS:  NamedType{"bool"},
	GTR:  NamedType{"bool"},
	NEQ:  NamedType{"bool"},
	LEQ:  NamedType{"bool"},
	GEQ:  NamedType{"bool"},
}

func (x *BinOpExpression) type_(env *TypeEnv) Type {
	if ty, exist := operatorResultType[x.Operator]; exist {
		return ty
	} else {
		panic("Unknown operator")
	}
}

func (x *TimeoutRecvExpression) type_(env *TypeEnv) Type {
	return NamedType{Name: "bool"}
}

func (x *TimeoutPeekExpression) type_(env *TypeEnv) Type {
	return NamedType{Name: "bool"}
}

func (x *NonblockRecvExpression) type_(env *TypeEnv) Type {
	return NamedType{Name: "bool"}
}

func (x *NonblockPeekExpression) type_(env *TypeEnv) Type {
	return NamedType{Name: "bool"}
}

func (x *ArrayExpression) type_(env *TypeEnv) Type {
	if len(x.Elems) == 0 {
		panic("An array should have at least one element")
	}
	// Every element of an array has the same type.
	return x.Elems[0].type_(env)
}

func (x *IdentifierExpression) typecheck(env *TypeEnv) error {
	if env.Lookup(x.Name) == nil {
		return fmt.Errorf("Undefined variable %s", x.Name)
	}
	return nil
}

func (x *NumberExpression) typecheck(env *TypeEnv) error {
	// Number expressions are always valid.
	return nil
}

func (x *NotExpression) typecheck(env *TypeEnv) error {
	if err := x.SubExpr.typecheck(env); err != nil {
		return err
	}
	if !x.SubExpr.type_(env).equal(NamedType{"bool"}) {
		return fmt.Errorf("Expect %s to have type bool, but got %s",
			x.SubExpr, x.SubExpr.type_(env))
	}
	return nil
}

func (x *UnarySubExpression) typecheck(env *TypeEnv) error {
	if err := x.SubExpr.typecheck(env); err != nil {
		return err
	}
	if !x.SubExpr.type_(env).equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to have type int, but got %s",
			x.SubExpr, x.SubExpr.type_(env))
	}
	return nil
}

func (x *ParenExpression) typecheck(env *TypeEnv) error {
	if err := x.SubExpr.typecheck(env); err != nil {
		return err
	}
	return nil
}

var operatorOperandType = map[int]Type{
	ADD:  NamedType{"int"},
	SUB:  NamedType{"int"},
	MUL:  NamedType{"int"},
	QUO:  NamedType{"int"},
	REM:  NamedType{"int"},
	AND:  NamedType{"int"},
	OR:   NamedType{"int"},
	XOR:  NamedType{"int"},
	SHL:  NamedType{"int"},
	SHR:  NamedType{"int"},
	LAND: NamedType{"bool"},
	LOR:  NamedType{"bool"},
	EQL:  nil,
	LSS:  NamedType{"int"},
	GTR:  NamedType{"int"},
	NEQ:  NamedType{"int"},
	LEQ:  NamedType{"int"},
	GEQ:  NamedType{"int"},
}

func (x *BinOpExpression) typecheck(env *TypeEnv) error {
	if err := x.LHS.typecheck(env); err != nil {
		return err
	}
	if err := x.RHS.typecheck(env); err != nil {
		return err
	}
	if ty, exist := operatorOperandType[x.Operator]; exist {
		if ty != nil {
			lhsType := x.LHS.type_(env)
			if !lhsType.equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s",
					x.LHS, ty, lhsType)
			}
			rhsType := x.RHS.type_(env)
			if !rhsType.equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s",
					x.RHS, ty, rhsType)
			}
		} else {
			lhsType := x.LHS.type_(env)
			rhsType := x.RHS.type_(env)
			if !lhsType.equal(rhsType) {
				return fmt.Errorf("Expect %s and %s to have the same type but got %s and %s",
					x.LHS, x.RHS, lhsType, rhsType)
			}
		}
	} else {
		panic("Unknown operator")
	}
	return nil
}

func channelRecvOrPeekCheck(chExpr ChanRecvExpr, env *TypeEnv) error {
	if err := chExpr.RecvChannel().typecheck(env); err != nil {
		return err
	}
	for _, arg := range chExpr.RecvArgs() {
		if err := arg.typecheck(env); err != nil {
			return err
		}
	}

	var elemTypes []Type
	switch ty := chExpr.RecvChannel().type_(env).(type) {
	case HandshakeChannelType:
		elemTypes = ty.Elems
	case BufferedChannelType:
		elemTypes = ty.Elems
	default:
		return fmt.Errorf("Expect the first argument of %s to be a channel but got %s",
			chExpr, chExpr.RecvChannel().type_(env))
	}

	if len(elemTypes) != len(chExpr.RecvArgs()) {
		return fmt.Errorf("Expect the arugments of %s to have %d elements",
			chExpr, len(elemTypes))
	}
	for i := 0; i < len(elemTypes); i++ {
		if !elemTypes[i].equal(chExpr.RecvArgs()[i].type_(env)) {
			return fmt.Errorf("Expect the argument %s to be a %s", chExpr.RecvArgs()[i], elemTypes[i])
		}
		if _, isIdentExpr := chExpr.RecvArgs()[i].(*IdentifierExpression); !isIdentExpr {
			return fmt.Errorf("Expect the argument %s to be an identifier", chExpr.RecvArgs()[i])
		}
	}
	return nil
}

func (x *TimeoutRecvExpression) typecheck(env *TypeEnv) error {
	return channelRecvOrPeekCheck(x, env)
}

func (x *TimeoutPeekExpression) typecheck(env *TypeEnv) error {
	return channelRecvOrPeekCheck(x, env)
}

func (x *NonblockRecvExpression) typecheck(env *TypeEnv) error {
	return channelRecvOrPeekCheck(x, env)
}

func (x *NonblockPeekExpression) typecheck(env *TypeEnv) error {
	return channelRecvOrPeekCheck(x, env)
}

func (x *ArrayExpression) typecheck(env *TypeEnv) error {
	ty := x.Elems[0].type_(env)
	for _, elem := range x.Elems {
		if err := elem.typecheck(env); err != nil {
			return err
		}
		if !ty.equal(elem.type_(env)) {
			return fmt.Errorf("Expect %s to be a %s", elem, ty)
		}
	}
	return nil
}
