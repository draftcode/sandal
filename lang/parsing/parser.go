
//line parser.go.y:3
package parsing
import __yyfmt__ "fmt"
//line parser.go.y:3
		
import (
	"log"
	data "github.com/draftcode/sandal/lang/data"
)

type token struct {
	tok int
	lit string
	pos data.Pos
}

//line parser.go.y:17
type yySymType struct{
	yys int
	definitions []data.Definition
	definition  data.Definition
	statements  []data.Statement
	statement   data.Statement
	expressions []data.Expression
	expression  data.Expression
	parameters  []data.Parameter
	parameter   data.Parameter
	typetypes   []data.Type
	typetype    data.Type
	identifiers []string
	tags        []string
	tag         string
	blocks      []data.BlockStatement
	initvars    []data.InitVar
	initvar     data.InitVar
	ltlexpr     data.LtlExpression
	ltlatom     data.LtlAtomExpression

	tok         token
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
const SKIP = 57402
const TRUE = 57403
const FALSE = 57404
const LTL = 57405
const THEN = 57406
const IFF = 57407
const UNARY = 57408

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
	"SKIP",
	"TRUE",
	"FALSE",
	"LTL",
	"THEN",
	"IFF",
	" {",
	" }",
	" (",
	" )",
	" [",
	" ]",
	" ,",
	" :",
	" ;",
	" U",
	" V",
	" S",
	" T",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:732


type lexerWrapper struct {
	s           *Scanner
	definitions []data.Definition
	recentLit   string
	recentPos   data.Pos
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	for tok == COMMENT {
		tok, lit, pos = l.s.Scan()
	}
	if tok == EOF {
		return 0
	}
	lval.tok = token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *lexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []data.Definition {
	l := lexerWrapper{s: s}
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

const yyNprod = 139
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 975

var yyAct = []int{

	152, 197, 243, 164, 161, 158, 100, 159, 79, 99,
	315, 151, 198, 329, 327, 48, 64, 323, 45, 310,
	295, 213, 5, 7, 5, 24, 6, 66, 67, 68,
	69, 44, 199, 87, 34, 309, 91, 308, 307, 293,
	314, 230, 231, 232, 233, 234, 235, 236, 237, 238,
	239, 280, 85, 273, 63, 102, 229, 84, 257, 98,
	36, 90, 294, 255, 252, 106, 11, 61, 250, 65,
	201, 202, 163, 203, 204, 93, 95, 205, 94, 96,
	206, 207, 208, 121, 109, 104, 58, 209, 210, 88,
	89, 144, 145, 146, 200, 35, 92, 228, 97, 124,
	122, 306, 212, 154, 64, 83, 81, 37, 38, 39,
	40, 41, 42, 43, 120, 66, 67, 68, 69, 25,
	107, 57, 192, 52, 300, 299, 298, 169, 170, 171,
	172, 173, 174, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 166, 254, 168, 53, 222,
	167, 247, 54, 221, 220, 27, 27, 63, 211, 188,
	189, 190, 191, 63, 194, 195, 219, 103, 214, 218,
	61, 62, 65, 82, 216, 246, 245, 162, 65, 150,
	149, 148, 147, 51, 328, 26, 26, 166, 28, 168,
	23, 325, 167, 320, 305, 301, 297, 270, 256, 211,
	227, 211, 240, 242, 226, 223, 225, 64, 224, 249,
	60, 217, 153, 64, 108, 80, 319, 59, 66, 67,
	68, 69, 56, 244, 66, 67, 68, 69, 196, 211,
	259, 260, 261, 262, 263, 264, 265, 266, 267, 268,
	269, 258, 157, 155, 63, 211, 274, 123, 271, 211,
	278, 9, 11, 10, 22, 12, 21, 275, 276, 277,
	281, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 248, 13, 211, 296, 20, 279, 14, 251, 241,
	302, 136, 137, 138, 139, 140, 215, 47, 141, 142,
	143, 50, 46, 44, 64, 11, 304, 30, 12, 32,
	19, 18, 17, 16, 312, 66, 67, 68, 69, 316,
	160, 1, 211, 317, 15, 13, 55, 49, 31, 321,
	211, 322, 29, 165, 8, 211, 326, 4, 313, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 3, 105, 141, 142, 143, 136,
	137, 138, 139, 140, 2, 0, 141, 142, 143, 0,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 0, 292, 141, 142, 143,
	136, 137, 138, 139, 140, 0, 291, 141, 142, 143,
	0, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 140, 0, 290, 141, 142,
	143, 136, 137, 138, 139, 140, 0, 289, 141, 142,
	143, 0, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 0, 288, 141,
	142, 143, 136, 137, 138, 139, 140, 0, 287, 141,
	142, 143, 0, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 0, 286,
	141, 142, 143, 136, 137, 138, 139, 140, 0, 285,
	141, 142, 143, 0, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 0,
	284, 141, 142, 143, 136, 137, 138, 139, 140, 0,
	283, 141, 142, 143, 0, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	0, 282, 141, 142, 143, 136, 137, 138, 139, 140,
	0, 253, 141, 142, 143, 0, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 128, 129, 130, 131,
	0, 0, 134, 135, 63, 0, 136, 137, 138, 139,
	140, 0, 125, 141, 142, 143, 0, 61, 62, 65,
	193, 0, 0, 0, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 136, 137, 138, 139, 140, 0,
	156, 141, 142, 143, 64, 0, 0, 0, 0, 119,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 0,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 0,
	0, 0, 0, 0, 0, 0, 187, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 110, 111, 112, 113,
	114, 115, 116, 117, 118, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	0, 0, 141, 142, 143, 136, 137, 138, 139, 140,
	0, 0, 141, 142, 143, 0, 0, 0, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 86, 87,
	0, 0, 91, 0, 324, 0, 0, 0, 136, 137,
	138, 139, 140, 0, 311, 141, 142, 143, 0, 0,
	0, 0, 0, 0, 86, 87, 0, 90, 91, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 0,
	0, 93, 95, 0, 94, 96, 0, 272, 86, 87,
	0, 0, 91, 90, 0, 88, 89, 0, 0, 0,
	0, 0, 92, 0, 97, 101, 0, 93, 95, 0,
	94, 96, 0, 0, 0, 0, 303, 90, 0, 0,
	0, 88, 89, 0, 0, 0, 0, 0, 92, 0,
	97, 93, 95, 0, 94, 96, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 88, 89, 0, 0, 0,
	0, 0, 92, 0, 97, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	0, 0, 141, 142, 143, 136, 0, 138, 139, 140,
	0, 0, 141, 142, 143, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 318, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 138, 139, 140,
	0, 0, 141, 142, 143,
}
var yyPact = []int{

	214, -1000, 214, -1000, -1000, -1000, -1000, -1000, -1000, 299,
	298, 297, 296, 209, 190, -1000, 188, 122, 115, 120,
	295, 27, 288, 287, 151, -1000, 52, 82, 287, 155,
	-1000, 49, 13, 143, -1000, 27, 27, 27, 27, 27,
	27, 27, 27, 27, -79, 148, 34, 104, -1000, 33,
	115, 834, 115, 115, 784, 98, 11, 295, 116, 147,
	10, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	620, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 289,
	9, 288, 181, 287, -1000, 568, -1000, -1000, -1000, -1000,
	834, 834, 834, 114, 113, 112, 111, 834, -1000, 145,
	31, 177, 609, 176, -1000, -1000, -84, 109, -2, -1000,
	149, 40, -48, -1000, 230, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 257, -1000, -1000, 834, 834, 834, 834,
	834, 834, 834, 834, 834, 834, 834, 834, 834, 834,
	834, 834, 834, 834, -1000, -1000, 647, 834, 834, 834,
	834, 51, 578, -1000, 115, 115, 162, 28, -1000, -1000,
	-84, 282, 834, -1000, 144, 257, -1000, -1000, -1000, 617,
	617, -1000, -1000, -1000, -1000, 617, 617, -1000, -1000, 938,
	908, 812, 812, 812, 812, 812, 812, -1000, 97, 85,
	84, 80, -1000, 834, -1000, 141, 115, 137, 28, 24,
	28, 275, 834, 157, 108, 107, 83, 205, -6, 274,
	-10, 537, -1000, -1000, -1000, -1000, 77, -11, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 131, -16, -1000, 28, 834,
	834, 834, 834, 834, 834, 834, 834, 834, 834, 834,
	130, 115, 771, -21, 28, 834, 834, 834, 28, 221,
	-1000, -23, -1000, -1000, -84, -1000, -1000, -1000, -1000, 527,
	496, 486, 455, 445, 414, 404, 373, 363, 332, 322,
	-35, -12, 28, -1000, 129, 57, 56, 55, 128, 810,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 834, 127, 29, -36, -37,
	-39, -55, 738, 834, 254, -34, 157, -1000, -1000, -1000,
	-1000, 28, 898, -1000, -1000, 150, -1000, 126, 834, 28,
	-57, 728, 124, -1000, 28, -60, 117, -1000, -61, -1000,
}
var yyPgo = []int{

	0, 311, 364, 354, 327, 21, 26, 23, 324, 3,
	323, 322, 297, 318, 1, 12, 0, 18, 287, 15,
	317, 11, 9, 6, 5, 7, 310, 2, 664, 34,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 3,
	4, 9, 9, 10, 10, 10, 5, 6, 7, 8,
	8, 11, 11, 12, 12, 12, 13, 13, 14, 14,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 29,
	29, 17, 17, 17, 18, 18, 19, 19, 19, 20,
	21, 21, 21, 22, 22, 22, 23, 23, 23, 23,
	23, 24, 24, 25, 25, 26, 27, 27, 27,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 6,
	9, 0, 2, 1, 1, 1, 6, 9, 5, 6,
	5, 0, 1, 1, 2, 3, 4, 7, 0, 2,
	3, 4, 4, 6, 6, 10, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 3, 5, 5,
	5, 5, 8, 11, 2, 3, 2, 2, 1, 1,
	1, 1, 1, 1, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 2, 2, 2, 2, 1,
	3, 1, 2, 3, 0, 1, 1, 2, 3, 2,
	1, 2, 3, 1, 2, 3, 1, 3, 4, 6,
	7, 0, 1, 1, 2, 2, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, 37,
	39, 38, 41, 58, 63, -1, 4, 4, 4, 4,
	66, 66, 66, 68, -23, 4, 70, 40, 68, -11,
	-12, -13, 4, -28, -29, 68, 33, 80, 81, 82,
	83, 84, 85, 86, 4, -17, 4, -18, -19, -20,
	4, 32, 71, 66, 70, -18, 67, 72, 73, 74,
	67, 27, 28, 14, 64, 29, 75, 76, 77, 78,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, 87,
	67, 72, 69, 72, -23, -16, 4, 5, 61, 62,
	33, 8, 68, 47, 50, 48, 51, 70, -23, -22,
	-23, 71, -16, 69, 74, -12, -23, 4, 67, 74,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, 69,
	-29, 74, -17, 66, -19, 74, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 27, 28, 29, 30,
	31, 34, 35, 36, -16, -16, -16, 68, 68, 68,
	68, -21, -16, 67, 72, 66, 71, 66, -24, -25,
	-26, 88, 68, 74, -9, -10, -5, -6, -7, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 69, -21, -21,
	-21, -21, 71, 72, -22, -22, 66, -14, -15, 4,
	66, 42, 43, 45, 46, 49, 52, 53, 54, 59,
	60, -16, 74, -5, -25, 4, -21, 67, -9, 69,
	69, 69, 69, -21, 67, -22, 67, -14, 73, 32,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	-14, 4, -16, -27, 66, 68, 68, 68, 66, 4,
	74, 4, 74, 74, 69, 74, 67, 74, -15, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	67, -23, 66, 74, -14, -21, -21, -21, -14, 55,
	74, -24, 74, 74, 74, 74, 74, 74, 74, 74,
	74, 74, 74, 74, 74, 32, -14, 67, 69, 69,
	69, 67, -16, 56, -16, 67, 72, 74, 74, 74,
	74, 66, -16, 74, 74, 44, -27, -14, 57, 66,
	67, -16, -14, 74, 66, 67, -14, 74, 67, 74,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 8, 0,
	0, 0, 0, 0, 0, 2, 0, 0, 0, 0,
	21, 0, 0, 114, 0, 126, 0, 0, 114, 0,
	22, 23, 0, 0, 90, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 0, 111, 0, 115, 116,
	0, 0, 0, 0, 0, 0, 0, 24, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 101, 102, 103, 104, 105, 106, 107, 108, 0,
	0, 112, 0, 117, 119, 0, 60, 61, 62, 63,
	0, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	123, 0, 0, 0, 18, 25, 131, 126, 0, 20,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 91,
	110, 9, 113, 11, 118, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 0, 0, 0, 0,
	0, 0, 120, 128, 124, 0, 0, 28, 26, 132,
	133, 0, 0, 19, 0, 11, 13, 14, 15, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 66, 0, 0,
	0, 0, 89, 121, 125, 0, 0, 0, 28, 60,
	28, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 59, 134, 135, 0, 0, 12, 85,
	86, 87, 88, 122, 129, 0, 0, 29, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 0, 0, 0, 28, 0,
	54, 0, 56, 57, 131, 10, 130, 17, 30, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 28, 47, 0, 0, 0, 0, 0, 0,
	55, 27, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 31, 32, 0, 0, 136, 0, 0,
	0, 0, 0, 0, 0, 0, 137, 48, 49, 50,
	51, 28, 0, 33, 34, 0, 138, 0, 0, 28,
	0, 0, 0, 52, 28, 0, 0, 35, 0, 53,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	68, 69, 3, 3, 72, 3, 87, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 73, 74,
	3, 3, 3, 3, 88, 3, 3, 3, 3, 3,
	82, 81, 85, 3, 3, 3, 3, 3, 3, 86,
	3, 3, 3, 77, 78, 75, 76, 3, 80, 83,
	84, 70, 3, 71, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 66, 3, 67,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 79,
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
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
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
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
		//line parser.go.y:146
		{
			yyVAL.definitions = []data.Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:153
		{
			yyVAL.definitions = append([]data.Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
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
		yyVAL.definition = yyS[yypt-0].definition
	case 9:
		//line parser.go.y:170
		{
			yyVAL.definition = data.DataDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 10:
		//line parser.go.y:176
		{
			yyVAL.definition = data.ModuleDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 11:
		//line parser.go.y:182
		{
			yyVAL.definitions = nil
		}
	case 12:
		//line parser.go.y:186
		{
			yyVAL.definitions = append([]data.Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
		}
	case 13:
		yyVAL.definition = yyS[yypt-0].definition
	case 14:
		yyVAL.definition = yyS[yypt-0].definition
	case 15:
		yyVAL.definition = yyS[yypt-0].definition
	case 16:
		//line parser.go.y:197
		{
			yyVAL.definition = data.ConstantDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 17:
		//line parser.go.y:203
		{
			yyVAL.definition = data.ProcDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 18:
		//line parser.go.y:209
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 19:
		//line parser.go.y:215
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-3].ltlexpr}
		}
	case 20:
		//line parser.go.y:219
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-2].ltlexpr}
		}
	case 21:
		//line parser.go.y:225
		{
			yyVAL.initvars = nil
		}
	case 22:
		//line parser.go.y:229
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 23:
		//line parser.go.y:235
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 24:
		//line parser.go.y:239
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 25:
		//line parser.go.y:243
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 26:
		//line parser.go.y:248
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-3].tok.lit, Type: yyS[yypt-1].typetype, Tags: yyS[yypt-0].tags}
		}
	case 27:
		//line parser.go.y:252
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-6].tok.lit, ProcDefName: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions, Tags: yyS[yypt-0].tags}
		}
	case 28:
		//line parser.go.y:258
		{
			yyVAL.statements = nil
		}
	case 29:
		//line parser.go.y:262
		{
			yyVAL.statements = append([]data.Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 30:
		//line parser.go.y:268
		{
			yyVAL.statement = data.LabelledStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 31:
		//line parser.go.y:272
		{
			yyVAL.statement = data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 32:
		//line parser.go.y:276
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 33:
		//line parser.go.y:280
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:284
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 35:
		//line parser.go.y:288
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 36:
		//line parser.go.y:292
		{
			yyVAL.statement = data.AssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:296
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:300
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 39:
		//line parser.go.y:304
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:308
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:312
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:316
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:320
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:324
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 45:
		//line parser.go.y:328
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 46:
		//line parser.go.y:332
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 47:
		//line parser.go.y:336
		{
			yyVAL.statement = data.ChoiceStatement{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 48:
		//line parser.go.y:340
		{
			yyVAL.statement = data.RecvStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 49:
		//line parser.go.y:344
		{
			yyVAL.statement = data.PeekStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 50:
		//line parser.go.y:348
		{
			yyVAL.statement = data.SendStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 51:
		//line parser.go.y:352
		{
			yyVAL.statement = data.ForStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 52:
		//line parser.go.y:356
		{
			yyVAL.statement = data.ForInStatement{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 53:
		//line parser.go.y:360
		{
			yyVAL.statement = data.ForInRangeStatement{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 54:
		//line parser.go.y:364
		{
			yyVAL.statement = data.BreakStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 55:
		//line parser.go.y:368
		{
			yyVAL.statement = data.GotoStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 56:
		//line parser.go.y:372
		{
			yyVAL.statement = data.SkipStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 57:
		//line parser.go.y:376
		{
			yyVAL.statement = data.ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 58:
		//line parser.go.y:380
		{
			yyVAL.statement = data.NullStatement{Pos: yyS[yypt-0].tok.pos}
		}
	case 59:
		//line parser.go.y:384
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Statement)
		}
	case 60:
		//line parser.go.y:389
		{
			yyVAL.expression = data.IdentifierExpression{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 61:
		//line parser.go.y:393
		{
			yyVAL.expression = data.NumberExpression{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 62:
		//line parser.go.y:397
		{
			yyVAL.expression = data.TrueExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 63:
		//line parser.go.y:401
		{
			yyVAL.expression = data.FalseExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 64:
		//line parser.go.y:405
		{
			yyVAL.expression = data.NotExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:409
		{
			yyVAL.expression = data.UnarySubExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:413
		{
			yyVAL.expression = data.ParenExpression{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 67:
		//line parser.go.y:417
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:421
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:425
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:429
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:433
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:437
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:441
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:445
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:449
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:453
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:457
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:461
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:465
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:469
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:473
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:477
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 83:
		//line parser.go.y:481
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 84:
		//line parser.go.y:485
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 85:
		//line parser.go.y:489
		{
			yyVAL.expression = data.TimeoutRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 86:
		//line parser.go.y:493
		{
			yyVAL.expression = data.TimeoutPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 87:
		//line parser.go.y:497
		{
			yyVAL.expression = data.NonblockRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 88:
		//line parser.go.y:501
		{
			yyVAL.expression = data.NonblockPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 89:
		//line parser.go.y:505
		{
			yyVAL.expression = data.ArrayExpression{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 90:
		//line parser.go.y:512
		{
			yyVAL.ltlexpr = yyS[yypt-0].ltlatom
		}
	case 91:
		//line parser.go.y:516
		{
			yyVAL.ltlexpr = data.ParenLtlExpression{SubExpr: yyS[yypt-1].ltlexpr}
		}
	case 92:
		//line parser.go.y:520
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "&", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 93:
		//line parser.go.y:524
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "|", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 94:
		//line parser.go.y:528
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "^", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 95:
		//line parser.go.y:532
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "->", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 96:
		//line parser.go.y:536
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "=", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 97:
		//line parser.go.y:540
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "U", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 98:
		//line parser.go.y:544
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "V", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 99:
		//line parser.go.y:548
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "S", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 100:
		//line parser.go.y:552
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "T", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 101:
		//line parser.go.y:556
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "!", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 102:
		//line parser.go.y:560
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "X", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 103:
		//line parser.go.y:564
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "G", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 104:
		//line parser.go.y:568
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "F", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 105:
		//line parser.go.y:572
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "Y", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 106:
		//line parser.go.y:576
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "Z", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 107:
		//line parser.go.y:580
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "H", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 108:
		//line parser.go.y:584
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "O", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 109:
		//line parser.go.y:589
		{
			yyVAL.ltlatom = data.LtlAtomExpression{Names: []string{yyS[yypt-0].tok.lit}}
		}
	case 110:
		//line parser.go.y:593
		{
			yyVAL.ltlatom = data.LtlAtomExpression{Names: append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].ltlatom.Names...)}
		}
	case 111:
		//line parser.go.y:601
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 112:
		//line parser.go.y:605
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 113:
		//line parser.go.y:609
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 114:
		//line parser.go.y:615
		{
			yyVAL.parameters = nil
		}
	case 115:
		//line parser.go.y:619
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 116:
		//line parser.go.y:625
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 117:
		//line parser.go.y:629
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 118:
		//line parser.go.y:633
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 119:
		//line parser.go.y:639
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 120:
		//line parser.go.y:645
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-0].expression}
		}
	case 121:
		//line parser.go.y:649
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-1].expression}
		}
	case 122:
		//line parser.go.y:653
		{
			yyVAL.expressions = append([]data.Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 123:
		//line parser.go.y:659
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 124:
		//line parser.go.y:663
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 125:
		//line parser.go.y:667
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 126:
		//line parser.go.y:672
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 127:
		//line parser.go.y:676
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 128:
		//line parser.go.y:680
		{
			yyVAL.typetype = data.HandshakeChannelType{Elems: yyS[yypt-1].typetypes}
		}
	case 129:
		//line parser.go.y:684
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 130:
		//line parser.go.y:688
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 131:
		//line parser.go.y:694
		{
			yyVAL.tags = nil
		}
	case 132:
		//line parser.go.y:698
		{
			yyVAL.tags = yyS[yypt-0].tags
		}
	case 133:
		//line parser.go.y:704
		{
			yyVAL.tags = []string{yyS[yypt-0].tag}
		}
	case 134:
		//line parser.go.y:708
		{
			yyVAL.tags = append([]string{yyS[yypt-1].tag}, yyS[yypt-0].tags...)
		}
	case 135:
		//line parser.go.y:714
		{
			yyVAL.tag = yyS[yypt-0].tok.lit
		}
	case 136:
		//line parser.go.y:720
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-2].tok.pos, Statements: yyS[yypt-1].statements}}
		}
	case 137:
		//line parser.go.y:724
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}}
		}
	case 138:
		//line parser.go.y:728
		{
			yyVAL.blocks = append([]data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
