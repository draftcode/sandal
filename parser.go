
//line parser.go.y:3
package sandal
import __yyfmt__ "fmt"
//line parser.go.y:3
		
import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

//line parser.go.y:16
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
const NONBLOCK_RECV = 57390
const PEEK = 57391
const TIMEOUT_PEEK = 57392
const NONBLOCK_PEEK = 57393
const SEND = 57394
const FOR = 57395
const BREAK = 57396
const IN = 57397
const RANGE = 57398
const TO = 57399
const INIT = 57400
const GOTO = 57401
const UNSTABLE = 57402
const SKIP = 57403
const UNARY = 57404

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
	"NONBLOCK_RECV",
	"PEEK",
	"TIMEOUT_PEEK",
	"NONBLOCK_PEEK",
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

//line parser.go.y:573


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

const yyNprod = 107
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 912

var yyAct = []int{

	41, 84, 130, 43, 5, 27, 5, 228, 56, 7,
	243, 6, 193, 53, 126, 124, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 185, 60, 22, 226,
	225, 68, 134, 82, 66, 121, 80, 61, 135, 224,
	223, 259, 219, 62, 218, 217, 113, 114, 115, 211,
	136, 125, 214, 122, 119, 118, 117, 116, 28, 88,
	128, 87, 258, 133, 24, 79, 131, 67, 86, 139,
	140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
	122, 26, 20, 21, 25, 213, 279, 122, 122, 122,
	277, 155, 273, 23, 267, 159, 162, 163, 164, 165,
	166, 167, 168, 169, 170, 171, 172, 173, 174, 175,
	176, 177, 178, 179, 263, 150, 247, 122, 122, 122,
	122, 246, 156, 157, 158, 127, 138, 129, 245, 29,
	44, 244, 240, 46, 212, 190, 198, 196, 188, 187,
	161, 154, 137, 93, 91, 278, 275, 271, 152, 269,
	265, 264, 181, 182, 183, 184, 256, 252, 45, 215,
	250, 221, 242, 10, 236, 220, 216, 31, 32, 192,
	33, 34, 48, 50, 35, 49, 51, 36, 37, 38,
	151, 83, 44, 123, 39, 46, 40, 122, 30, 90,
	42, 47, 65, 230, 52, 268, 233, 234, 85, 232,
	255, 231, 237, 235, 199, 239, 194, 189, 19, 18,
	45, 160, 63, 83, 44, 241, 10, 46, 59, 11,
	58, 54, 227, 249, 48, 50, 55, 49, 51, 92,
	8, 10, 9, 230, 11, 81, 12, 251, 253, 232,
	254, 231, 45, 47, 17, 260, 52, 197, 89, 16,
	15, 12, 14, 64, 261, 120, 48, 50, 266, 49,
	51, 57, 229, 270, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 272, 47, 4, 3, 52, 132,
	276, 2, 0, 0, 105, 106, 107, 108, 109, 0,
	0, 110, 111, 112, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 0,
	0, 110, 111, 112, 1, 0, 0, 13, 186, 0,
	0, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 105, 106, 107, 108, 109, 0, 238, 110, 111,
	112, 105, 106, 107, 108, 109, 0, 0, 110, 111,
	112, 0, 0, 0, 0, 95, 96, 97, 98, 99,
	100, 101, 102, 103, 104, 97, 98, 99, 100, 0,
	0, 103, 104, 0, 195, 105, 106, 107, 108, 109,
	0, 180, 110, 111, 112, 0, 0, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 0, 257, 110, 111, 112, 105, 106, 107,
	108, 109, 0, 0, 110, 111, 112, 0, 0, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 0,
	0, 0, 0, 0, 0, 210, 0, 0, 0, 105,
	106, 107, 108, 109, 0, 209, 110, 111, 112, 0,
	0, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 105, 106, 107, 108, 109, 0, 208, 110, 111,
	112, 105, 106, 107, 108, 109, 0, 0, 110, 111,
	112, 0, 0, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 0, 0, 0, 0, 0, 0, 207,
	0, 0, 0, 105, 106, 107, 108, 109, 0, 206,
	110, 111, 112, 0, 0, 95, 96, 97, 98, 99,
	100, 101, 102, 103, 104, 95, 96, 97, 98, 99,
	100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
	0, 205, 110, 111, 112, 105, 106, 107, 108, 109,
	0, 0, 110, 111, 112, 0, 0, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 0, 0, 0,
	0, 0, 0, 204, 0, 0, 0, 105, 106, 107,
	108, 109, 0, 203, 110, 111, 112, 0, 0, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	106, 107, 108, 109, 0, 202, 110, 111, 112, 105,
	106, 107, 108, 109, 0, 0, 110, 111, 112, 0,
	0, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 0, 0, 0, 0, 0, 0, 201, 0, 0,
	0, 105, 106, 107, 108, 109, 0, 200, 110, 111,
	112, 0, 0, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 0, 191,
	110, 111, 112, 105, 106, 107, 108, 109, 0, 0,
	110, 111, 112, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 0, 105, 106, 107, 108, 109, 0, 274,
	110, 111, 112, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 0, 83, 44, 0, 0, 46, 0,
	0, 0, 0, 105, 106, 107, 108, 109, 0, 248,
	110, 111, 112, 0, 83, 44, 0, 0, 46, 0,
	0, 0, 0, 45, 0, 95, 96, 97, 98, 99,
	100, 101, 102, 103, 104, 0, 0, 48, 50, 153,
	49, 51, 0, 45, 0, 105, 222, 107, 108, 109,
	0, 0, 110, 111, 112, 0, 47, 48, 50, 52,
	49, 51, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 0, 0, 0, 0, 47, 0, 0, 52,
	0, 0, 105, 106, 107, 108, 109, 0, 0, 110,
	111, 112, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 262, 0, 107, 108, 109, 0, 0, 110,
	111, 112,
}
var yyPact = []int{

	193, -1000, 193, -1000, -1000, -1000, -1000, -1000, 248, 246,
	245, 240, 146, -1000, 145, 16, 24, 15, 125, 217,
	216, 186, -1000, -43, -26, 172, 216, 128, 125, -1,
	125, 231, 800, 135, 2, -5, -7, 185, 79, 225,
	78, 696, -1000, -1000, -1000, 800, 800, 800, -9, -10,
	-11, -12, 800, 119, -56, -16, -1000, -57, 24, 800,
	24, 24, 209, -31, -17, 77, -1000, 125, 800, 800,
	800, 800, 800, 800, 800, 800, 800, 800, 800, 800,
	116, 24, 766, -1000, 76, 125, 800, 800, 800, 125,
	156, -1000, 75, -1000, -1000, 800, 800, 800, 800, 800,
	800, 800, 800, 800, 800, 800, 800, 800, 800, 800,
	800, 800, 800, -1000, -1000, 334, 800, 800, 800, 800,
	-44, -1000, 257, 74, 217, 144, 216, -1000, 664, -1000,
	105, -59, 143, 324, 24, 177, 141, -1000, -1000, 632,
	622, 590, 558, 548, 516, 484, 474, 442, 410, 400,
	-18, 69, 20, 125, -1000, 102, -22, -23, -25, 101,
	780, -1000, 376, 376, -1000, -1000, -1000, -1000, 376, 376,
	-1000, -1000, 875, 808, 297, 297, 297, 297, 297, 297,
	-1000, -27, -28, -37, -38, -1000, 800, -1000, -1000, 178,
	-1000, -1000, -1000, 24, 24, 140, 100, 139, 287, 125,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 67, -1000, -1000, 800, 98, -61, 66, 63, 56,
	51, 736, 800, -1000, -1000, -1000, -1000, -1000, 96, 178,
	-1000, -1000, -1000, -1000, 93, 24, -1000, 24, 137, 92,
	-1000, 368, -3, 135, -1000, -1000, -1000, -1000, 125, 845,
	49, -1000, -1000, 87, 86, 24, 29, -1000, -1000, 132,
	-1000, 85, 800, -1000, -1000, -1000, 83, -1000, 125, 27,
	706, -1000, 82, -1000, 125, 25, 81, -1000, 21, -1000,
}
var yyPgo = []int{

	0, 324, 281, 277, 276, 3, 11, 9, 7, 262,
	5, 58, 0, 13, 226, 8, 261, 255, 35, 2,
	66, 1,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 3, 4,
	8, 8, 9, 9, 9, 5, 6, 7, 10, 10,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 13, 14, 14, 15, 15, 15, 16, 17, 17,
	18, 18, 18, 19, 19, 19, 20, 20, 20, 20,
	20, 20, 20, 20, 21, 21, 21,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 6, 9, 5, 0, 2,
	3, 4, 4, 6, 6, 10, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 3, 5, 5,
	5, 5, 8, 11, 2, 3, 5, 2, 2, 1,
	1, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 4, 4, 3, 1,
	2, 3, 0, 1, 1, 2, 3, 2, 0, 1,
	1, 2, 3, 1, 2, 3, 1, 3, 4, 5,
	6, 7, 7, 8, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 63, 63,
	66, -20, 4, 69, 40, 60, 66, -10, -11, 4,
	63, 42, 43, 45, 46, 49, 52, 53, 54, 59,
	61, -12, 65, -5, 5, 33, 8, 66, 47, 50,
	48, 51, 69, -13, 4, -14, -15, -16, 4, 32,
	70, 63, 69, 40, -14, 64, -10, 68, 32, 17,
	18, 19, 20, 21, 22, 23, 24, 25, 26, 66,
	-10, 4, -12, 4, -21, 63, 66, 66, 66, 63,
	4, 65, 4, 65, 65, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 27, 28, 29, 30, 31,
	34, 35, 36, -12, -12, -12, 66, 66, 66, 66,
	-17, -18, -12, 64, 71, 67, 71, -20, -12, -20,
	-19, -20, 70, -12, 63, 69, 67, 65, -11, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-18, 64, -20, 63, 65, -10, -18, -18, -18, -10,
	55, 65, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	67, -18, -18, -18, -18, 70, 71, 65, -13, 63,
	-15, 65, 64, 71, 63, 70, -19, 70, -12, 63,
	65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 67, 65, 65, 32, -10, 64, 67, 67, 67,
	64, -12, 56, 67, 67, 67, 67, -18, -8, -9,
	-5, -6, -7, -19, -19, 63, 64, 63, 70, -10,
	65, -12, 64, 71, 65, 65, 65, 65, 63, -12,
	64, -8, 64, -19, -19, 63, 64, 65, 65, 44,
	-21, -10, 57, 65, 64, 64, -19, 65, 63, 64,
	-12, 64, -10, 65, 63, 64, -10, 65, 64, 65,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	82, 0, 96, 0, 0, 0, 82, 0, 18, 51,
	18, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 50, 52, 0, 0, 0, 0, 0,
	0, 0, 88, 0, 79, 0, 83, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 51, 0, 18, 0, 0, 0, 18,
	0, 44, 0, 47, 48, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 53, 54, 0, 0, 0, 0, 0,
	0, 89, 90, 0, 80, 0, 85, 87, 0, 97,
	0, 93, 0, 0, 0, 0, 0, 17, 20, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 18, 37, 0, 0, 0, 0, 0,
	0, 45, 56, 57, 58, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71, 72, 73,
	55, 0, 0, 0, 0, 78, 91, 8, 81, 10,
	86, 15, 98, 94, 0, 0, 0, 0, 0, 18,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 0, 21, 22, 0, 0, 104, 0, 0, 0,
	0, 0, 0, 74, 75, 76, 77, 92, 0, 10,
	12, 13, 14, 95, 0, 0, 99, 0, 0, 0,
	46, 0, 0, 105, 38, 39, 40, 41, 18, 0,
	0, 11, 100, 0, 0, 0, 0, 23, 24, 0,
	106, 0, 0, 9, 101, 102, 0, 16, 18, 0,
	0, 103, 0, 42, 18, 0, 0, 25, 0, 43,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 67, 3, 3, 71, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 68, 65,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 70, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 3, 64,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62,
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
		//line parser.go.y:127
		{
			yyVAL.definitions = []Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:134
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
		//line parser.go.y:150
		{
			yyVAL.definition = &DataDefinition{Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:156
		{
			yyVAL.definition = &ModuleDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:162
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:166
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
		//line parser.go.y:177
		{
			yyVAL.definition = &ConstantDefinition{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:183
		{
			yyVAL.definition = &ProcDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:189
		{
			yyVAL.definition = &InitBlock{Statements: yyS[yypt-2].statements}
		}
	case 18:
		//line parser.go.y:195
		{
			yyVAL.statements = nil
		}
	case 19:
		//line parser.go.y:199
		{
			yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 20:
		//line parser.go.y:205
		{
			yyVAL.statement = &LabelledStatement{Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 21:
		//line parser.go.y:209
		{
			yyVAL.statement = &BlockStatement{Statements: yyS[yypt-2].statements}
		}
	case 22:
		//line parser.go.y:213
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 23:
		//line parser.go.y:217
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 24:
		//line parser.go.y:221
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 25:
		//line parser.go.y:225
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 26:
		//line parser.go.y:229
		{
			yyVAL.statement = &AssignmentStatement{Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 27:
		//line parser.go.y:233
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: ADD, Expr: yyS[yypt-1].expression}
		}
	case 28:
		//line parser.go.y:237
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SUB, Expr: yyS[yypt-1].expression}
		}
	case 29:
		//line parser.go.y:241
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: MUL, Expr: yyS[yypt-1].expression}
		}
	case 30:
		//line parser.go.y:245
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: QUO, Expr: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:249
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: REM, Expr: yyS[yypt-1].expression}
		}
	case 32:
		//line parser.go.y:253
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: AND, Expr: yyS[yypt-1].expression}
		}
	case 33:
		//line parser.go.y:257
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: OR, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:261
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: XOR, Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:265
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHL, Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:269
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHR, Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:273
		{
			yyVAL.statement = &ChoiceStatement{Blocks: yyS[yypt-1].blocks}
		}
	case 38:
		//line parser.go.y:277
		{
			yyVAL.statement = &RecvStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 39:
		//line parser.go.y:281
		{
			yyVAL.statement = &PeekStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 40:
		//line parser.go.y:285
		{
			yyVAL.statement = &SendStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 41:
		//line parser.go.y:289
		{
			yyVAL.statement = &ForStatement{Statements: yyS[yypt-2].statements}
		}
	case 42:
		//line parser.go.y:293
		{
			yyVAL.statement = &ForInStatement{Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 43:
		//line parser.go.y:297
		{
			yyVAL.statement = &ForInRangeStatement{Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 44:
		//line parser.go.y:301
		{
			yyVAL.statement = &BreakStatement{}
		}
	case 45:
		//line parser.go.y:305
		{
			yyVAL.statement = &GotoStatement{Label: yyS[yypt-1].tok.lit}
		}
	case 46:
		//line parser.go.y:309
		{
			yyVAL.statement = &CallStatement{Name: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions}
		}
	case 47:
		//line parser.go.y:313
		{
			yyVAL.statement = &SkipStatement{}
		}
	case 48:
		//line parser.go.y:317
		{
			yyVAL.statement = &ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 49:
		//line parser.go.y:321
		{
			yyVAL.statement = &NullStatement{}
		}
	case 50:
		//line parser.go.y:325
		{
			yyVAL.statement = yyS[yypt-0].definition.(Statement)
		}
	case 51:
		//line parser.go.y:330
		{
			yyVAL.expression = &IdentifierExpression{Name: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:334
		{
			yyVAL.expression = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 53:
		//line parser.go.y:338
		{
			yyVAL.expression = &NotExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 54:
		//line parser.go.y:342
		{
			yyVAL.expression = &UnarySubExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 55:
		//line parser.go.y:346
		{
			yyVAL.expression = &ParenExpression{SubExpr: yyS[yypt-1].expression}
		}
	case 56:
		//line parser.go.y:350
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ADD, RHS: yyS[yypt-0].expression}
		}
	case 57:
		//line parser.go.y:354
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SUB, RHS: yyS[yypt-0].expression}
		}
	case 58:
		//line parser.go.y:358
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: MUL, RHS: yyS[yypt-0].expression}
		}
	case 59:
		//line parser.go.y:362
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: QUO, RHS: yyS[yypt-0].expression}
		}
	case 60:
		//line parser.go.y:366
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: REM, RHS: yyS[yypt-0].expression}
		}
	case 61:
		//line parser.go.y:370
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: AND, RHS: yyS[yypt-0].expression}
		}
	case 62:
		//line parser.go.y:374
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: OR, RHS: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:378
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: XOR, RHS: yyS[yypt-0].expression}
		}
	case 64:
		//line parser.go.y:382
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHL, RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:386
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHR, RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:390
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LAND, RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:394
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LOR, RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:398
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: EQL, RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:402
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LSS, RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:406
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GTR, RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:410
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: NEQ, RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:414
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LEQ, RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:418
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GEQ, RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:422
		{
			yyVAL.expression = &TimeoutRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 75:
		//line parser.go.y:426
		{
			yyVAL.expression = &TimeoutPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 76:
		//line parser.go.y:430
		{
			yyVAL.expression = &NonblockRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 77:
		//line parser.go.y:434
		{
			yyVAL.expression = &NonblockPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 78:
		//line parser.go.y:438
		{
			yyVAL.expression = &ArrayExpression{Elems: yyS[yypt-1].expressions}
		}
	case 79:
		//line parser.go.y:446
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 80:
		//line parser.go.y:450
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 81:
		//line parser.go.y:454
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 82:
		//line parser.go.y:460
		{
			yyVAL.parameters = nil
		}
	case 83:
		//line parser.go.y:464
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 84:
		//line parser.go.y:470
		{
			yyVAL.parameters = []Parameter{yyS[yypt-0].parameter}
		}
	case 85:
		//line parser.go.y:474
		{
			yyVAL.parameters = []Parameter{yyS[yypt-1].parameter}
		}
	case 86:
		//line parser.go.y:478
		{
			yyVAL.parameters = append([]Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 87:
		//line parser.go.y:484
		{
			yyVAL.parameter = Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 88:
		//line parser.go.y:490
		{
			yyVAL.expressions = nil
		}
	case 89:
		//line parser.go.y:494
		{
			yyVAL.expressions = yyS[yypt-0].expressions
		}
	case 90:
		//line parser.go.y:500
		{
			yyVAL.expressions = []Expression{yyS[yypt-0].expression}
		}
	case 91:
		//line parser.go.y:504
		{
			yyVAL.expressions = []Expression{yyS[yypt-1].expression}
		}
	case 92:
		//line parser.go.y:508
		{
			yyVAL.expressions = append([]Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 93:
		//line parser.go.y:514
		{
			yyVAL.typetypes = []Type{yyS[yypt-0].typetype}
		}
	case 94:
		//line parser.go.y:518
		{
			yyVAL.typetypes = []Type{yyS[yypt-1].typetype}
		}
	case 95:
		//line parser.go.y:522
		{
			yyVAL.typetypes = append([]Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 96:
		//line parser.go.y:527
		{
			yyVAL.typetype = &NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 97:
		//line parser.go.y:531
		{
			yyVAL.typetype = &ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 98:
		//line parser.go.y:535
		{
			yyVAL.typetype = &HandshakeChannelType{IsUnstable: false, Elems: yyS[yypt-1].typetypes}
		}
	case 99:
		//line parser.go.y:539
		{
			yyVAL.typetype = &HandshakeChannelType{IsUnstable: true, Elems: yyS[yypt-1].typetypes}
		}
	case 100:
		//line parser.go.y:543
		{
			yyVAL.typetype = &BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 101:
		//line parser.go.y:547
		{
			yyVAL.typetype = &BufferedChannelType{IsUnstable: false, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 102:
		//line parser.go.y:551
		{
			yyVAL.typetype = &BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 103:
		//line parser.go.y:555
		{
			yyVAL.typetype = &BufferedChannelType{IsUnstable: true, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 104:
		//line parser.go.y:561
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-1].statements}}
		}
	case 105:
		//line parser.go.y:565
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-2].statements}}
		}
	case 106:
		//line parser.go.y:569
		{
			yyVAL.blocks = append([]BlockStatement{BlockStatement{Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
