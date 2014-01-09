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
		ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{Pos{}, []InitVar{
		ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
		ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{Pos{}, []InitVar{
		ChannelVar{Pos{}, "a", NamedType{"int"}, nil},
	}}, newTypeEnv())

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{NamedType{"int"}}})
		typeEnv.add("a", NamedType{"int"})
		typeEnv.add("b", NamedType{"bool"})
		expectValid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}}, nil},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "a", []Expression{IdentifierExpression{Pos{}, "a"}}, nil},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}, IdentifierExpression{Pos{}, "a"}}, nil},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "b"}}, nil},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "c"}}, nil},
		}}, typeEnv)
		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "a"}}, nil},
		}}, newTypeEnv())
	}

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{HandshakeChannelType{[]Type{NamedType{"int"}}}}})
		expectValid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expression{IdentifierExpression{Pos{}, "ch"}}, nil},
			ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
		}}, typeEnv)
	}
}
