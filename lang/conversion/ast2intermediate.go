package conversion

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
	"strconv"
	"strings"
)

func convertASTToIntModule(defs []Definition) (error, []intModule) {
	converter := newIntModConverter()
	for _, def := range defs {
		switch def := def.(type) {
		case *DataDefinition:
			// TODO
		case *ModuleDefinition:
			// TODO
		case *ConstantDefinition:
			converter.env.add(def.Name, intInternalConstant{
				Type: def.Type,
				Expr: def.Expr,
			})
		case *ProcDefinition:
			converter.env.add(def.Name, intInternalProcDef{
				Def: def,
			})
		case *InitBlock:
			// Do nothing
		}
	}
	for _, def := range defs {
		switch def := def.(type) {
		case *DataDefinition:
		case *ModuleDefinition:
		case *ConstantDefinition:
		case *ProcDefinition:
		case *InitBlock:
			converter.convertInitBlock(def)
		}
	}
	if err := converter.buildMainModule(); err != nil {
		return err, nil
	}
	return nil, converter.modules
}

// ========================================

type intModConverter struct {
	env      *varEnv
	channels []intInternalChannelVal
	procs    []intInternalProcVal
	modules  []intModule
	pid      int
}

func newIntModConverter() (converter *intModConverter) {
	converter = new(intModConverter)
	converter.env = newVarEnv()
	return
}

func (x *intModConverter) pushEnv() {
	x.env = newVarEnvFromUpper(x.env)
}

func (x *intModConverter) popEnv() {
	x.env = x.env.upper
}

func (x *intModConverter) convertInitBlock(def *InitBlock) error {
	x.pushEnv()
	defer x.popEnv()
	for _, initVar := range def.Vars {
		switch initVar := initVar.(type) {
		case InstanceVar:
			// Do nothing
		case ChannelVar:
			err, chVal := x.buildChannelVal(initVar.Name, initVar.Type)
			if err != nil {
				return err
			}
			x.env.add(initVar.Name, chVal)
		default:
			panic("Unknown InitVar")
		}
	}
	for _, initVar := range def.Vars {
		switch initVar := initVar.(type) {
		case InstanceVar:
			err, _ := x.buildProcVal(initVar)
			if err != nil {
				return err
			}
		case ChannelVar:
			// Do nothing
		}
	}
	return nil
}

// Convert expressions into NuSMV's basic_expr
func (x *intModConverter) convertBasicExpr(expr Expression) (error, []string) {
	switch expr := expr.(type) {
	case *IdentifierExpression:
		val := x.env.lookup(expr.Name)
		if val == nil {
			// Since it is typechecked. This shouldn't be happened.
			panic("Undefined variable")
		}
		switch val.(type) {
		case intInternalChannelVal:
			return nil, []string{fmt.Sprintf("__pid%d_%s", x.pid, expr.Name)}
		case intInternalConstant:
			// TODO
			panic("Not implemented")
		default:
			return nil, []string{expr.Name}
		}
	case *NumberExpression:
		return nil, []string{expr.Lit}
	case *NotExpression:
		err, basicExprs := x.convertBasicExpr(expr.SubExpr)
		if err != nil {
			return err, nil
		} else if len(basicExprs) != 1 {
			return fmt.Errorf("Expect %s to be basic_expr", expr.SubExpr.String()), nil
		}
		basicExprs[0] = "!" + basicExprs[0]
		return nil, basicExprs
	case *UnarySubExpression:
		err, basicExprs := x.convertBasicExpr(expr.SubExpr)
		if err != nil {
			return err, nil
		} else if len(basicExprs) != 1 {
			return fmt.Errorf("Expect %s to be basic_expr", expr.SubExpr.String()), nil
		}
		basicExprs[0] = "-" + basicExprs[0]
		return nil, basicExprs
	case *ParenExpression:
		err, basicExprs := x.convertBasicExpr(expr.SubExpr)
		if err != nil {
			return err, nil
		}
		for i, basicExpr := range basicExprs {
			basicExprs[i] = "(" + basicExpr + ")"
		}
		return nil, basicExprs
	case *BinOpExpression:
		err, basicExprsLHS := x.convertBasicExpr(expr.LHS)
		if err != nil {
			return err, nil
		} else if len(basicExprsLHS) != 1 {
			return fmt.Errorf("Expect %s to be basic_expr", expr.LHS.String()), nil
		}
		err, basicExprsRHS := x.convertBasicExpr(expr.RHS)
		if err != nil {
			return err, nil
		} else if len(basicExprsRHS) != 1 {
			return fmt.Errorf("Expect %s to be basic_expr", expr.RHS.String()), nil
		}
		return nil, []string{basicExprsLHS[0] + expr.Operator + basicExprsRHS[0]}
	case *TimeoutRecvExpression:
		return fmt.Errorf("timeout_recv cannot be appeared"), nil
	case *TimeoutPeekExpression:
		return fmt.Errorf("timeout_peek cannot be appeared"), nil
	case *NonblockRecvExpression:
		return fmt.Errorf("nonblock_recv cannot be appeared"), nil
	case *NonblockPeekExpression:
		return fmt.Errorf("nonblock_peek cannot be appeared"), nil
	case *ArrayExpression:
		basicExprs := []string{strconv.Itoa(len(expr.Elems))}
		for _, expr := range expr.Elems {
			err, basicSubExprs := x.convertBasicExpr(expr)
			if err != nil {
				return err, nil
			} else if len(basicSubExprs) != 1 {
				return fmt.Errorf("Expect %s to be basic_expr", expr.String()), nil
			}
			basicExprs = append(basicExprs, basicSubExprs...)
		}
		return nil, basicExprs
	default:
		panic("Unknown expression")
	}
}

