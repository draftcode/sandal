// vim: noet sw=8 sts=8
%{
package lang

import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}
%}

%union{
	definitions []Definition
	definition  Definition
	statements  []Statement
	statement   Statement
	expressions []Expression
	expression  Expression
	parameters  []Parameter
	parameter   Parameter
	typetypes   []Type
	typetype    Type
	identifiers []string
	blocks      []BlockStatement

	tok         Token
}

%type<definitions> spec
%type<definition> toplevel_body
%type<definition> data_def module_def const_def proc_def init_block
%type<definitions> module_body_zero
%type<definition> module_body
%type<statements> statements_zero
%type<statement> statement
%type<expression> expr

%type<identifiers> idents_one
%type<parameters> parameters_zero parameters_one
%type<parameter> parameter
%type<expressions> arguments_one
%type<typetypes> types_one
%type<typetype> type
%type<blocks> blocks_one

%token<tok> IDENTIFIER
%token<tok> NUMBER
%token<tok> COMMENT

%token<tok> ADD // +
%token<tok> SUB // -
%token<tok> MUL // *
%token<tok> QUO // /
%token<tok> REM // %

%token<tok> AND // &
%token<tok> OR  // |
%token<tok> XOR // ^
%token<tok> SHL // <<
%token<tok> SHR // >>

%token<tok> ADD_ASSIGN // +=
%token<tok> SUB_ASSIGN // -=
%token<tok> MUL_ASSIGN // *=
%token<tok> QUO_ASSIGN // /=
%token<tok> REM_ASSIGN // %=

%token<tok> AND_ASSIGN // &=
%token<tok> OR_ASSIGN  // |=
%token<tok> XOR_ASSIGN // ^=
%token<tok> SHL_ASSIGN // <<=
%token<tok> SHR_ASSIGN // >>=

%token<tok> LAND // &&
%token<tok> LOR  // ||

%token<tok> EQL    // ==
%token<tok> LSS    // <
%token<tok> GTR    // >
%token<tok> ASSIGN // =
%token<tok> NOT    // !

%token<tok> NEQ // !=
%token<tok> LEQ // <=
%token<tok> GEQ // >=

%token<tok> DATA
%token<tok> CONST
%token<tok> MODULE
%token<tok> CHANNEL
%token<tok> PROC
%token<tok> VAR
%token<tok> IF
%token<tok> ELSE
%token<tok> CHOICE
%token<tok> RECV
%token<tok> TIMEOUT_RECV
%token<tok> NONBLOCK_RECV
%token<tok> PEEK
%token<tok> TIMEOUT_PEEK
%token<tok> NONBLOCK_PEEK
%token<tok> SEND
%token<tok> FOR
%token<tok> BREAK
%token<tok> IN
%token<tok> RANGE
%token<tok> TO
%token<tok> INIT
%token<tok> GOTO
%token<tok> UNSTABLE
%token<tok> SKIP

%left LOR
%left LAND
%left EQL NEQ LSS LEQ GTR GEQ
%left ADD SUB OR XOR
%left MUL QUO REM SHL SHR AND
%right UNARY

%%

