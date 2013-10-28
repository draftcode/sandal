package sandal

type (
	Definition interface {
		definition()
	}

	Statement interface {
		statement()
	}

	Expression interface {
		expression()
	}
)

// ========================================
// Definitions

type (
	DataDefinition struct {
		Name  string
		Elems []string
	}

	ModuleDefinition struct {
		Name        string
		Parameters  []Parameter
		Definitions []Definition
	}

	// ConstantDefinition is a definition but also is a statement.
	ConstantDefinition struct {
		Name string
		Expr Expression
	}

	ProcDefinition struct {
		Name       string
		Parameters []Parameter
		Statements []Statement
	}

	InitBlock struct {
		Statements []Statement
	}
)

func (x *DataDefinition) definition()     {}
func (x *ModuleDefinition) definition()   {}
func (x *ConstantDefinition) definition() {}
func (x *ConstantDefinition) statement()  {}
func (x *ProcDefinition) definition()     {}
func (x *InitBlock) definition()          {}

// ========================================
// Statements

type (
	LabelledStatement struct {
		Label     string
		Statement Statement
	}

	BlockStatement struct {
		Statements []Statement
	}

	VarDeclStatement struct {
		Name        string
		Type        Type
		Initializer Expression
	}

	IfStatement struct {
		Condition   Expression
		TrueBranch  []Statement
		FalseBranch []Statement
	}

	AssignmentStatement struct {
		Variable string
		Expr     Expression
	}

	OpAssignmentStatement struct {
		Variable string
		Operator int
		Expr     Expression
	}

	ChoiceStatement struct {
		Blocks []BlockStatement
	}

	RecvStatement struct {
		Channel Expression
		Args    []Expression
	}

	PeekStatement struct {
		Channel Expression
		Args    []Expression
	}

	SendStatement struct {
		Channel Expression
		Args    []Expression
	}

	ForStatement struct {
		Statements []Statement
	}

	ForInStatement struct {
		Variable   string
		Container  Expression
		Statements []Statement
	}

	ForInRangeStatement struct {
		Variable   string
		FromExpr   Expression
		ToExpr     Expression
		Statements []Statement
	}

	BreakStatement struct {
	}

	GotoStatement struct {
		Label string
	}

	CallStatement struct {
		Name string
		Args []Expression
	}

	SkipStatement struct {
	}

	ExprStatement struct {
		Expr Expression
	}

	NullStatement struct {
	}
)

func (x *LabelledStatement) statement()     {}
func (x *BlockStatement) statement()        {}
func (x *VarDeclStatement) statement()      {}
func (x *IfStatement) statement()           {}
func (x *AssignmentStatement) statement()   {}
func (x *OpAssignmentStatement) statement() {}
func (x *ChoiceStatement) statement()       {}
func (x *RecvStatement) statement()         {}
func (x *PeekStatement) statement()         {}
func (x *SendStatement) statement()         {}
func (x *ForStatement) statement()          {}
func (x *ForInStatement) statement()        {}
func (x *ForInRangeStatement) statement()   {}
func (x *BreakStatement) statement()        {}
func (x *GotoStatement) statement()         {}
func (x *CallStatement) statement()         {}
func (x *SkipStatement) statement()         {}
func (x *ExprStatement) statement()         {}
func (x *NullStatement) statement()         {}

// ========================================
// Expressions

type (
	IdentifierExpression struct {
		Name string
	}

	NumberExpression struct {
		Lit string
	}

	NotExpression struct {
		SubExpr Expression
	}

	UnarySubExpression struct {
		SubExpr Expression
	}

	ParenExpression struct {
		SubExpr Expression
	}

	BinOpExpression struct {
		LHS      Expression
		Operator int
		RHS      Expression
	}

	TimeoutRecvExpression struct {
		Channel Expression
		Args    []Expression
	}

	TimeoutPeekExpression struct {
		Channel Expression
		Args    []Expression
	}

	ArrayExpression struct {
		Elems []Expression
	}
)

func (x *IdentifierExpression) expression()  {}
func (x *NumberExpression) expression()      {}
func (x *NotExpression) expression()         {}
func (x *UnarySubExpression) expression()    {}
func (x *ParenExpression) expression()       {}
func (x *BinOpExpression) expression()       {}
func (x *TimeoutRecvExpression) expression() {}
func (x *TimeoutPeekExpression) expression() {}
func (x *ArrayExpression) expression()       {}

// ========================================
// Misc

type (
	Parameter struct {
		Name string
		Type Type
	}

	Type interface {
		typetype()
	}

	NamedType struct {
		Name string
	}

	SetType struct {
		SetType Type
	}

	ChannelType struct {
		IsUnstable bool
		Elems      []Type
	}
)

func (x *NamedType) typetype()   {}
func (x *SetType) typetype()     {}
func (x *ChannelType) typetype() {}
