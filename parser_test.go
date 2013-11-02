package sandal

import (
	"reflect"
	"testing"
)

func parse(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune(src), 0)
	definitions := Parse(s)
	if !reflect.DeepEqual(definitions, expect) {
		t.Errorf("Expect %+#v \n but got %+#v", expect, definitions)
	}
}

func TestDataDefinition(t *testing.T) {
	parse(t, "data Maybe { Just, Nothing };",
		[]Definition{&DataDefinition{"Maybe", []string{"Just", "Nothing"}}})
}

func TestParseModuleDefinition(t *testing.T) {
	parse(t, "module A(ch channel [] { bool }, chs []channel [] { bit }) { init { ; }; };",
		[]Definition{&ModuleDefinition{"A",
			[]Parameter{Parameter{"ch", &ChannelType{false, nil, []Type{&NamedType{"bool"}}}},
				Parameter{"chs", &SetType{&ChannelType{false, nil, []Type{&NamedType{"bit"}}}}}},
			[]Definition{&InitBlock{[]Statement{&NullStatement{}}}}}})
}

func TestParseConstantDefinition(t *testing.T) {
	parse(t, "const a = 1;", []Definition{&ConstantDefinition{"a", &NumberExpression{"1"}}})
}

func TestParseProcDefinition(t *testing.T) {
	parse(t, "proc A(ch channel [] { bool }, chs []channel [] { bit }) { ; };",
		[]Definition{&ProcDefinition{"A",
			[]Parameter{Parameter{"ch", &ChannelType{false, nil, []Type{&NamedType{"bool"}}}},
				Parameter{"chs", &SetType{&ChannelType{false, nil, []Type{&NamedType{"bit"}}}}}},
			[]Statement{&NullStatement{}}}})
}

func TestParseInitBlock(t *testing.T) {
	parse(t, "init { };", []Definition{&InitBlock{}})
	parse(t, "init { skip; };", []Definition{&InitBlock{[]Statement{&SkipStatement{}}}})
}

