
//line parser.go.y:3
package sandal
import __yyfmt__ "fmt"
//line parser.go.y:3
		
import (
	"log")

type Token struct {
	tok int
	lit string
	pos Position
}

//line parser.go.y:15
type yySymType struct{
	yys int
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

const IDENTIFIER = 57346
const NUMBER = 57347
const COMMENT = 57348
const ADD = 57349
const SUB = 57350
const MUL = 57351
const QUO = 57352
const REM = 57353
const AND = 57354
const OR = 57355
const XOR = 57356
const SHL = 57357
const SHR = 57358
const ADD_ASSIGN = 57359
const SUB_ASSIGN = 57360
const MUL_ASSIGN = 57361
const QUO_ASSIGN = 57362
const REM_ASSIGN = 57363
const AND_ASSIGN = 57364
const OR_ASSIGN = 57365
const XOR_ASSIGN = 57366
const SHL_ASSIGN = 57367
const SHR_ASSIGN = 57368
const LAND = 57369
const LOR = 57370
const EQL = 57371
const LSS = 57372
const GTR = 57373
const ASSIGN = 57374
const NOT = 57375
const NEQ = 57376
const LEQ = 57377
const GEQ = 57378
const DATA = 57379
const CONST = 57380
const MODULE = 57381
const CHANNEL = 57382
const PROC = 57383
const VAR = 57384
const IF = 57385
const ELSE = 57386
const CHOICE = 57387
const RECV = 57388
const TIMEOUT_RECV = 57389
const PEEK = 57390
const TIMEOUT_PEEK = 57391
const SEND = 57392
const FOR = 57393
const BREAK = 57394
const IN = 57395
const RANGE = 57396
const TO = 57397
const INIT = 57398
const GOTO = 57399
const UNSTABLE = 57400
const SKIP = 57401
const UNARY = 57402

var yyToknames = []string{
	"IDENTIFIER",
	"NUMBER",
	"COMMENT",
	"ADD",
	"SUB",
	"MUL",
	"QUO",
	"REM",
	"AND",
	"OR",
	"XOR",
	"SHL",
	"SHR",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"QUO_ASSIGN",
	"REM_ASSIGN",
	"AND_ASSIGN",
	"OR_ASSIGN",
	"XOR_ASSIGN",
	"SHL_ASSIGN",
	"SHR_ASSIGN",
	"LAND",
	"LOR",
	"EQL",
	"LSS",
	"GTR",
	"ASSIGN",
	"NOT",
	"NEQ",
	"LEQ",
	"GEQ",
	"DATA",
	"CONST",
	"MODULE",
	"CHANNEL",
	"PROC",
	"VAR",
	"IF",
	"ELSE",
	"CHOICE",
	"RECV",
	"TIMEOUT_RECV",
	"PEEK",
	"TIMEOUT_PEEK",
	"SEND",
	"FOR",
	"BREAK",
	"IN",
	"RANGE",
	"TO",
	"INIT",
	"GOTO",
	"UNSTABLE",
	"SKIP",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:546


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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 101
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 835

var yyAct = []int{

	37, 210, 74, 39, 5, 23, 5, 7, 50, 6,
	204, 47, 54, 40, 227, 217, 42, 114, 112, 174,
	116, 168, 53, 202, 201, 197, 196, 195, 189, 73,
	57, 211, 71, 121, 109, 113, 107, 106, 78, 232,
	77, 41, 103, 104, 105, 192, 76, 110, 22, 20,
	250, 248, 244, 239, 236, 44, 118, 45, 231, 221,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 110, 43, 220, 119, 46, 191, 110, 110, 110,
	219, 140, 218, 117, 115, 144, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
	161, 162, 163, 164, 137, 135, 214, 110, 110, 190,
	170, 141, 142, 143, 25, 40, 146, 139, 42, 122,
	83, 81, 249, 173, 171, 246, 241, 238, 229, 226,
	224, 216, 198, 194, 136, 111, 54, 40, 80, 56,
	42, 166, 167, 41, 193, 240, 199, 75, 10, 212,
	177, 175, 27, 28, 172, 29, 30, 44, 31, 45,
	32, 33, 34, 19, 18, 41, 24, 35, 145, 36,
	110, 26, 21, 38, 43, 176, 206, 46, 52, 44,
	208, 45, 207, 213, 10, 49, 200, 11, 8, 10,
	9, 48, 11, 215, 82, 79, 43, 72, 17, 46,
	16, 223, 12, 15, 203, 14, 209, 12, 55, 206,
	108, 51, 205, 208, 228, 207, 225, 4, 1, 3,
	233, 13, 2, 0, 0, 123, 0, 0, 234, 237,
	0, 87, 88, 89, 90, 0, 242, 93, 94, 0,
	0, 0, 0, 0, 0, 0, 243, 0, 0, 0,
	0, 247, 85, 86, 87, 88, 89, 90, 91, 92,
	93, 94, 85, 86, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 97, 98, 99, 0, 0, 100,
	101, 102, 95, 96, 97, 98, 99, 0, 0, 100,
	101, 102, 0, 0, 0, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 169, 95, 96, 97, 98, 99,
	165, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	0, 230, 100, 101, 102, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 188, 0, 0, 0, 95, 96, 97, 98, 99,
	0, 187, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	0, 186, 100, 101, 102, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 185, 0, 0, 0, 95, 96, 97, 98, 99,
	0, 184, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	0, 183, 100, 101, 102, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 182, 0, 0, 0, 95, 96, 97, 98, 99,
	0, 181, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	0, 180, 100, 101, 102, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 179, 0, 0, 0, 95, 96, 97, 98, 99,
	0, 178, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	0, 120, 100, 101, 102, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 0, 0, 0, 0, 0,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 245,
	95, 96, 97, 98, 99, 0, 0, 100, 101, 102,
	0, 0, 0, 0, 0, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 222, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 0, 0, 0, 0, 138,
	0, 0, 0, 0, 0, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 58, 95, 96, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 235, 0, 95, 0, 97, 98, 99,
	0, 0, 100, 101, 102, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 0, 0, 97, 98, 99,
	0, 0, 100, 101, 102,
}
var yyPact = []int{

	151, -1000, 151, -1000, -1000, -1000, -1000, -1000, 201, 199,
	196, 194, 103, -1000, 102, -15, 140, -16, 110, 187,
	174, 8, 174, 77, 110, 698, 110, 193, 8, 86,
	-18, -24, -26, 134, 58, 190, 57, 598, -1000, -1000,
	-1000, 8, 8, 8, -27, -28, 8, 73, -51, -30,
	-1000, -52, 16, 568, -1000, -32, 56, -1000, 110, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 72, 16, 678, 54, 110, 8, 8, 8, 110,
	115, -1000, 53, -1000, -1000, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, -1000, -1000, 255, 8, 8, -47, -1000,
	245, 47, 187, 93, 174, -1000, -1000, -49, 90, 135,
	-1000, 89, -1000, -1000, 538, 528, 498, 468, 458, 428,
	398, 388, 358, 328, 318, -37, 46, 13, 110, -1000,
	71, -38, -39, -40, 70, 132, -1000, 222, 222, -1000,
	-1000, -1000, -1000, 222, 222, -1000, -1000, 798, 768, 808,
	808, 808, 808, 808, 808, -1000, -41, -42, -1000, 8,
	-1000, -1000, 146, -1000, 16, 16, 88, 110, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 43,
	-1000, -1000, 8, 69, -54, 19, 17, 10, -4, 643,
	8, -1000, -1000, -1000, 68, 146, -1000, -1000, -1000, -1000,
	67, -55, 16, 66, -1000, 288, -5, 86, -1000, -1000,
	-1000, -1000, 110, 738, -9, -1000, -1000, 16, 65, -10,
	-1000, -1000, 84, -1000, 64, 8, -1000, -1000, -1000, -1000,
	110, -11, 608, 63, -1000, 110, -12, 60, -1000, -13,
	-1000,
}
var yyPgo = []int{

	0, 218, 222, 219, 217, 3, 9, 7, 10, 212,
	5, 166, 0, 11, 185, 8, 211, 210, 34, 1,
	31, 2,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 3, 4,
	8, 8, 9, 9, 9, 5, 6, 7, 10, 10,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 13, 13, 13,
	14, 14, 15, 15, 15, 16, 17, 17, 18, 18,
	18, 19, 19, 19, 20, 20, 20, 20, 21, 21,
	21,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 5, 9, 5, 0, 2,
	3, 4, 4, 6, 6, 10, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 3, 5, 5,
	5, 5, 8, 11, 2, 3, 5, 2, 2, 1,
	1, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 3, 1, 2, 3,
	0, 1, 1, 2, 3, 2, 0, 1, 1, 2,
	3, 1, 2, 3, 1, 3, 4, 5, 3, 4,
	5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 56, -1, 4, 4, 4, 4, 61, 61,
	64, 32, 64, -10, -11, 4, 61, 42, 43, 45,
	46, 48, 50, 51, 52, 57, 59, -12, 63, -5,
	5, 33, 8, 64, 47, 49, 67, -13, 4, -14,
	-15, -16, 4, -12, 4, -14, 62, -10, 66, 32,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	64, -10, 4, -12, -21, 61, 64, 64, 64, 61,
	4, 63, 4, 63, 63, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 27, 28, 29, 30, 31,
	34, 35, 36, -12, -12, -12, 64, 64, -17, -18,
	-12, 62, 69, 65, 69, -20, 4, 67, 40, 58,
	63, 65, 63, -11, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -18, 62, -20, 61, 63,
	-10, -18, -18, -18, -10, 53, 63, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, 65, -18, -18, 68, 69,
	63, -13, 61, -15, 68, 61, 40, 61, 63, 63,
	63, 63, 63, 63, 63, 63, 63, 63, 63, 65,
	63, 63, 32, -10, 62, 65, 65, 65, 62, -12,
	54, 65, 65, -18, -8, -9, -5, -6, -7, -20,
	-19, -20, 61, -10, 63, -12, 62, 69, 63, 63,
	63, 63, 61, -12, 62, -8, 62, 69, -19, 62,
	63, 63, 44, -21, -10, 55, 63, -19, 62, 63,
	61, 62, -12, -10, 63, 61, 62, -10, 63, 62,
	63,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	80, 0, 80, 0, 18, 51, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 49, 50,
	52, 0, 0, 0, 0, 0, 86, 0, 77, 0,
	81, 82, 0, 0, 51, 0, 0, 19, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 18, 0, 0, 0, 18,
	0, 44, 0, 47, 48, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 53, 54, 0, 0, 0, 0, 87,
	88, 0, 78, 0, 83, 85, 94, 0, 0, 0,
	15, 0, 17, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 37,
	0, 0, 0, 0, 0, 0, 45, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 72, 73, 55, 0, 0, 76, 89,
	8, 79, 10, 84, 0, 0, 0, 18, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 0,
	21, 22, 0, 0, 98, 0, 0, 0, 0, 0,
	0, 74, 75, 90, 0, 10, 12, 13, 14, 95,
	0, 91, 0, 0, 46, 0, 0, 99, 38, 39,
	40, 41, 18, 0, 0, 11, 96, 92, 0, 0,
	23, 24, 0, 100, 0, 0, 9, 93, 97, 16,
	18, 0, 0, 0, 42, 18, 0, 0, 25, 0,
	43,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	64, 65, 3, 3, 69, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 66, 63,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 67, 3, 68, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 62,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:124
		{
			yyVAL.definitions = []Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:131
		{
			yyVAL.definitions = append([]Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 3:
		yyVAL.definition = yyS[yypt-0].definition
	case 4:
		yyVAL.definition = yyS[yypt-0].definition
	case 5:
		yyVAL.definition = yyS[yypt-0].definition
	case 6:
		yyVAL.definition = yyS[yypt-0].definition
	case 7:
		yyVAL.definition = yyS[yypt-0].definition
	case 8:
		//line parser.go.y:147
		{
			yyVAL.definition = &DataDefinition{Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:153
		{
			yyVAL.definition = &ModuleDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:159
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:163
		{
			yyVAL.definitions = append([]Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
		}
	case 12:
		yyVAL.definition = yyS[yypt-0].definition
	case 13:
		yyVAL.definition = yyS[yypt-0].definition
	case 14:
		yyVAL.definition = yyS[yypt-0].definition
	case 15:
		//line parser.go.y:174
		{
			yyVAL.definition = &ConstantDefinition{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:180
		{
			yyVAL.definition = &ProcDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:186
		{
			yyVAL.definition = &InitBlock{Statements: yyS[yypt-2].statements}
		}
	case 18:
		//line parser.go.y:192
		{
			yyVAL.statements = nil
		}
	case 19:
		//line parser.go.y:196
		{
			yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 20:
		//line parser.go.y:202
		{
			yyVAL.statement = &LabelledStatement{Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 21:
		//line parser.go.y:206
		{
			yyVAL.statement = &BlockStatement{Statements: yyS[yypt-2].statements}
		}
	case 22:
		//line parser.go.y:210
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 23:
		//line parser.go.y:214
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 24:
		//line parser.go.y:218
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 25:
		//line parser.go.y:222
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 26:
		//line parser.go.y:226
		{
			yyVAL.statement = &AssignmentStatement{Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 27:
		//line parser.go.y:230
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: ADD, Expr: yyS[yypt-1].expression}
		}
	case 28:
		//line parser.go.y:234
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SUB, Expr: yyS[yypt-1].expression}
		}
	case 29:
		//line parser.go.y:238
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: MUL, Expr: yyS[yypt-1].expression}
		}
	case 30:
		//line parser.go.y:242
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: QUO, Expr: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:246
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: REM, Expr: yyS[yypt-1].expression}
		}
	case 32:
		//line parser.go.y:250
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: AND, Expr: yyS[yypt-1].expression}
		}
	case 33:
		//line parser.go.y:254
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: OR, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:258
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: XOR, Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:262
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHL, Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:266
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHR, Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:270
		{
			yyVAL.statement = &ChoiceStatement{Blocks: yyS[yypt-1].blocks}
		}
	case 38:
		//line parser.go.y:274
		{
			yyVAL.statement = &RecvStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 39:
		//line parser.go.y:278
		{
			yyVAL.statement = &PeekStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 40:
		//line parser.go.y:282
		{
			yyVAL.statement = &SendStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 41:
		//line parser.go.y:286
		{
			yyVAL.statement = &ForStatement{Statements: yyS[yypt-2].statements}
		}
	case 42:
		//line parser.go.y:290
		{
			yyVAL.statement = &ForInStatement{Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 43:
		//line parser.go.y:294
		{
			yyVAL.statement = &ForInRangeStatement{Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 44:
		//line parser.go.y:298
		{
			yyVAL.statement = &BreakStatement{}
		}
	case 45:
		//line parser.go.y:302
		{
			yyVAL.statement = &GotoStatement{Label: yyS[yypt-1].tok.lit}
		}
	case 46:
		//line parser.go.y:306
		{
			yyVAL.statement = &CallStatement{Name: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions}
		}
	case 47:
		//line parser.go.y:310
		{
			yyVAL.statement = &SkipStatement{}
		}
	case 48:
		//line parser.go.y:314
		{
			yyVAL.statement = &ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 49:
		//line parser.go.y:318
		{
			yyVAL.statement = &NullStatement{}
		}
	case 50:
		//line parser.go.y:322
		{
			yyVAL.statement = yyS[yypt-0].definition.(Statement)
		}
	case 51:
		//line parser.go.y:327
		{
			yyVAL.expression = &IdentifierExpression{Name: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:331
		{
			yyVAL.expression = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 53:
		//line parser.go.y:335
		{
			yyVAL.expression = &NotExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 54:
		//line parser.go.y:339
		{
			yyVAL.expression = &UnarySubExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 55:
		//line parser.go.y:343
		{
			yyVAL.expression = &ParenExpression{SubExpr: yyS[yypt-1].expression}
		}
	case 56:
		//line parser.go.y:347
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ADD, RHS: yyS[yypt-0].expression}
		}
	case 57:
		//line parser.go.y:351
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SUB, RHS: yyS[yypt-0].expression}
		}
	case 58:
		//line parser.go.y:355
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: MUL, RHS: yyS[yypt-0].expression}
		}
	case 59:
		//line parser.go.y:359
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: QUO, RHS: yyS[yypt-0].expression}
		}
	case 60:
		//line parser.go.y:363
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: REM, RHS: yyS[yypt-0].expression}
		}
	case 61:
		//line parser.go.y:367
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: AND, RHS: yyS[yypt-0].expression}
		}
	case 62:
		//line parser.go.y:371
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: OR, RHS: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:375
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: XOR, RHS: yyS[yypt-0].expression}
		}
	case 64:
		//line parser.go.y:379
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHL, RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:383
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHR, RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:387
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LAND, RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:391
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LOR, RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:395
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: EQL, RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:399
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LSS, RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:403
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GTR, RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:407
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: NEQ, RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:411
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LEQ, RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:415
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GEQ, RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:419
		{
			yyVAL.expression = &TimeoutRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 75:
		//line parser.go.y:423
		{
			yyVAL.expression = &TimeoutPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 76:
		//line parser.go.y:427
		{
			yyVAL.expression = &ArrayExpression{Elems: yyS[yypt-1].expressions}
		}
	case 77:
		//line parser.go.y:435
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 78:
		//line parser.go.y:439
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 79:
		//line parser.go.y:443
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 80:
		//line parser.go.y:449
		{
			yyVAL.parameters = nil
		}
	case 81:
		//line parser.go.y:453
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 82:
		//line parser.go.y:459
		{
			yyVAL.parameters = []Parameter{yyS[yypt-0].parameter}
		}
	case 83:
		//line parser.go.y:463
		{
			yyVAL.parameters = []Parameter{yyS[yypt-1].parameter}
		}
	case 84:
		//line parser.go.y:467
		{
			yyVAL.parameters = append([]Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 85:
		//line parser.go.y:473
		{
			yyVAL.parameter = Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 86:
		//line parser.go.y:479
		{
			yyVAL.expressions = nil
		}
	case 87:
		//line parser.go.y:483
		{
			yyVAL.expressions = yyS[yypt-0].expressions
		}
	case 88:
		//line parser.go.y:489
		{
			yyVAL.expressions = []Expression{yyS[yypt-0].expression}
		}
	case 89:
		//line parser.go.y:493
		{
			yyVAL.expressions = []Expression{yyS[yypt-1].expression}
		}
	case 90:
		//line parser.go.y:497
		{
			yyVAL.expressions = append([]Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 91:
		//line parser.go.y:503
		{
			yyVAL.typetypes = []Type{yyS[yypt-0].typetype}
		}
	case 92:
		//line parser.go.y:507
		{
			yyVAL.typetypes = []Type{yyS[yypt-1].typetype}
		}
	case 93:
		//line parser.go.y:511
		{
			yyVAL.typetypes = append([]Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 94:
		//line parser.go.y:516
		{
			yyVAL.typetype = &NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 95:
		//line parser.go.y:520
		{
			yyVAL.typetype = &SetType{SetType: yyS[yypt-0].typetype}
		}
	case 96:
		//line parser.go.y:524
		{
			yyVAL.typetype = &ChannelType{IsUnstable: false, Elems: yyS[yypt-1].typetypes}
		}
	case 97:
		//line parser.go.y:528
		{
			yyVAL.typetype = &ChannelType{IsUnstable: true, Elems: yyS[yypt-1].typetypes}
		}
	case 98:
		//line parser.go.y:534
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-1].statements}}
		}
	case 99:
		//line parser.go.y:538
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-2].statements}}
		}
	case 100:
		//line parser.go.y:542
		{
			yyVAL.blocks = append([]BlockStatement{BlockStatement{Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
