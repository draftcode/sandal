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

func (x *LabelledStatement) typecheck(env *TypeEnv) error {
	return x.Statement.typecheck(env)
}
func (x *BlockStatement) typecheck(env *TypeEnv) error {
	for _, stmt := range x.Statements {
		if err := stmt.typecheck(env); err != nil {
			return err
		}
	}
	return nil
}
func (x *VarDeclStatement) typecheck(env *TypeEnv) error {
	if x.Initializer != nil {
		if err := x.Initializer.typecheck(env); err != nil {
			return err
		}
	}
	return nil
}
func (x *IfStatement) typecheck(env *TypeEnv) error {
	if err := x.Condition.typecheck(env); err != nil {
		return err
	}
	for _, stmt := range x.TrueBranch {
		if err := stmt.typecheck(env); err != nil {
			return err
		}
	}
	for _, stmt := range x.FalseBranch {
		if err := stmt.typecheck(env); err != nil {
			return err
		}
	}
	return nil
}
func (x *AssignmentStatement) typecheck(env *TypeEnv) error {
	if err := x.Expr.typecheck(env); err != nil {
		return err
	}
	if ty := env.Lookup(x.Variable); ty != nil {
		if !ty.equal(x.Expr.type_(env)) {
			return fmt.Errorf("Expect %s to be a type %s", x.Expr, ty)
		}
	} else {
		return fmt.Errorf("Undefined variable %s", x.Variable)
	}
	return nil
}
func (x *OpAssignmentStatement) typecheck(env *TypeEnv) error { return nil } // TODO
func (x *ChoiceStatement) typecheck(env *TypeEnv) error {
	for _, block := range x.Blocks {
		if err := block.typecheck(env); err != nil {
			return err
		}
	}
	return nil
}
func (x *RecvStatement) typecheck(env *TypeEnv) error { return channelRecvOrPeekCheck(x, env) }
func (x *PeekStatement) typecheck(env *TypeEnv) error { return channelRecvOrPeekCheck(x, env) }
func (x *SendStatement) typecheck(env *TypeEnv) error         { return nil }
func (x *ForStatement) typecheck(env *TypeEnv) error          { return nil }
func (x *ForInStatement) typecheck(env *TypeEnv) error        { return nil }
func (x *ForInRangeStatement) typecheck(env *TypeEnv) error   { return nil }
func (x *BreakStatement) typecheck(env *TypeEnv) error        { return nil }
func (x *GotoStatement) typecheck(env *TypeEnv) error         { return nil }
func (x *CallStatement) typecheck(env *TypeEnv) error         { return nil }

// Always valid
func (x *SkipStatement) typecheck(env *TypeEnv) error { return nil }
func (x *ExprStatement) typecheck(env *TypeEnv) error { return nil }
func (x *NullStatement) typecheck(env *TypeEnv) error { return nil }

func (x *ConstantDefinition) typeexec(env *TypeEnv)    {}
func (x *LabelledStatement) typeexec(env *TypeEnv)     {}
func (x *BlockStatement) typeexec(env *TypeEnv)        {}
func (x *VarDeclStatement) typeexec(env *TypeEnv)      {}
func (x *IfStatement) typeexec(env *TypeEnv)           {}
func (x *AssignmentStatement) typeexec(env *TypeEnv)   {}
func (x *OpAssignmentStatement) typeexec(env *TypeEnv) {}
func (x *ChoiceStatement) typeexec(env *TypeEnv)       {}
func (x *RecvStatement) typeexec(env *TypeEnv)         {}
func (x *PeekStatement) typeexec(env *TypeEnv)         {}
func (x *SendStatement) typeexec(env *TypeEnv)         {}
func (x *ForStatement) typeexec(env *TypeEnv)          {}
func (x *ForInStatement) typeexec(env *TypeEnv)        {}
func (x *ForInRangeStatement) typeexec(env *TypeEnv)   {}
func (x *BreakStatement) typeexec(env *TypeEnv)        {}
func (x *GotoStatement) typeexec(env *TypeEnv)         {}
func (x *CallStatement) typeexec(env *TypeEnv)         {}
func (x *SkipStatement) typeexec(env *TypeEnv)         {}
func (x *ExprStatement) typeexec(env *TypeEnv)         {}
func (x *NullStatement) typeexec(env *TypeEnv)         {}

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

// ========================================

func channelRecvOrPeekCheck(ch ChanRecvPoll, env *TypeEnv) error {
	if err := ch.RecvChannel().typecheck(env); err != nil {
		return err
	}
	for _, arg := range ch.RecvArgs() {
		if err := arg.typecheck(env); err != nil {
			return err
		}
	}

	var elemTypes []Type
	switch ty := ch.RecvChannel().type_(env).(type) {
	case HandshakeChannelType:
		elemTypes = ty.Elems
	case BufferedChannelType:
		elemTypes = ty.Elems
	default:
		return fmt.Errorf("Expect the first argument of %s to be a channel but got %s",
			ch, ch.RecvChannel().type_(env))
	}

	if len(elemTypes) != len(ch.RecvArgs()) {
		return fmt.Errorf("Expect the arugments of %s to have %d elements",
			ch, len(elemTypes))
	}
	for i := 0; i < len(elemTypes); i++ {
		if !elemTypes[i].equal(ch.RecvArgs()[i].type_(env)) {
			return fmt.Errorf("Expect the argument %s to be a %s", ch.RecvArgs()[i], elemTypes[i])
		}
		if _, isIdentExpr := ch.RecvArgs()[i].(*IdentifierExpression); !isIdentExpr {
			return fmt.Errorf("Expect the argument %s to be an identifier", ch.RecvArgs()[i])
		}
	}
	return nil
}
