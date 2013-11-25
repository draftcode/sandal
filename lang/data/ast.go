package data

type (
	Definition interface {
		definition()
	}

	Statement interface {
		statement()
		String() string
	}

	Expression interface {
		expression()
		String() string
	}

	// For type-checking
	ChanExpr interface {
		ChannelExpr() Expression
		ArgExprs() []Expression
		String() string
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
		Type Type
		Expr Expression
	}

	ProcDefinition struct {
		Name       string
		Parameters []Parameter
		Statements []Statement
	}

	InitBlock struct {
		Vars []InitVar
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
		Operator string
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
func (x *SkipStatement) statement()         {}
func (x *ExprStatement) statement()         {}
func (x *NullStatement) statement()         {}

func (x *RecvStatement) ChannelExpr() Expression { return x.Channel }
func (x *PeekStatement) ChannelExpr() Expression { return x.Channel }
func (x *SendStatement) ChannelExpr() Expression { return x.Channel }
func (x *RecvStatement) ArgExprs() []Expression  { return x.Args }
func (x *PeekStatement) ArgExprs() []Expression  { return x.Args }
func (x *SendStatement) ArgExprs() []Expression  { return x.Args }

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
		Operator string
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

	NonblockRecvExpression struct {
		Channel Expression
		Args    []Expression
	}

	NonblockPeekExpression struct {
		Channel Expression
		Args    []Expression
	}

	ArrayExpression struct {
		Elems []Expression
	}
)

func (x *TimeoutRecvExpression) ChannelExpr() Expression  { return x.Channel }
func (x *TimeoutPeekExpression) ChannelExpr() Expression  { return x.Channel }
func (x *NonblockRecvExpression) ChannelExpr() Expression { return x.Channel }
func (x *NonblockPeekExpression) ChannelExpr() Expression { return x.Channel }
func (x *TimeoutRecvExpression) ArgExprs() []Expression   { return x.Args }
func (x *TimeoutPeekExpression) ArgExprs() []Expression   { return x.Args }
func (x *NonblockRecvExpression) ArgExprs() []Expression  { return x.Args }
func (x *NonblockPeekExpression) ArgExprs() []Expression  { return x.Args }

func (x *IdentifierExpression) expression()   {}
func (x *NumberExpression) expression()       {}
func (x *NotExpression) expression()          {}
func (x *UnarySubExpression) expression()     {}
func (x *ParenExpression) expression()        {}
func (x *BinOpExpression) expression()        {}
func (x *TimeoutRecvExpression) expression()  {}
func (x *TimeoutPeekExpression) expression()  {}
func (x *NonblockRecvExpression) expression() {}
func (x *NonblockPeekExpression) expression() {}
func (x *ArrayExpression) expression()        {}

// ========================================
// Misc

type (
	Parameter struct {
		Name string
		Type Type
	}

	InitVar interface {
		initvar()
		VarName() string
	}

	ChannelVar struct {
		Name string
		Type Type
	}

	InstanceVar struct {
		Name       string
		ModuleName string
		Args       []Expression
	}

	Type interface {
		typetype()
		Equal(Type) bool
		String() string
	}

	NamedType struct {
		Name string
	}

	CallableType struct {
		Parameters []Type
	}

	ArrayType struct {
		ElemType Type
	}

	HandshakeChannelType struct {
		IsUnstable bool
		Elems      []Type
	}

	BufferedChannelType struct {
		IsUnstable bool
		BufferSize Expression
		Elems      []Type
	}
)

func (x ChannelVar) initvar()         {}
func (x InstanceVar) initvar()        {}
func (x ChannelVar) VarName() string  { return x.Name }
func (x InstanceVar) VarName() string { return x.Name }

func (x NamedType) typetype()            {}
func (x CallableType) typetype()         {}
func (x ArrayType) typetype()            {}
func (x HandshakeChannelType) typetype() {}
func (x BufferedChannelType) typetype()  {}

func (x NamedType) Equal(ty Type) bool {
	if ty, b := ty.(NamedType); b {
		return (ty.Name == x.Name)
	} else {
		return false
	}
}

func (x CallableType) Equal(ty Type) bool {
	if ty, b := ty.(CallableType); b {
		if len(ty.Parameters) != len(x.Parameters) {
			return false
		}
		for i := 0; i < len(x.Parameters); i++ {
			if !ty.Parameters[i].Equal(x.Parameters[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x ArrayType) Equal(ty Type) bool {
	if ty, b := ty.(ArrayType); b {
		return ty.ElemType.Equal(x.ElemType)
	} else {
		return false
	}
}

func (x HandshakeChannelType) Equal(ty Type) bool {
	if ty, b := ty.(HandshakeChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x BufferedChannelType) Equal(ty Type) bool {
	if ty, b := ty.(BufferedChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
