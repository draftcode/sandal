package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestConstantDefinitionTypeCheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := NumberExpression{"1"}

	expectValid(t, ConstantDefinition{"a", intType, numberExpr}, newTypeEnv())
	expectInvalid(t, ConstantDefinition{"a", boolType, numberExpr}, newTypeEnv())
}

func TestInitBlockTypeCheck(t *testing.T) {
	expectValid(t, InitBlock{[]InitVar{
		ChannelVar{"ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{[]InitVar{
		ChannelVar{"ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
		ChannelVar{"ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
	}}, newTypeEnv())
	expectInvalid(t, InitBlock{[]InitVar{
		ChannelVar{"a", NamedType{"int"}},
	}}, newTypeEnv())

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{NamedType{"int"}}})
		typeEnv.add("a", NamedType{"int"})
		typeEnv.add("b", NamedType{"bool"})
		expectValid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "a", []Expression{IdentifierExpression{"a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"a"}, IdentifierExpression{"a"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"b"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"c"}}},
		}}, typeEnv)
		expectInvalid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"a"}}},
		}}, newTypeEnv())
	}

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{HandshakeChannelType{false, []Type{NamedType{"int"}}}}})
		expectValid(t, InitBlock{[]InitVar{
			InstanceVar{"proc1", "A", []Expression{IdentifierExpression{"ch"}}},
			ChannelVar{"ch", HandshakeChannelType{false, []Type{NamedType{"int"}}}},
		}}, typeEnv)
	}
}
