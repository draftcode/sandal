package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestIdentifierExpressionTypecheck(t *testing.T) {
	expr := &IdentifierExpression{"a"}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestNotExpressionTypecheck(t *testing.T) {
	expr := &NotExpression{&IdentifierExpression{"a"}}
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
	expr := &UnarySubExpression{&IdentifierExpression{"a"}}
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
	expr := &ParenExpression{&IdentifierExpression{"a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestBinOpExpressionTypecheck(t *testing.T) {
	{
		expr := &BinOpExpression{&IdentifierExpression{"a"}, "+", &IdentifierExpression{"b"}}
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
		expr := &BinOpExpression{&IdentifierExpression{"a"}, "==", &IdentifierExpression{"b"}}
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
	chExp := &IdentifierExpression{"ch"}
	{
		expr := &TimeoutRecvExpression{chExp, []Expression{&IdentifierExpression{"a"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			env.add("a", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", &NamedType{"int"})
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
	}
	{
		expr := &TimeoutRecvExpression{chExp, []Expression{&NumberExpression{"1"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
	}
}

func TestArrayExpressionTypecheck(t *testing.T) {
	expr := &ArrayExpression{[]Expression{&IdentifierExpression{"a"}, &NumberExpression{"1"}}}
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
