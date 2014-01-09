package data

import (
	"fmt"
)

type (
	Pos struct {
		Line   int
		Column int
	}

	Definition interface {
		Position() Pos
		definition()
	}

	Statement interface {
		Position() Pos
		statement()
		String() string
	}

	Expression interface {
		Position() Pos
		expression()
		String() string
	}

	// For type-checking
	ChanExpr interface {
		Position() Pos
		ChannelExpr() Expression
		ArgExprs() []Expression
		String() string
	}
)

func (x Pos) String() string {
	return fmt.Sprintf("Line: %d, Column %d", x.Line, x.Column)
}

// ========================================
// Definitions

type (
	DataDefinition struct {
		Pos   Pos
		Name  string
		Elems []string
	}

	ModuleDefinition struct {
		Pos         Pos
		Name        string
		Parameters  []Parameter
		Definitions []Definition
	}

	// ConstantDefinition is a definition but also is a statement.
	ConstantDefinition struct {
		Pos  Pos
		Name string
		Type Type
		Expr Expression
	}

	ProcDefinition struct {
		Pos        Pos
		Name       string
		Parameters []Parameter
		Statements []Statement
	}

	InitBlock struct {
		Pos  Pos
		Vars []InitVar
	}
)

func (x DataDefinition) definition()     {}
func (x ModuleDefinition) definition()   {}
func (x ConstantDefinition) definition() {}
func (x ConstantDefinition) statement()  {}
func (x ProcDefinition) definition()     {}
func (x InitBlock) definition()          {}

func (x DataDefinition) Position() Pos     { return x.Pos }
func (x ModuleDefinition) Position() Pos   { return x.Pos }
func (x ConstantDefinition) Position() Pos { return x.Pos }
func (x ProcDefinition) Position() Pos     { return x.Pos }
func (x InitBlock) Position() Pos          { return x.Pos }

// ========================================
// Statements

type (
	LabelledStatement struct {
		Pos       Pos
		Label     string
		Statement Statement
	}

	BlockStatement struct {
		Pos        Pos
		Statements []Statement
	}

	VarDeclStatement struct {
		Pos         Pos
		Name        string
		Type        Type
		Initializer Expression
	}

	IfStatement struct {
		Pos         Pos
		Condition   Expression
		TrueBranch  []Statement
		FalseBranch []Statement
	}

	AssignmentStatement struct {
		Pos      Pos
		Variable string
		Expr     Expression
	}

	OpAssignmentStatement struct {
		Pos      Pos
		Variable string
		Operator string
		Expr     Expression
	}

	ChoiceStatement struct {
		Pos    Pos
		Blocks []BlockStatement
	}

	RecvStatement struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	PeekStatement struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	SendStatement struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	ForStatement struct {
		Pos        Pos
		Statements []Statement
	}

	ForInStatement struct {
		Pos        Pos
		Variable   string
		Container  Expression
		Statements []Statement
	}

	ForInRangeStatement struct {
		Pos        Pos
		Variable   string
		FromExpr   Expression
		ToExpr     Expression
		Statements []Statement
	}

	BreakStatement struct {
		Pos Pos
	}

	GotoStatement struct {
		Pos   Pos
		Label string
	}

	SkipStatement struct {
		Pos Pos
	}

	ExprStatement struct {
		Expr Expression
	}

	NullStatement struct {
		Pos Pos
	}
)

func (x LabelledStatement) statement()     {}
func (x BlockStatement) statement()        {}
func (x VarDeclStatement) statement()      {}
func (x IfStatement) statement()           {}
func (x AssignmentStatement) statement()   {}
func (x OpAssignmentStatement) statement() {}
func (x ChoiceStatement) statement()       {}
func (x RecvStatement) statement()         {}
func (x PeekStatement) statement()         {}
func (x SendStatement) statement()         {}
func (x ForStatement) statement()          {}
func (x ForInStatement) statement()        {}
func (x ForInRangeStatement) statement()   {}
func (x BreakStatement) statement()        {}
func (x GotoStatement) statement()         {}
func (x SkipStatement) statement()         {}
func (x ExprStatement) statement()         {}
func (x NullStatement) statement()         {}

func (x LabelledStatement) Position() Pos     { return x.Pos }
func (x BlockStatement) Position() Pos        { return x.Pos }
func (x VarDeclStatement) Position() Pos      { return x.Pos }
func (x IfStatement) Position() Pos           { return x.Pos }
func (x AssignmentStatement) Position() Pos   { return x.Pos }
func (x OpAssignmentStatement) Position() Pos { return x.Pos }
func (x ChoiceStatement) Position() Pos       { return x.Pos }
func (x RecvStatement) Position() Pos         { return x.Pos }
func (x PeekStatement) Position() Pos         { return x.Pos }
func (x SendStatement) Position() Pos         { return x.Pos }
func (x ForStatement) Position() Pos          { return x.Pos }
func (x ForInStatement) Position() Pos        { return x.Pos }
func (x ForInRangeStatement) Position() Pos   { return x.Pos }
func (x BreakStatement) Position() Pos        { return x.Pos }
func (x GotoStatement) Position() Pos         { return x.Pos }
func (x SkipStatement) Position() Pos         { return x.Pos }
func (x ExprStatement) Position() Pos         { return x.Expr.Position() }
func (x NullStatement) Position() Pos         { return x.Pos }

