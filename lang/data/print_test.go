package data

import (
	"testing"
)

type Printable interface {
	String() string
}

func expectString(t *testing.T, printable Printable, expected string) {
	if s := printable.String(); s != expected {
		t.Errorf("Expect %q to be %q", s, expected)
	}
}

func TestDefinitionStringify(t *testing.T) {
	expectString(t, ConstantDefinition{Pos{}, "a", NamedType{"int"}, NumberExpression{Pos{}, "1"}}, "const a int = 1;")
}

func TestStatementStringify(t *testing.T) {
	expectString(t, LabelledStatement{Pos{}, "label", ExprStatement{NumberExpression{Pos{}, "1"}}}, "label: 1;")
	expectString(t, BlockStatement{Pos{}, []Statement{SkipStatement{}, NullStatement{}}}, "{ skip; ; };")
	expectString(t, VarDeclStatement{Pos{}, "a", NamedType{"int"}, NumberExpression{Pos{}, "1"}}, "var a int = 1;")
	expectString(t, VarDeclStatement{Pos{}, "a", NamedType{"int"}, nil}, "var a int;")
	expectString(t, IfStatement{Pos{}, IdentifierExpression{Pos{}, "a"}, []Statement{SkipStatement{}}, nil}, "if a { skip; };")
	expectString(t, IfStatement{Pos{}, IdentifierExpression{Pos{}, "a"}, []Statement{SkipStatement{}}, []Statement{}}, "if a { skip; } else {  };")
	expectString(t, IfStatement{Pos{}, IdentifierExpression{Pos{}, "a"}, []Statement{SkipStatement{}}, []Statement{SkipStatement{}}}, "if a { skip; } else { skip; };")
	expectString(t, AssignmentStatement{Pos{}, "a", NumberExpression{Pos{}, "1"}}, "a = 1;")
	expectString(t, OpAssignmentStatement{Pos{}, "a", "+", NumberExpression{Pos{}, "1"}}, "a += 1;")
	expectString(t,
		ChoiceStatement{Pos{}, []BlockStatement{
			BlockStatement{Pos{}, []Statement{SkipStatement{}}},
			BlockStatement{Pos{}, []Statement{ExprStatement{NumberExpression{Pos{}, "1"}}}}}},
		"choice { skip; }, { 1; };")
	expectString(t, RecvStatement{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "recv(ch, a);")
	expectString(t, PeekStatement{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "peek(ch, a);")
	expectString(t, SendStatement{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "send(ch, a);")
	expectString(t, ForStatement{Pos{}, []Statement{SkipStatement{}}}, "for { skip; };")
	expectString(t, ForInStatement{Pos{}, "ch", IdentifierExpression{Pos{}, "chs"}, []Statement{SkipStatement{}}}, "for ch in chs { skip; };")
	expectString(t, ForInRangeStatement{Pos{}, "i", NumberExpression{Pos{}, "1"}, NumberExpression{Pos{}, "5"}, []Statement{SkipStatement{}}}, "for i in range 1 to 5 { skip; };")
	expectString(t, BreakStatement{}, "break;")
	expectString(t, GotoStatement{Pos{}, "label"}, "goto label;")
	expectString(t, SkipStatement{}, "skip;")
	expectString(t, ExprStatement{NumberExpression{Pos{}, "1"}}, "1;")
	expectString(t, NullStatement{}, ";")
}

func TestExpressionStringify(t *testing.T) {
	expectString(t, IdentifierExpression{Pos{}, "a"}, "a")
	expectString(t, NumberExpression{Pos{}, "1"}, "1")
	expectString(t, TrueExpression{Pos{}}, "true")
	expectString(t, FalseExpression{Pos{}}, "false")
	expectString(t, NotExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}, "!a")
	expectString(t, UnarySubExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}, "-a")
	expectString(t, ParenExpression{Pos{}, IdentifierExpression{Pos{}, "a"}}, "(a)")
	expectString(t, BinOpExpression{IdentifierExpression{Pos{}, "a"}, "+", IdentifierExpression{Pos{}, "b"}}, "a+b")
	expectString(t, TimeoutRecvExpression{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "timeout_recv(ch, a)")
	expectString(t, TimeoutPeekExpression{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "timeout_peek(ch, a)")
	expectString(t, NonblockRecvExpression{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "nonblock_recv(ch, a)")
	expectString(t, NonblockPeekExpression{Pos{}, IdentifierExpression{Pos{}, "ch"}, []Expression{IdentifierExpression{Pos{}, "a"}}}, "nonblock_peek(ch, a)")
	expectString(t, ArrayExpression{Pos{}, []Expression{IdentifierExpression{Pos{}, "a"}, IdentifierExpression{Pos{}, "b"}}}, "[a, b]")
}

func TestTypeStringify(t *testing.T) {
	expectString(t, NamedType{"int"}, "int")
	expectString(t, CallableType{[]Type{NamedType{"int"}, NamedType{"bool"}}}, "callable(int, bool)")
	expectString(t, ArrayType{NamedType{"int"}}, "[]int")
	expectString(t, HandshakeChannelType{false, []Type{NamedType{"int"}}}, "channel {int}")
	expectString(t, HandshakeChannelType{true, []Type{NamedType{"int"}}}, "unstable channel {int}")
	expectString(t, BufferedChannelType{false, IdentifierExpression{Pos{}, "a"}, []Type{NamedType{"int"}}}, "channel [a] {int}")
	expectString(t, BufferedChannelType{true, IdentifierExpression{Pos{}, "a"}, []Type{NamedType{"int"}}}, "unstable channel [a] {int}")
	expectString(t, BufferedChannelType{false, nil, []Type{NamedType{"int"}}}, "channel [] {int}")
	expectString(t, BufferedChannelType{true, nil, []Type{NamedType{"int"}}}, "unstable channel [] {int}")
}
