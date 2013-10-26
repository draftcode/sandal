// vim: noet sw=8 sts=8
%{
package sandal
%}

%%

spec	: toplevel
	| toplevel spec

toplevel
	: data_def
	| const_def
	| module_def
	| proc_def
	| init_block

data_def
	: DATA IDENTIFIER '{' idents_one '}' ';'

const_def
	: CONST IDENTIFIER '=' expr ';' /* This should be a const expression. */

module_def
	: MODULE IDENTIFIER '(' parameters_zero ')' '{' module_body '}' ';'

module_body
	: const_def
	| proc_def
	| init_block

proc_def
	: PROC IDENTIFIER '(' parameters_zero ')' '{' statements_zero '}' ';'

init_block
	: INIT '{' statements_zero '}'

statements_zero
	:
	| statement statements

statement
	: IDENTIFIER ':' statement /* no semicolon */
	| '{' statements_zero '}' ';'
	| VAR IDENTIFIER type ';'
	| VAR IDENTIFIER type '=' expr ';'
	| IF expr '{' statements_zero '}' ';'
	| IF expr '{' statements_zero '}' ELSE '{' statements_zero '}' ';'
	| IDENTIFIER ASSIGN expr ';'
	| IDENTIFIER ADD_ASSIGN expr ';'
	| IDENTIFIER SUB_ASSIGN expr ';'
	| IDENTIFIER MUL_ASSIGN expr ';'
	| IDENTIFIER QUO_ASSIGN expr ';'
	| IDENTIFIER REM_ASSIGN expr ';'
	| IDENTIFIER AND_ASSIGN expr ';'
	| IDENTIFIER OR_ASSIGN expr ';'
	| IDENTIFIER XOR_ASSIGN expr ';'
	| IDENTIFIER SHL_ASSIGN expr ';'
	| IDENTIFIER SHR_ASSIGN expr ';'
	| CHOICE blocks_one ';'
	| RECV '(' arguments_one ')' ';'
	| PEEK '(' arguments_one ')' ';'
	| SEND '(' arguments_one ')' ';'
	| FOR '{' statements_zero '}' ';'
	| FOR IDENTIFIER IN expr '{' statements_zero '}' ';'
	| FOR IDENTIFIER IN RANGE expr TO expr '{' statements_zero '}' ';'
	| BREAK ';'
	| GOTO IDENTIFIER ';'
	| IDENTIFIER '(' arguments_one ')' ';'
	| SKIP ';'
	| expr ';'
	| ';'

expr	: IDENTIFIER
	| NUMBER
	| NOT expr
	| SUB expr
	| '(' expr ')'
	| expr ADD expr
	| expr SUB expr
	| expr MUL expr
	| expr QUO expr
	| expr REM expr
	| expr AND expr
	| expr OR expr
	| expr XOR expr
	| expr SHL expr
	| expr SHR expr
	| expr LAND expr
	| expr LOR expr
	| expr EQL expr
	| expr LSS expr
	| expr GTR expr
	| expr NEQ expr
	| expr LEQ expr
	| expr GEQ expr

/* ======================================== */

idents_one
	: IDENTIFIER
	| IDENTIFIER ',' idents_one

parameters_zero
	:
	| parameters_one

parameters_one
	: parameter
	| parameter ',' parameters_one

parameter
	: IDENTIFIER type

arguments_one
	: expr
	| expr ',' expr

type	: BIT
	| BYTE
	| CHANNEL '{' idents_one '}'
	| UNSTABLE CHANNEL '{' idents_one '}'

blocks_one
	: '{' statements_zero '}'
	| '{' statements_zero '}' ',' blocks_one
