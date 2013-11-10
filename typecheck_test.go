package sandal

import (
	"testing"
)

type TypeCheckable interface {
	typecheck(*TypeEnv) error
	String() string
}

func expectValid(t *testing.T, x TypeCheckable, env *TypeEnv) {
	if err := x.typecheck(env); err != nil {
		t.Errorf("Expect %q to be valid, but got an error %q", x, err)
	}
}

func expectInvalid(t *testing.T, x TypeCheckable, env *TypeEnv) {
	if err := x.typecheck(env); err == nil {
		t.Errorf("Expect %q to be invalid", x)
	}
}

// ========================================
// Typecheck of definitions

func TestConstantDefinitionTypecheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := &NumberExpression{"1"}

	expectValid(t, &ConstantDefinition{"a", intType, numberExpr}, NewTypeEnv())
	expectInvalid(t, &ConstantDefinition{"a", boolType, numberExpr}, NewTypeEnv())
}

// ========================================
// Typecheck of expression

func TestIdentifierExpressionTypecheck(t *testing.T) {
	expr := &IdentifierExpression{"a"}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, NewTypeEnv())
}

func TestNotExpressionTypecheck(t *testing.T) {
	expr := &NotExpression{&IdentifierExpression{"a"}}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		expectValid(t, expr, env)
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, NewTypeEnv())
}

func TestUnarySubExpressionTypecheck(t *testing.T) {
	expr := &UnarySubExpression{&IdentifierExpression{"a"}}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, NewTypeEnv())
}

func TestParenExpressionTypecheck(t *testing.T) {
	expr := &ParenExpression{&IdentifierExpression{"a"}}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, NewTypeEnv())
}

func TestBinOpExpressionTypecheck(t *testing.T) {
	{
		expr := &BinOpExpression{&IdentifierExpression{"a"}, ADD, &IdentifierExpression{"b"}}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			env.Add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			env.Add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
		expectInvalid(t, expr, NewTypeEnv())
	}
	{
		expr := &BinOpExpression{&IdentifierExpression{"a"}, EQL, &IdentifierExpression{"b"}}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			env.Add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			env.Add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
	}
}

func TestTimeoutRecvExpressionTypecheck(t *testing.T) {
	chExp := &IdentifierExpression{"ch"}
	{
		expr := &TimeoutRecvExpression{chExp, []Expression{&IdentifierExpression{"a"}}}
		{
			env := NewTypeEnv()
			env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			env.Add("a", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := NewTypeEnv()
			env.Add("ch", &NamedType{"int"})
			env.Add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
	}
	{
		expr := &TimeoutRecvExpression{chExp, []Expression{&NumberExpression{"1"}}}
		{
			env := NewTypeEnv()
			env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
	}
}

func TestArrayExpressionTypecheck(t *testing.T) {
	expr := &ArrayExpression{[]Expression{&IdentifierExpression{"a"}, &NumberExpression{"1"}}}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
}
