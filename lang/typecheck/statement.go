package typecheck

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

// ========================================
// typeCheckStatement

func typeCheckStatement(x Statement, env *typeEnv) error {
	switch x := x.(type) {
	case *ConstantDefinition:
		return typeCheckConstantDefinition(x, env)
	case *LabelledStatement:
		return typeCheckLabelledStatement(x, env)
	case *BlockStatement:
		return typeCheckBlockStatement(x, env)
	case *VarDeclStatement:
		return typeCheckVarDeclStatement(x, env)
	case *IfStatement:
		return typeCheckIfStatement(x, env)
	case *AssignmentStatement:
		return typeCheckAssignmentStatement(x, env)
	case *OpAssignmentStatement:
		return typeCheckOpAssignmentStatement(x, env)
	case *ChoiceStatement:
		return typeCheckChoiceStatement(x, env)
	case *RecvStatement:
		return typeCheckRecvStatement(x, env)
	case *PeekStatement:
		return typeCheckPeekStatement(x, env)
	case *SendStatement:
		return typeCheckSendStatement(x, env)
	case *ForStatement:
		return typeCheckForStatement(x, env)
	case *ForInStatement:
		return typeCheckForInStatement(x, env)
	case *ForInRangeStatement:
		return typeCheckForInRangeStatement(x, env)
	case *BreakStatement:
		return typeCheckBreakStatement(x, env)
	case *GotoStatement:
		return typeCheckGotoStatement(x, env)
	case *SkipStatement:
		return typeCheckSkipStatement(x, env)
	case *ExprStatement:
		return typeCheckExprStatement(x, env)
	case *NullStatement:
		return typeCheckNullStatement(x, env)
	}
	panic("Unknown Statement")
}

func typeCheckStatements(stmts []Statement, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	for _, stmt := range stmts {
		if err := typeCheckStatement(stmt, env); err != nil {
			return err
		}
		switch s := stmt.(type) {
		case *ConstantDefinition:
			env.add(s.Name, s.Type)
		case *VarDeclStatement:
			env.add(s.Name, s.Type)
		}
	}
	return nil
}

func typeCheckLabelledStatement(x *LabelledStatement, env *typeEnv) error {
	return typeCheckStatement(x.Statement, env)
}
func typeCheckBlockStatement(x *BlockStatement, env *typeEnv) error {
	return typeCheckStatements(x.Statements, env)
}
func typeCheckVarDeclStatement(x *VarDeclStatement, env *typeEnv) error {
	if x.Initializer != nil {
		if err := typeCheckExpression(x.Initializer, env); err != nil {
			return err
		}
	}
	return nil
}
func typeCheckIfStatement(x *IfStatement, env *typeEnv) error {
	if err := typeCheckExpression(x.Condition, env); err != nil {
		return err
	}
	if err := typeCheckStatements(x.TrueBranch, env); err != nil {
		return err
	}
	if err := typeCheckStatements(x.FalseBranch, env); err != nil {
		return err
	}
	return nil
}
func typeCheckAssignmentStatement(x *AssignmentStatement, env *typeEnv) error {
	if err := typeCheckExpression(x.Expr, env); err != nil {
		return err
	}
	if ty := env.lookup(x.Variable); ty != nil {
		if !typeOfExpression(x.Expr, env).Equal(ty) {
			return fmt.Errorf("Expect %s to be a type %s", x.Expr, ty)
		}
	} else {
		return fmt.Errorf("Undefined variable %s", x.Variable)
	}
	return nil
}
func typeCheckOpAssignmentStatement(x *OpAssignmentStatement, env *typeEnv) error {
	return typeCheckExpression(
		&BinOpExpression{&IdentifierExpression{x.Variable}, x.Operator, x.Expr},
		env,
	)
}
func typeCheckChoiceStatement(x *ChoiceStatement, env *typeEnv) error {
	for _, block := range x.Blocks {
		if err := typeCheckStatement(&block, env); err != nil {
			return err
		}
	}
	return nil
}
func typeCheckRecvStatement(x *RecvStatement, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}
func typeCheckPeekStatement(x *PeekStatement, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}
func typeCheckSendStatement(x *SendStatement, env *typeEnv) error {
	return channelExprCheck(x, env, false)
}
func typeCheckForStatement(x *ForStatement, env *typeEnv) error {
	return typeCheckStatements(x.Statements, env)
}
func typeCheckForInStatement(x *ForInStatement, env *typeEnv) error {
	if err := typeCheckExpression(x.Container, env); err != nil {
		return err
	}
	if ty, isArrayType := typeOfExpression(x.Container, env).(ArrayType); isArrayType {
		blockEnv := newTypeEnvFromUpper(env)
		blockEnv.add(x.Variable, ty.ElemType)
		return typeCheckStatements(x.Statements, blockEnv)
	} else {
		return fmt.Errorf("Expect %s to be an array", x.Container)
	}
}
func typeCheckForInRangeStatement(x *ForInRangeStatement, env *typeEnv) error {
	if err := typeCheckExpression(x.FromExpr, env); err != nil {
		return err
	}
	if err := typeCheckExpression(x.ToExpr, env); err != nil {
		return err
	}
	if !typeOfExpression(x.FromExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to be an int", x.FromExpr)
	}
	if !typeOfExpression(x.ToExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to be an int", x.ToExpr)
	}
	blockEnv := newTypeEnvFromUpper(env)
	blockEnv.add(x.Variable, NamedType{"int"})
	return typeCheckStatements(x.Statements, blockEnv)
}
func typeCheckBreakStatement(x *BreakStatement, env *typeEnv) error { return nil }
func typeCheckGotoStatement(x *GotoStatement, env *typeEnv) error   { return nil }
func typeCheckSkipStatement(x *SkipStatement, env *typeEnv) error   { return nil }
func typeCheckExprStatement(x *ExprStatement, env *typeEnv) error   { return nil }
func typeCheckNullStatement(x *NullStatement, env *typeEnv) error   { return nil }
