package sandal

import (
	"testing"
)

func TestConstantDefinitionTypecheck(t *testing.T) {
	intType := &NamedType{"int"}
	numberExpr := &NumberExpression{"1"}

	{
		def := &ConstantDefinition{"a", intType, numberExpr}
		if err := def.typecheck(NewTypeEnv()); err != nil {
			t.Errorf("Expect \"const a int = 1\" to be valid, but got an error %s", err.Error())
		}
	}
}