spec	: toplevel_body
	{
		$$ = []Definition{$1}
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}
	| toplevel_body spec
	{
		$$ = append([]Definition{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}

toplevel_body
	: data_def
	| module_def
	| const_def
	| proc_def
	| init_block

data_def
	: DATA IDENTIFIER '{' idents_one '}' ';'
	{
		$$ = &DataDefinition{Name: $2.lit, Elems: $4}
	}

module_def
	: MODULE IDENTIFIER '(' parameters_zero ')' '{' module_body_zero '}' ';'
	{
		$$ = &ModuleDefinition{Name: $2.lit, Parameters: $4, Definitions: $7}
	}

module_body_zero
	:
	{
		$$ = nil
	}
	| module_body module_body_zero
	{
		$$ = append([]Definition{$1}, $2...)
	}

module_body
	: const_def
	| proc_def
	| init_block

const_def
	: CONST IDENTIFIER type ASSIGN expr ';' /* This should be a const expression. */
	{
		$$ = &ConstantDefinition{Name: $2.lit, Type: $3, Expr: $5}
	}

proc_def
	: PROC IDENTIFIER '(' parameters_zero ')' '{' statements_zero '}' ';'
	{
		$$ = &ProcDefinition{Name: $2.lit, Parameters: $4, Statements: $7}
	}

init_block
	: INIT '{' statements_zero '}' ';'
	{
		$$ = &InitBlock{Statements: $3}
	}

statements_zero
	:
	{
		$$ = nil
	}
	| statement statements_zero
	{
		$$ = append([]Statement{$1}, $2...)
	}

statement
	: IDENTIFIER ':' statement /* no semicolon */
	{
		$$ = &LabelledStatement{Label: $1.lit, Statement: $3}
	}
	| '{' statements_zero '}' ';'
	{
		$$ = &BlockStatement{Statements: $2}
	}
	| VAR IDENTIFIER type ';'
	{
		$$ = &VarDeclStatement{Name: $2.lit, Type: $3}
	}
	| VAR IDENTIFIER type ASSIGN expr ';'
	{
		$$ = &VarDeclStatement{Name: $2.lit, Type: $3, Initializer: $5}
	}
	| IF expr '{' statements_zero '}' ';'
	{
		$$ = &IfStatement{Condition: $2, TrueBranch: $4}
	}
	| IF expr '{' statements_zero '}' ELSE '{' statements_zero '}' ';'
	{
		$$ = &IfStatement{Condition: $2, TrueBranch: $4, FalseBranch: $8}
	}
	| IDENTIFIER ASSIGN expr ';'
	{
		$$ = &AssignmentStatement{Variable: $1.lit, Expr: $3}
	}
	| IDENTIFIER ADD_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: ADD, Expr: $3}
	}
	| IDENTIFIER SUB_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: SUB, Expr: $3}
	}
	| IDENTIFIER MUL_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: MUL, Expr: $3}
	}
	| IDENTIFIER QUO_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: QUO, Expr: $3}
	}
	| IDENTIFIER REM_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: REM, Expr: $3}
	}
	| IDENTIFIER AND_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: AND, Expr: $3}
	}
	| IDENTIFIER OR_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: OR, Expr: $3}
	}
	| IDENTIFIER XOR_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: XOR, Expr: $3}
	}
	| IDENTIFIER SHL_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: SHL, Expr: $3}
	}
	| IDENTIFIER SHR_ASSIGN expr ';'
	{
		$$ = &OpAssignmentStatement{Variable: $1.lit, Operator: SHR, Expr: $3}
	}
	| CHOICE blocks_one ';'
	{
		$$ = &ChoiceStatement{Blocks: $2}
	}
	| RECV '(' arguments_one ')' ';'
	{
		$$ = &RecvStatement{Channel: $3[0], Args: $3[1:]}
	}
	| PEEK '(' arguments_one ')' ';'
	{
		$$ = &PeekStatement{Channel: $3[0], Args: $3[1:]}
	}
	| SEND '(' arguments_one ')' ';'
	{
		$$ = &SendStatement{Channel: $3[0], Args: $3[1:]}
	}
	| FOR '{' statements_zero '}' ';'
	{
		$$ = &ForStatement{Statements: $3}
	}
	| FOR IDENTIFIER IN expr '{' statements_zero '}' ';'
	{
		$$ = &ForInStatement{Variable: $2.lit, Container: $4, Statements: $6}
	}
	| FOR IDENTIFIER IN RANGE expr TO expr '{' statements_zero '}' ';'
	{
		$$ = &ForInRangeStatement{Variable: $2.lit, FromExpr: $5, ToExpr: $7, Statements: $9}
	}
	| BREAK ';'
	{
		$$ = &BreakStatement{}
	}
	| GOTO IDENTIFIER ';'
	{
		$$ = &GotoStatement{Label: $2.lit}
	}
	| IDENTIFIER '(' arguments_one ')' ';'
	{
		$$ = &CallStatement{Name: $1.lit, Args: $3}
	}
	| SKIP ';'
	{
		$$ = &SkipStatement{}
	}
	| expr ';'
	{
		$$ = &ExprStatement{Expr: $1}
	}
	| ';'
	{
		$$ = &NullStatement{}
	}
	| const_def
	{
		$$ = $1.(Statement)
	}

