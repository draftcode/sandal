package conversion_deprecated

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
		case DataDefinition:
			for _, elem := range def.Elems {
				converter.env.add(elem, intInternalLiteral{
					Lit:  elem,
					Type: NamedType{def.Name},
				})
			}
			converter.env.add(def.Name, intInternalDataTypeDef{
				Elems: def.Elems,
			})
		case ModuleDefinition:
			// TODO
		case ConstantDefinition:
			converter.env.add(def.Name, intInternalConstantDef{
				Type: def.Type,
				Expr: def.Expr,
			})
		case ProcDefinition:
			converter.env.add(def.Name, intInternalProcDef{
				Def: def,
			})
		case InitBlock:
			// Do nothing
		case LtlSpec:
			converter.convertLtlSpec(def)
		}
	}
	for _, def := range defs {
		switch def := def.(type) {
		case InitBlock:
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
	channels []intInternalObj
	procs    []intInternalProcVar
	modules  []intModule
	ltls     []string
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

func (x *intModConverter) convertInitBlock(def InitBlock) error {
	x.pushEnv()
	defer x.popEnv()
	for _, initVar := range def.Vars {
		switch initVar := initVar.(type) {
		case InstanceVar:
			// Do nothing
		case ChannelVar:
			err, chVar := x.buildChannelVar(initVar.Name, initVar.Type, initVar.Tags)
			if err != nil {
				return err
			}
			x.env.add(initVar.Name, chVar)
		default:
			panic("Unknown InitVar")
		}
	}
	for _, initVar := range def.Vars {
		switch initVar := initVar.(type) {
		case InstanceVar:
			err := x.buildProcVar(initVar)
			if err != nil {
				return err
			}
		case ChannelVar:
			// Do nothing
		}
	}
	return nil
}

func (x *intModConverter) convertLtlSpec(def LtlSpec) error {
	x.ltls = append(x.ltls, convertLtlExpression(def.Expr))
	return nil
}

func (x *intModConverter) buildMainModule() error {
	if len(x.procs) == 0 {
		return fmt.Errorf("No running procs")
	}

	module := intMainModule{}
	// Vars
	for _, chVar := range x.channels {
		switch chVar := chVar.(type) {
		case intInternalHandshakeChannelVar:
			module.Vars = append(module.Vars, intVar{
				Name: chVar.RealName,
				Type: chVar.ModuleName,
			})
		case intInternalBufferedChannelVar:
			module.Vars = append(module.Vars, intVar{
				Name: chVar.RealName,
				Type: chVar.ModuleName,
			})
		default:
			panic("Unknown channel value")
		}
	}
	for _, procVal := range x.procs {
		args := []string{}
		for _, arg := range procVal.Args {
			if arrayArg, isArrayLit := arg.(intInternalArrayLiteral); isArrayLit {
				args = append(args, arrayArg.ArgString()...)
			} else {
				args = append(args, arg.String())
			}
		}
		module.Vars = append(module.Vars, intVar{
			Name: procVal.Name,
			Type: fmt.Sprintf("process %s(%s)", procVal.ModuleName, argJoin(args)),
		})
	}

	// LtlSpecs
	module.LtlSpecs = x.ltls

	x.modules = append(x.modules, module)
	return nil
}

func (x *intModConverter) buildChannelVar(name string, ty Type, tags []string) (error, intInternalObj) {
	chNumber := len(x.channels)
	var mod intModule
	var chVar intInternalObj
	switch ty := ty.(type) {
	case HandshakeChannelType:
		types := []string{}
		zeroValues := []string{}
		for _, elem := range ty.Elems {
			types = append(types, convertTypeToString(elem, x.env))
			zeroValues = append(zeroValues, zeroValueOfType(elem, x.env))
		}
		moduleName := fmt.Sprintf("HandshakeChannel%d", chNumber)
		mod = intHandshakeChannel{
			Name:      moduleName,
			ValueType: types,
			ZeroValue: zeroValues,
		}
		chVar = intInternalHandshakeChannelVar{
			ModuleName: moduleName,
			RealName:   name,
			Type:       ty,
			Tags:       tags,
			Pids:       make(map[int]bool),
		}
	case BufferedChannelType:
		types := []string{}
		zeroValues := []string{}
		for _, elem := range ty.Elems {
			types = append(types, convertTypeToString(elem, x.env))
			zeroValues = append(zeroValues, zeroValueOfType(elem, x.env))
		}
		moduleName := fmt.Sprintf("BufferedChannel%d", chNumber)
		mod = intBufferedChannel{
			Name:      moduleName,
			Length:    x.calculateConstExpression(ty.BufferSize),
			ValueType: types,
			ZeroValue: zeroValues,
		}
		chVar = intInternalBufferedChannelVar{
			ModuleName: moduleName,
			RealName:   name,
			Type:       ty,
			Tags:       tags,
			Pids:       make(map[int]bool),
		}
	default:
		panic("Unknown channel type")
	}
	x.modules = append(x.modules, mod)
	x.channels = append(x.channels, chVar)
	return nil, chVar
}

func (x *intModConverter) buildProcVar(initVar InstanceVar) error {
	// Find intInternalProcDef from ProcDefName
	intVal := x.env.lookup(initVar.ProcDefName)
	if intVal == nil {
		panic(initVar.ProcDefName + " should be found in env")
	}
	var intProcDef intInternalProcDef
	if def, ok := intVal.(intInternalProcDef); ok {
		intProcDef = def
	} else {
		panic(initVar.ProcDefName + " should be a intInternalProcDef")
	}

	x.pid = len(x.procs)
	args := []intInternalExpressionObj{}
	for _, arg := range initVar.Args {
		args = append(args, expressionToInternalObj(arg, x.env))
	}
	moduleName := fmt.Sprintf("__pid%d_%s", x.pid, initVar.ProcDefName)
	x.instantiateProcDef(intProcDef, moduleName, args, initVar.Tags)
	procvar := intInternalProcVar{
		Name:       initVar.Name,
		ModuleName: moduleName,
		Def:        intProcDef,
		Args:       args,
		Pid:        x.pid,
	}
	x.procs = append(x.procs, procvar)
	return nil
}

func (x *intModConverter) instantiateProcDef(def intInternalProcDef, moduleName string, args []intInternalExpressionObj, tags []string) {
	x.pushEnv()
	defer x.popEnv()
	vars := []intVar{}
	params := []string{}
	defaults := make(map[string]string)

	processBufferedChannel := func(paramName string, moduleName string, numElems int) {
		defaults[paramName+".send_filled"] = "FALSE"
		defaults[paramName+".recv_received"] = "FALSE"
		for i := 0; i < numElems; i++ {
			defaults[fmt.Sprintf("%s.send_value_%d", paramName, i)] = fmt.Sprintf("%s.value_%d", paramName, i)
		}
		vars = append(vars, intVar{
			Name: paramName,
			Type: fmt.Sprintf("%sProxy(__orig_%s)", moduleName, paramName),
		})
		params = append(params, "__orig_" + paramName)
	}
	processHandshakeChannel := func(paramName string, moduleName string, numElems int) {
		defaults[paramName+".send_leaving"] = "FALSE"
		processBufferedChannel(paramName, moduleName, numElems)
	}

	for idx, arg := range args {
		param := def.Def.Parameters[idx]
		switch arg := arg.(type) {
		case intInternalArrayLiteral:
			for i := 0; i < len(arg.Elems); i++ {
				paramName := fmt.Sprintf("__elem%d_%s", i, param.Name)
				switch elem := arg.Elems[i].(type) {
				case intInternalHandshakeChannelVar:
					processHandshakeChannel(paramName, elem.ModuleName, len(elem.Type.Elems))
				case intInternalBufferedChannelVar:
					processBufferedChannel(paramName, elem.ModuleName, len(elem.Type.Elems))
				default:
					params = append(params, paramName)
				}
			}
			x.env.add(param.Name, intInternalArrayVar{param.Name, arg})
		case intInternalHandshakeChannelVar:
			processHandshakeChannel(param.Name, arg.ModuleName, len(arg.Type.Elems))
			x.env.add(param.Name, intInternalPrimitiveVar{param.Name, param.Type, arg})
		case intInternalBufferedChannelVar:
			processBufferedChannel(param.Name, arg.ModuleName, len(arg.Type.Elems))
			x.env.add(param.Name, intInternalPrimitiveVar{param.Name, param.Type, arg})
		case intInternalLiteral, intInternalNot, intInternalUnarySub, intInternalParen, intInternalBinOp:
			params = append(params, param.Name)
			x.env.add(param.Name, intInternalPrimitiveVar{param.Name, param.Type, nil})
		default:
			panic("unexpected")
		}
	}
	vars, initState, trans := x.convertStatements(def.Def.Statements, defaults, tags, vars)

	x.modules = append(x.modules, intProcModule{
		Name:      moduleName,
		Args:      params,
		Vars:      vars,
		InitState: initState,
		Trans:     trans,
		Defaults:  defaults,
	})
}

func convertTypeToString(ty Type, env *varEnv) string {
	// TODO
	switch ty := ty.(type) {
	case NamedType:
		switch ty.Name {
		case "bool":
			return "boolean"
		case "int":
			return "0..8"
		default:
			switch intObj := env.lookup(ty.Name).(type) {
			case intInternalDataTypeDef:
				return "{" + argJoin(intObj.Elems) + "}"
			default:
				panic("unknown type")
			}
		}
	default:
		return ty.String()
	}
}

func zeroValueOfType(ty Type, env *varEnv) string {
	// TODO
	switch ty := ty.(type) {
	case NamedType:
		switch ty.Name {
		case "bool":
			return "FALSE"
		case "int":
			return "0"
		default:
			switch intObj := env.lookup(ty.Name).(type) {
			case intInternalDataTypeDef:
				return intObj.Elems[0]
			default:
				panic("unknown type")
			}
		}
	default:
		panic("not implemented")
	}
}

func (x *intModConverter) calculateConstExpression(expr Expression) int {
	switch expr := expr.(type) {
	case NumberExpression:
		i, err := strconv.Atoi(expr.Lit)
		if err != nil {
			panic("Expect " + expr.Lit + " to be converted to integer")
		}
		return i
	default:
		panic("not implemented")
	}
	return 0
}

// ========================================

func argJoin(args []string) string {
	return strings.Join(args, ", ")
}

func convertLtlExpression(expr LtlExpression) string {
	switch expr := expr.(type) {
	case LtlAtomExpression:
		return strings.Join(expr.Names, ".")
	case ParenLtlExpression:
		return "(" + convertLtlExpression(expr.SubExpr) + ")"
	case UnOpLtlExpression:
		return expr.Operator + convertLtlExpression(expr.SubExpr)
	case BinOpLtlExpression:
		return convertLtlExpression(expr.LHS) + expr.Operator + convertLtlExpression(expr.RHS)
	default:
		panic("unknown ltl expression")
	}
}
