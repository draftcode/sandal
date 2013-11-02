
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

//line parser.go.y:564


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

const yyNprod = 105
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 950

var yyAct = []int{

	37, 76, 246, 212, 7, 23, 39, 5, 6, 5,
	52, 49, 56, 40, 256, 225, 42, 118, 116, 180,
	174, 220, 55, 181, 210, 209, 208, 207, 203, 75,
	59, 247, 73, 202, 113, 201, 195, 125, 117, 241,
	111, 41, 105, 106, 107, 110, 109, 108, 80, 114,
	79, 198, 78, 22, 20, 44, 46, 272, 45, 47,
	240, 270, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 114, 43, 261, 251, 48, 236, 114,
	114, 114, 245, 144, 197, 229, 119, 148, 151, 152,
	153, 154, 155, 156, 157, 158, 159, 160, 161, 162,
	163, 164, 165, 166, 167, 168, 141, 139, 228, 114,
	114, 114, 114, 145, 146, 147, 227, 226, 222, 196,
	176, 150, 143, 126, 85, 120, 83, 271, 177, 179,
	269, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 267, 265, 170, 171, 172, 173, 264, 199, 255,
	205, 97, 98, 99, 100, 101, 253, 238, 102, 103,
	104, 122, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 232, 224, 204, 200, 114, 61, 140, 115,
	82, 123, 219, 216, 58, 214, 259, 215, 252, 221,
	121, 249, 248, 77, 234, 175, 183, 178, 19, 223,
	18, 149, 8, 10, 9, 24, 11, 231, 182, 21,
	211, 72, 217, 60, 51, 54, 50, 233, 216, 84,
	214, 237, 215, 12, 74, 17, 10, 242, 16, 11,
	15, 25, 40, 14, 112, 42, 243, 57, 53, 81,
	213, 89, 90, 91, 92, 254, 12, 95, 96, 4,
	1, 257, 258, 13, 3, 2, 0, 0, 260, 263,
	41, 0, 266, 0, 0, 10, 127, 0, 268, 27,
	28, 0, 29, 30, 44, 46, 31, 45, 47, 32,
	33, 34, 0, 0, 0, 0, 35, 0, 36, 0,
	26, 0, 38, 43, 0, 0, 48, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	100, 101, 0, 0, 102, 103, 104, 97, 98, 99,
	100, 101, 0, 0, 102, 103, 104, 0, 0, 0,
	0, 0, 0, 0, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 0, 0, 0, 0, 0, 0,
	250, 0, 0, 0, 97, 98, 99, 100, 101, 0,
	235, 102, 103, 104, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 97, 98, 99, 100, 101, 0,
	0, 102, 103, 104, 169, 0, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 0, 239, 102, 103, 104, 97, 98, 99, 100,
	101, 0, 0, 102, 103, 104, 0, 0, 87, 88,
	89, 90, 91, 92, 93, 94, 95, 96, 0, 0,
	0, 0, 0, 0, 194, 0, 0, 0, 97, 98,
	99, 100, 101, 0, 193, 102, 103, 104, 0, 0,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 0, 192, 102, 103, 104,
	97, 98, 99, 100, 101, 0, 0, 102, 103, 104,
	0, 0, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 0, 0, 0, 0, 0, 0, 191, 0,
	0, 0, 97, 98, 99, 100, 101, 0, 190, 102,
	103, 104, 0, 0, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 97, 98, 99, 100, 101, 0,
	189, 102, 103, 104, 97, 98, 99, 100, 101, 0,
	0, 102, 103, 104, 0, 0, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 0, 0, 0, 0,
	0, 0, 188, 0, 0, 0, 97, 98, 99, 100,
	101, 0, 187, 102, 103, 104, 0, 0, 87, 88,
	89, 90, 91, 92, 93, 94, 95, 96, 87, 88,
	89, 90, 91, 92, 93, 94, 95, 96, 97, 98,
	99, 100, 101, 0, 186, 102, 103, 104, 97, 98,
	99, 100, 101, 0, 0, 102, 103, 104, 0, 0,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	0, 0, 0, 0, 0, 0, 185, 0, 0, 0,
	97, 98, 99, 100, 101, 0, 184, 102, 103, 104,
	0, 0, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 0, 124, 102,
	103, 104, 97, 98, 99, 100, 101, 0, 0, 102,
	103, 104, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 0, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 97, 98, 99, 100, 101, 0, 262, 102,
	103, 104, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 0, 56, 40, 0, 0, 42, 0, 0,
	0, 0, 97, 98, 99, 100, 101, 0, 230, 102,
	103, 104, 0, 0, 0, 0, 0, 56, 40, 0,
	0, 42, 41, 0, 56, 40, 0, 0, 42, 0,
	0, 0, 0, 0, 0, 0, 44, 46, 142, 45,
	47, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	0, 0, 0, 41, 0, 43, 0, 0, 48, 218,
	44, 46, 0, 45, 47, 0, 0, 44, 46, 206,
	45, 47, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 0, 48, 0, 0, 0, 43, 0, 0, 48,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 0, 0, 102, 103, 104,
	97, 0, 99, 100, 101, 0, 0, 102, 103, 104,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	244, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 99, 100, 101, 0, 0, 102, 103, 104,
}
var yyPact = []int{

	165, -1000, 165, -1000, -1000, -1000, -1000, -1000, 229, 226,
	224, 221, 137, -1000, 135, -12, 177, -13, 227, 212,
	211, 810, 211, 120, 227, 145, 227, 220, 810, 130,
	-14, -16, -18, 176, 61, 215, 59, 695, -1000, -1000,
	-1000, 810, 810, 810, -19, -20, -21, -26, 810, 115,
	-53, -29, -1000, -54, 121, 663, -1000, -30, 58, -1000,
	227, 810, 810, 810, 810, 810, 810, 810, 810, 810,
	810, 810, 810, 114, 121, 765, 57, 227, 810, 810,
	810, 227, 146, -1000, 56, -1000, -1000, 810, 810, 810,
	810, 810, 810, 810, 810, 810, 810, 810, 810, 810,
	810, 810, 810, 810, 810, -1000, -1000, 337, 810, 810,
	810, 810, -50, -1000, 124, 55, 212, 134, 211, -1000,
	-1000, -51, -46, 168, -1000, 133, -1000, -1000, 631, 621,
	589, 557, 547, 515, 483, 473, 441, 409, 399, -31,
	54, 19, 227, -1000, 111, -32, -34, -39, 110, 803,
	-1000, 232, 232, -1000, -1000, -1000, -1000, 232, 232, -1000,
	-1000, 913, 883, 377, 377, 377, 377, 377, 377, -1000,
	-40, -41, -42, -43, -1000, 810, -1000, -1000, 188, -1000,
	121, 779, -48, 227, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 53, -1000, -1000, 810, 109,
	-56, 52, 51, 43, 20, 735, 810, -1000, -1000, -1000,
	-1000, -1000, 108, 188, -1000, -1000, -1000, -1000, 131, 300,
	8, 93, -1000, 367, -5, 130, -1000, -1000, -1000, -1000,
	227, 873, 17, -1000, 121, 129, 128, 290, 11, -1000,
	-1000, 125, -1000, 92, 810, -1000, 85, -57, 121, 121,
	123, -1000, 227, 10, 705, -1000, 121, 83, 78, 121,
	77, -1000, 227, -1000, -1000, -1000, 66, -4, 63, -1000,
	-1000, -8, -1000,
}
var yyPgo = []int{

	0, 250, 255, 254, 249, 6, 8, 4, 3, 240,
	5, 205, 0, 11, 214, 10, 238, 234, 34, 2,
	31, 1,
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
	20, 20, 21, 21, 21,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 5, 9, 5, 0, 2,
	3, 4, 4, 6, 6, 10, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 3, 5, 5,
	5, 5, 8, 11, 2, 3, 5, 2, 2, 1,
	1, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 4, 4, 3, 1,
	2, 3, 0, 1, 1, 2, 3, 2, 0, 1,
	1, 2, 3, 1, 2, 3, 1, 3, 6, 7,
	7, 8, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 63, 63,
	66, 32, 66, -10, -11, 4, 63, 42, 43, 45,
	46, 49, 52, 53, 54, 59, 61, -12, 65, -5,
	5, 33, 8, 66, 47, 50, 48, 51, 69, -13,
	4, -14, -15, -16, 4, -12, 4, -14, 64, -10,
	68, 32, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 66, -10, 4, -12, -21, 63, 66, 66,
	66, 63, 4, 65, 4, 65, 65, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 27, 28, 29,
	30, 31, 34, 35, 36, -12, -12, -12, 66, 66,
	66, 66, -17, -18, -12, 64, 71, 67, 71, -20,
	4, 69, 40, 60, 65, 67, 65, -11, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -18,
	64, -20, 63, 65, -10, -18, -18, -18, -10, 55,
	65, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 67,
	-18, -18, -18, -18, 70, 71, 65, -13, 63, -15,
	70, 69, 40, 63, 65, 65, 65, 65, 65, 65,
	65, 65, 65, 65, 65, 67, 65, 65, 32, -10,
	64, 67, 67, 67, 64, -12, 56, 67, 67, 67,
	67, -18, -8, -9, -5, -6, -7, -20, 70, -12,
	69, -10, 65, -12, 64, 71, 65, 65, 65, 65,
	63, -12, 64, -8, 63, 70, 70, -12, 64, 65,
	65, 44, -21, -10, 57, 65, -19, -20, 63, 63,
	70, 65, 63, 64, -12, 64, 71, -19, -19, 63,
	-10, 65, 63, -19, 64, 64, -19, 64, -10, 64,
	65, 64, 65,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	82, 0, 82, 0, 18, 51, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 49, 50,
	52, 0, 0, 0, 0, 0, 0, 0, 88, 0,
	79, 0, 83, 84, 0, 0, 51, 0, 0, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 18, 0, 0,
	0, 18, 0, 44, 0, 47, 48, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 53, 54, 0, 0, 0,
	0, 0, 0, 89, 90, 0, 80, 0, 85, 87,
	96, 0, 0, 0, 15, 0, 17, 20, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 37, 0, 0, 0, 0, 0, 0,
	45, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 55,
	0, 0, 0, 0, 78, 91, 8, 81, 10, 86,
	0, 0, 0, 18, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 0, 21, 22, 0, 0,
	102, 0, 0, 0, 0, 0, 0, 74, 75, 76,
	77, 92, 0, 10, 12, 13, 14, 97, 0, 0,
	0, 0, 46, 0, 0, 103, 38, 39, 40, 41,
	18, 0, 0, 11, 0, 0, 0, 0, 0, 23,
	24, 0, 104, 0, 0, 9, 0, 93, 0, 0,
	0, 16, 18, 0, 0, 98, 94, 0, 0, 0,
	0, 42, 18, 95, 99, 100, 0, 0, 0, 101,
	25, 0, 43,
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
		//line parser.go.y:126
		{
			yyVAL.definitions = []Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:133
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
		//line parser.go.y:149
		{
			yyVAL.definition = &DataDefinition{Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:155
		{
			yyVAL.definition = &ModuleDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:161
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:165
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
		//line parser.go.y:176
		{
			yyVAL.definition = &ConstantDefinition{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:182
		{
			yyVAL.definition = &ProcDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:188
		{
			yyVAL.definition = &InitBlock{Statements: yyS[yypt-2].statements}
		}
	case 18:
		//line parser.go.y:194
		{
			yyVAL.statements = nil
		}
	case 19:
		//line parser.go.y:198
		{
			yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 20:
		//line parser.go.y:204
		{
			yyVAL.statement = &LabelledStatement{Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 21:
		//line parser.go.y:208
		{
			yyVAL.statement = &BlockStatement{Statements: yyS[yypt-2].statements}
		}
	case 22:
		//line parser.go.y:212
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 23:
		//line parser.go.y:216
		{
			yyVAL.statement = &VarDeclStatement{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 24:
		//line parser.go.y:220
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 25:
		//line parser.go.y:224
		{
			yyVAL.statement = &IfStatement{Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 26:
		//line parser.go.y:228
		{
			yyVAL.statement = &AssignmentStatement{Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 27:
		//line parser.go.y:232
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: ADD, Expr: yyS[yypt-1].expression}
		}
	case 28:
		//line parser.go.y:236
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SUB, Expr: yyS[yypt-1].expression}
		}
	case 29:
		//line parser.go.y:240
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: MUL, Expr: yyS[yypt-1].expression}
		}
	case 30:
		//line parser.go.y:244
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: QUO, Expr: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:248
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: REM, Expr: yyS[yypt-1].expression}
		}
	case 32:
		//line parser.go.y:252
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: AND, Expr: yyS[yypt-1].expression}
		}
	case 33:
		//line parser.go.y:256
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: OR, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:260
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: XOR, Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:264
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHL, Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:268
		{
			yyVAL.statement = &OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: SHR, Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:272
		{
			yyVAL.statement = &ChoiceStatement{Blocks: yyS[yypt-1].blocks}
		}
	case 38:
		//line parser.go.y:276
		{
			yyVAL.statement = &RecvStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 39:
		//line parser.go.y:280
		{
			yyVAL.statement = &PeekStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 40:
		//line parser.go.y:284
		{
			yyVAL.statement = &SendStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 41:
		//line parser.go.y:288
		{
			yyVAL.statement = &ForStatement{Statements: yyS[yypt-2].statements}
		}
	case 42:
		//line parser.go.y:292
		{
			yyVAL.statement = &ForInStatement{Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 43:
		//line parser.go.y:296
		{
			yyVAL.statement = &ForInRangeStatement{Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 44:
		//line parser.go.y:300
		{
			yyVAL.statement = &BreakStatement{}
		}
	case 45:
		//line parser.go.y:304
		{
			yyVAL.statement = &GotoStatement{Label: yyS[yypt-1].tok.lit}
		}
	case 46:
		//line parser.go.y:308
		{
			yyVAL.statement = &CallStatement{Name: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions}
		}
	case 47:
		//line parser.go.y:312
		{
			yyVAL.statement = &SkipStatement{}
		}
	case 48:
		//line parser.go.y:316
		{
			yyVAL.statement = &ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 49:
		//line parser.go.y:320
		{
			yyVAL.statement = &NullStatement{}
		}
	case 50:
		//line parser.go.y:324
		{
			yyVAL.statement = yyS[yypt-0].definition.(Statement)
		}
	case 51:
		//line parser.go.y:329
		{
			yyVAL.expression = &IdentifierExpression{Name: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:333
		{
			yyVAL.expression = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 53:
		//line parser.go.y:337
		{
			yyVAL.expression = &NotExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 54:
		//line parser.go.y:341
		{
			yyVAL.expression = &UnarySubExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 55:
		//line parser.go.y:345
		{
			yyVAL.expression = &ParenExpression{SubExpr: yyS[yypt-1].expression}
		}
	case 56:
		//line parser.go.y:349
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ADD, RHS: yyS[yypt-0].expression}
		}
	case 57:
		//line parser.go.y:353
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SUB, RHS: yyS[yypt-0].expression}
		}
	case 58:
		//line parser.go.y:357
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: MUL, RHS: yyS[yypt-0].expression}
		}
	case 59:
		//line parser.go.y:361
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: QUO, RHS: yyS[yypt-0].expression}
		}
	case 60:
		//line parser.go.y:365
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: REM, RHS: yyS[yypt-0].expression}
		}
	case 61:
		//line parser.go.y:369
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: AND, RHS: yyS[yypt-0].expression}
		}
	case 62:
		//line parser.go.y:373
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: OR, RHS: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:377
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: XOR, RHS: yyS[yypt-0].expression}
		}
	case 64:
		//line parser.go.y:381
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHL, RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:385
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: SHR, RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:389
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LAND, RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:393
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LOR, RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:397
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: EQL, RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:401
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LSS, RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:405
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GTR, RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:409
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: NEQ, RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:413
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: LEQ, RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:417
		{
			yyVAL.expression = &BinOpExpression{LHS: yyS[yypt-2].expression, Operator: GEQ, RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:421
		{
			yyVAL.expression = &TimeoutRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 75:
		//line parser.go.y:425
		{
			yyVAL.expression = &TimeoutPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 76:
		//line parser.go.y:429
		{
			yyVAL.expression = &NonblockRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 77:
		//line parser.go.y:433
		{
			yyVAL.expression = &NonblockPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 78:
		//line parser.go.y:437
		{
			yyVAL.expression = &ArrayExpression{Elems: yyS[yypt-1].expressions}
		}
	case 79:
		//line parser.go.y:445
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 80:
		//line parser.go.y:449
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 81:
		//line parser.go.y:453
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 82:
		//line parser.go.y:459
		{
			yyVAL.parameters = nil
		}
	case 83:
		//line parser.go.y:463
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 84:
		//line parser.go.y:469
		{
			yyVAL.parameters = []Parameter{yyS[yypt-0].parameter}
		}
	case 85:
		//line parser.go.y:473
		{
			yyVAL.parameters = []Parameter{yyS[yypt-1].parameter}
		}
	case 86:
		//line parser.go.y:477
		{
			yyVAL.parameters = append([]Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 87:
		//line parser.go.y:483
		{
			yyVAL.parameter = Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 88:
		//line parser.go.y:489
		{
			yyVAL.expressions = nil
		}
	case 89:
		//line parser.go.y:493
		{
			yyVAL.expressions = yyS[yypt-0].expressions
		}
	case 90:
		//line parser.go.y:499
		{
			yyVAL.expressions = []Expression{yyS[yypt-0].expression}
		}
	case 91:
		//line parser.go.y:503
		{
			yyVAL.expressions = []Expression{yyS[yypt-1].expression}
		}
	case 92:
		//line parser.go.y:507
		{
			yyVAL.expressions = append([]Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 93:
		//line parser.go.y:513
		{
			yyVAL.typetypes = []Type{yyS[yypt-0].typetype}
		}
	case 94:
		//line parser.go.y:517
		{
			yyVAL.typetypes = []Type{yyS[yypt-1].typetype}
		}
	case 95:
		//line parser.go.y:521
		{
			yyVAL.typetypes = append([]Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 96:
		//line parser.go.y:526
		{
			yyVAL.typetype = &NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 97:
		//line parser.go.y:530
		{
			yyVAL.typetype = &SetType{SetType: yyS[yypt-0].typetype}
		}
	case 98:
		//line parser.go.y:534
		{
			yyVAL.typetype = &ChannelType{IsUnstable: false, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 99:
		//line parser.go.y:538
		{
			yyVAL.typetype = &ChannelType{IsUnstable: false, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 100:
		//line parser.go.y:542
		{
			yyVAL.typetype = &ChannelType{IsUnstable: true, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 101:
		//line parser.go.y:546
		{
			yyVAL.typetype = &ChannelType{IsUnstable: true, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 102:
		//line parser.go.y:552
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-1].statements}}
		}
	case 103:
		//line parser.go.y:556
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-2].statements}}
		}
	case 104:
		//line parser.go.y:560
		{
			yyVAL.blocks = append([]BlockStatement{BlockStatement{Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
