package parsing

import (
	. "github.com/draftcode/sandal/lang/data"
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
	parse(t, "module A(ch channel { bool }, chs []channel { bit }) { init { }; };",
		[]Definition{&ModuleDefinition{"A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{false, []Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{false, []Type{NamedType{"bit"}}}}}},
			[]Definition{&InitBlock{}}}})
}

func TestParseConstantDefinition(t *testing.T) {
	parse(t, "const a int = 1;", []Definition{&ConstantDefinition{"a", NamedType{"int"}, &NumberExpression{"1"}}})
}

func TestParseProcDefinition(t *testing.T) {
	parse(t, "proc A(ch channel { bool }, chs []channel { bit }) { ; };",
		[]Definition{&ProcDefinition{"A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{false, []Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{false, []Type{NamedType{"bit"}}}}}},
			[]Statement{&NullStatement{}}}})
}

func TestParseInitBlock(t *testing.T) {
	parse(t, "init { };", []Definition{&InitBlock{}})
	parse(t, "init { a : M(b) };",
		[]Definition{&InitBlock{[]InitVar{
			InstanceVar{"a", "M", []Expression{&IdentifierExpression{"b"}}},
		}}})
	parse(t, "init { a : channel { bool } };",
		[]Definition{&InitBlock{[]InitVar{
			ChannelVar{"a", HandshakeChannelType{false, []Type{NamedType{"bool"}}}},
		}}})
}

