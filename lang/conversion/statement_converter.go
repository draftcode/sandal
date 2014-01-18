package conversion

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

func (x *intModConverter) convertStatements(statements []Statement, defaults map[string]string, tags []string) ([]intVar, intState, []intTransition) {
	converter := newIntStatementConverter(x.env, defaults, tags)

	for _, stmt := range statements {
		converter.convertStatement(stmt)
	}

	return converter.vars, "state0", converter.trans
}

// ========================================
// Statement conversion

type intStatementConverter struct {
	env           *varEnv
	vars          []intVar
	defaults      map[string]string
	trans         []intTransition
	currentState  intState
	nextStateNum  int
	labelToState  map[string]intState
	breakToState  intState
	tags          []string
	unstable      bool
	unstableState intState
}

func newIntStatementConverter(upper *varEnv, defaults map[string]string, tags []string) *intStatementConverter {
	x := new(intStatementConverter)
	x.env = newVarEnvFromUpper(upper)
	x.defaults = defaults
	x.currentState = "state0"
	x.nextStateNum = 1
	x.labelToState = make(map[string]intState)
	x.tags = tags

	if hasTag(x.tags, "unstable") || hasTag(x.tags, "reboot") {
		x.unstable = true
		x.unstableState = x.genNextState()

		if hasTag(x.tags, "reboot") {
			x.trans = append(x.trans, intTransition{
				FromState: x.unstableState,
				NextState: "state0",
				Condition: "",
			})
		}
	}
	return x
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (x *intStatementConverter) convertStatement(stmt Statement) {
	if x.unstable {
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: x.unstableState,
			Condition: "",
		})
	}

	switch stmt := stmt.(type) {
	case ConstantDefinition:
		x.convertConstantDefinition(stmt)
	case LabelledStatement:
		x.convertLabelledStatement(stmt)
	case BlockStatement:
		x.convertBlockStatement(stmt)
	case VarDeclStatement:
		x.convertVarDeclStatement(stmt)
	case IfStatement:
		x.convertIfStatement(stmt)
	case AssignmentStatement:
		x.convertAssignmentStatement(stmt)
	case OpAssignmentStatement:
		x.convertOpAssignmentStatement(stmt)
	case ChoiceStatement:
		x.convertChoiceStatement(stmt)
	case RecvStatement:
		x.convertRecvStatement(stmt)
	case PeekStatement:
		x.convertPeekStatement(stmt)
	case SendStatement:
		x.convertSendStatement(stmt)
	case ForStatement:
		x.convertForStatement(stmt)
	case ForInStatement:
		x.convertForInStatement(stmt)
	case ForInRangeStatement:
		x.convertForInRangeStatement(stmt)
	case BreakStatement:
		x.convertBreakStatement(stmt)
	case GotoStatement:
		x.convertGotoStatement(stmt)
	case SkipStatement:
		x.convertSkipStatement(stmt)
	case ExprStatement:
		x.convertExprStatement(stmt)
	case NullStatement:
		x.convertNullStatement(stmt)
	}
}

func (x *intStatementConverter) hasRealName(realName string) bool {
	for _, intvar := range x.vars {
		if intvar.Name == realName {
			return true
		}
	}
	return false
}

func (x *intStatementConverter) genRealName(name string) string {
	realName := name
	if x.hasRealName(realName) {
		i := 2
		realName = fmt.Sprintf("%s_%d", name, i)
		for x.hasRealName(realName) {
			i += 1
			realName = fmt.Sprintf("%s_%d", name, i)
		}
	}
	return realName
}

// ========================================

func (x *intStatementConverter) convertConstantDefinition(stmt ConstantDefinition) {
	panic("not implemented")
}
func (x *intStatementConverter) convertLabelledStatement(stmt LabelledStatement) {
	x.labelToState[stmt.Label] = x.currentState
	x.convertStatement(stmt.Statement)
}
func (x *intStatementConverter) convertBlockStatement(stmt BlockStatement) {
	nextState := x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Statements {
		x.convertStatement(stmt)
	}
	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}
