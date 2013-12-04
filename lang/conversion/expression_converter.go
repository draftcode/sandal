package conversion

import (
	. "github.com/draftcode/sandal/lang/data"
)

func (x *intStatementConverter) convertExpression(expr Expression) string {
	// TODO
	switch expr := expr.(type) {
	case IdentifierExpression:
		if val := x.env.lookup(expr.Name); val != nil {
		} else if expr.Name == "true" {
			return "TRUE"
		} else if expr.Name == "false" {
			return "FALSE"
		}
	case NumberExpression:
	case NotExpression:
	case UnarySubExpression:
	case ParenExpression:
	case BinOpExpression:
	case TimeoutRecvExpression:
	case TimeoutPeekExpression:
	case NonblockRecvExpression:
	case NonblockPeekExpression:
	case ArrayExpression:
	}
	return expr.String()
}