func (x RecvStatement) ChannelExpr() Expression { return x.Channel }
func (x PeekStatement) ChannelExpr() Expression { return x.Channel }
func (x SendStatement) ChannelExpr() Expression { return x.Channel }
func (x RecvStatement) ArgExprs() []Expression  { return x.Args }
func (x PeekStatement) ArgExprs() []Expression  { return x.Args }
func (x SendStatement) ArgExprs() []Expression  { return x.Args }

// ========================================
// Expressions

type (
	IdentifierExpression struct {
		Pos  Pos
		Name string
	}

	NumberExpression struct {
		Pos Pos
		Lit string
	}

	TrueExpression struct {
		Pos Pos
	}

	FalseExpression struct {
		Pos Pos
	}

	NotExpression struct {
		Pos     Pos
		SubExpr Expression
	}

	UnarySubExpression struct {
		Pos     Pos
		SubExpr Expression
	}

	ParenExpression struct {
		Pos     Pos
		SubExpr Expression
	}

	BinOpExpression struct {
		LHS      Expression
		Operator string
		RHS      Expression
	}

	TimeoutRecvExpression struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	TimeoutPeekExpression struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	NonblockRecvExpression struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	NonblockPeekExpression struct {
		Pos     Pos
		Channel Expression
		Args    []Expression
	}

	ArrayExpression struct {
		Pos   Pos
		Elems []Expression
	}
)

func (x TimeoutRecvExpression) ChannelExpr() Expression  { return x.Channel }
func (x TimeoutPeekExpression) ChannelExpr() Expression  { return x.Channel }
func (x NonblockRecvExpression) ChannelExpr() Expression { return x.Channel }
func (x NonblockPeekExpression) ChannelExpr() Expression { return x.Channel }
func (x TimeoutRecvExpression) ArgExprs() []Expression   { return x.Args }
func (x TimeoutPeekExpression) ArgExprs() []Expression   { return x.Args }
func (x NonblockRecvExpression) ArgExprs() []Expression  { return x.Args }
func (x NonblockPeekExpression) ArgExprs() []Expression  { return x.Args }

func (x IdentifierExpression) expression()   {}
func (x NumberExpression) expression()       {}
func (x TrueExpression) expression()         {}
func (x FalseExpression) expression()        {}
func (x NotExpression) expression()          {}
func (x UnarySubExpression) expression()     {}
func (x ParenExpression) expression()        {}
func (x BinOpExpression) expression()        {}
func (x TimeoutRecvExpression) expression()  {}
func (x TimeoutPeekExpression) expression()  {}
func (x NonblockRecvExpression) expression() {}
func (x NonblockPeekExpression) expression() {}
func (x ArrayExpression) expression()        {}

func (x IdentifierExpression) Position() Pos   { return x.Pos }
func (x NumberExpression) Position() Pos       { return x.Pos }
func (x TrueExpression) Position() Pos         { return x.Pos }
func (x FalseExpression) Position() Pos        { return x.Pos }
func (x NotExpression) Position() Pos          { return x.Pos }
func (x UnarySubExpression) Position() Pos     { return x.Pos }
func (x ParenExpression) Position() Pos        { return x.Pos }
func (x BinOpExpression) Position() Pos        { return x.LHS.Position() }
func (x TimeoutRecvExpression) Position() Pos  { return x.Pos }
func (x TimeoutPeekExpression) Position() Pos  { return x.Pos }
func (x NonblockRecvExpression) Position() Pos { return x.Pos }
func (x NonblockPeekExpression) Position() Pos { return x.Pos }
func (x ArrayExpression) Position() Pos        { return x.Pos }

// ========================================
// Misc

type (
	Parameter struct {
		Name string
		Type Type
	}

	InitVar interface {
		Position() Pos
		initvar()
		VarName() string
	}

	ChannelVar struct {
		Pos  Pos
		Name string
		Type Type
		Tags []string
	}

	InstanceVar struct {
		Pos         Pos
		Name        string
		ProcDefName string
		Args        []Expression
		Tags        []string
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
		Elems []Type
	}

	BufferedChannelType struct {
		BufferSize Expression
		Elems      []Type
	}
)

func (x ChannelVar) initvar()         {}
func (x InstanceVar) initvar()        {}
func (x ChannelVar) Position() Pos    { return x.Pos }
func (x InstanceVar) Position() Pos   { return x.Pos }
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
