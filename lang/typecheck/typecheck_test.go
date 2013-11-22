package typecheck

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func universalTypeCheck(x interface{}, env *typeEnv) error {
	switch x := x.(type) {
	case Definition:
		return typeCheckDefinition(x, env)
	case Statement:
		return typeCheckStatement(x, env)
	case Expression:
		return typeCheckExpression(x, env)
	}
	panic("Unknown value")
}

func expectValid(t *testing.T, x interface{}, env *typeEnv) {
	if err := universalTypeCheck(x, env); err != nil {
		t.Errorf("Expect %q to be valid, but got an error %q", x, err)
	}
}

func expectInvalid(t *testing.T, x interface{}, env *typeEnv) {
	if err := universalTypeCheck(x, env); err == nil {
		t.Errorf("Expect %q to be invalid", x)
	}
}
