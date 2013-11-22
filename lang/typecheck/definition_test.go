package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestConstantDefinitionTypeCheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := &NumberExpression{"1"}

	expectValid(t, &ConstantDefinition{"a", intType, numberExpr}, newTypeEnv())
	expectInvalid(t, &ConstantDefinition{"a", boolType, numberExpr}, newTypeEnv())
}

func TestInitBlockTypeCheck(t *testing.T) {
	t.Errorf("Not implemented")
}
