
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
const TRUE = 57404
const FALSE = 57405
const UNARY = 57406

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
	"TRUE",
	"FALSE",
	" {",
	" }",
	" (",
	" )",
	" [",
	" ]",
	" ,",
	" :",
	" ;",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:608


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

const yyNprod = 113
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 913

var yyAct = []int{

	107, 153, 199, 273, 117, 106, 66, 253, 287, 285,
	281, 268, 154, 267, 266, 265, 251, 238, 231, 215,
	211, 208, 206, 21, 169, 5, 7, 5, 6, 76,
	65, 272, 72, 45, 264, 109, 49, 47, 51, 44,
	145, 68, 69, 50, 39, 64, 70, 252, 40, 258,
	38, 257, 74, 256, 210, 176, 175, 99, 100, 101,
	186, 187, 188, 189, 190, 191, 192, 193, 194, 195,
	22, 114, 174, 173, 71, 185, 48, 203, 202, 201,
	116, 105, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	112, 104, 103, 119, 102, 121, 24, 120, 141, 142,
	143, 144, 26, 20, 184, 286, 167, 52, 53, 283,
	278, 57, 170, 172, 263, 259, 25, 255, 239, 228,
	213, 237, 212, 182, 23, 178, 171, 150, 205, 108,
	147, 148, 46, 119, 43, 121, 56, 120, 277, 200,
	181, 151, 177, 149, 115, 167, 183, 167, 196, 198,
	59, 61, 110, 60, 62, 78, 19, 18, 41, 261,
	37, 207, 33, 34, 197, 54, 55, 36, 32, 58,
	179, 63, 180, 30, 28, 167, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 226, 227, 216, 204, 42,
	17, 167, 232, 16, 229, 167, 236, 233, 234, 235,
	31, 15, 214, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 79, 8, 10, 9, 14, 11, 73,
	35, 167, 254, 91, 92, 93, 94, 95, 260, 29,
	96, 97, 98, 10, 75, 12, 11, 155, 53, 1,
	27, 57, 13, 118, 262, 4, 3, 2, 77, 0,
	0, 0, 270, 12, 0, 0, 0, 274, 0, 0,
	167, 275, 0, 0, 0, 152, 56, 279, 167, 280,
	24, 10, 0, 167, 284, 157, 158, 0, 159, 160,
	59, 61, 161, 60, 62, 162, 163, 164, 0, 0,
	25, 0, 165, 0, 166, 54, 55, 156, 23, 58,
	0, 63, 0, 0, 0, 168, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 0, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 0, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 271, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 250, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 249, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 248, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 247, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 246, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 245, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 244, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 243, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 242, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 241, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 240, 0, 96, 97, 98, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 209, 0, 96, 97, 98, 91, 92, 93, 94,
	95, 80, 0, 96, 97, 98, 0, 0, 0, 0,
	0, 0, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 0, 0, 83, 84, 85, 86, 0, 146,
	89, 90, 91, 92, 93, 94, 95, 0, 111, 96,
	97, 98, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 91, 92, 93, 94, 95, 0, 0, 96,
	97, 98, 140, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 0, 282,
	96, 97, 98, 91, 92, 93, 94, 95, 0, 0,
	96, 97, 98, 52, 53, 0, 0, 57, 0, 0,
	0, 52, 53, 0, 0, 57, 0, 0, 0, 0,
	269, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	230, 0, 56, 0, 52, 53, 0, 0, 57, 0,
	56, 0, 0, 0, 0, 0, 59, 61, 0, 60,
	62, 0, 0, 0, 59, 61, 0, 60, 62, 0,
	0, 54, 55, 56, 0, 58, 0, 63, 113, 54,
	55, 0, 0, 58, 0, 63, 67, 59, 61, 0,
	60, 62, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 54, 55, 0, 0, 58, 0, 63, 0,
	0, 0, 91, 92, 93, 94, 95, 0, 0, 96,
	97, 98, 0, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 276, 91, 0, 93, 94, 95, 0, 0,
	96, 97, 98, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 93, 94, 95, 0, 0,
	96, 97, 98,
}
var yyPact = []int{

	187, -1000, 187, -1000, -1000, -1000, -1000, -1000, 223, 207,
	199, 196, 103, -1000, 102, 47, 66, 46, 179, 174,
	173, 138, -1000, -19, -20, 128, 173, 79, -1000, -31,
	-38, 77, -33, 9, -1000, -34, 66, 770, 66, 66,
	747, -22, 7, -40, 179, 240, -43, 174, 101, 173,
	-1000, 559, -1000, -1000, -1000, -1000, 770, 770, 770, 38,
	36, 35, 15, 770, -1000, 74, -35, 98, 599, 66,
	739, 90, -1000, -1000, -1000, 14, -1000, -1000, 205, -1000,
	-1000, 770, 770, 770, 770, 770, 770, 770, 770, 770,
	770, 770, 770, 770, 770, 770, 770, 770, 770, -1000,
	-1000, 635, 770, 770, 770, 770, -29, 589, -1000, 66,
	66, 89, 72, 87, 206, 243, 770, 71, 205, -1000,
	-1000, -1000, 645, 645, -1000, -1000, -1000, -1000, 645, 645,
	-1000, -1000, 876, 846, 675, 675, 675, 675, 675, 675,
	-1000, 6, 5, -11, -12, -1000, 770, -1000, 70, 66,
	-1000, 66, 86, 68, 243, 43, 243, 170, 770, 85,
	13, 12, 11, 134, -50, 167, -51, 549, -1000, -1000,
	-13, -52, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 67,
	65, 66, -53, -1000, 243, 770, 770, 770, 770, 770,
	770, 770, 770, 770, 770, 770, 64, 66, 706, -54,
	243, 770, 770, 770, 243, 76, -1000, -55, -1000, -1000,
	-1000, -1000, -1000, -1000, 63, -1000, -1000, 519, 509, 479,
	469, 439, 429, 399, 389, 359, 349, 319, -56, -25,
	243, -1000, 62, -14, -16, -18, 60, 113, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 770, 59, -36, -57, -58, -59, -61,
	696, 770, 309, -41, 85, -1000, -1000, -1000, -1000, 243,
	815, -1000, -1000, 84, -1000, 55, 770, 243, -62, 665,
	54, -1000, 243, -63, 50, -1000, -64, -1000,
}
var yyPgo = []int{

	0, 249, 257, 256, 255, 24, 28, 26, 4, 253,
	250, 184, 239, 1, 12, 0, 210, 172, 173, 230,
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
	15, 15, 15, 15, 15, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 20, 20, 20, 21,
	21, 21, 22, 22, 22, 22, 22, 22, 22, 22,
	23, 23, 23,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 6, 9, 5, 0, 1,
	1, 2, 3, 3, 6, 0, 2, 3, 4, 4,
	6, 6, 10, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 3, 5, 5, 5, 5, 8,
	11, 2, 3, 2, 2, 1, 1, 1, 1, 1,
	1, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 4, 4, 4, 3, 1, 2, 3,
	0, 1, 1, 2, 3, 2, 1, 2, 3, 1,
	2, 3, 1, 3, 4, 5, 6, 7, 7, 8,
	3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 64, 64,
	66, -22, 4, 68, 40, 60, 66, -10, -11, -12,
	4, -16, 4, -17, -18, -19, 4, 32, 69, 64,
	68, 40, -17, 65, 70, 71, 65, 70, 67, 70,
	-22, -15, 4, 5, 62, 63, 33, 8, 66, 47,
	50, 48, 51, 68, -22, -21, -22, 69, -15, 64,
	68, 67, 72, -11, -22, 4, 72, -16, 64, -18,
	72, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 27, 28, 29, 30, 31, 34, 35, 36, -15,
	-15, -15, 66, 66, 66, 66, -20, -15, 65, 70,
	64, 69, -21, 69, -15, 64, 66, -8, -9, -5,
	-6, -7, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	67, -20, -20, -20, -20, 69, 70, -21, -21, 64,
	65, 64, 69, -13, -14, 4, 64, 42, 43, 45,
	46, 49, 52, 53, 54, 59, 61, -15, 72, -5,
	-20, 65, -8, 67, 67, 67, 67, -20, 65, -21,
	-21, 64, 65, -13, 71, 32, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, -13, 4, -15, -23,
	64, 66, 66, 66, 64, 4, 72, 4, 72, 72,
	67, 72, 65, 65, -21, 72, -14, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, 65, -22,
	64, 72, -13, -20, -20, -20, -13, 55, 72, 65,
	72, 72, 72, 72, 72, 72, 72, 72, 72, 72,
	72, 72, 72, 32, -13, 65, 67, 67, 67, 65,
	-15, 56, -15, 65, 70, 72, 72, 72, 72, 64,
	-15, 72, 72, 44, -23, -13, 57, 64, 65, -15,
	-13, 72, 64, 65, -13, 72, 65, 72,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	90, 0, 102, 0, 0, 0, 90, 0, 19, 20,
	0, 0, 87, 0, 91, 92, 0, 0, 0, 0,
	0, 0, 0, 0, 21, 0, 0, 88, 0, 93,
	95, 0, 57, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 0, 0, 103, 0, 99, 0, 0, 0,
	0, 0, 17, 22, 23, 102, 8, 89, 10, 94,
	15, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 61,
	62, 0, 0, 0, 0, 0, 0, 96, 104, 100,
	0, 0, 0, 0, 0, 25, 0, 0, 10, 12,
	13, 14, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	63, 0, 0, 0, 0, 86, 97, 101, 0, 0,
	105, 0, 0, 0, 25, 57, 25, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 55, 56,
	0, 0, 11, 82, 83, 84, 85, 98, 106, 0,
	0, 0, 0, 26, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	25, 0, 0, 0, 25, 0, 51, 0, 53, 54,
	24, 9, 107, 108, 0, 16, 27, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	25, 44, 0, 0, 0, 0, 0, 0, 52, 109,
	33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 28, 29, 0, 0, 110, 0, 0, 0, 0,
	0, 0, 0, 0, 111, 45, 46, 47, 48, 25,
	0, 30, 31, 0, 112, 0, 0, 25, 0, 0,
	0, 49, 25, 0, 0, 32, 0, 50,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 67, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 71, 72,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 68, 3, 69, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 65,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 73,
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
		//line parser.go.y:135
		{
			yyVAL.definitions = []data.Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:142
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
		//line parser.go.y:158
		{
			yyVAL.definition = data.DataDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:164
		{
			yyVAL.definition = data.ModuleDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:170
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:174
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
		//line parser.go.y:185
		{
			yyVAL.definition = data.ConstantDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:191
		{
			yyVAL.definition = data.ProcDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:197
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 18:
		//line parser.go.y:203
		{
			yyVAL.initvars = nil
		}
	case 19:
		//line parser.go.y:207
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 20:
		//line parser.go.y:213
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 21:
		//line parser.go.y:217
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 22:
		//line parser.go.y:221
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 23:
		//line parser.go.y:226
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-2].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 24:
		//line parser.go.y:230
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-5].tok.lit, ProcDefName: yyS[yypt-3].tok.lit, Args: yyS[yypt-1].expressions}
		}
	case 25:
		//line parser.go.y:236
		{
			yyVAL.statements = nil
		}
	case 26:
		//line parser.go.y:240
		{
			yyVAL.statements = append([]data.Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 27:
		//line parser.go.y:246
		{
			yyVAL.statement = data.LabelledStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 28:
		//line parser.go.y:250
		{
			yyVAL.statement = data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 29:
		//line parser.go.y:254
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 30:
		//line parser.go.y:258
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:262
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 32:
		//line parser.go.y:266
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 33:
		//line parser.go.y:270
		{
			yyVAL.statement = data.AssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:274
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:278
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:282
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:286
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:290
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 39:
		//line parser.go.y:294
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:298
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:302
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:306
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:310
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:314
		{
			yyVAL.statement = data.ChoiceStatement{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 45:
		//line parser.go.y:318
		{
			yyVAL.statement = data.RecvStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 46:
		//line parser.go.y:322
		{
			yyVAL.statement = data.PeekStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 47:
		//line parser.go.y:326
		{
			yyVAL.statement = data.SendStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 48:
		//line parser.go.y:330
		{
			yyVAL.statement = data.ForStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 49:
		//line parser.go.y:334
		{
			yyVAL.statement = data.ForInStatement{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 50:
		//line parser.go.y:338
		{
			yyVAL.statement = data.ForInRangeStatement{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 51:
		//line parser.go.y:342
		{
			yyVAL.statement = data.BreakStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 52:
		//line parser.go.y:346
		{
			yyVAL.statement = data.GotoStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 53:
		//line parser.go.y:350
		{
			yyVAL.statement = data.SkipStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 54:
		//line parser.go.y:354
		{
			yyVAL.statement = data.ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 55:
		//line parser.go.y:358
		{
			yyVAL.statement = data.NullStatement{Pos: yyS[yypt-0].tok.pos}
		}
	case 56:
		//line parser.go.y:362
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Statement)
		}
	case 57:
		//line parser.go.y:367
		{
			yyVAL.expression = data.IdentifierExpression{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 58:
		//line parser.go.y:371
		{
			yyVAL.expression = data.NumberExpression{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 59:
		//line parser.go.y:375
		{
			yyVAL.expression = data.TrueExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 60:
		//line parser.go.y:379
		{
			yyVAL.expression = data.FalseExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 61:
		//line parser.go.y:383
		{
			yyVAL.expression = data.NotExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 62:
		//line parser.go.y:387
		{
			yyVAL.expression = data.UnarySubExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:391
		{
			yyVAL.expression = data.ParenExpression{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 64:
		//line parser.go.y:395
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:399
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:403
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:407
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:411
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:415
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:419
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:423
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:427
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:431
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:435
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:439
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:443
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:447
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:451
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:455
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:459
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:463
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:467
		{
			yyVAL.expression = data.TimeoutRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 83:
		//line parser.go.y:471
		{
			yyVAL.expression = data.TimeoutPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 84:
		//line parser.go.y:475
		{
			yyVAL.expression = data.NonblockRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 85:
		//line parser.go.y:479
		{
			yyVAL.expression = data.NonblockPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 86:
		//line parser.go.y:483
		{
			yyVAL.expression = data.ArrayExpression{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 87:
		//line parser.go.y:491
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 88:
		//line parser.go.y:495
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 89:
		//line parser.go.y:499
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 90:
		//line parser.go.y:505
		{
			yyVAL.parameters = nil
		}
	case 91:
		//line parser.go.y:509
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 92:
		//line parser.go.y:515
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 93:
		//line parser.go.y:519
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 94:
		//line parser.go.y:523
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 95:
		//line parser.go.y:529
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 96:
		//line parser.go.y:535
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-0].expression}
		}
	case 97:
		//line parser.go.y:539
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-1].expression}
		}
	case 98:
		//line parser.go.y:543
		{
			yyVAL.expressions = append([]data.Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 99:
		//line parser.go.y:549
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 100:
		//line parser.go.y:553
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 101:
		//line parser.go.y:557
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 102:
		//line parser.go.y:562
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 103:
		//line parser.go.y:566
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 104:
		//line parser.go.y:570
		{
			yyVAL.typetype = data.HandshakeChannelType{IsUnstable: false, Elems: yyS[yypt-1].typetypes}
		}
	case 105:
		//line parser.go.y:574
		{
			yyVAL.typetype = data.HandshakeChannelType{IsUnstable: true, Elems: yyS[yypt-1].typetypes}
		}
	case 106:
		//line parser.go.y:578
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 107:
		//line parser.go.y:582
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: false, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 108:
		//line parser.go.y:586
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 109:
		//line parser.go.y:590
		{
			yyVAL.typetype = data.BufferedChannelType{IsUnstable: true, BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 110:
		//line parser.go.y:596
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-2].tok.pos, Statements: yyS[yypt-1].statements}}
		}
	case 111:
		//line parser.go.y:600
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}}
		}
	case 112:
		//line parser.go.y:604
		{
			yyVAL.blocks = append([]data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
