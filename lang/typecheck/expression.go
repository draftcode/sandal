package typecheck

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

// ========================================
// typeOfExpression

func typeOfExpression(x Expression, env *typeEnv) Type {
	switch x := x.(type) {
	case IdentifierExpression:
		return typeOfIdentifierExpression(x, env)
	case NumberExpression:
		return typeOfNumberExpression(x, env)
	case NotExpression:
		return typeOfNotExpression(x, env)
	case UnarySubExpression:
		return typeOfUnarySubExpression(x, env)
	case ParenExpression:
		return typeOfParenExpression(x, env)
	case BinOpExpression:
		return typeOfBinOpExpression(x, env)
	case TimeoutRecvExpression:
		return typeOfTimeoutRecvExpression(x, env)
	case TimeoutPeekExpression:
		return typeOfTimeoutPeekExpression(x, env)
	case NonblockRecvExpression:
		return typeOfNonblockRecvExpression(x, env)
	case NonblockPeekExpression:
		return typeOfNonblockPeekExpression(x, env)
	case ArrayExpression:
		return typeOfArrayExpression(x, env)
	}
	panic("Unknown Expression")
}

func typeOfIdentifierExpression(x IdentifierExpression, env *typeEnv) Type {
	return env.lookup(x.Name)
}

func typeOfNumberExpression(x NumberExpression, env *typeEnv) Type {
	return NamedType{Name: "int"}
}

func typeOfNotExpression(x NotExpression, env *typeEnv) Type {
	return typeOfExpression(x.SubExpr, env)
}

func typeOfUnarySubExpression(x UnarySubExpression, env *typeEnv) Type {
	return typeOfExpression(x.SubExpr, env)
}

func typeOfParenExpression(x ParenExpression, env *typeEnv) Type {
	return typeOfExpression(x.SubExpr, env)
}

var operatorResultType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": NamedType{"bool"},
	"<":  NamedType{"bool"},
	">":  NamedType{"bool"},
	"!=": NamedType{"bool"},
	"<=": NamedType{"bool"},
	">=": NamedType{"bool"},
}

func typeOfBinOpExpression(x BinOpExpression, env *typeEnv) Type {
	if ty, exist := operatorResultType[x.Operator]; exist {
		return ty
	} else {
		panic("Unknown operator: " + x.Operator)
	}
}

func typeOfTimeoutRecvExpression(x TimeoutRecvExpression, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfTimeoutPeekExpression(x TimeoutPeekExpression, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfNonblockRecvExpression(x NonblockRecvExpression, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfNonblockPeekExpression(x NonblockPeekExpression, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfArrayExpression(x ArrayExpression, env *typeEnv) Type {
	if len(x.Elems) == 0 {
		panic("An array should have at least one element")
	}
	// Every element of an array has the same type.
	return typeOfExpression(x.Elems[0], env)
}

// ========================================
// typeCheckExpression

func typeCheckExpression(x Expression, env *typeEnv) error {
	switch x := x.(type) {
	case IdentifierExpression:
		return typeCheckIdentifierExpression(x, env)
	case NumberExpression:
		return typeCheckNumberExpression(x, env)
	case NotExpression:
		return typeCheckNotExpression(x, env)
	case UnarySubExpression:
		return typeCheckUnarySubExpression(x, env)
	case ParenExpression:
		return typeCheckParenExpression(x, env)
	case BinOpExpression:
		return typeCheckBinOpExpression(x, env)
	case TimeoutRecvExpression:
		return typeCheckTimeoutRecvExpression(x, env)
	case TimeoutPeekExpression:
		return typeCheckTimeoutPeekExpression(x, env)
	case NonblockRecvExpression:
		return typeCheckNonblockRecvExpression(x, env)
	case NonblockPeekExpression:
		return typeCheckNonblockPeekExpression(x, env)
	case ArrayExpression:
		return typeCheckArrayExpression(x, env)
	}
	panic("Unknown Expression")
}

func typeCheckIdentifierExpression(x IdentifierExpression, env *typeEnv) error {
	if env.lookup(x.Name) == nil {
		return fmt.Errorf("Undefined variable %s", x.Name)
	}
	return nil
}

func typeCheckNumberExpression(x NumberExpression, env *typeEnv) error {
	// Number expressions are always valid.
	return nil
}

func typeCheckNotExpression(x NotExpression, env *typeEnv) error {
	if err := typeCheckExpression(x.SubExpr, env); err != nil {
		return err
	}
	if !typeOfExpression(x.SubExpr, env).Equal(NamedType{"bool"}) {
		return fmt.Errorf("Expect %s to have type bool, but got %s",
			x.SubExpr, typeOfExpression(x.SubExpr, env))
	}
	return nil
}

func typeCheckUnarySubExpression(x UnarySubExpression, env *typeEnv) error {
	if err := typeCheckExpression(x.SubExpr, env); err != nil {
		return err
	}
	if !typeOfExpression(x.SubExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to have type int, but got %s",
			x.SubExpr, typeOfExpression(x.SubExpr, env))
	}
	return nil
}

func typeCheckParenExpression(x ParenExpression, env *typeEnv) error {
	if err := typeCheckExpression(x.SubExpr, env); err != nil {
		return err
	}
	return nil
}

var operatorOperandType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": nil,
	"<":  NamedType{"int"},
	">":  NamedType{"int"},
	"!=": NamedType{"int"},
	"<=": NamedType{"int"},
	">=": NamedType{"int"},
}

func typeCheckBinOpExpression(x BinOpExpression, env *typeEnv) error {
	if err := typeCheckExpression(x.LHS, env); err != nil {
		return err
	}
	if err := typeCheckExpression(x.RHS, env); err != nil {
		return err
	}
	if ty, exist := operatorOperandType[x.Operator]; exist {
		if ty != nil {
			lhsType := typeOfExpression(x.LHS, env)
			if !lhsType.Equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s",
					x.LHS, ty, lhsType)
			}
			rhsType := typeOfExpression(x.RHS, env)
			if !rhsType.Equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s",
					x.RHS, ty, rhsType)
			}
		} else {
			lhsType := typeOfExpression(x.LHS, env)
			rhsType := typeOfExpression(x.RHS, env)
			if !lhsType.Equal(rhsType) {
				return fmt.Errorf("Expect %s and %s to have the same type but got %s and %s",
					x.LHS, x.RHS, lhsType, rhsType)
			}
		}
	} else {
		panic("Unknown operator: " + x.Operator)
	}
	return nil
}

func typeCheckTimeoutRecvExpression(x TimeoutRecvExpression, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckTimeoutPeekExpression(x TimeoutPeekExpression, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckNonblockRecvExpression(x NonblockRecvExpression, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckNonblockPeekExpression(x NonblockPeekExpression, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckArrayExpression(x ArrayExpression, env *typeEnv) error {
	ty := typeOfExpression(x.Elems[0], env)
	for _, elem := range x.Elems {
		if err := typeCheckExpression(elem, env); err != nil {
			return err
		}
		if !typeOfExpression(elem, env).Equal(ty) {
			return fmt.Errorf("Expect %s to be a %s", elem, ty)
		}
	}
	return nil
}