expr	: IDENTIFIER
	{
		$$ = &IdentifierExpression{Name: $1.lit}
	}
	| NUMBER
	{
		$$ = &NumberExpression{Lit: $1.lit}
	}
	| NOT expr      %prec UNARY
	{
		$$ = &NotExpression{SubExpr: $2}
	}
	| SUB expr      %prec UNARY
	{
		$$ = &UnarySubExpression{SubExpr: $2}
	}
	| '(' expr ')'
	{
		$$ = &ParenExpression{SubExpr: $2}
	}
	| expr ADD expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: ADD, RHS: $3}
	}
	| expr SUB expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: SUB, RHS: $3}
	}
	| expr MUL expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: MUL, RHS: $3}
	}
	| expr QUO expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: QUO, RHS: $3}
	}
	| expr REM expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: REM, RHS: $3}
	}
	| expr AND expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: AND, RHS: $3}
	}
	| expr OR expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: OR, RHS: $3}
	}
	| expr XOR expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: XOR, RHS: $3}
	}
	| expr SHL expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: SHL, RHS: $3}
	}
	| expr SHR expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: SHR, RHS: $3}
	}
	| expr LAND expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: LAND, RHS: $3}
	}
	| expr LOR expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: LOR, RHS: $3}
	}
	| expr EQL expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: EQL, RHS: $3}
	}
	| expr LSS expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: LSS, RHS: $3}
	}
	| expr GTR expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: GTR, RHS: $3}
	}
	| expr NEQ expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: NEQ, RHS: $3}
	}
	| expr LEQ expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: LEQ, RHS: $3}
	}
	| expr GEQ expr
	{
		$$ = &BinOpExpression{LHS: $1, Operator: GEQ, RHS: $3}
	}
	| TIMEOUT_RECV '(' arguments_one ')'
	{
		$$ = &TimeoutRecvExpression{Channel: $3[0], Args: $3[1:]}
	}
	| TIMEOUT_PEEK '(' arguments_one ')'
	{
		$$ = &TimeoutPeekExpression{Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_RECV '(' arguments_one ')'
	{
		$$ = &NonblockRecvExpression{Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_PEEK '(' arguments_one ')'
	{
		$$ = &NonblockPeekExpression{Channel: $3[0], Args: $3[1:]}
	}
	| '[' arguments_one ']'
	{
		$$ = &ArrayExpression{Elems: $2}
	}

/* ======================================== */

idents_one
	: IDENTIFIER
	{
		$$ = []string{$1.lit}
	}
	| IDENTIFIER ','
	{
		$$ = []string{$1.lit}
	}
	| IDENTIFIER ',' idents_one
	{
		$$ = append([]string{$1.lit}, $3...)
	}

parameters_zero
	:
	{
		$$ = nil
	}
	| parameters_one
	{
		$$ = $1
	}

parameters_one
	: parameter
	{
		$$ = []Parameter{$1}
	}
	| parameter ','
	{
		$$ = []Parameter{$1}
	}
	| parameter ',' parameters_one
	{
		$$ = append([]Parameter{$1}, $3...)
	}

parameter
	: IDENTIFIER type
	{
		$$ = Parameter{Name: $1.lit, Type: $2}
	}

arguments_one
	: expr
	{
		$$ = []Expression{$1}
	}
	| expr ','
	{
		$$ = []Expression{$1}
	}
	| expr ',' arguments_one
	{
		$$ = append([]Expression{$1}, $3...)
	}

types_one
	: type
	{
		$$ = []Type{$1}
	}
	| type ','
	{
		$$ = []Type{$1}
	}
	| type ',' types_one
	{
		$$ = append([]Type{$1}, $3...)
	}

type	: IDENTIFIER
	{
		$$ = NamedType{Name: $1.lit}
	}
	| '[' ']' type
	{
		$$ = ArrayType{ElemType: $3}
	}
	| CHANNEL '{' types_one '}'
	{
		$$ = HandshakeChannelType{IsUnstable: false, Elems: $3}
	}
	| UNSTABLE CHANNEL '{' types_one '}'
	{
		$$ = HandshakeChannelType{IsUnstable: true, Elems: $4}
	}
	| CHANNEL '[' ']' '{' types_one '}'
	{
		$$ = BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: $5}
	}
	| CHANNEL '[' expr ']' '{' types_one '}'
	{
		$$ = BufferedChannelType{IsUnstable: false, BufferSize: $3, Elems: $6}
	}
	| UNSTABLE CHANNEL '[' ']' '{' types_one '}'
	{
		$$ = BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: $6}
	}
	| UNSTABLE CHANNEL '[' expr ']' '{' types_one '}'
	{
		$$ = BufferedChannelType{IsUnstable: true, BufferSize: $4, Elems: $7}
	}

blocks_one
	: '{' statements_zero '}'
	{
		$$ = []BlockStatement{BlockStatement{Statements: $2}}
	}
	| '{' statements_zero '}' ','
	{
		$$ = []BlockStatement{BlockStatement{Statements: $2}}
	}
	| '{' statements_zero '}' ',' blocks_one
	{
		$$ = append([]BlockStatement{BlockStatement{Statements: $2}}, $5...)
	}

%%

type LexerWrapper struct {
	s           *Scanner
	definitions []Definition
	recentLit   string
	recentPos   Position
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	for tok == COMMENT {
		tok, lit, pos = l.s.Scan()
	}
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []Definition {
	l := LexerWrapper{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.definitions
}
