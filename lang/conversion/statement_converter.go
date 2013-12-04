package conversion

import (
	"fmt"
	. "github.com/draftcode/sandal/lang/data"
)

func (x *intModConverter) convertStatements(statements []Statement) ([]intVar, intState, map[intState][]intTransition) {
	converter := newIntStatementConverter(x.env)

	for _, stmt := range statements {
		switch stmt := stmt.(type) {
		case ConstantDefinition:
			converter.convertConstantDefinition(stmt)
		case LabelledStatement:
			converter.convertLabelledStatement(stmt)
		case BlockStatement:
			converter.convertBlockStatement(stmt)
		case VarDeclStatement:
			converter.convertVarDeclStatement(stmt)
		case IfStatement:
			converter.convertIfStatement(stmt)
		case AssignmentStatement:
			converter.convertAssignmentStatement(stmt)
		case OpAssignmentStatement:
			converter.convertOpAssignmentStatement(stmt)
		case ChoiceStatement:
			converter.convertChoiceStatement(stmt)
		case RecvStatement:
			converter.convertRecvStatement(stmt)
		case PeekStatement:
			converter.convertPeekStatement(stmt)
		case SendStatement:
			converter.convertSendStatement(stmt)
		case ForStatement:
			converter.convertForStatement(stmt)
		case ForInStatement:
			converter.convertForInStatement(stmt)
		case ForInRangeStatement:
			converter.convertForInRangeStatement(stmt)
		case BreakStatement:
			converter.convertBreakStatement(stmt)
		case GotoStatement:
			converter.convertGotoStatement(stmt)
		case SkipStatement:
			converter.convertSkipStatement(stmt)
		case ExprStatement:
			converter.convertExprStatement(stmt)
		case NullStatement:
			converter.convertNullStatement(stmt)
		}
	}

	return converter.vars, "state0", converter.trans
}

// ========================================
// Statement conversion

type intStatementConverter struct {
	env          *varEnv
	vars         []intVar
	trans        map[intState][]intTransition
	currentState intState
	nextStateNum int
	labelToState map[string]intState
}

func newIntStatementConverter(upper *varEnv) *intStatementConverter {
	converter := new(intStatementConverter)
	converter.env = newVarEnvFromUpper(upper)
	converter.trans = make(map[intState][]intTransition)
	converter.currentState = "state0"
	converter.nextStateNum = 1
	converter.labelToState = make(map[string]intState)
	return converter
}

// ========================================

func (x *intStatementConverter) convertConstantDefinition(stmt ConstantDefinition) {}
func (x *intStatementConverter) convertLabelledStatement(stmt LabelledStatement)   {}
func (x *intStatementConverter) convertBlockStatement(stmt BlockStatement)         {}
func (x *intStatementConverter) convertVarDeclStatement(stmt VarDeclStatement) {
	nextState := x.genNextState()
	x.vars = append(x.vars, intVar{stmt.Name, convertTypeToString(stmt.Type)})
	actions := make(map[intState][]intAssign)
	if stmt.Initializer != nil {
		actions[nextState] = []intAssign{
			{LHS: stmt.Name, RHS: x.convertExpression(stmt.Initializer)},
		}
	} else {
		actions[nextState] = []intAssign{}
	}
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Condition: "",
		Actions:   actions,
	})
	x.currentState = nextState
}
func (x *intStatementConverter) convertIfStatement(stmt IfStatement)                     {}
func (x *intStatementConverter) convertAssignmentStatement(stmt AssignmentStatement)     {}
func (x *intStatementConverter) convertOpAssignmentStatement(stmt OpAssignmentStatement) {}
func (x *intStatementConverter) convertChoiceStatement(stmt ChoiceStatement)             {}
func (x *intStatementConverter) convertRecvStatement(stmt RecvStatement)                 {}
func (x *intStatementConverter) convertPeekStatement(stmt PeekStatement)                 {}
func (x *intStatementConverter) convertSendStatement(stmt SendStatement) {
	args := []string{}
	for _, arg := range stmt.Args {
		args = append(args, x.convertExpression(arg))
	}

	nextState := x.genNextState()
	ch := x.convertExpression(stmt.Channel)
	actions := make(map[intState][]intAssign)
	actions[nextState] = []intAssign{
		{LHS: ch + ".next_filled", RHS: "TRUE"},
		{LHS: ch + ".next_received", RHS: "FALSE"},
	}
	for i, arg := range args {
		actions[nextState] = append(actions[nextState], intAssign{
			LHS: fmt.Sprintf("%s.next_value_%d", ch, i),
			RHS: arg,
		})
	}
	x.trans[x.currentState] = append(x.trans[x.currentState], intTransition{
		Condition: "!" + ch + ".filled", // TODO: naive
		Actions:   actions,
	})
	x.currentState = nextState

}
func (x *intStatementConverter) convertForStatement(stmt ForStatement)                   {}
func (x *intStatementConverter) convertForInStatement(stmt ForInStatement)               {}
func (x *intStatementConverter) convertForInRangeStatement(stmt ForInRangeStatement)     {}
func (x *intStatementConverter) convertBreakStatement(stmt BreakStatement)               {}
func (x *intStatementConverter) convertGotoStatement(stmt GotoStatement)                 {}
func (x *intStatementConverter) convertSkipStatement(stmt SkipStatement)                 {}
func (x *intStatementConverter) convertExprStatement(stmt ExprStatement)                 {}
func (x *intStatementConverter) convertNullStatement(stmt NullStatement)                 {}

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
