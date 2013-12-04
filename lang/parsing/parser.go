
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
	pos Position
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
	blocks      []data.BlockStatement
	initvars    []data.InitVar
	initvar     data.InitVar

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

//line parser.go.y:597


type lexerWrapper struct {
	s           *Scanner
	definitions []data.Definition
	recentLit   string
	recentPos   Position
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

const yyNprod = 111
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 972

var yyAct = []int{

	105, 151, 197, 143, 115, 104, 64, 67, 39, 38,
	45, 262, 152, 256, 68, 40, 107, 49, 47, 44,
	255, 254, 208, 21, 167, 5, 7, 5, 6, 174,
	63, 173, 172, 171, 69, 48, 201, 200, 51, 251,
	199, 66, 114, 50, 103, 62, 102, 101, 100, 271,
	26, 20, 72, 285, 283, 97, 98, 99, 184, 185,
	186, 187, 188, 189, 190, 191, 192, 193, 22, 112,
	270, 279, 250, 183, 266, 265, 73, 264, 263, 249,
	120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 110, 236,
	229, 117, 213, 119, 24, 118, 139, 140, 141, 142,
	182, 209, 24, 206, 165, 52, 53, 204, 74, 55,
	168, 170, 70, 284, 25, 281, 276, 261, 28, 257,
	253, 237, 25, 226, 23, 211, 235, 210, 145, 146,
	180, 117, 23, 119, 54, 118, 176, 169, 148, 203,
	175, 106, 46, 165, 181, 165, 194, 196, 57, 59,
	43, 58, 60, 275, 198, 179, 149, 147, 113, 108,
	76, 19, 18, 71, 34, 41, 10, 56, 177, 11,
	178, 61, 111, 165, 215, 216, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 214, 12, 37, 205, 165,
	230, 33, 227, 165, 234, 231, 232, 233, 202, 31,
	212, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 8, 10, 9, 77, 11, 195, 36, 42, 165,
	252, 89, 90, 91, 92, 93, 258, 32, 94, 95,
	96, 30, 12, 17, 16, 153, 53, 15, 14, 55,
	1, 35, 260, 13, 29, 27, 116, 75, 4, 3,
	268, 2, 0, 0, 0, 272, 0, 0, 165, 273,
	0, 0, 144, 0, 54, 277, 165, 278, 0, 10,
	0, 165, 282, 155, 156, 0, 157, 158, 57, 59,
	159, 58, 60, 160, 161, 162, 0, 0, 0, 0,
	163, 0, 164, 0, 154, 0, 166, 56, 0, 0,
	0, 61, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 0, 0, 94,
	95, 96, 89, 90, 91, 92, 93, 0, 0, 94,
	95, 96, 81, 82, 83, 84, 0, 0, 87, 88,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	0, 0, 0, 0, 0, 0, 150, 0, 0, 0,
	89, 90, 91, 92, 93, 0, 109, 94, 95, 96,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 91, 92, 93, 0, 0, 94, 95, 96,
	138, 0, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 0, 269, 94,
	95, 96, 89, 90, 91, 92, 93, 0, 0, 94,
	95, 96, 0, 0, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 0, 0, 0, 0, 0, 0,
	248, 0, 0, 0, 89, 90, 91, 92, 93, 0,
	247, 94, 95, 96, 0, 0, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 91, 92,
	93, 0, 246, 94, 95, 96, 89, 90, 91, 92,
	93, 0, 0, 94, 95, 96, 0, 0, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 0, 0,
	0, 0, 0, 0, 245, 0, 0, 0, 89, 90,
	91, 92, 93, 0, 244, 94, 95, 96, 0, 0,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 91, 92, 93, 0, 243, 94, 95, 96,
	89, 90, 91, 92, 93, 0, 0, 94, 95, 96,
	0, 0, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 0, 0, 0, 0, 0, 0, 242, 0,
	0, 0, 89, 90, 91, 92, 93, 0, 241, 94,
	95, 96, 0, 0, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 0,
	240, 94, 95, 96, 89, 90, 91, 92, 93, 0,
	0, 94, 95, 96, 0, 0, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 0, 0, 0, 0,
	0, 0, 239, 0, 0, 0, 89, 90, 91, 92,
	93, 0, 238, 94, 95, 96, 0, 0, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 0, 207, 94, 95, 96, 89, 90,
	91, 92, 93, 0, 0, 94, 95, 96, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 0, 0,
	0, 0, 0, 0, 0, 0, 78, 0, 89, 90,
	91, 92, 93, 0, 280, 94, 95, 96, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 0, 52,
	53, 0, 0, 55, 0, 0, 0, 0, 89, 90,
	91, 92, 93, 0, 267, 94, 95, 96, 0, 0,
	0, 0, 0, 0, 0, 52, 53, 0, 54, 55,
	0, 0, 0, 0, 0, 52, 53, 0, 0, 55,
	0, 0, 57, 59, 228, 58, 60, 0, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 0, 0, 0,
	0, 56, 0, 0, 54, 61, 65, 0, 57, 59,
	0, 58, 60, 0, 0, 0, 0, 259, 57, 59,
	0, 58, 60, 0, 0, 0, 0, 56, 0, 0,
	0, 61, 0, 0, 0, 0, 0, 56, 0, 0,
	0, 61, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 0, 0, 94,
	95, 96, 89, 0, 91, 92, 93, 0, 0, 94,
	95, 96, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 274, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 91, 92, 93, 0, 0, 94,
	95, 96,
}
var yyPact = []int{

	184, -1000, 184, -1000, -1000, -1000, -1000, -1000, 244, 243,
	240, 239, 109, -1000, 108, -15, 64, -16, 237, 233,
	223, 165, -1000, -62, -55, 135, 223, 96, -1000, -49,
	-59, 88, -50, -32, -1000, -51, 64, 831, 64, 64,
	795, -56, -33, 57, 237, 72, 53, 233, 107, 223,
	-1000, 711, -1000, -1000, 831, 831, 831, -18, -19, -20,
	-22, 831, -1000, 87, -52, 106, 315, 64, 111, 105,
	-1000, -1000, -1000, -24, -1000, -1000, 138, -1000, -1000, 831,
	831, 831, 831, 831, 831, 831, 831, 831, 831, 831,
	831, 831, 831, 831, 831, 831, 831, -1000, -1000, 353,
	831, 831, 831, 831, -68, 204, -1000, 64, 64, 104,
	84, 103, 305, 241, 831, 83, 138, -1000, -1000, -1000,
	343, 343, -1000, -1000, -1000, -1000, 343, 343, -1000, -1000,
	935, 905, 393, 393, 393, 393, 393, 393, -1000, -34,
	-35, -36, -38, -1000, 831, -1000, 82, 64, -1000, 64,
	102, 76, 241, 41, 241, 222, 831, 101, -26, -29,
	-30, 145, 52, 194, 48, 679, -1000, -1000, -45, 46,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 73, 71, 64,
	37, -1000, 241, 831, 831, 831, 831, 831, 831, 831,
	831, 831, 831, 831, 69, 64, 781, 35, 241, 831,
	831, 831, 241, 81, -1000, 34, -1000, -1000, -1000, -1000,
	-1000, -1000, 67, -1000, -1000, 647, 637, 605, 573, 563,
	531, 499, 489, 457, 425, 415, 14, 7, 241, -1000,
	66, -46, -47, -54, 65, 821, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 831, 63, -57, 13, 12, 10, 9, 751, 831,
	383, 5, 101, -1000, -1000, -1000, -1000, 241, 895, -1000,
	-1000, 100, -1000, 62, 831, 241, 6, 721, 61, -1000,
	241, -11, 59, -1000, -12, -1000,
}
var yyPgo = []int{

	0, 250, 261, 259, 258, 24, 28, 26, 4, 256,
	255, 128, 254, 1, 12, 0, 209, 201, 174, 251,
	5, 30, 6, 2,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 3, 4,
	8, 8, 9, 9, 9, 5, 6, 7, 10, 10,
	11, 11, 11, 12, 12, 13, 13, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 16, 16, 16, 17, 17,
	18, 18, 18, 19, 20, 20, 20, 21, 21, 21,
	22, 22, 22, 22, 22, 22, 22, 22, 23, 23,
	23,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 6, 9, 5, 0, 1,
	1, 2, 3, 3, 6, 0, 2, 3, 4, 4,
	6, 6, 10, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 3, 5, 5, 5, 5, 8,
	11, 2, 3, 2, 2, 1, 1, 1, 1, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 4, 4, 3, 1, 2, 3, 0, 1,
	1, 2, 3, 2, 1, 2, 3, 1, 2, 3,
	1, 3, 4, 5, 6, 7, 7, 8, 3, 4,
	5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 63, 63,
	66, -22, 4, 70, 40, 60, 66, -10, -11, -12,
	4, -16, 4, -17, -18, -19, 4, 32, 71, 63,
	70, 40, -17, 64, 68, 69, 64, 68, 67, 68,
	-22, -15, 4, 5, 33, 8, 66, 47, 50, 48,
	51, 70, -22, -21, -22, 71, -15, 63, 70, 67,
	65, -11, -22, 4, 65, -16, 63, -18, 65, 7,
	8, 9, 10, 11, 12, 13, 14, 15, 16, 27,
	28, 29, 30, 31, 34, 35, 36, -15, -15, -15,
	66, 66, 66, 66, -20, -15, 64, 68, 63, 71,
	-21, 71, -15, 63, 66, -8, -9, -5, -6, -7,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, 67, -20,
	-20, -20, -20, 71, 68, -21, -21, 63, 64, 63,
	71, -13, -14, 4, 63, 42, 43, 45, 46, 49,
	52, 53, 54, 59, 61, -15, 65, -5, -20, 64,
	-8, 67, 67, 67, 67, -20, 64, -21, -21, 63,
	64, -13, 69, 32, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 26, -13, 4, -15, -23, 63, 66,
	66, 66, 63, 4, 65, 4, 65, 65, 67, 65,
	64, 64, -21, 65, -14, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, 64, -22, 63, 65,
	-13, -20, -20, -20, -13, 55, 65, 64, 65, 65,
	65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 32, -13, 64, 67, 67, 67, 64, -15, 56,
	-15, 64, 68, 65, 65, 65, 65, 63, -15, 65,
	65, 44, -23, -13, 57, 63, 64, -15, -13, 65,
	63, 64, -13, 65, 64, 65,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	88, 0, 100, 0, 0, 0, 88, 0, 19, 20,
	0, 0, 85, 0, 89, 90, 0, 0, 0, 0,
	0, 0, 0, 0, 21, 0, 0, 86, 0, 91,
	93, 0, 57, 58, 0, 0, 0, 0, 0, 0,
	0, 0, 101, 0, 97, 0, 0, 0, 0, 0,
	17, 22, 23, 100, 8, 87, 10, 92, 15, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 60, 0,
	0, 0, 0, 0, 0, 94, 102, 98, 0, 0,
	0, 0, 0, 25, 0, 0, 10, 12, 13, 14,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 61, 0,
	0, 0, 0, 84, 95, 99, 0, 0, 103, 0,
	0, 0, 25, 57, 25, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 55, 56, 0, 0,
	11, 80, 81, 82, 83, 96, 104, 0, 0, 0,
	0, 26, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 25, 0,
	0, 0, 25, 0, 51, 0, 53, 54, 24, 9,
	105, 106, 0, 16, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 25, 44,
	0, 0, 0, 0, 0, 0, 52, 107, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 28,
	29, 0, 0, 108, 0, 0, 0, 0, 0, 0,
	0, 0, 109, 45, 46, 47, 48, 25, 0, 30,
	31, 0, 110, 0, 0, 25, 0, 0, 0, 49,
	25, 0, 0, 32, 0, 50,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 67, 3, 3, 68, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 69, 65,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 70, 3, 71, 3, 3, 3, 3, 3, 3,
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
		//line parser.go.y:132
		{
			yyVAL.definitions = []data.Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:139
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
		//line parser.go.y:155
		{
			yyVAL.definition = data.DataDefinition{Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:161
		{
			yyVAL.definition = data.ModuleDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:167
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:171
		{
			yyVAL.definitions = append([]data.Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
		}
	case 12:
		yyVAL.definition = yyS[yypt-0].definition
	case 13:
		yyVAL.definition = yyS[yypt-0].definition
	case 14:
		yyVAL.definition = yyS[yypt-0].definition
	case 15:
		//line parser.go.y:182
		{
			yyVAL.definition = data.ConstantDefinition{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:188
		{
			yyVAL.definition = data.ProcDefinition{Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:194
		{
			yyVAL.definition = data.InitBlock{Vars: yyS[yypt-2].initvars}
		}
	case 18:
		//line parser.go.y:200
		{
			yyVAL.initvars = nil
		}
	case 19:
		//line parser.go.y:204
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 20:
		//line parser.go.y:210
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 21:
		//line parser.go.y:214
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 22:
		//line parser.go.y:218
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 23:
		//line parser.go.y:223
		{
			yyVAL.initvar = data.ChannelVar{Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 24:
		//line parser.go.y:227
		{
			yyVAL.initvar = data.InstanceVar{Name: yyS[yypt-5].tok.lit, ProcDefName: yyS[yypt-3].tok.lit, Args: yyS[yypt-1].expressions}
		}
	case 25:
		//line parser.go.y:233
		{
			yyVAL.statements = nil
		}
	case 26:
		//line parser.go.y:237
		{
			yyVAL.statements = append([]data.Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 27:
		//line parser.go.y:243
		{
			yyVAL.statement = data.LabelledStatement{Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 28:
		//line parser.go.y:247
		{
			yyVAL.statement = data.BlockStatement{Statements: yyS[yypt-2].statements}
		}
	case 29:
		//line parser.go.y:251
		{
			yyVAL.statement = data.VarDeclStatement{Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 30:
		//line parser.go.y:255
		{
			yyVAL.statement = data.VarDeclStatement{Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:259
		{
			yyVAL.statement = data.IfStatement{Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 32:
		//line parser.go.y:263
		{
			yyVAL.statement = data.IfStatement{Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 33:
		//line parser.go.y:267
		{
			yyVAL.statement = data.AssignmentStatement{Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:271
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:275
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:279
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:283
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:287
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 39:
		//line parser.go.y:291
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:295
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:299
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:303
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:307
		{
			yyVAL.statement = data.OpAssignmentStatement{Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:311
		{
			yyVAL.statement = data.ChoiceStatement{Blocks: yyS[yypt-1].blocks}
		}
	case 45:
		//line parser.go.y:315
		{
			yyVAL.statement = data.RecvStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 46:
		//line parser.go.y:319
		{
			yyVAL.statement = data.PeekStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 47:
		//line parser.go.y:323
		{
			yyVAL.statement = data.SendStatement{Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 48:
		//line parser.go.y:327
		{
			yyVAL.statement = data.ForStatement{Statements: yyS[yypt-2].statements}
		}
	case 49:
		//line parser.go.y:331
		{
			yyVAL.statement = data.ForInStatement{Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 50:
		//line parser.go.y:335
		{
			yyVAL.statement = data.ForInRangeStatement{Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 51:
		//line parser.go.y:339
		{
			yyVAL.statement = data.BreakStatement{}
		}
	case 52:
		//line parser.go.y:343
		{
			yyVAL.statement = data.GotoStatement{Label: yyS[yypt-1].tok.lit}
		}
	case 53:
		//line parser.go.y:347
		{
			yyVAL.statement = data.SkipStatement{}
		}
	case 54:
		//line parser.go.y:351
		{
			yyVAL.statement = data.ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 55:
		//line parser.go.y:355
		{
			yyVAL.statement = data.NullStatement{}
		}
	case 56:
		//line parser.go.y:359
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Statement)
		}
	case 57:
		//line parser.go.y:364
		{
			yyVAL.expression = data.IdentifierExpression{Name: yyS[yypt-0].tok.lit}
		}
	case 58:
		//line parser.go.y:368
		{
			yyVAL.expression = data.NumberExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 59:
		//line parser.go.y:372
		{
			yyVAL.expression = data.NotExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 60:
		//line parser.go.y:376
		{
			yyVAL.expression = data.UnarySubExpression{SubExpr: yyS[yypt-0].expression}
		}
	case 61:
		//line parser.go.y:380
		{
			yyVAL.expression = data.ParenExpression{SubExpr: yyS[yypt-1].expression}
		}
	case 62:
		//line parser.go.y:384
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:388
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 64:
		//line parser.go.y:392
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:396
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:400
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:404
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:408
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:412
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:416
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:420
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:424
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:428
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:432
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:436
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:440
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:444
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:448
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:452
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:456
		{
			yyVAL.expression = data.TimeoutRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 81:
		//line parser.go.y:460
		{
			yyVAL.expression = data.TimeoutPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 82:
		//line parser.go.y:464
		{
			yyVAL.expression = data.NonblockRecvExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 83:
		//line parser.go.y:468
		{
			yyVAL.expression = data.NonblockPeekExpression{Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 84:
		//line parser.go.y:472
		{
			yyVAL.expression = data.ArrayExpression{Elems: yyS[yypt-1].expressions}
		}
	case 85:
		//line parser.go.y:480
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 86:
		//line parser.go.y:484
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 87:
		//line parser.go.y:488
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 88:
		//line parser.go.y:494
		{
			yyVAL.parameters = nil
		}
	case 89:
		//line parser.go.y:498
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 90:
		//line parser.go.y:504
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 91:
		//line parser.go.y:508
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 92:
		//line parser.go.y:512
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 93:
		//line parser.go.y:518
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 94:
		//line parser.go.y:524
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-0].expression}
		}
	case 95:
		//line parser.go.y:528
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-1].expression}
		}
	case 96:
		//line parser.go.y:532
		{
			yyVAL.expressions = append([]data.Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 97:
		//line parser.go.y:538
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 98:
		//line parser.go.y:542
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 99:
		//line parser.go.y:546
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 100:
		//line parser.go.y:551
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 101:
		//line parser.go.y:555
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 102:
		//line parser.go.y:559
		{
			yyVAL.typetype = data.HandshakeChannelType{IsUnstable: false, Elems: yyS[yypt-1].typetypes}
		}
	case 103:
		//line parser.go.y:563
		{
			yyVAL.typetype = data.HandshakeChannelType{IsUnstable: true, Elems: yyS[yypt-1].typetypes}
		}
	case 104:
		//line parser.go.y:567
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 105:
		//line parser.go.y:571
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: false, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 106:
		//line parser.go.y:575
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 107:
		//line parser.go.y:579
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: true, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 108:
		//line parser.go.y:585
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Statements: yyS[yypt-1].statements}}
		}
	case 109:
		//line parser.go.y:589
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Statements: yyS[yypt-2].statements}}
		}
	case 110:
		//line parser.go.y:593
		{
			yyVAL.blocks = append([]data.BlockStatement{data.BlockStatement{Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
