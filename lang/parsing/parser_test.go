package parsing

import (
	"github.com/cookieo9/go-misc/pp"
	. "github.com/draftcode/sandal/lang/data"
	"reflect"
	"testing"
)

func parse(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune(src), 0)
	definitions := Parse(s)
	expectPP := pp.PP(expect)
	actualPP := pp.PP(definitions)
	if expectPP != actualPP {
		t.Errorf("\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}

func TestDataDefinition(t *testing.T) {
	parse(t, "data Maybe { Just, Nothing };",
		[]Definition{DataDefinition{Pos{1, 1}, "Maybe", []string{"Just", "Nothing"}}})
}

func TestParseModuleDefinition(t *testing.T) {
	parse(t, "module A(ch channel { bool }, chs []channel { bit }) { init { }; };",
		[]Definition{ModuleDefinition{Pos{1, 1}, "A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{[]Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{[]Type{NamedType{"bit"}}}}}},
			[]Definition{InitBlock{Pos: Pos{1, 56}}}}})
}

func TestParseConstantDefinition(t *testing.T) {
	parse(t, "const a int = 1;", []Definition{ConstantDefinition{Pos{1, 1}, "a", NamedType{"int"}, NumberExpression{Pos{1, 15}, "1"}}})
}

func TestParseProcDefinition(t *testing.T) {
	parse(t, "proc A(ch channel { bool }, chs []channel { bit }) { ; };",
		[]Definition{ProcDefinition{Pos{1, 1}, "A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{[]Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{[]Type{NamedType{"bit"}}}}}},
			[]Statement{NullStatement{Pos{1, 54}}}}})
}

func TestParseInitBlock(t *testing.T) {
	parse(t, "init { };", []Definition{InitBlock{Pos: Pos{1, 1}}})
	parse(t, "init { a : M(b) };",
		[]Definition{InitBlock{Pos{1, 1}, []InitVar{
			InstanceVar{Pos{1, 8}, "a", "M", []Expression{IdentifierExpression{Pos{1, 14}, "b"}}, []string{}},
		}}})
	parse(t, "init { a : M(b) @unstable };",
		[]Definition{InitBlock{Pos{1, 1}, []InitVar{
			InstanceVar{Pos{1, 8}, "a", "M", []Expression{IdentifierExpression{Pos{1, 14}, "b"}}, []string{"unstable"}},
		}}})
	parse(t, "init { a : channel { bool } };",
		[]Definition{InitBlock{Pos{1, 1}, []InitVar{
			ChannelVar{Pos{1, 8}, "a", HandshakeChannelType{[]Type{NamedType{"bool"}}}, []string{}},
		}}})
	parse(t, "init { a : channel { bool } @unstable };",
		[]Definition{InitBlock{Pos{1, 1}, []InitVar{
			ChannelVar{Pos{1, 8}, "a", HandshakeChannelType{[]Type{NamedType{"bool"}}}, []string{"unstable"}},
		}}})
}

const parseBlockOffset = 11

func parseInBlock(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { "+src+" }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isProcDef := definitions[0].(ProcDefinition); isProcDef {
		if len(procDef.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDefinition", src)
			return
		}
		if !reflect.DeepEqual(procDef.Statements[0], expect) {
			t.Errorf("\nExpected %s\nGot      %s", pp.PP(expect), pp.PP(procDef.Statements[0]))
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in ProcDefinition", src)
		return
	}
}

func TestParseStatement(t *testing.T) {
	parseInBlock(t, "test: ;", LabelledStatement{Pos{1, 1+parseBlockOffset}, "test", NullStatement{Pos{1, 7+parseBlockOffset}}})
	parseInBlock(t, "{ ; };", BlockStatement{Pos{1, 1+parseBlockOffset}, []Statement{NullStatement{Pos{1, 3+parseBlockOffset}}}})
	parseInBlock(t, "var abc bool;", VarDeclStatement{Pos{1, 1+parseBlockOffset}, "abc", NamedType{"bool"}, nil})
	parseInBlock(t, "var abc bool = false;", VarDeclStatement{Pos{1, 1+parseBlockOffset}, "abc", NamedType{"bool"}, FalseExpression{Pos{1, 16+parseBlockOffset}}})
	parseInBlock(t, "if false { ; };", IfStatement{Pos{1, 1+parseBlockOffset}, FalseExpression{Pos{1, 4+parseBlockOffset}}, []Statement{NullStatement{Pos{1, 12+parseBlockOffset}}}, nil})
	parseInBlock(t, "if false { ; } else { skip; };", IfStatement{Pos{1, 1+parseBlockOffset}, FalseExpression{Pos{1, 4+parseBlockOffset}}, []Statement{NullStatement{Pos{1, 12+parseBlockOffset}}}, []Statement{SkipStatement{Pos{1, 23+parseBlockOffset}}}})

	parseInBlock(t, "a=b;", AssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", IdentifierExpression{Pos{1, 3+parseBlockOffset}, "b"}})
	parseInBlock(t, "a+=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "+", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a-=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "-", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a*=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "*", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a/=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "/", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a%=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "%", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a&=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "&", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a|=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "|", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a^=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "^", IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}})
	parseInBlock(t, "a<<=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", "<<", IdentifierExpression{Pos{1, 5+parseBlockOffset}, "b"}})
	parseInBlock(t, "a>>=b;", OpAssignmentStatement{Pos{1, 1+parseBlockOffset}, "a", ">>", IdentifierExpression{Pos{1, 5+parseBlockOffset}, "b"}})

	parseInBlock(t, "choice { ; }, { skip; };", ChoiceStatement{Pos{1, 1+parseBlockOffset}, []BlockStatement{BlockStatement{Pos{1, 8+parseBlockOffset}, []Statement{NullStatement{Pos{1, 10+parseBlockOffset}}}}, BlockStatement{Pos{1, 15+parseBlockOffset}, []Statement{SkipStatement{Pos{1, 17+parseBlockOffset}}}}}})
	parseInBlock(t, "recv(ch, 1, 2);", RecvStatement{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 6+parseBlockOffset}, "ch"}, []Expression{NumberExpression{Pos{1, 10+parseBlockOffset}, "1"}, NumberExpression{Pos{1, 13+parseBlockOffset}, "2"}}})
	parseInBlock(t, "peek(ch);", PeekStatement{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 6+parseBlockOffset}, "ch"}, []Expression{}})
	parseInBlock(t, "send(ch, 1, 2);", SendStatement{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 6+parseBlockOffset}, "ch"}, []Expression{NumberExpression{Pos{1, 10+parseBlockOffset}, "1"}, NumberExpression{Pos{1, 13+parseBlockOffset}, "2"}}})
	parseInBlock(t, "for { ; };", ForStatement{Pos{1, 1+parseBlockOffset}, []Statement{NullStatement{Pos{1, 7+parseBlockOffset}}}})
	parseInBlock(t, "for ch in chs { ; };", ForInStatement{Pos{1, 1+parseBlockOffset}, "ch", IdentifierExpression{Pos{1, 11+parseBlockOffset}, "chs"}, []Statement{NullStatement{Pos{1, 17+parseBlockOffset}}}})
	parseInBlock(t, "for i in range 1 to 5 { ; };", ForInRangeStatement{Pos{1, 1+parseBlockOffset}, "i", NumberExpression{Pos{1, 16+parseBlockOffset}, "1"}, NumberExpression{Pos{1, 21+parseBlockOffset}, "5"}, []Statement{NullStatement{Pos{1, 25+parseBlockOffset}}}})
	parseInBlock(t, "break;", BreakStatement{Pos{1, 1+parseBlockOffset}})
	parseInBlock(t, "goto here;", GotoStatement{Pos{1, 1+parseBlockOffset}, "here"})
	parseInBlock(t, "skip;", SkipStatement{Pos{1, 1+parseBlockOffset}})
	parseInBlock(t, ";", NullStatement{Pos{1, 1+parseBlockOffset}})
	parseInBlock(t, "1;", ExprStatement{NumberExpression{Pos{1, 1+parseBlockOffset}, "1"}})
	parseInBlock(t, "const a int = 1;", ConstantDefinition{Pos{1, 1+parseBlockOffset}, "a", NamedType{"int"}, NumberExpression{Pos{1, 15+parseBlockOffset}, "1"}})
}

