package typecheck

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

func TypeCheck(defs []Definition) error {
	return typeCheckDefinitions(defs, newTypeEnv())
}

// ========================================

type typeEnv struct {
	upper *typeEnv
	scope map[string]Type
}

func newTypeEnv() (ret *typeEnv) {
	ret = new(typeEnv)
	ret.scope = make(map[string]Type)
	return
}

func newTypeEnvFromUpper(upper *typeEnv) (ret *typeEnv) {
	ret = newTypeEnv()
	ret.upper = upper
	return
}

func (env *typeEnv) add(name string, ty Type) {
	env.scope[name] = ty
}

func (env *typeEnv) lookup(name string) Type {
	if ty, found := env.scope[name]; found {
		return ty
	}
	if env.upper != nil {
		return env.upper.lookup(name)
	} else {
		return nil
	}
}

func channelExprCheck(ch ChanExpr, env *typeEnv, recvOrPeek bool) error {
	chExpr := ch.ChannelExpr()
	args := ch.ArgExprs()
	if err := typeCheckExpression(chExpr, env); err != nil {
		return err
	}
	for _, arg := range args {
		if err := typeCheckExpression(arg, env); err != nil {
			return err
		}
	}

	var elemTypes []Type
	switch ty := typeOfExpression(chExpr, env).(type) {
	case HandshakeChannelType:
		elemTypes = ty.Elems
	case BufferedChannelType:
		elemTypes = ty.Elems
	default:
		return fmt.Errorf("Expect the first argument of %s to be a channel but got %s",
			ch, typeOfExpression(chExpr, env))
	}

	if len(elemTypes) != len(args) {
		return fmt.Errorf("Expect the arugments of %s to have %d elements",
			ch, len(elemTypes))
	}
	for i := 0; i < len(elemTypes); i++ {
		if !typeOfExpression(args[i], env).Equal(elemTypes[i]) {
			return fmt.Errorf("Expect the argument %s to be a %s", args[i], elemTypes[i])
		}
		if recvOrPeek {
			if _, isIdentExpr := args[i].(IdentifierExpression); !isIdentExpr {
				return fmt.Errorf("Expect the argument %s to be an identifier", args[i])
			}
		}
	}
	return nil
}