func (x *intStatementConverter) convertVarDeclStatement(stmt VarDeclStatement) {
	nextState := x.genNextState()

	realName := x.genRealName(stmt.Name)
	nextRealName := fmt.Sprintf("next(%s)", realName)
	if stmt.Initializer != nil {
		intExprObj := expressionToInternalObj(stmt.Initializer, x.env)
		x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, nextRealName)...)
	} else {
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.vars = append(x.vars, intVar{realName, convertTypeToString(stmt.Type, x.env)})
	x.env.add(stmt.Name, intInternalPrimitiveVar{realName, stmt.Type, nil})
	x.defaults[nextRealName] = realName
	x.currentState = nextState
}
func (x *intStatementConverter) convertIfStatement(stmt IfStatement) {
	nextState := x.genNextState()
	trueBranchState := x.genNextState()
	falseBranchState := x.genNextState()

	{
		intExprObj := expressionToInternalObj(stmt.Condition, x.env)
		if intExprObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: trueBranchState,
			Condition: intExprObj.String(),
		})
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: falseBranchState,
			Condition: "!(" + intExprObj.String() + ")",
		})
	}
	{
		x.currentState = trueBranchState
		x.pushEnv()
		for _, stmt := range stmt.TrueBranch {
			x.convertStatement(stmt)
		}
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	{
		x.currentState = falseBranchState
		x.pushEnv()
		for _, stmt := range stmt.FalseBranch {
			x.convertStatement(stmt)
		}
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.currentState = nextState
}
func (x *intStatementConverter) convertAssignmentStatement(stmt AssignmentStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}
func (x *intStatementConverter) convertOpAssignmentStatement(stmt OpAssignmentStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(BinOpExpression{
		IdentifierExpression{Name: stmt.Variable}, stmt.Operator, stmt.Expr,
	}, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}
func (x *intStatementConverter) convertChoiceStatement(stmt ChoiceStatement) {
	nextState := x.genNextState()
	currentState := x.currentState
	for _, block := range stmt.Blocks {
		choicedState := x.genNextState()
		x.trans = append(x.trans, intTransition{
			FromState: currentState,
			NextState: choicedState,
		})
		x.currentState = choicedState
		x.pushEnv()
		x.convertStatement(block)
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.currentState = nextState
}
func (x *intStatementConverter) convertRecvStatement(stmt RecvStatement) {
	nextState := x.genNextState()

	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := []intAssign{}
	switch chType.(type) {
	case HandshakeChannelType:
		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("%s.next_received", ch),
			RHS: "TRUE",
		})
		for i, arg := range args {
			actions = append(actions, intAssign{
				LHS: fmt.Sprintf("next(%s)", arg),
				RHS: fmt.Sprintf("%s.value_%d", ch, i),
			})
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
			Condition: fmt.Sprintf("%s.filled & !%s.received", ch, ch),
			Actions:   actions,
		})
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
	}
	x.currentState = nextState
}
func (x *intStatementConverter) convertPeekStatement(stmt PeekStatement) {
	panic("not implemented")
}
func (x *intStatementConverter) convertSendStatement(stmt SendStatement) {
	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := []intAssign{}
	switch chType.(type) {
	case HandshakeChannelType:
		chVar := resolveRealObj(ch).(intInternalHandshakeChannelProxyVar)
		firstState := x.currentState
		secondState := x.genNextState()
		lastState := x.genNextState()

		// Generate the first state transition
		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("%s.next_filled", ch),
			RHS: "TRUE",
		})
		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("%s.next_received", ch),
			RHS: "FALSE",
		})
		for i, arg := range args {
			actions = append(actions, intAssign{
				LHS: fmt.Sprintf("%s.next_value_%d", ch, i),
				RHS: arg.String(),
			})
		}
		x.trans = append(x.trans, intTransition{
			FromState: firstState,
			NextState: secondState,
			Condition: fmt.Sprintf("!(%s.filled)", ch),
			Actions:   actions,
		})

		// Generate the second state transition
		x.trans = append(x.trans, intTransition{
			FromState: secondState,
			NextState: lastState,
			Condition: fmt.Sprintf("(%s.filled) & (%s.received)", ch, ch),
			Actions: []intAssign{
				{LHS: fmt.Sprintf("%s.next_filled", ch), RHS: "FALSE"},
			},
		})

		// Inject drop fault
		if hasTag(chVar.ChannelVar.Tags, "drop") {
			x.trans = append(x.trans, intTransition{
				FromState: firstState,
				NextState: lastState,
			})
		}

		x.currentState = lastState
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
	}
}
func (x *intStatementConverter) convertForStatement(stmt ForStatement) {
	savedCurrentState := x.currentState
	savedBreakState := x.breakToState
	x.breakToState = x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Statements {
		x.convertStatement(stmt)
	}
	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: savedCurrentState,
	})
	x.currentState = x.breakToState
	x.breakToState = savedBreakState
}
func (x *intStatementConverter) convertForInStatement(stmt ForInStatement) {
	switch container := expressionToInternalObj(stmt.Container, x.env).(type) {
	case intInternalArrayVar:
		savedBreakState := x.breakToState
		x.breakToState = x.genNextState()
		for i, elem := range container.RealLiteral.Elems {
			x.pushEnv()
			x.env.add(stmt.Variable, intInternalPrimitiveVar{
				fmt.Sprintf("__elem%d_%s", i, container.RealName),
				elem.GetType(),
				elem,
			})
			for _, stmt := range stmt.Statements {
				x.convertStatement(stmt)
			}
			x.popEnv()
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: x.breakToState,
		})
		x.currentState = x.breakToState
		x.breakToState = savedBreakState
	default:
		// TODO
		panic("not implemented")
	}
}
func (x *intStatementConverter) convertForInRangeStatement(stmt ForInRangeStatement) {
	panic("not implemented")
}
func (x *intStatementConverter) convertBreakStatement(stmt BreakStatement) {
	nextState := x.genNextState()
	if x.breakToState == "" {
		panic("Invalid break statement")
	}
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: x.breakToState,
	})
	x.currentState = nextState
}
func (x *intStatementConverter) convertGotoStatement(stmt GotoStatement) {
	nextState := x.genNextState()
	jumpState := x.labelToState[stmt.Label]
	if jumpState == "" {
		panic("Invalid jump label")
	}
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: jumpState,
	})
	x.currentState = nextState
}
func (x *intStatementConverter) convertSkipStatement(stmt SkipStatement) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}
func (x *intStatementConverter) convertExprStatement(stmt ExprStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, "")...)
	x.currentState = nextState
}
func (x *intStatementConverter) convertNullStatement(stmt NullStatement) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

// ========================================

func (x *intStatementConverter) genNextState() (state intState) {
	state = intState(fmt.Sprintf("state%d", x.nextStateNum))
	x.nextStateNum++
	return
}

func (x *intStatementConverter) pushEnv() {
	x.env = newVarEnvFromUpper(x.env)
}

func (x *intStatementConverter) popEnv() {
	x.env = x.env.upper
}
