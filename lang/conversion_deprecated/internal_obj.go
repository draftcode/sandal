package conversion_deprecated

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

type (
	intInternalObj interface {
		intinternalobj()
	}
)

// ========================================
// intInternalObj

type (
	intInternalConstantDef struct {
		Type Type
		Expr Expression
	}

	intInternalDataTypeDef struct {
		Elems []string
	}

	intInternalProcDef struct {
		Def ProcDefinition
	}

	intInternalProcVar struct {
		Name       string
		ModuleName string
		Def        intInternalProcDef
		Args       []intInternalExpressionObj
		Pid        int
	}
)

func (x intInternalConstantDef) intinternalobj() {}
func (x intInternalDataTypeDef) intinternalobj() {}
func (x intInternalProcDef) intinternalobj()     {}
func (x intInternalProcVar) intinternalobj()     {}

// ========================================
// intInternalExpressionObj

type (
	intInternalExpressionObj interface {
		intInternalObj
		Steps() int
		Transition(fromState, nextState intState, varName string) []intTransition
		String() string
		GetType() Type
	}

	intInternalPrimitiveVar struct {
		RealName string
		Type     Type
		RealObj  intInternalExpressionObj
	}

	intInternalArrayVar struct {
		RealName    string
		RealLiteral intInternalArrayLiteral
	}

	intInternalLiteral struct {
		Lit  string
		Type Type
	}

	intInternalNot struct {
		Sub intInternalExpressionObj
	}

	intInternalUnarySub struct {
		Sub intInternalExpressionObj
	}

	intInternalParen struct {
		Sub intInternalExpressionObj
	}

	intInternalBinOp struct {
		LHS intInternalExpressionObj
		Op  string
		RHS intInternalExpressionObj
	}

	intInternalTimeoutRecv struct {
		Channel intInternalExpressionObj
		Args    []intInternalExpressionObj
	}

	intInternalTimeoutPeek struct {
		Channel intInternalExpressionObj
		Args    []intInternalExpressionObj
	}

	intInternalNonblockRecv struct {
		Channel intInternalExpressionObj
		Args    []intInternalExpressionObj
	}

	intInternalNonblockPeek struct {
		Channel intInternalExpressionObj
		Args    []intInternalExpressionObj
	}

	intInternalArrayLiteral struct {
		Elems []intInternalExpressionObj
	}

	intInternalHandshakeChannelVar struct {
		ModuleName string
		RealName   string
		Type       HandshakeChannelType
		Tags       []string
		Pids       map[int]bool
	}

	intInternalBufferedChannelVar struct {
		ModuleName string
		RealName   string
		Type       BufferedChannelType
		Tags       []string
		Pids       map[int]bool
	}
)

func (x intInternalPrimitiveVar) intinternalobj()        {}
func (x intInternalArrayVar) intinternalobj()            {}
func (x intInternalLiteral) intinternalobj()             {}
func (x intInternalNot) intinternalobj()                 {}
func (x intInternalUnarySub) intinternalobj()            {}
func (x intInternalParen) intinternalobj()               {}
func (x intInternalBinOp) intinternalobj()               {}
func (x intInternalTimeoutRecv) intinternalobj()         {}
func (x intInternalTimeoutPeek) intinternalobj()         {}
func (x intInternalNonblockRecv) intinternalobj()        {}
func (x intInternalNonblockPeek) intinternalobj()        {}
func (x intInternalArrayLiteral) intinternalobj()        {}
func (x intInternalHandshakeChannelVar) intinternalobj() {}
func (x intInternalBufferedChannelVar) intinternalobj()  {}

// ========================================
// Steps
// Steps requried to determine the evaluated value of expression.
// TODO: This should be checked beforehand.

func (x intInternalPrimitiveVar) Steps() int { return 0 }
func (x intInternalArrayVar) Steps() int     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternalLiteral) Steps() int      { return 0 }
func (x intInternalNot) Steps() int          { return x.Sub.Steps() }
func (x intInternalUnarySub) Steps() int     { return x.Sub.Steps() }
func (x intInternalParen) Steps() int        { return x.Sub.Steps() }
func (x intInternalBinOp) Steps() int        { return x.LHS.Steps() + x.RHS.Steps() }
func (x intInternalTimeoutRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternalTimeoutPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternalNonblockRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternalNonblockPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternalArrayLiteral) Steps() int {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalHandshakeChannelVar) Steps() int { return 0 }
func (x intInternalBufferedChannelVar) Steps() int  { return 0 }

// ========================================
// String
// Used for converting internal objects to NuSMV expression.

var operatorConversionTable = map[string]string{
	"+":  "+",
	"-":  "-",
	"*":  "*",
	"/":  "/",
	"%":  "mod",
	"&":  "&",
	"|":  "|",
	"^":  "xor",
	"<<": "<<",
	">>": ">>",
	"&&": "&",
	"||": "|",
	"==": "=",
	"<":  "<",
	">":  ">",
	"!=": "!=",
	"<=": "<=",
	">=": ">=",
}

