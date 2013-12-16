package conversion

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

func (x *intModConverter) convertStatements(statements []Statement, defaults map[string]string) ([]intVar, intState, map[intState][]intTransition) {
	converter := newIntStatementConverter(x.env, defaults)

	for _, stmt := range statements {
		converter.convertStatement(stmt)
	}

	return converter.vars, "state0", converter.trans
}

// ========================================
// Statement conversion

type intStatementConverter struct {
	env          *varEnv
	vars         []intVar
	defaults     map[string]string
	trans        map[intState][]intTransition
	currentState intState
	nextStateNum int
	labelToState map[string]intState
	breakToState intState
}

func newIntStatementConverter(upper *varEnv, defaults map[string]string) *intStatementConverter {
	converter := new(intStatementConverter)
	converter.env = newVarEnvFromUpper(upper)
	converter.defaults = defaults
	converter.trans = make(map[intState][]intTransition)
	converter.currentState = "state0"
	converter.nextStateNum = 1
	converter.labelToState = make(map[string]intState)
	return converter
}

func (x *intStatementConverter) convertStatement(stmt Statement) {
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

func (x *intStatementConverter) convertConstantDefinition(stmt ConstantDefinition) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertLabelledStatement(stmt LabelledStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertBlockStatement(stmt BlockStatement) error {
	nextState := x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Statements {
		x.convertStatement(stmt)
	}
	x.popEnv()
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Actions: map[intState][]intAssign{
			nextState: nil,
		},
	})
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertVarDeclStatement(stmt VarDeclStatement) error {
	nextState := x.genNextState()

	realName := x.genRealName(stmt.Name)
	nextRealName := fmt.Sprintf("next(%s)", realName)
	var condition string = ""
	actions := make(map[intState][]intAssign)
	if stmt.Initializer != nil {
		intExprObj := expressionToInternalObj(stmt.Initializer, x.env)
		condition = intExprObj.Condition()
		actions[nextState] = intExprObj.Assignments(nextRealName)
	} else {
		actions[nextState] = []intAssign{}
	}
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Condition: condition,
		Actions:   actions,
	})
	x.vars = append(x.vars, intVar{realName, convertTypeToString(stmt.Type, x.env)})
	x.env.add(stmt.Name, intInternalPrimitiveVar{realName, stmt.Type})
	x.defaults[nextRealName] = realName
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertIfStatement(stmt IfStatement) error {
	nextState := x.genNextState()
	trueBranchState := x.genNextState()
	falseBranchState := x.genNextState()

	{
		intExprObj := expressionToInternalObj(stmt.Condition, x.env)
		if intExprObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Condition: intExprObj.String(),
			Actions: map[intState][]intAssign{
				trueBranchState: nil,
			},
		})
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Condition: "!(" + intExprObj.String() + ")",
			Actions: map[intState][]intAssign{
				falseBranchState: nil,
			},
		})
	}
	{
		x.currentState = trueBranchState
		x.pushEnv()
		for _, stmt := range stmt.TrueBranch {
			x.convertStatement(stmt)
		}
		x.popEnv()
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Actions: map[intState][]intAssign{
				nextState: nil,
			},
		})
	}
	{
		x.currentState = falseBranchState
		x.pushEnv()
		for _, stmt := range stmt.FalseBranch {
			x.convertStatement(stmt)
		}
		x.popEnv()
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Actions: map[intState][]intAssign{
				nextState: nil,
			},
		})
	}
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertAssignmentStatement(stmt AssignmentStatement) error {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Condition: intExprObj.Condition(),
		Actions: map[intState][]intAssign{
			nextState: intExprObj.Assignments(fmt.Sprintf("next(%s)", stmt.Variable)),
		},
	})
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertOpAssignmentStatement(stmt OpAssignmentStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertChoiceStatement(stmt ChoiceStatement) error {
	nextState := x.genNextState()
	currentState := x.currentState
	for _, block := range stmt.Blocks {
		choicedState := x.genNextState()
		x.trans[currentState] = append(x.trans[currentState], intTransition{
			Actions: map[intState][]intAssign{
				choicedState: nil,
			},
		})
		x.currentState = choicedState
		x.pushEnv()
		x.convertStatement(block)
		x.popEnv()
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Actions: map[intState][]intAssign{
				nextState: nil,
			},
		})
	}
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertRecvStatement(stmt RecvStatement) error {
	nextState := x.genNextState()

	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := make(map[intState][]intAssign)
	switch chType.(type) {
	case HandshakeChannelType:
		actions[nextState] = []intAssign{
			{LHS: fmt.Sprintf("%s.next_received", ch), RHS: "TRUE"},
		}
		for i, arg := range args {
			actions[nextState] = append(actions[nextState], intAssign{
				LHS: fmt.Sprintf("next(%s)", arg),
				RHS: fmt.Sprintf("%s.value_%d", ch, i),
			})
		}
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Condition: fmt.Sprintf("%s.filled & !%s.received", ch, ch),
			Actions:   actions,
		})
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
	}
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertPeekStatement(stmt PeekStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertSendStatement(stmt SendStatement) error {
	nextState := x.genNextState()

	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := make(map[intState][]intAssign)
	switch chType.(type) {
	case HandshakeChannelType:
		actions[nextState] = []intAssign{
			{LHS: fmt.Sprintf("%s.next_filled", ch), RHS: "TRUE"},
			{LHS: fmt.Sprintf("%s.next_received", ch), RHS: "FALSE"},
		}
		for i, arg := range args {
			actions[nextState] = append(actions[nextState], intAssign{
				LHS: fmt.Sprintf("%s.next_value_%d", ch, i),
				RHS: arg.String(),
			})
		}
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Condition: fmt.Sprintf("!(%s.filled)", ch),
			Actions:   actions,
		})
		x.currentState = nextState
		nextState = x.genNextState()
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Condition: fmt.Sprintf("(%s.filled) & (%s.received)", ch, ch),
			Actions: map[intState][]intAssign{
				nextState: []intAssign{
					{LHS: fmt.Sprintf("%s.next_filled", ch), RHS: "FALSE"},
				},
			},
		})
	case BufferedChannelType:
		panic("Not Implemented")
	default:
		panic("unknown channel type")
	}
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertForStatement(stmt ForStatement) error { panic("not implemented") }
func (x *intStatementConverter) convertForInStatement(stmt ForInStatement) error {
	switch container := expressionToInternalObj(stmt.Container, x.env).(type) {
	case intInternalArrayVar:
		savedBreakState := x.breakToState
		x.breakToState = x.genNextState()
		for i, elem := range container.RealLiteral.Elems {
			x.pushEnv()
			x.env.add(stmt.Variable, intInternalPrimitiveVar{
				fmt.Sprintf("__elem%d_%s", i, container.RealName),
				elem.GetType(),
			})
			for _, stmt := range stmt.Statements {
				x.convertStatement(stmt)
			}
			x.popEnv()
		}
		x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
			Actions: map[intState][]intAssign{
				x.breakToState: nil,
			},
		})
		x.currentState = x.breakToState
		x.breakToState = savedBreakState
	default:
		// TODO
		panic("not implemented")
	}
	return nil
}
func (x *intStatementConverter) convertForInRangeStatement(stmt ForInRangeStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertBreakStatement(stmt BreakStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertGotoStatement(stmt GotoStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertSkipStatement(stmt SkipStatement) error {
	nextState := x.genNextState()
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Actions: map[intState][]intAssign{
			nextState: nil,
		},
	})
	x.currentState = nextState
	return nil
}
func (x *intStatementConverter) convertExprStatement(stmt ExprStatement) error {
	panic("not implemented")
}
func (x *intStatementConverter) convertNullStatement(stmt NullStatement) error {
	panic("not implemented")
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
