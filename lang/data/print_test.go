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
	expectString(t, ConstantDefinition{"a", NamedType{"int"}, NumberExpression{"1"}}, "const a int = 1;")
}

func TestStatementStringify(t *testing.T) {
	expectString(t, LabelledStatement{"label", ExprStatement{NumberExpression{"1"}}}, "label: 1;")
	expectString(t, BlockStatement{[]Statement{SkipStatement{}, NullStatement{}}}, "{ skip; ; };")
	expectString(t, VarDeclStatement{"a", NamedType{"int"}, NumberExpression{"1"}}, "var a int = 1;")
	expectString(t, VarDeclStatement{"a", NamedType{"int"}, nil}, "var a int;")
	expectString(t, IfStatement{IdentifierExpression{"a"}, []Statement{SkipStatement{}}, nil}, "if a { skip; };")
	expectString(t, IfStatement{IdentifierExpression{"a"}, []Statement{SkipStatement{}}, []Statement{}}, "if a { skip; } else {  };")
	expectString(t, IfStatement{IdentifierExpression{"a"}, []Statement{SkipStatement{}}, []Statement{SkipStatement{}}}, "if a { skip; } else { skip; };")
	expectString(t, AssignmentStatement{"a", NumberExpression{"1"}}, "a = 1;")
	expectString(t, OpAssignmentStatement{"a", "+", NumberExpression{"1"}}, "a += 1;")
	expectString(t,
		ChoiceStatement{[]BlockStatement{
			BlockStatement{[]Statement{SkipStatement{}}},
			BlockStatement{[]Statement{ExprStatement{NumberExpression{"1"}}}}}},
		"choice { skip; }, { 1; };")
	expectString(t, RecvStatement{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "recv(ch, a);")
	expectString(t, PeekStatement{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "peek(ch, a);")
	expectString(t, SendStatement{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "send(ch, a);")
	expectString(t, ForStatement{[]Statement{SkipStatement{}}}, "for { skip; };")
	expectString(t, ForInStatement{"ch", IdentifierExpression{"chs"}, []Statement{SkipStatement{}}}, "for ch in chs { skip; };")
	expectString(t, ForInRangeStatement{"i", NumberExpression{"1"}, NumberExpression{"5"}, []Statement{SkipStatement{}}}, "for i in range 1 to 5 { skip; };")
	expectString(t, BreakStatement{}, "break;")
	expectString(t, GotoStatement{"label"}, "goto label;")
	expectString(t, SkipStatement{}, "skip;")
	expectString(t, ExprStatement{NumberExpression{"1"}}, "1;")
	expectString(t, NullStatement{}, ";")
}

func TestExpressionStringify(t *testing.T) {
	expectString(t, IdentifierExpression{"a"}, "a")
	expectString(t, NumberExpression{"1"}, "1")
	expectString(t, NotExpression{IdentifierExpression{"a"}}, "!a")
	expectString(t, UnarySubExpression{IdentifierExpression{"a"}}, "-a")
	expectString(t, ParenExpression{IdentifierExpression{"a"}}, "(a)")
	expectString(t, BinOpExpression{IdentifierExpression{"a"}, "+", IdentifierExpression{"b"}}, "a+b")
	expectString(t, TimeoutRecvExpression{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "timeout_recv(ch, a)")
	expectString(t, TimeoutPeekExpression{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "timeout_peek(ch, a)")
	expectString(t, NonblockRecvExpression{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "nonblock_recv(ch, a)")
	expectString(t, NonblockPeekExpression{IdentifierExpression{"ch"}, []Expression{IdentifierExpression{"a"}}}, "nonblock_peek(ch, a)")
	expectString(t, ArrayExpression{[]Expression{IdentifierExpression{"a"}, IdentifierExpression{"b"}}}, "[a, b]")
}

func TestTypeStringify(t *testing.T) {
	expectString(t, NamedType{"int"}, "int")
	expectString(t, CallableType{[]Type{NamedType{"int"}, NamedType{"bool"}}}, "callable(int, bool)")
	expectString(t, ArrayType{NamedType{"int"}}, "[]int")
	expectString(t, HandshakeChannelType{false, []Type{NamedType{"int"}}}, "channel {int}")
	expectString(t, HandshakeChannelType{true, []Type{NamedType{"int"}}}, "unstable channel {int}")
	expectString(t, BufferedChannelType{false, IdentifierExpression{"a"}, []Type{NamedType{"int"}}}, "channel [a] {int}")
	expectString(t, BufferedChannelType{true, IdentifierExpression{"a"}, []Type{NamedType{"int"}}}, "unstable channel [a] {int}")
	expectString(t, BufferedChannelType{false, nil, []Type{NamedType{"int"}}}, "channel [] {int}")
	expectString(t, BufferedChannelType{true, nil, []Type{NamedType{"int"}}}, "unstable channel [] {int}")
}