func (x intInternalPrimitiveVar) String() string { return x.RealName }
func (x intInternalArrayVar) String() string     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternalLiteral) String() string      { return x.Lit }
func (x intInternalNot) String() string          { return "!" + x.Sub.String() }
func (x intInternalUnarySub) String() string     { return "-" + x.Sub.String() }
func (x intInternalParen) String() string        { return "(" + x.Sub.String() + ")" }
func (x intInternalBinOp) String() string {
	// TODO: this cannot encode nonblock_recv(...) && nonblock_recv(...)
	return x.LHS.String() + operatorConversionTable[x.Op] + x.RHS.String()
}
func (x intInternalTimeoutRecv) String() string {
	panic("timeout_recv cannot directly be expressed in NuSMV")
}
func (x intInternalTimeoutPeek) String() string {
	panic("timeout_peek cannot directly be expressed in NuSMV")
}
func (x intInternalNonblockRecv) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x intInternalNonblockPeek) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x intInternalArrayLiteral) String() string {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalHandshakeChannelVar) String() string { return x.RealName }
func (x intInternalBufferedChannelVar) String() string  { return x.RealName }

func (x intInternalArrayLiteral) ArgString() (ret []string) {
	for _, elem := range x.Elems {
		ret = append(ret, elem.String())
	}
	return ret
}

// ========================================
// Transition

func assignByString(x intInternalExpressionObj, fromState, nextState intState, varName string) []intTransition {
	if varName == "" {
		return []intTransition{{FromState: fromState, NextState: nextState}}
	} else {
		return []intTransition{{
			FromState: fromState,
			NextState: nextState,
			Actions: []intAssign{
				{LHS: varName, RHS: x.String()},
			},
		}}
	}
}

func (x intInternalPrimitiveVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalArrayVar) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("ArrayVar cannot directly be expressed in NuSMV")
}
func (x intInternalLiteral) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalNot) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalUnarySub) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalParen) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalBinOp) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalTimeoutRecv) Transition(fromState, nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{FromState: fromState, NextState: nextState}
	timeoutTrans := intTransition{FromState: fromState, NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready & !%s.received", x.Channel, x.Channel)
	case BufferedChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready", x.Channel)
	default:
		panic("unknown channel type")
	}
	recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
		LHS: fmt.Sprintf("%s.recv_received", x.Channel),
		RHS: "TRUE",
	})
	for i, arg := range x.Args {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("next(%s)", arg),
			RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
		})
	}
	if varName != "" {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: varName, RHS: "TRUE",
		})
		timeoutTrans.Actions = append(timeoutTrans.Actions, intAssign{
			LHS: varName, RHS: "FALSE",
		})
	}
	return []intTransition{recvedTrans, timeoutTrans}
}
func (x intInternalTimeoutPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternalNonblockRecv) Transition(fromState, nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{FromState: fromState, NextState: nextState}
	notRecvedTrans := intTransition{FromState: fromState, NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready & !%s.received", x.Channel, x.Channel)
	case BufferedChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready", x.Channel)
	default:
		panic("unknown channel type")
	}
	notRecvedTrans.Condition = "!(" + recvedTrans.Condition + ")"
	recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
		LHS: fmt.Sprintf("%s.recv_received", x.Channel),
		RHS: "TRUE",
	})
	for i, arg := range x.Args {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("next(%s)", arg),
			RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
		})
	}
	if varName != "" {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: varName, RHS: "TRUE",
		})
		notRecvedTrans.Actions = append(notRecvedTrans.Actions, intAssign{
			LHS: varName, RHS: "FALSE",
		})
	}
	return []intTransition{recvedTrans, notRecvedTrans}
}
func (x intInternalNonblockPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternalArrayLiteral) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalHandshakeChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternalBufferedChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}

// ========================================
// GetType

var operatorResultType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": NamedType{"bool"},
	"<":  NamedType{"bool"},
	">":  NamedType{"bool"},
	"!=": NamedType{"bool"},
	"<=": NamedType{"bool"},
	">=": NamedType{"bool"},
}

func (x intInternalPrimitiveVar) GetType() Type             { return x.Type }
func (x intInternalArrayVar) GetType() Type                 { return x.RealLiteral.GetType() }
func (x intInternalLiteral) GetType() Type                  { return x.Type }
func (x intInternalNot) GetType() Type                      { return x.Sub.GetType() }
func (x intInternalUnarySub) GetType() Type                 { return x.Sub.GetType() }
func (x intInternalParen) GetType() Type                    { return x.Sub.GetType() }
func (x intInternalBinOp) GetType() Type                    { return operatorResultType[x.Op] }
func (x intInternalTimeoutRecv) GetType() Type              { return NamedType{"bool"} }
func (x intInternalTimeoutPeek) GetType() Type              { return NamedType{"bool"} }
func (x intInternalNonblockRecv) GetType() Type             { return NamedType{"bool"} }
func (x intInternalNonblockPeek) GetType() Type             { return NamedType{"bool"} }
func (x intInternalArrayLiteral) GetType() Type             { return ArrayType{x.Elems[0].GetType()} }
func (x intInternalHandshakeChannelVar) GetType() Type      { return x.Type }
func (x intInternalBufferedChannelVar) GetType() Type       { return x.Type }

// ========================================

func resolveRealObj(obj intInternalExpressionObj) intInternalExpressionObj {
	for {
		if primVarObj, isPrimVarObj := obj.(intInternalPrimitiveVar); isPrimVarObj {
			if primVarObj.RealObj != nil {
				obj = primVarObj.RealObj
			} else {
				return obj
			}
		} else {
			return obj
		}
	}
}
