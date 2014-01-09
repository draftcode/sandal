package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestIdentifierExpressionTypecheck(t *testing.T) {
	expr := IdentifierExpression{Pos{}, "a"}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestNotExpressionTypecheck(t *testing.T) {
	expr := NotExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestUnarySubExpressionTypecheck(t *testing.T) {
	expr := UnarySubExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestParenExpressionTypecheck(t *testing.T) {
	expr := ParenExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestBinOpExpressionTypecheck(t *testing.T) {
	{
		expr := BinOpExpression{IdentifierExpression{Pos{}, "a"}, "+", IdentifierExpression{Pos{}, "b"}}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
		expectInvalid(t, expr, newTypeEnv())
	}
	{
		expr := BinOpExpression{IdentifierExpression{Pos{}, "a"}, "==", IdentifierExpression{Pos{}, "b"}}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
	}
}

func TestTimeoutRecvExpressionTypecheck(t *testing.T) {
	chExp := IdentifierExpression{Pos{}, "ch"}
	{
		expr := TimeoutRecvExpression{Pos{}, chExp, []Expression{IdentifierExpression{Pos{}, "a"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			env.add("a", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", NamedType{"int"})
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
	}
	{
		expr := TimeoutRecvExpression{Pos{}, chExp, []Expression{NumberExpression{Pos{}, "1"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
	}
}

func TestArrayExpressionTypecheck(t *testing.T) {
	expr := ArrayExpression{Pos{}, []Expression{IdentifierExpression{Pos{}, "a"}, NumberExpression{Pos{}, "1"}}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
}