func parseInBlock(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { "+src+" }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isProcDef := definitions[0].(*ProcDefinition); isProcDef {
		if len(procDef.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDefinition", src)
			return
		}
		if !reflect.DeepEqual(procDef.Statements[0], expect) {
			t.Errorf("Expect %+#v \n but got %+#v",
				expect, procDef.Statements[0])
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in ProcDefinition", src)
		return
	}
}

func TestParseStatement(t *testing.T) {
	parseInBlock(t, "test: ;", &LabelledStatement{"test", &NullStatement{}})
	parseInBlock(t, "{ ; };", &BlockStatement{[]Statement{&NullStatement{}}})
	parseInBlock(t, "var abc bool;", &VarDeclStatement{"abc", NamedType{"bool"}, nil})
	parseInBlock(t, "var abc bool = false;", &VarDeclStatement{"abc", NamedType{"bool"}, &IdentifierExpression{"false"}})
	parseInBlock(t, "if false { ; };", &IfStatement{&IdentifierExpression{"false"}, []Statement{&NullStatement{}}, nil})
	parseInBlock(t, "if false { ; } else { skip; };", &IfStatement{&IdentifierExpression{"false"}, []Statement{&NullStatement{}}, []Statement{&SkipStatement{}}})

	bExp := &IdentifierExpression{"b"}
	parseInBlock(t, "a=b;", &AssignmentStatement{"a", bExp})
	parseInBlock(t, "a+=b;", &OpAssignmentStatement{"a", "+", bExp})
	parseInBlock(t, "a-=b;", &OpAssignmentStatement{"a", "-", bExp})
	parseInBlock(t, "a*=b;", &OpAssignmentStatement{"a", "*", bExp})
	parseInBlock(t, "a/=b;", &OpAssignmentStatement{"a", "/", bExp})
	parseInBlock(t, "a%=b;", &OpAssignmentStatement{"a", "%", bExp})
	parseInBlock(t, "a&=b;", &OpAssignmentStatement{"a", "&", bExp})
	parseInBlock(t, "a|=b;", &OpAssignmentStatement{"a", "|", bExp})
	parseInBlock(t, "a^=b;", &OpAssignmentStatement{"a", "^", bExp})
	parseInBlock(t, "a<<=b;", &OpAssignmentStatement{"a", "<<", bExp})
	parseInBlock(t, "a>>=b;", &OpAssignmentStatement{"a", ">>", bExp})

	parseInBlock(t, "choice { ; }, { skip; };", &ChoiceStatement{[]BlockStatement{BlockStatement{[]Statement{&NullStatement{}}}, BlockStatement{[]Statement{&SkipStatement{}}}}})
	parseInBlock(t, "recv(ch, 1, 2);", &RecvStatement{&IdentifierExpression{"ch"}, []Expression{&NumberExpression{"1"}, &NumberExpression{"2"}}})
	parseInBlock(t, "peek(ch);", &PeekStatement{&IdentifierExpression{"ch"}, []Expression{}})
	parseInBlock(t, "send(ch, 1, 2);", &SendStatement{&IdentifierExpression{"ch"}, []Expression{&NumberExpression{"1"}, &NumberExpression{"2"}}})
	parseInBlock(t, "for { ; };", &ForStatement{[]Statement{&NullStatement{}}})
	parseInBlock(t, "for ch in chs { ; };", &ForInStatement{"ch", &IdentifierExpression{"chs"}, []Statement{&NullStatement{}}})
	parseInBlock(t, "for i in range 1 to 5 { ; };", &ForInRangeStatement{"i", &NumberExpression{"1"}, &NumberExpression{"5"}, []Statement{&NullStatement{}}})
	parseInBlock(t, "break;", &BreakStatement{})
	parseInBlock(t, "goto here;", &GotoStatement{"here"})
	parseInBlock(t, "skip;", &SkipStatement{})
	parseInBlock(t, ";", &NullStatement{})
	parseInBlock(t, "1;", &ExprStatement{&NumberExpression{"1"}})
	parseInBlock(t, "const a int = 1;", &ConstantDefinition{"a", NamedType{"int"}, &NumberExpression{"1"}})
}

func TestParseExpression(t *testing.T) {
	parseInBlock(t, "abc;", &ExprStatement{&IdentifierExpression{"abc"}})
	parseInBlock(t, "123;", &ExprStatement{&NumberExpression{"123"}})
	parseInBlock(t, "!abc;", &ExprStatement{&NotExpression{&IdentifierExpression{"abc"}}})
	parseInBlock(t, "-abc;", &ExprStatement{&UnarySubExpression{&IdentifierExpression{"abc"}}})
	parseInBlock(t, "(abc);", &ExprStatement{&ParenExpression{&IdentifierExpression{"abc"}}})

	aExp := &IdentifierExpression{"a"}
	bExp := &IdentifierExpression{"b"}
	parseInBlock(t, "a+b;", &ExprStatement{&BinOpExpression{aExp, "+", bExp}})
	parseInBlock(t, "a-b;", &ExprStatement{&BinOpExpression{aExp, "-", bExp}})
	parseInBlock(t, "a*b;", &ExprStatement{&BinOpExpression{aExp, "*", bExp}})
	parseInBlock(t, "a/b;", &ExprStatement{&BinOpExpression{aExp, "/", bExp}})
	parseInBlock(t, "a%b;", &ExprStatement{&BinOpExpression{aExp, "%", bExp}})
	parseInBlock(t, "a&b;", &ExprStatement{&BinOpExpression{aExp, "&", bExp}})
	parseInBlock(t, "a|b;", &ExprStatement{&BinOpExpression{aExp, "|", bExp}})
	parseInBlock(t, "a^b;", &ExprStatement{&BinOpExpression{aExp, "^", bExp}})
	parseInBlock(t, "a<<b;", &ExprStatement{&BinOpExpression{aExp, "<<", bExp}})
	parseInBlock(t, "a>>b;", &ExprStatement{&BinOpExpression{aExp, ">>", bExp}})
	parseInBlock(t, "a&&b;", &ExprStatement{&BinOpExpression{aExp, "&&", bExp}})
	parseInBlock(t, "a||b;", &ExprStatement{&BinOpExpression{aExp, "||", bExp}})
	parseInBlock(t, "a==b;", &ExprStatement{&BinOpExpression{aExp, "==", bExp}})
	parseInBlock(t, "a<b;", &ExprStatement{&BinOpExpression{aExp, "<", bExp}})
	parseInBlock(t, "a>b;", &ExprStatement{&BinOpExpression{aExp, ">", bExp}})
	parseInBlock(t, "a!=b;", &ExprStatement{&BinOpExpression{aExp, "!=", bExp}})
	parseInBlock(t, "a<=b;", &ExprStatement{&BinOpExpression{aExp, "<=", bExp}})
	parseInBlock(t, "a>=b;", &ExprStatement{&BinOpExpression{aExp, ">=", bExp}})

	parseInBlock(t, "timeout_recv(ch);", &ExprStatement{&TimeoutRecvExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInBlock(t, "timeout_peek(ch);", &ExprStatement{&TimeoutPeekExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInBlock(t, "nonblock_recv(ch);", &ExprStatement{&NonblockRecvExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInBlock(t, "nonblock_peek(ch);", &ExprStatement{&NonblockPeekExpression{&IdentifierExpression{"ch"}, []Expression{}}})
	parseInBlock(t, "[a, b];", &ExprStatement{&ArrayExpression{[]Expression{aExp, bExp}}})
}

func parseType(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { var a "+src+"; }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isInitBlock := definitions[0].(*ProcDefinition); isInitBlock {
		if len(procDef.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDefinition", src)
			return
		}
		if stmt, isVarDecl := procDef.Statements[0].(*VarDeclStatement); isVarDecl {
			if !reflect.DeepEqual(stmt.Type, expect) {
				t.Errorf("Expect %+#v \n but got %+#v", expect, stmt.Type)
				return
			}
		} else {
			t.Errorf("Expect %q to be parsed in ProcDefinition", src)
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in ProcDefinition", src)
		return
	}
}

func TestParseType(t *testing.T) {
	parseType(t, "bool", NamedType{"bool"})
	parseType(t, "[]bool", ArrayType{NamedType{"bool"}})
	parseType(t, "channel { bool }", HandshakeChannelType{false, []Type{NamedType{"bool"}}})
	parseType(t, "unstable channel { bool }", HandshakeChannelType{true, []Type{NamedType{"bool"}}})
	parseType(t, "channel [] { bool }",
		BufferedChannelType{false, nil, []Type{NamedType{"bool"}}})
	parseType(t, "channel [1+2] { bool }",
		BufferedChannelType{false, &BinOpExpression{&NumberExpression{"1"}, "+", &NumberExpression{"2"}}, []Type{NamedType{"bool"}}})
	parseType(t, "unstable channel [] { bool }",
		BufferedChannelType{true, nil, []Type{NamedType{"bool"}}})
	parseType(t, "unstable channel [1+2] { bool }",
		BufferedChannelType{true, &BinOpExpression{&NumberExpression{"1"}, "+", &NumberExpression{"2"}}, []Type{NamedType{"bool"}}})
}
