package sandal

import (
	"testing"
)

func TestConstantDefinitionTypecheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := &NumberExpression{"1"}

	{
		def := &ConstantDefinition{"a", intType, numberExpr}
		if err := def.typecheck(NewTypeEnv()); err != nil {
			t.Errorf("Expect \"const a int = 1\" to be valid, but got an error %q", err.Error())
		}
	}
	{
		def := &ConstantDefinition{"a", boolType, numberExpr}
		if err := def.typecheck(NewTypeEnv()); err == nil {
			t.Error("Expect \"const a int = 1\" not to be valid")
		}
	}
}

func TestIdentifierExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expr := &IdentifierExpression{"a"}
		if err := expr.typecheck(env); err != nil {
			t.Errorf("Expect \"a\" to be valid, but got an error %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		expr := &IdentifierExpression{"a"}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"a\" not to be valid")
		}
	}
}

func TestNotExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		expr := &NotExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err != nil {
			t.Errorf("Expect \"!a\" to be valid, but got %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expr := &NotExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"!a\" not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		expr := &NotExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"!a\" not to be valid")
		}
	}
}

func TestUnarySubExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expr := &UnarySubExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err != nil {
			t.Errorf("Expect \"-a\" to be valid, but got %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		expr := &UnarySubExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"-a\" not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		expr := &UnarySubExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"-a\" not to be valid")
		}
	}
}

func TestParenExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expr := &ParenExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err != nil {
			t.Errorf("Expect \"(a)\" to be valid, but got %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		expr := &ParenExpression{&IdentifierExpression{"a"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"(a)\" not to be valid")
		}
	}
}

func TestBinOpExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		expr := &BinOpExpression{&IdentifierExpression{"a"}, ADD, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"a+b\" not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		expr := &BinOpExpression{&IdentifierExpression{"a"}, ADD, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"a+b\" not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		env.Add("b", NamedType{"bool"})
		expr := &BinOpExpression{&IdentifierExpression{"a"}, ADD, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err == nil {
			t.Error("Expect \"a+b\" not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		env.Add("b", NamedType{"int"})
		expr := &BinOpExpression{&IdentifierExpression{"a"}, ADD, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err != nil {
			t.Error("Expect \"a+b\" to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		env.Add("b", NamedType{"int"})
		expr := &BinOpExpression{&IdentifierExpression{"a"}, EQL, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err != nil {
			t.Errorf("Expect \"a==b\" to be valid, but got %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		env.Add("b", NamedType{"bool"})
		expr := &BinOpExpression{&IdentifierExpression{"a"}, EQL, &IdentifierExpression{"b"}}
		if err := expr.typecheck(env); err == nil {
			t.Errorf("Expect \"a==b\" not to be valid")
		}
	}
}

func TestTimeoutRecvExpressionTypecheck(t *testing.T) {
	chExp := &IdentifierExpression{"ch"}
	argExp := []Expression{&IdentifierExpression{"a"}}
	{
		env := NewTypeEnv()
		env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
		env.Add("a", NamedType{"int"})
		if err := (&TimeoutRecvExpression{chExp, argExp}).typecheck(env); err != nil {
			t.Errorf("Expect timeout_recv(ch, a) to be valid, but got %q", err.Error())
		}
	}
	{
		env := NewTypeEnv()
		env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
		if err := (&TimeoutRecvExpression{chExp, argExp}).typecheck(env); err == nil {
			t.Error("Expect timeout_recv(ch, a) not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		if err := (&TimeoutRecvExpression{chExp, argExp}).typecheck(env); err == nil {
			t.Error("Expect timeout_recv(ch, a) not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("ch", &NamedType{"int"})
		env.Add("a", NamedType{"int"})
		if err := (&TimeoutRecvExpression{chExp, argExp}).typecheck(env); err == nil {
			t.Error("Expect timeout_recv(ch, a) not to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("ch", HandshakeChannelType{false, []Type{NamedType{"int"}}})
		if err := (&TimeoutRecvExpression{chExp, []Expression{&NumberExpression{"1"}}}).typecheck(env); err == nil {
			t.Error("Expect timeout_recv(ch, 1) not to be valid")
		}
	}
}

func TestArrayExpressionTypecheck(t *testing.T) {
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"int"})
		if err := (&ArrayExpression{[]Expression{&IdentifierExpression{"a"}, &NumberExpression{"1"}}}).typecheck(env); err != nil {
			t.Error("Expect [a, 1] to be valid")
		}
	}
	{
		env := NewTypeEnv()
		env.Add("a", NamedType{"bool"})
		if err := (&ArrayExpression{[]Expression{&IdentifierExpression{"a"}, &NumberExpression{"1"}}}).typecheck(env); err == nil {
			t.Error("Expect [a, 1] not to be valid")
		}
	}
}
