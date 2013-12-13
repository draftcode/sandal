package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestConstantDefinitionTypeCheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := NumberExpression{Pos{}, "1"}

	expectValid(t, ConstantDefinition{Pos{}, "a", intType, numberExpr}, newTypeEnv())
	expectInvalid(t, ConstantDefinition{Pos{}, "a", boolType, numberExpr}, newTypeEnv())
}

func TestInitBlockTypeCheck(t *testing.T) {
	expectValid(t, InitBlock{Pos{}, []InitVar{
		ChannelVar{Pos{}, "ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{Pos{}, []InitVar{
		ChannelVar{Pos{}, "ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
		ChannelVar{Pos{}, "ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{Pos{}, []InitVar{
		ChannelVar{Pos{}, "a", NamedType{"int"}},
	}}, newTypeEnv())

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{NamedType{"int"}}})
		typeEnv.add("a", NamedType{"int"})
		typeEnv.add("b", NamedType{"bool"})
		expectValid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "a", []Expression{IdentifierExpression{Pos{}, "a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}, IdentifierExpression{Pos{}, "a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "b"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "c"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}}},
		}}, newTypeEnv())
	}

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{HandshakeChannelType{false, []Type{NamedType{"int"}}}}})
		expectValid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "ch"}}},
			ChannelVar{Pos{}, "ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
		}}, typeEnv)
	}
}