func TestParseExpression(t *testing.T) {
	parseInBlock(t, "abc;", ExprStatement{IdentifierExpression{Pos{1, 1+parseBlockOffset}, "abc"}})
	parseInBlock(t, "123;", ExprStatement{NumberExpression{Pos{1, 1+parseBlockOffset}, "123"}})
	parseInBlock(t, "true;", ExprStatement{TrueExpression{Pos{1, 1+parseBlockOffset}}})
	parseInBlock(t, "false;", ExprStatement{FalseExpression{Pos{1, 1+parseBlockOffset}}})
	parseInBlock(t, "!abc;", ExprStatement{NotExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 2+parseBlockOffset}, "abc"}}})
	parseInBlock(t, "-abc;", ExprStatement{UnarySubExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 2+parseBlockOffset}, "abc"}}})
	parseInBlock(t, "(abc);", ExprStatement{ParenExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 2+parseBlockOffset}, "abc"}}})

	aExp := IdentifierExpression{Pos{1, 1+parseBlockOffset}, "a"}
	bExp := IdentifierExpression{Pos{1, 4+parseBlockOffset}, "b"}
	parseInBlock(t, "a+ b;", ExprStatement{BinOpExpression{aExp, "+", bExp}})
	parseInBlock(t, "a- b;", ExprStatement{BinOpExpression{aExp, "-", bExp}})
	parseInBlock(t, "a* b;", ExprStatement{BinOpExpression{aExp, "*", bExp}})
	parseInBlock(t, "a/ b;", ExprStatement{BinOpExpression{aExp, "/", bExp}})
	parseInBlock(t, "a% b;", ExprStatement{BinOpExpression{aExp, "%", bExp}})
	parseInBlock(t, "a& b;", ExprStatement{BinOpExpression{aExp, "&", bExp}})
	parseInBlock(t, "a| b;", ExprStatement{BinOpExpression{aExp, "|", bExp}})
	parseInBlock(t, "a^ b;", ExprStatement{BinOpExpression{aExp, "^", bExp}})
	parseInBlock(t, "a<<b;", ExprStatement{BinOpExpression{aExp, "<<", bExp}})
	parseInBlock(t, "a>>b;", ExprStatement{BinOpExpression{aExp, ">>", bExp}})
	parseInBlock(t, "a&&b;", ExprStatement{BinOpExpression{aExp, "&&", bExp}})
	parseInBlock(t, "a||b;", ExprStatement{BinOpExpression{aExp, "||", bExp}})
	parseInBlock(t, "a==b;", ExprStatement{BinOpExpression{aExp, "==", bExp}})
	parseInBlock(t, "a< b;", ExprStatement{BinOpExpression{aExp, "<", bExp}})
	parseInBlock(t, "a> b;", ExprStatement{BinOpExpression{aExp, ">", bExp}})
	parseInBlock(t, "a!=b;", ExprStatement{BinOpExpression{aExp, "!=", bExp}})
	parseInBlock(t, "a<=b;", ExprStatement{BinOpExpression{aExp, "<=", bExp}})
	parseInBlock(t, "a>=b;", ExprStatement{BinOpExpression{aExp, ">=", bExp}})

	parseInBlock(t, "timeout_recv(ch);", ExprStatement{TimeoutRecvExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 14+parseBlockOffset}, "ch"}, []Expression{}}})
	parseInBlock(t, "timeout_peek(ch);", ExprStatement{TimeoutPeekExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 14+parseBlockOffset}, "ch"}, []Expression{}}})
	parseInBlock(t, "nonblock_recv(ch);", ExprStatement{NonblockRecvExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 15+parseBlockOffset}, "ch"}, []Expression{}}})
	parseInBlock(t, "nonblock_peek(ch);", ExprStatement{NonblockPeekExpression{Pos{1, 1+parseBlockOffset}, IdentifierExpression{Pos{1, 15+parseBlockOffset}, "ch"}, []Expression{}}})
	parseInBlock(t, "[a, b];", ExprStatement{ArrayExpression{Pos{1, 1+parseBlockOffset}, []Expression{
		IdentifierExpression{Pos{1, 2+parseBlockOffset}, "a"},
		IdentifierExpression{Pos{1, 5+parseBlockOffset}, "b"},
	}}})
}

const parseTypeOffset = 17

func parseType(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { var a "+src+"; }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isInitBlock := definitions[0].(ProcDefinition); isInitBlock {
		if len(procDef.Statements) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDefinition", src)
			return
		}
		if stmt, isVarDecl := procDef.Statements[0].(VarDeclStatement); isVarDecl {
			if !reflect.DeepEqual(stmt.Type, expect) {
				t.Errorf("\nExpected %s\nGot      %s", pp.PP(expect), pp.PP(stmt.Type))
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
	parseType(t, "channel { bool }", HandshakeChannelType{[]Type{NamedType{"bool"}}})
	parseType(t, "channel [] { bool }",
		BufferedChannelType{nil, []Type{NamedType{"bool"}}})
	parseType(t, "channel [1+2] { bool }",
		BufferedChannelType{BinOpExpression{NumberExpression{Pos{1, 10+parseTypeOffset}, "1"}, "+", NumberExpression{Pos{1, 12+parseTypeOffset}, "2"}}, []Type{NamedType{"bool"}}})
}
