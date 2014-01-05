package conversion

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
		Transition(nextState intState, varName string) []intTransition
		String() string
		GetType() Type
	}

	intInternalPrimitiveVar struct {
		RealName string
		Type     Type
	}

	intInternalHandshakeChannelProxyVar struct {
		RealName   string
		ChannelVar intInternalHandshakeChannelVar
	}

	intInternalBufferedChannelProxyVar struct {
		RealName   string
		ChannelVar intInternalBufferedChannelVar
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
		Pids       map[int]bool
	}

	intInternalBufferedChannelVar struct {
		ModuleName string
		RealName   string
		Type       BufferedChannelType
		Pids       map[int]bool
	}
)

func (x intInternalPrimitiveVar) intinternalobj()             {}
func (x intInternalHandshakeChannelProxyVar) intinternalobj() {}
func (x intInternalBufferedChannelProxyVar) intinternalobj()  {}
func (x intInternalArrayVar) intinternalobj()                 {}
func (x intInternalLiteral) intinternalobj()                  {}
func (x intInternalNot) intinternalobj()                      {}
func (x intInternalUnarySub) intinternalobj()                 {}
func (x intInternalParen) intinternalobj()                    {}
func (x intInternalBinOp) intinternalobj()                    {}
func (x intInternalTimeoutRecv) intinternalobj()              {}
func (x intInternalTimeoutPeek) intinternalobj()              {}
func (x intInternalNonblockRecv) intinternalobj()             {}
func (x intInternalNonblockPeek) intinternalobj()             {}
func (x intInternalArrayLiteral) intinternalobj()             {}
func (x intInternalHandshakeChannelVar) intinternalobj()      {}
func (x intInternalBufferedChannelVar) intinternalobj()       {}

// ========================================
// Steps
// Steps requried to determine the evaluated value of expression.
// TODO: This should be checked beforehand.

func (x intInternalPrimitiveVar) Steps() int             { return 0 }
func (x intInternalHandshakeChannelProxyVar) Steps() int { return 0 }
func (x intInternalBufferedChannelProxyVar) Steps() int  { return 0 }
func (x intInternalArrayVar) Steps() int                 { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternalLiteral) Steps() int                  { return 0 }
func (x intInternalNot) Steps() int                      { return x.Sub.Steps() }
func (x intInternalUnarySub) Steps() int                 { return x.Sub.Steps() }
func (x intInternalParen) Steps() int                    { return x.Sub.Steps() }
func (x intInternalBinOp) Steps() int                    { return x.LHS.Steps() + x.RHS.Steps() }
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
func (x intInternalHandshakeChannelVar) Steps() int {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalBufferedChannelVar) Steps() int {
	panic("Array literals cannot directly be expressed in NuSMV")
}

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

func (x intInternalPrimitiveVar) String() string             { return x.RealName }
func (x intInternalHandshakeChannelProxyVar) String() string { return x.RealName }
func (x intInternalBufferedChannelProxyVar) String() string  { return x.RealName }
func (x intInternalArrayVar) String() string                 { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternalLiteral) String() string                  { return x.Lit }
func (x intInternalNot) String() string                      { return "!" + x.Sub.String() }
func (x intInternalUnarySub) String() string                 { return "-" + x.Sub.String() }
func (x intInternalParen) String() string                    { return "(" + x.Sub.String() + ")" }
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
func (x intInternalHandshakeChannelVar) String() string {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalBufferedChannelVar) String() string {
	panic("Array literals cannot directly be expressed in NuSMV")
}

func (x intInternalArrayLiteral) ArgString() (ret []string) {
	for _, elem := range x.Elems {
		ret = append(ret, elem.String())
	}
	return ret
}

// ========================================
// Transition

func assignByString(x intInternalExpressionObj, nextState intState, varName string) []intTransition {
	if varName == "" {
		return []intTransition{{NextState: nextState}}
	} else {
		return []intTransition{{
			NextState: nextState,
			Actions: []intAssign{
				{LHS: varName, RHS: x.String()},
			},
		}}
	}
}

