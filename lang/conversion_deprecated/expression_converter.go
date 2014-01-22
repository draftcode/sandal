package conversion_deprecated

import (
	. "github.com/draftcode/sandal/lang/data"
)

func expressionToInternalObj(expr Expression, env *varEnv) intInternalExpressionObj {
	// This function does not return nil.
	switch expr := expr.(type) {
	case IdentifierExpression:
		intObj := env.lookup(expr.Name)
		if intExprObj, isExprObj := intObj.(intInternalExpressionObj); isExprObj {
			return intExprObj
		} else {
			panic("Referenced name is not expression")
		}
	case NumberExpression:
		return intInternalLiteral{Lit: expr.Lit, Type: NamedType{"number"}}
	case TrueExpression:
		return intInternalLiteral{Lit: "TRUE", Type: NamedType{"bool"}}
	case FalseExpression:
		return intInternalLiteral{Lit: "FALSE", Type: NamedType{"bool"}}
	case NotExpression:
		return intInternalNot{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case UnarySubExpression:
		return intInternalUnarySub{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case ParenExpression:
		return intInternalParen{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case BinOpExpression:
		intObjLHS := expressionToInternalObj(expr.LHS, env)
		intObjRHS := expressionToInternalObj(expr.RHS, env)
		return intInternalBinOp{LHS: intObjLHS, Op: expr.Operator, RHS: intObjRHS}
	case TimeoutRecvExpression:
		ch, args := convertChannelExpr(expr, env)
		return intInternalTimeoutRecv{Channel: ch, Args: args}
	case TimeoutPeekExpression:
		ch, args := convertChannelExpr(expr, env)
		return intInternalTimeoutPeek{Channel: ch, Args: args}
	case NonblockRecvExpression:
		ch, args := convertChannelExpr(expr, env)
		return intInternalNonblockRecv{Channel: ch, Args: args}
	case NonblockPeekExpression:
		ch, args := convertChannelExpr(expr, env)
		return intInternalNonblockPeek{Channel: ch, Args: args}
	case ArrayExpression:
		elems := []intInternalExpressionObj{}
		for _, subExpr := range expr.Elems {
			elems = append(elems, expressionToInternalObj(subExpr, env))
		}
		return intInternalArrayLiteral{Elems: elems}
	default:
		panic("Unknown Expression")
	}
}

func convertChannelExpr(expr ChanExpr, env *varEnv) (ch intInternalExpressionObj, args []intInternalExpressionObj) {
	ch = expressionToInternalObj(expr.ChannelExpr(), env)
	if ch.Steps() != 0 {
		panic("Steps constraint violation")
	}
	for _, arg := range expr.ArgExprs() {
		argObj := expressionToInternalObj(arg, env)
		if argObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		args = append(args, argObj)
	}
	return
}
