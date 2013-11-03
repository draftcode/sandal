
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

//line parser.go.y:563


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

const yyLast = 911

var yyAct = []int{

	41, 84, 129, 43, 5, 27, 5, 227, 56, 7,
	242, 6, 192, 53, 125, 123, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 184, 60, 22, 225,
	224, 68, 133, 82, 66, 120, 80, 61, 134, 223,
	222, 258, 218, 62, 217, 216, 113, 114, 115, 210,
	135, 124, 213, 121, 119, 118, 117, 116, 28, 88,
	127, 87, 257, 132, 24, 79, 130, 67, 86, 138,
	139, 140, 141, 142, 143, 144, 145, 146, 147, 148,
	121, 26, 20, 21, 25, 212, 278, 121, 121, 121,
	276, 154, 272, 23, 266, 158, 161, 162, 163, 164,
	165, 166, 167, 168, 169, 170, 171, 172, 173, 174,
	175, 176, 177, 178, 262, 149, 246, 121, 121, 121,
	121, 245, 155, 156, 157, 126, 137, 128, 29, 44,
	244, 243, 46, 239, 189, 197, 195, 187, 211, 186,
	160, 153, 136, 93, 91, 277, 274, 270, 151, 268,
	264, 263, 180, 181, 182, 183, 255, 45, 214, 251,
	220, 249, 10, 241, 235, 219, 31, 32, 215, 33,
	34, 48, 50, 35, 49, 51, 36, 37, 38, 191,
	83, 44, 150, 39, 46, 40, 121, 30, 90, 42,
	47, 122, 229, 52, 65, 232, 233, 267, 231, 85,
	230, 254, 236, 234, 238, 198, 193, 188, 19, 45,
	18, 159, 83, 44, 240, 10, 46, 59, 11, 63,
	58, 226, 248, 48, 50, 55, 49, 51, 54, 8,
	10, 9, 229, 11, 92, 12, 250, 252, 231, 253,
	230, 45, 47, 81, 259, 52, 196, 89, 17, 16,
	12, 15, 64, 260, 14, 48, 50, 265, 49, 51,
	57, 228, 269, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 271, 47, 4, 3, 52, 131, 275,
	2, 0, 0, 105, 106, 107, 108, 109, 0, 0,
	110, 111, 112, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 0, 0,
	110, 111, 112, 1, 0, 0, 13, 185, 0, 0,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	105, 106, 107, 108, 109, 0, 237, 110, 111, 112,
	105, 106, 107, 108, 109, 0, 0, 110, 111, 112,
	0, 0, 0, 0, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 97, 98, 99, 100, 0, 0,
	103, 104, 0, 194, 105, 106, 107, 108, 109, 0,
	179, 110, 111, 112, 0, 0, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 107, 108,
	109, 0, 256, 110, 111, 112, 105, 106, 107, 108,
	109, 0, 0, 110, 111, 112, 0, 0, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 0, 0,
	0, 0, 0, 0, 209, 0, 0, 0, 105, 106,
	107, 108, 109, 0, 208, 110, 111, 112, 0, 0,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	105, 106, 107, 108, 109, 0, 207, 110, 111, 112,
	105, 106, 107, 108, 109, 0, 0, 110, 111, 112,
	0, 0, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 0, 0, 0, 0, 0, 0, 206, 0,
	0, 0, 105, 106, 107, 108, 109, 0, 205, 110,
	111, 112, 0, 0, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 0,
	204, 110, 111, 112, 105, 106, 107, 108, 109, 0,
	0, 110, 111, 112, 0, 0, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 0, 0, 0, 0,
	0, 0, 203, 0, 0, 0, 105, 106, 107, 108,
	109, 0, 202, 110, 111, 112, 0, 0, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 0, 201, 110, 111, 112, 105, 106,
	107, 108, 109, 0, 0, 110, 111, 112, 0, 0,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	0, 0, 0, 0, 0, 0, 200, 0, 0, 0,
	105, 106, 107, 108, 109, 0, 199, 110, 111, 112,
	0, 0, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 105, 106, 107, 108, 109, 0, 190, 110,
	111, 112, 105, 106, 107, 108, 109, 0, 0, 110,
	111, 112, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	94, 0, 105, 106, 107, 108, 109, 0, 273, 110,
	111, 112, 95, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 0, 83, 44, 0, 0, 46, 0, 0,
	0, 0, 105, 106, 107, 108, 109, 0, 247, 110,
	111, 112, 0, 83, 44, 0, 0, 46, 0, 0,
	0, 0, 45, 0, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 0, 0, 48, 50, 152, 49,
	51, 0, 45, 0, 105, 221, 107, 108, 109, 0,
	0, 110, 111, 112, 0, 47, 48, 50, 52, 49,
	51, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 0, 0, 0, 0, 47, 0, 0, 52, 0,
	0, 105, 106, 107, 108, 109, 0, 0, 110, 111,
	112, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 261, 0, 107, 108, 109, 0, 0, 110, 111,
	112,
}
var yyPact = []int{

	192, -1000, 192, -1000, -1000, -1000, -1000, -1000, 250, 247,
	245, 244, 147, -1000, 145, 16, 24, 15, 124, 224,
	216, 185, -1000, -43, -26, 179, 216, 130, 124, -1,
	124, 239, 799, 136, 2, -5, -7, 184, 79, 230,
	78, 695, -1000, -1000, -1000, 799, 799, 799, -9, -10,
	-11, -12, 799, 127, -56, -16, -1000, -57, 24, 799,
	24, 24, 208, -31, -17, 77, -1000, 124, 799, 799,
	799, 799, 799, 799, 799, 799, 799, 799, 799, 799,
	118, 24, 765, -1000, 76, 124, 799, 799, 799, 124,
	156, -1000, 75, -1000, -1000, 799, 799, 799, 799, 799,
	799, 799, 799, 799, 799, 799, 799, 799, 799, 799,
	799, 799, 799, -1000, -1000, 333, 799, 799, 799, 799,
	-44, 256, 74, 224, 144, 216, -1000, 663, -1000, 115,
	-59, 143, 323, 24, 176, 142, -1000, -1000, 631, 621,
	589, 557, 547, 515, 483, 473, 441, 409, 399, -18,
	73, 20, 124, -1000, 104, -22, -23, -25, 101, 779,
	-1000, 375, 375, -1000, -1000, -1000, -1000, 375, 375, -1000,
	-1000, 874, 807, 296, 296, 296, 296, 296, 296, -1000,
	-27, -28, -37, -38, -1000, 799, -1000, -1000, 177, -1000,
	-1000, -1000, 24, 24, 140, 100, 139, 286, 124, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	68, -1000, -1000, 799, 99, -61, 66, 65, 56, 51,
	735, 799, -1000, -1000, -1000, -1000, -1000, 97, 177, -1000,
	-1000, -1000, -1000, 95, 24, -1000, 24, 138, 92, -1000,
	367, -3, 136, -1000, -1000, -1000, -1000, 124, 844, 49,
	-1000, -1000, 87, 86, 24, 29, -1000, -1000, 134, -1000,
	85, 799, -1000, -1000, -1000, 83, -1000, 124, 27, 705,
	-1000, 82, -1000, 124, 25, 81, -1000, 21, -1000,
}
var yyPgo = []int{

	0, 323, 280, 276, 275, 3, 11, 9, 7, 261,
	5, 58, 0, 13, 225, 8, 260, 35, 2, 66,
	1,
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
	17, 18, 18, 18, 19, 19, 19, 19, 19, 19,
	19, 19, 20, 20, 20,
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
	2, 3, 0, 1, 1, 2, 3, 2, 1, 2,
	3, 1, 2, 3, 1, 3, 4, 5, 6, 7,
	7, 8, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 63, 63,
	66, -19, 4, 69, 40, 60, 66, -10, -11, 4,
	63, 42, 43, 45, 46, 49, 52, 53, 54, 59,
	61, -12, 65, -5, 5, 33, 8, 66, 47, 50,
	48, 51, 69, -13, 4, -14, -15, -16, 4, 32,
	70, 63, 69, 40, -14, 64, -10, 68, 32, 17,
	18, 19, 20, 21, 22, 23, 24, 25, 26, 66,
	-10, 4, -12, 4, -20, 63, 66, 66, 66, 63,
	4, 65, 4, 65, 65, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 27, 28, 29, 30, 31,
	34, 35, 36, -12, -12, -12, 66, 66, 66, 66,
	-17, -12, 64, 71, 67, 71, -19, -12, -19, -18,
	-19, 70, -12, 63, 69, 67, 65, -11, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -17,
	64, -19, 63, 65, -10, -17, -17, -17, -10, 55,
	65, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 67,
	-17, -17, -17, -17, 70, 71, 65, -13, 63, -15,
	65, 64, 71, 63, 70, -18, 70, -12, 63, 65,
	65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
	67, 65, 65, 32, -10, 64, 67, 67, 67, 64,
	-12, 56, 67, 67, 67, 67, -17, -8, -9, -5,
	-6, -7, -18, -18, 63, 64, 63, 70, -10, 65,
	-12, 64, 71, 65, 65, 65, 65, 63, -12, 64,
	-8, 64, -18, -18, 63, 64, 65, 65, 44, -20,
	-10, 57, 65, 64, 64, -18, 65, 63, 64, -12,
	64, -10, 65, 63, 64, -10, 65, 64, 65,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	82, 0, 94, 0, 0, 0, 82, 0, 18, 51,
	18, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 50, 52, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 79, 0, 83, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 51, 0, 18, 0, 0, 0, 18,
	0, 44, 0, 47, 48, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 53, 54, 0, 0, 0, 0, 0,
	0, 88, 0, 80, 0, 85, 87, 0, 95, 0,
	91, 0, 0, 0, 0, 0, 17, 20, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 37, 0, 0, 0, 0, 0, 0,
	45, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 55,
	0, 0, 0, 0, 78, 89, 8, 81, 10, 86,
	15, 96, 92, 0, 0, 0, 0, 0, 18, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	0, 21, 22, 0, 0, 102, 0, 0, 0, 0,
	0, 0, 74, 75, 76, 77, 90, 0, 10, 12,
	13, 14, 93, 0, 0, 97, 0, 0, 0, 46,
	0, 0, 103, 38, 39, 40, 41, 18, 0, 0,
	11, 98, 0, 0, 0, 0, 23, 24, 0, 104,
	0, 0, 9, 99, 100, 0, 16, 18, 0, 0,
	101, 0, 42, 18, 0, 0, 25, 0, 43,
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
			yyVAL.expressions = []Expression{yyS[yypt-0].expression}
		}
	case 89:
		//line parser.go.y:494
		{
			yyVAL.expressions = []Expression{yyS[yypt-1].expression}
		}
	case 90:
		//line parser.go.y:498
		{
			yyVAL.expressions = append([]Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 91:
		//line parser.go.y:504
		{
			yyVAL.typetypes = []Type{yyS[yypt-0].typetype}
		}
	case 92:
		//line parser.go.y:508
		{
			yyVAL.typetypes = []Type{yyS[yypt-1].typetype}
		}
	case 93:
		//line parser.go.y:512
		{
			yyVAL.typetypes = append([]Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 94:
		//line parser.go.y:517
		{
			yyVAL.typetype = NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 95:
		//line parser.go.y:521
		{
			yyVAL.typetype = ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 96:
		//line parser.go.y:525
		{
			yyVAL.typetype = HandshakeChannelType{IsUnstable: false, Elems: yyS[yypt-1].typetypes}
		}
	case 97:
		//line parser.go.y:529
		{
			yyVAL.typetype = HandshakeChannelType{IsUnstable: true, Elems: yyS[yypt-1].typetypes}
		}
	case 98:
		//line parser.go.y:533
		{
			yyVAL.typetype = BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 99:
		//line parser.go.y:537
		{
			yyVAL.typetype = BufferedChannelType{IsUnstable: false, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 100:
		//line parser.go.y:541
		{
			yyVAL.typetype = BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 101:
		//line parser.go.y:545
		{
			yyVAL.typetype = BufferedChannelType{IsUnstable: true, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 102:
		//line parser.go.y:551
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-1].statements}}
		}
	case 103:
		//line parser.go.y:555
		{
			yyVAL.blocks = []BlockStatement{BlockStatement{Statements: yyS[yypt-2].statements}}
		}
	case 104:
		//line parser.go.y:559
		{
			yyVAL.blocks = append([]BlockStatement{BlockStatement{Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