func (x intInternalPrimitiveVar) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalHandshakeChannelProxyVar) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalBufferedChannelProxyVar) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalArrayVar) Transition(nextState intState, varName string) []intTransition {
	panic("ArrayVar cannot directly be expressed in NuSMV")
}
func (x intInternalLiteral) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalNot) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalUnarySub) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalParen) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalBinOp) Transition(nextState intState, varName string) []intTransition {
	return assignByString(x, nextState, varName)
}
func (x intInternalTimeoutRecv) Transition(nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{NextState: nextState}
	timeoutTrans := intTransition{NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.filled & !%s.received", x.Channel, x.Channel)
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("%s.next_received", x.Channel),
			RHS: "TRUE",
		})
		for i, arg := range x.Args {
			recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
				LHS: fmt.Sprintf("next(%s)", arg),
				RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
			})
		}
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
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
func (x intInternalTimeoutPeek) Transition(nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternalNonblockRecv) Transition(nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{NextState: nextState}
	notRecvedTrans := intTransition{NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.filled & !%s.received", x.Channel, x.Channel)
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("%s.next_received", x.Channel),
			RHS: "TRUE",
		})
		for i, arg := range x.Args {
			recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
				LHS: fmt.Sprintf("next(%s)", arg),
				RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
			})
		}
		notRecvedTrans.Condition = fmt.Sprintf("!(%s.filled & !%s.received)", x.Channel, x.Channel)
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
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
func (x intInternalNonblockPeek) Transition(nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternalArrayLiteral) Transition(nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalHandshakeChannelVar) Transition(nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternalBufferedChannelVar) Transition(nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
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
func (x intInternalHandshakeChannelProxyVar) GetType() Type { return x.ChannelVar.GetType() }
func (x intInternalBufferedChannelProxyVar) GetType() Type  { return x.ChannelVar.GetType() }
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
// Channel Proxy Conversion

func (x intInternalHandshakeChannelVar) Proxy(pid int) intInternalHandshakeChannelProxyVar {
	x.Pids[pid] = true
	return intInternalHandshakeChannelProxyVar{
		RealName:   fmt.Sprintf("__pid%d_%s", pid, x.RealName),
		ChannelVar: x,
	}
}
func (x intInternalBufferedChannelVar) Proxy(pid int) intInternalBufferedChannelProxyVar {
	x.Pids[pid] = true
	return intInternalBufferedChannelProxyVar{
		RealName:   fmt.Sprintf("__pid%d_%s", pid, x.RealName),
		ChannelVar: x,
	}
}

func changeToProxy(intExprObj intInternalExpressionObj, pid int) intInternalExpressionObj {
	switch intExprObj := intExprObj.(type) {
	case intInternalPrimitiveVar:
		return intExprObj
	case intInternalHandshakeChannelProxyVar:
		panic("unexpected")
	case intInternalBufferedChannelProxyVar:
		panic("unexpected")
	case intInternalArrayVar:
		panic("unexpected")
	case intInternalLiteral:
		return intExprObj
	case intInternalNot:
		return intInternalNot{Sub: changeToProxy(intExprObj.Sub, pid)}
	case intInternalUnarySub:
		return intInternalUnarySub{Sub: changeToProxy(intExprObj.Sub, pid)}
	case intInternalParen:
		return intInternalParen{Sub: changeToProxy(intExprObj.Sub, pid)}
	case intInternalBinOp:
		return intInternalBinOp{
			LHS: changeToProxy(intExprObj.LHS, pid),
			Op:  intExprObj.Op,
			RHS: changeToProxy(intExprObj.RHS, pid),
		}
	case intInternalTimeoutRecv:
		panic("unexpected")
	case intInternalTimeoutPeek:
		panic("unexpected")
	case intInternalNonblockRecv:
		panic("unexpected")
	case intInternalNonblockPeek:
		panic("unexpected")
	case intInternalArrayLiteral:
		elems := []intInternalExpressionObj{}
		for _, elem := range intExprObj.Elems {
			elems = append(elems, changeToProxy(elem, pid))
		}
		return intInternalArrayLiteral{Elems: elems}
	case intInternalHandshakeChannelVar:
		return intExprObj.Proxy(pid)
	case intInternalBufferedChannelVar:
		return intExprObj.Proxy(pid)
	default:
		panic("unexpected")
	}
}