func (x *intModConverter) buildMainModule() error {
	if len(x.procs) == 0 {
		return fmt.Errorf("No running procs")
	}
	pids := make([]string, len(x.procs))
	for i, proc := range x.procs {
		pids[i] = strconv.Itoa(proc.Pid)
	}

	module := intMainModule{}
	// Vars
	for _, chVal := range x.channels {
		module.Vars = append(module.Vars, intVar{
			Name: chVal.Name,
			Type: fmt.Sprintf("%s(running_pid, %s_filled, %s_received, %s_value)",
				chVal.ModuleName, chVal.Name, chVal.Name, chVal.Name),
		})
		for _, pid := range pids {
			module.Vars = append(module.Vars, intVar{
				Name: fmt.Sprintf("__pid%s_%s", pid, chVal.Name),
				Type: fmt.Sprintf("%sProxy(%s)", chVal.ModuleName, chVal.Name),
			})
		}
	}
	for _, procVal := range x.procs {
		args := []string{"running_pid", strconv.Itoa(procVal.Pid)}
		args = append(args, procVal.Args...)
		module.Vars = append(module.Vars, intVar{
			Name: procVal.Name,
			Type: fmt.Sprintf("%s(%s)", procVal.ModuleName, argJoin(args)),
		})
	}
	module.Vars = append(module.Vars, intVar{"running_pid", "{" + argJoin(pids) + "}"})

	// Assigns
	module.Assigns = append(module.Assigns, intAssign{"running_pid", "{" + argJoin(pids) + "}"})

	// Defs
	for _, chVal := range x.channels {
		nextFilled := []string{}
		nextReceived := []string{}
		nextValue := []string{}
		for _, pid := range pids {
			nextFilled = append(nextFilled, fmt.Sprintf("__pid%s_%s.next_filled", pid, chVal.Name))
			nextReceived = append(nextReceived, fmt.Sprintf("__pid%s_%s.next_received", pid, chVal.Name))
			nextValue = append(nextValue, fmt.Sprintf("__pid%s_%s.next_value", pid, chVal.Name))
		}
		module.Defs = append(module.Defs, intAssign{chVal.Name + "_filled", "[" + argJoin(nextFilled) + "]"})
		module.Defs = append(module.Defs, intAssign{chVal.Name + "_received", "[" + argJoin(nextReceived) + "]"})
		module.Defs = append(module.Defs, intAssign{chVal.Name + "_value", "[" + argJoin(nextValue) + "]"})
	}

	x.modules = append(x.modules, module)
	return nil
}

func (x *intModConverter) buildChannelVal(name string, ty Type) (error, intInternalChannelVal) {
	chNumber := len(x.channels)
	chTypeName := ""
	switch ty.(type) {
	case HandshakeChannelType:
		chTypeName = "HandshakeChannel"
	case BufferedChannelType:
		chTypeName = "BufferedChannel"
	default:
		panic("Unknown channel type")
	}
	chModName := fmt.Sprintf("%s%d", chTypeName, chNumber)
	val := intInternalChannelVal{Name: name, ModuleName: chModName}
	x.channels = append(x.channels, val)

	var mod intModule
	switch ty := ty.(type) {
	case HandshakeChannelType:
		types := []string{}
		for _, elem := range ty.Elems {
			types = append(types, x.convertTypeToString(elem))
		}
		mod = intHandshakeChannel{
			Name: chModName,
			ValueType: types,
		}
	case BufferedChannelType:
		types := []string{}
		for _, elem := range ty.Elems {
			types = append(types, x.convertTypeToString(elem))
		}
		mod = intBufferedChannel{
			Name: chModName,
			Length: x.calculateConstExpression(ty.BufferSize),
			ValueType: types,
		}
	}
	x.modules = append(x.modules, mod)

	return nil, val
}

func (x *intModConverter) buildProcVal(initVar InstanceVar) (error, intInternalProcVal) {
	intVal := x.env.lookup(initVar.ProcDefName)
	if intVal == nil {
		panic(initVar.ProcDefName + " should be found in env")
	}
	var def Definition
	if intProcDef, ok := intVal.(intInternalProcDef); ok {
		def = intProcDef.Def
	} else {
		panic(initVar.ProcDefName + " should be a procdef")
	}
	x.pid = len(x.procs)
	args := []string{}
	for _, arg := range initVar.Args {
		err, basicExprs := x.convertBasicExpr(arg)
		if err != nil {
			return err, intInternalProcVal{}
		}
		args = append(args, basicExprs...)
	}
	val := intInternalProcVal{
		Name: initVar.Name,
		ModuleName: fmt.Sprintf("__pid%d_%s", x.pid, initVar.ProcDefName),
		Def: def,
		Args: args,
		Pid: x.pid,
	}
	x.procs = append(x.procs, val)
	return nil, val
}

func (x *intModConverter) convertTypeToString(ty Type) string {
	// TODO
	return ty.String()
}

func (x *intModConverter) calculateConstExpression(expr Expression) int {
	// TODO
	return 0
}

// ========================================

func argJoin(args []string) string {
	return strings.Join(args, ", ")
}