func parseInInitBlock(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("init { "+src+" }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if initBlock, isInitBlock := definitions[0].(*InitBlock); isInitBlock {
		if len(initBlock.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in InitBlock", src)
			return
		}
		if !reflect.DeepEqual(initBlock.Statements[0], expect) {
			t.Errorf("Expect %+#v \n but got %+#v",
				expect, initBlock.Statements[0])
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in InitBlock", src)
		return
	}
}

func TestParseStatement(t *testing.T) {
	parseInInitBlock(t, "test: ;", &LabelledStatement{"test", &NullStatement{}})
	parseInInitBlock(t, "{ ; };", &BlockStatement{[]Statement{&NullStatement{}}})
	parseInInitBlock(t, "var abc bool;", &VarDeclStatement{"abc", &NamedType{"bool"}, nil})
	parseInInitBlock(t, "var abc bool = false;", &VarDeclStatement{"abc", &NamedType{"bool"}, &IdentifierExpression{"false"}})
	parseInInitBlock(t, "if false { ; };", &IfStatement{&IdentifierExpression{"false"}, []Statement{&NullStatement{}}, nil})
	parseInInitBlock(t, "if false { ; } else { skip; };", &IfStatement{&IdentifierExpression{"false"}, []Statement{&NullStatement{}}, []Statement{&SkipStatement{}}})

	bExp := &IdentifierExpression{"b"}
	parseInInitBlock(t, "a=b;", &AssignmentStatement{"a", bExp})
	parseInInitBlock(t, "a+=b;", &OpAssignmentStatement{"a", ADD, bExp})
	parseInInitBlock(t, "a-=b;", &OpAssignmentStatement{"a", SUB, bExp})
	parseInInitBlock(t, "a*=b;", &OpAssignmentStatement{"a", MUL, bExp})
	parseInInitBlock(t, "a/=b;", &OpAssignmentStatement{"a", QUO, bExp})
	parseInInitBlock(t, "a%=b;", &OpAssignmentStatement{"a", REM, bExp})
	parseInInitBlock(t, "a&=b;", &OpAssignmentStatement{"a", AND, bExp})
	parseInInitBlock(t, "a|=b;", &OpAssignmentStatement{"a", OR, bExp})
	parseInInitBlock(t, "a^=b;", &OpAssignmentStatement{"a", XOR, bExp})
	parseInInitBlock(t, "a<<=b;", &OpAssignmentStatement{"a", SHL, bExp})
	parseInInitBlock(t, "a>>=b;", &OpAssignmentStatement{"a", SHR, bExp})

	parseInInitBlock(t, "choice { ; }, { skip; };", &ChoiceStatement{[]BlockStatement{BlockStatement{[]Statement{&NullStatement{}}}, BlockStatement{[]Statement{&SkipStatement{}}}}})
	parseInInitBlock(t, "recv(ch, 1, 2);", &RecvStatement{&IdentifierExpression{"ch"}, []Expression{&NumberExpression{"1"}, &NumberExpression{"2"}}})
	parseInInitBlock(t, "peek(ch);", &PeekStatement{&IdentifierExpression{"ch"}, []Expression{}})
	parseInInitBlock(t, "send(ch, 1, 2);", &SendStatement{&IdentifierExpression{"ch"}, []Expression{&NumberExpression{"1"}, &NumberExpression{"2"}}})
	parseInInitBlock(t, "for { ; };", &ForStatement{[]Statement{&NullStatement{}}})
	parseInInitBlock(t, "for ch in chs { ; };", &ForInStatement{"ch", &IdentifierExpression{"chs"}, []Statement{&NullStatement{}}})
	parseInInitBlock(t, "for i in range 1 to 5 { ; };", &ForInRangeStatement{"i", &NumberExpression{"1"}, &NumberExpression{"5"}, []Statement{&NullStatement{}}})
	parseInInitBlock(t, "break;", &BreakStatement{})
	parseInInitBlock(t, "goto here;", &GotoStatement{"here"})
	parseInInitBlock(t, "Mod(1, 2);", &CallStatement{"Mod", []Expression{&NumberExpression{"1"}, &NumberExpression{"2"}}})
	parseInInitBlock(t, "skip;", &SkipStatement{})
	parseInInitBlock(t, ";", &NullStatement{})
	parseInInitBlock(t, "1;", &ExprStatement{&NumberExpression{"1"}})
	parseInInitBlock(t, "const a = 1;", &ConstantDefinition{"a", &NumberExpression{"1"}})
}

func TestParseExpression(t *testing.T) {
	parseInInitBlock(t, "abc;", &ExprStatement{&IdentifierExpression{"abc"}})
	parseInInitBlock(t, "123;", &ExprStatement{&NumberExpression{"123"}})
	parseInInitBlock(t, "!abc;", &ExprStatement{&NotExpression{&IdentifierExpression{"abc"}}})
	parseInInitBlock(t, "-abc;", &ExprStatement{&UnarySubExpression{&IdentifierExpression{"abc"}}})
	parseInInitBlock(t, "(abc);", &ExprStatement{&ParenExpression{&IdentifierExpression{"abc"}}})

	aExp := &IdentifierExpression{"a"}
	bExp := &IdentifierExpression{"b"}
	parseInInitBlock(t, "a+b;", &ExprStatement{&BinOpExpression{aExp, ADD, bExp}})
	parseInInitBlock(t, "a-b;", &ExprStatement{&BinOpExpression{aExp, SUB, bExp}})
	parseInInitBlock(t, "a*b;", &ExprStatement{&BinOpExpression{aExp, MUL, bExp}})
	parseInInitBlock(t, "a/b;", &ExprStatement{&BinOpExpression{aExp, QUO, bExp}})
	parseInInitBlock(t, "a%b;", &ExprStatement{&BinOpExpression{aExp, REM, bExp}})
	parseInInitBlock(t, "a&b;", &ExprStatement{&BinOpExpression{aExp, AND, bExp}})
	parseInInitBlock(t, "a|b;", &ExprStatement{&BinOpExpression{aExp, OR, bExp}})
	parseInInitBlock(t, "a^b;", &ExprStatement{&BinOpExpression{aExp, XOR, bExp}})
	parseInInitBlock(t, "a<<b;", &ExprStatement{&BinOpExpression{aExp, SHL, bExp}})
	parseInInitBlock(t, "a>>b;", &ExprStatement{&BinOpExpression{aExp, SHR, bExp}})
	parseInInitBlock(t, "a&&b;", &ExprStatement{&BinOpExpression{aExp, LAND, bExp}})
	parseInInitBlock(t, "a||b;", &ExprStatement{&BinOpExpression{aExp, LOR, bExp}})
	parseInInitBlock(t, "a==b;", &ExprStatement{&BinOpExpression{aExp, EQL, bExp}})
	parseInInitBlock(t, "a<b;", &ExprStatement{&BinOpExpression{aExp, LSS, bExp}})
	parseInInitBlock(t, "a>b;", &ExprStatement{&BinOpExpression{aExp, GTR, bExp}})
	parseInInitBlock(t, "a!=b;", &ExprStatement{&BinOpExpression{aExp, NEQ, bExp}})
	parseInInitBlock(t, "a<=b;", &ExprStatement{&BinOpExpression{aExp, LEQ, bExp}})
	parseInInitBlock(t, "a>=b;", &ExprStatement{&BinOpExpression{aExp, GEQ, bExp}})

	parseInInitBlock(t, "timeout_recv(ch);", &ExprStatement{&TimeoutRecvExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInInitBlock(t, "timeout_peek(ch);", &ExprStatement{&TimeoutPeekExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInInitBlock(t, "nonblock_recv(ch);", &ExprStatement{&NonblockRecvExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInInitBlock(t, "nonblock_peek(ch);", &ExprStatement{&NonblockPeekExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInInitBlock(t, "[a, b];", &ExprStatement{&ArrayExpression{[]Expression{aExp, bExp}}})
}

func parseType(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("init { var a "+src+"; }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if initBlock, isInitBlock := definitions[0].(*InitBlock); isInitBlock {
		if len(initBlock.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in InitBlock", src)
			return
		}
		if stmt, isVarDecl := initBlock.Statements[0].(*VarDeclStatement); isVarDecl {
			if !reflect.DeepEqual(stmt.Type, expect) {
				t.Errorf("Expect %+#v \n but got %+#v", expect, stmt.Type)
				return
			}
		} else {
			t.Errorf("Expect %q to be parsed in InitBlock", src)
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in InitBlock", src)
		return
	}
}

func TestParseType(t *testing.T) {
	parseType(t, "bool", &NamedType{"bool"})
	parseType(t, "[]bool", &SetType{&NamedType{"bool"}})
	parseType(t, "channel [] { bool }",
		&ChannelType{false, nil, []Type{&NamedType{"bool"}}})
	parseType(t, "channel [1+2] { bool }",
		&ChannelType{false, &BinOpExpression{&NumberExpression{"1"}, ADD, &NumberExpression{"2"}}, []Type{&NamedType{"bool"}}})
	parseType(t, "unstable channel [] { bool }",
		&ChannelType{true, nil, []Type{&NamedType{"bool"}}})
	parseType(t, "unstable channel [1+2] { bool }",
		&ChannelType{true, &BinOpExpression{&NumberExpression{"1"}, ADD, &NumberExpression{"2"}}, []Type{&NamedType{"bool"}}})
}
