
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
const UNARY = 57405

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

//line parser.go.y:625


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

const yyNprod = 115
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1004

var yyAct = []int{

	103, 147, 193, 114, 112, 109, 64, 265, 110, 63,
	245, 102, 148, 279, 277, 273, 260, 259, 258, 257,
	243, 163, 5, 21, 5, 7, 6, 230, 223, 207,
	205, 202, 200, 72, 264, 68, 43, 49, 256, 22,
	66, 105, 48, 47, 62, 45, 42, 142, 38, 244,
	70, 37, 39, 250, 249, 95, 96, 97, 248, 204,
	172, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 71, 171, 170, 169, 24, 179, 67, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 116, 46, 197, 196,
	118, 117, 23, 195, 113, 101, 100, 24, 99, 161,
	138, 139, 140, 141, 178, 144, 145, 98, 25, 168,
	164, 20, 278, 275, 270, 166, 255, 251, 247, 220,
	206, 36, 176, 199, 23, 174, 167, 116, 104, 44,
	41, 118, 117, 269, 194, 146, 108, 106, 74, 161,
	177, 161, 190, 192, 19, 173, 175, 18, 10, 229,
	30, 11, 27, 201, 191, 33, 165, 32, 50, 51,
	35, 31, 55, 8, 10, 9, 29, 11, 12, 161,
	209, 210, 211, 212, 213, 214, 215, 216, 217, 218,
	219, 208, 198, 40, 12, 161, 224, 54, 221, 161,
	228, 17, 16, 15, 14, 69, 73, 225, 226, 227,
	231, 57, 59, 75, 58, 60, 1, 111, 34, 13,
	28, 26, 115, 161, 246, 52, 53, 4, 3, 56,
	252, 61, 65, 2, 0, 79, 80, 81, 82, 149,
	51, 85, 86, 55, 0, 0, 254, 0, 0, 0,
	0, 0, 0, 0, 262, 0, 0, 0, 0, 266,
	0, 0, 161, 267, 0, 0, 0, 0, 54, 271,
	161, 272, 0, 10, 0, 161, 276, 151, 152, 0,
	153, 154, 57, 59, 155, 58, 60, 156, 157, 158,
	0, 0, 0, 0, 159, 160, 52, 53, 150, 0,
	56, 0, 61, 0, 0, 0, 162, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 0, 0, 92, 93, 94, 87, 88, 89,
	90, 91, 0, 0, 92, 93, 94, 0, 0, 0,
	0, 0, 0, 0, 0, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 0, 0, 0, 0, 0,
	0, 263, 0, 0, 0, 87, 88, 89, 90, 91,
	0, 242, 92, 93, 94, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	0, 0, 92, 93, 94, 0, 0, 0, 0, 241,
	0, 0, 0, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 91, 0, 240,
	92, 93, 94, 87, 88, 89, 90, 91, 0, 0,
	92, 93, 94, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 78, 79, 80, 81, 82, 83, 84, 85,
	86, 0, 0, 0, 0, 0, 0, 239, 0, 0,
	0, 87, 88, 89, 90, 91, 0, 238, 92, 93,
	94, 77, 78, 79, 80, 81, 82, 83, 84, 85,
	86, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 87, 88, 89, 90, 91, 0, 0, 92, 93,
	94, 0, 0, 0, 0, 237, 0, 0, 0, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 86, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 0, 236, 92, 93, 94, 87,
	88, 89, 90, 91, 0, 0, 92, 93, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 0, 0, 0,
	0, 0, 0, 235, 0, 0, 0, 87, 88, 89,
	90, 91, 0, 234, 92, 93, 94, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 87, 88, 89,
	90, 91, 0, 0, 92, 93, 94, 0, 0, 0,
	0, 233, 0, 0, 0, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	0, 232, 92, 93, 94, 87, 88, 89, 90, 91,
	0, 0, 92, 93, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 0, 0, 0, 0, 0, 0, 203,
	0, 0, 0, 87, 88, 89, 90, 91, 0, 76,
	92, 93, 94, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 87, 88, 89, 90, 91, 0, 0,
	92, 93, 94, 0, 0, 143, 0, 0, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 91, 0, 107, 92, 93, 94, 87, 88,
	89, 90, 91, 0, 0, 92, 93, 94, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 137, 87, 88,
	89, 90, 91, 0, 274, 92, 93, 94, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 0, 50,
	51, 0, 0, 55, 0, 0, 0, 0, 87, 88,
	89, 90, 91, 0, 261, 92, 93, 94, 0, 0,
	50, 51, 0, 0, 55, 0, 0, 0, 54, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 57, 59, 222, 58, 60, 0, 0, 54,
	0, 253, 0, 0, 0, 0, 52, 53, 0, 0,
	56, 0, 61, 57, 59, 0, 58, 60, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 52, 53, 0,
	0, 56, 0, 61, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 86, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 91, 0,
	0, 92, 93, 94, 87, 0, 89, 90, 91, 0,
	0, 92, 93, 94, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 86, 268, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 89, 90, 91, 0,
	0, 92, 93, 94,
}
var yyPact = []int{

	136, -1000, 136, -1000, -1000, -1000, -1000, -1000, 200, 199,
	198, 197, 94, -1000, 91, 56, 35, 53, 172, 167,
	166, 99, -1000, -17, -15, 166, 76, -1000, -23, -34,
	75, -24, 31, -1000, -26, 35, 866, 35, 35, 164,
	11, -36, 172, 67, -38, 167, 85, 166, -1000, 658,
	-1000, -1000, -1000, -1000, 866, 866, 866, 52, 43, 41,
	40, 866, -1000, 74, -28, 84, 726, 83, -1000, -1000,
	-69, 39, -1000, -1000, 120, -1000, -1000, 866, 866, 866,
	866, 866, 866, 866, 866, 866, 866, 866, 866, 866,
	866, 866, 866, 866, 866, -1000, -1000, 761, 866, 866,
	866, 866, -21, 696, -1000, 35, 35, 82, 235, -1000,
	-1000, -69, 162, 866, 72, 120, -1000, -1000, -1000, 226,
	226, -1000, -1000, -1000, -1000, 226, 226, -1000, -1000, 967,
	937, 388, 388, 388, 388, 388, 388, -1000, 8, 7,
	6, -6, -1000, 866, -1000, 71, 35, 68, 235, 44,
	235, 160, 866, 81, 38, 34, 33, 129, -39, 159,
	-40, 648, -1000, -1000, -1000, -1000, -7, -41, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 66, -42, -1000, 235, 866,
	866, 866, 866, 866, 866, 866, 866, 866, 866, 866,
	65, 35, 831, -43, 235, 866, 866, 866, 235, 104,
	-1000, -44, -1000, -1000, -69, -1000, -1000, -1000, -1000, 610,
	580, 542, 532, 494, 464, 426, 416, 378, 348, 310,
	-51, -22, 235, -1000, 64, -8, -12, -13, 63, 845,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 866, 62, -31, -52, -53,
	-54, -55, 801, 866, 300, -37, 81, -1000, -1000, -1000,
	-1000, 235, 927, -1000, -1000, 80, -1000, 60, 866, 235,
	-56, 771, 59, -1000, 235, -57, 58, -1000, -58, -1000,
}
var yyPgo = []int{

	0, 216, 233, 228, 227, 21, 26, 25, 3, 222,
	221, 162, 220, 1, 12, 0, 160, 167, 165, 218,
	11, 9, 6, 5, 8, 217, 2,
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
	21, 21, 22, 22, 22, 22, 22, 23, 23, 24,
	24, 25, 26, 26, 26,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 6, 9,
	0, 2, 1, 1, 1, 6, 9, 5, 0, 1,
	1, 2, 3, 4, 7, 0, 2, 3, 4, 4,
	6, 6, 10, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 3, 5, 5, 5, 5, 8,
	11, 2, 3, 2, 2, 1, 1, 1, 1, 1,
	1, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 4, 4, 4, 3, 1, 2, 3,
	0, 1, 1, 2, 3, 2, 1, 2, 3, 1,
	2, 3, 1, 3, 4, 6, 7, 0, 1, 1,
	2, 2, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 37, 39,
	38, 41, 58, -1, 4, 4, 4, 4, 63, 63,
	65, -22, 4, 67, 40, 65, -10, -11, -12, 4,
	-16, 4, -17, -18, -19, 4, 32, 68, 63, 67,
	-17, 64, 69, 70, 64, 69, 66, 69, -22, -15,
	4, 5, 61, 62, 33, 8, 65, 47, 50, 48,
	51, 67, -22, -21, -22, 68, -15, 66, 71, -11,
	-22, 4, 71, -16, 63, -18, 71, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 27, 28, 29,
	30, 31, 34, 35, 36, -15, -15, -15, 65, 65,
	65, 65, -20, -15, 64, 69, 63, 68, 63, -23,
	-24, -25, 73, 65, -8, -9, -5, -6, -7, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, 66, -20, -20,
	-20, -20, 68, 69, -21, -21, 63, -13, -14, 4,
	63, 42, 43, 45, 46, 49, 52, 53, 54, 59,
	60, -15, 71, -5, -24, 4, -20, 64, -8, 66,
	66, 66, 66, -20, 64, -21, 64, -13, 70, 32,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	-13, 4, -15, -26, 63, 65, 65, 65, 63, 4,
	71, 4, 71, 71, 66, 71, 64, 71, -14, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	64, -22, 63, 71, -13, -20, -20, -20, -13, 55,
	71, -23, 71, 71, 71, 71, 71, 71, 71, 71,
	71, 71, 71, 71, 71, 32, -13, 64, 66, 66,
	66, 64, -15, 56, -15, 64, 69, 71, 71, 71,
	71, 63, -15, 71, 71, 44, -26, -13, 57, 63,
	64, -15, -13, 71, 63, 64, -13, 71, 64, 71,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 0, 0,
	0, 0, 0, 2, 0, 0, 0, 0, 18, 0,
	90, 0, 102, 0, 0, 90, 0, 19, 20, 0,
	0, 87, 0, 91, 92, 0, 0, 0, 0, 0,
	0, 0, 21, 0, 0, 88, 0, 93, 95, 0,
	57, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	0, 0, 103, 0, 99, 0, 0, 0, 17, 22,
	107, 102, 8, 89, 10, 94, 15, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 61, 62, 0, 0, 0,
	0, 0, 0, 96, 104, 100, 0, 0, 25, 23,
	108, 109, 0, 0, 0, 10, 12, 13, 14, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 63, 0, 0,
	0, 0, 86, 97, 101, 0, 0, 0, 25, 57,
	25, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 55, 56, 110, 111, 0, 0, 11, 82,
	83, 84, 85, 98, 105, 0, 0, 26, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 25, 0, 0, 0, 25, 0,
	51, 0, 53, 54, 107, 9, 106, 16, 27, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 25, 44, 0, 0, 0, 0, 0, 0,
	52, 24, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 28, 29, 0, 0, 112, 0, 0,
	0, 0, 0, 0, 0, 0, 113, 45, 46, 47,
	48, 25, 0, 30, 31, 0, 114, 0, 0, 25,
	0, 0, 0, 49, 25, 0, 0, 32, 0, 50,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	65, 66, 3, 3, 69, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 70, 71,
	3, 3, 3, 3, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 67, 3, 68, 3, 3, 3, 3, 3, 3,
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
	62, 72,
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
		//line parser.go.y:138
		{
			yyVAL.definitions = []data.Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:145
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
		//line parser.go.y:161
		{
			yyVAL.definition = data.DataDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 9:
		//line parser.go.y:167
		{
			yyVAL.definition = data.ModuleDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 10:
		//line parser.go.y:173
		{
			yyVAL.definitions = nil
		}
	case 11:
		//line parser.go.y:177
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
		//line parser.go.y:188
		{
			yyVAL.definition = data.ConstantDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 16:
		//line parser.go.y:194
		{
			yyVAL.definition = data.ProcDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 17:
		//line parser.go.y:200
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 18:
		//line parser.go.y:206
		{
			yyVAL.initvars = nil
		}
	case 19:
		//line parser.go.y:210
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 20:
		//line parser.go.y:216
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 21:
		//line parser.go.y:220
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 22:
		//line parser.go.y:224
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 23:
		//line parser.go.y:229
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-3].tok.lit, Type: yyS[yypt-1].typetype, Tags: yyS[yypt-0].tags}
		}
	case 24:
		//line parser.go.y:233
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-6].tok.lit, ProcDefName: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions, Tags: yyS[yypt-0].tags}
		}
	case 25:
		//line parser.go.y:239
		{
			yyVAL.statements = nil
		}
	case 26:
		//line parser.go.y:243
		{
			yyVAL.statements = append([]data.Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 27:
		//line parser.go.y:249
		{
			yyVAL.statement = data.LabelledStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 28:
		//line parser.go.y:253
		{
			yyVAL.statement = data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 29:
		//line parser.go.y:257
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 30:
		//line parser.go.y:261
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 31:
		//line parser.go.y:265
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 32:
		//line parser.go.y:269
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 33:
		//line parser.go.y:273
		{
			yyVAL.statement = data.AssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:277
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 35:
		//line parser.go.y:281
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 36:
		//line parser.go.y:285
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:289
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:293
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 39:
		//line parser.go.y:297
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:301
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:305
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:309
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:313
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:317
		{
			yyVAL.statement = data.ChoiceStatement{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 45:
		//line parser.go.y:321
		{
			yyVAL.statement = data.RecvStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 46:
		//line parser.go.y:325
		{
			yyVAL.statement = data.PeekStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 47:
		//line parser.go.y:329
		{
			yyVAL.statement = data.SendStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 48:
		//line parser.go.y:333
		{
			yyVAL.statement = data.ForStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 49:
		//line parser.go.y:337
		{
			yyVAL.statement = data.ForInStatement{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 50:
		//line parser.go.y:341
		{
			yyVAL.statement = data.ForInRangeStatement{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 51:
		//line parser.go.y:345
		{
			yyVAL.statement = data.BreakStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 52:
		//line parser.go.y:349
		{
			yyVAL.statement = data.GotoStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 53:
		//line parser.go.y:353
		{
			yyVAL.statement = data.SkipStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 54:
		//line parser.go.y:357
		{
			yyVAL.statement = data.ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 55:
		//line parser.go.y:361
		{
			yyVAL.statement = data.NullStatement{Pos: yyS[yypt-0].tok.pos}
		}
	case 56:
		//line parser.go.y:365
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Statement)
		}
	case 57:
		//line parser.go.y:370
		{
			yyVAL.expression = data.IdentifierExpression{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 58:
		//line parser.go.y:374
		{
			yyVAL.expression = data.NumberExpression{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 59:
		//line parser.go.y:378
		{
			yyVAL.expression = data.TrueExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 60:
		//line parser.go.y:382
		{
			yyVAL.expression = data.FalseExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 61:
		//line parser.go.y:386
		{
			yyVAL.expression = data.NotExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 62:
		//line parser.go.y:390
		{
			yyVAL.expression = data.UnarySubExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 63:
		//line parser.go.y:394
		{
			yyVAL.expression = data.ParenExpression{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 64:
		//line parser.go.y:398
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:402
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:406
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 67:
		//line parser.go.y:410
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:414
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:418
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:422
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:426
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:430
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:434
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:438
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:442
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:446
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:450
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:454
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:458
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:462
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:466
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:470
		{
			yyVAL.expression = data.TimeoutRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 83:
		//line parser.go.y:474
		{
			yyVAL.expression = data.TimeoutPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 84:
		//line parser.go.y:478
		{
			yyVAL.expression = data.NonblockRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 85:
		//line parser.go.y:482
		{
			yyVAL.expression = data.NonblockPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 86:
		//line parser.go.y:486
		{
			yyVAL.expression = data.ArrayExpression{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 87:
		//line parser.go.y:494
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 88:
		//line parser.go.y:498
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 89:
		//line parser.go.y:502
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 90:
		//line parser.go.y:508
		{
			yyVAL.parameters = nil
		}
	case 91:
		//line parser.go.y:512
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 92:
		//line parser.go.y:518
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 93:
		//line parser.go.y:522
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 94:
		//line parser.go.y:526
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 95:
		//line parser.go.y:532
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 96:
		//line parser.go.y:538
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-0].expression}
		}
	case 97:
		//line parser.go.y:542
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-1].expression}
		}
	case 98:
		//line parser.go.y:546
		{
			yyVAL.expressions = append([]data.Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 99:
		//line parser.go.y:552
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 100:
		//line parser.go.y:556
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 101:
		//line parser.go.y:560
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 102:
		//line parser.go.y:565
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 103:
		//line parser.go.y:569
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 104:
		//line parser.go.y:573
		{
			yyVAL.typetype = data.HandshakeChannelType{Elems: yyS[yypt-1].typetypes}
		}
	case 105:
		//line parser.go.y:577
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 106:
		//line parser.go.y:581
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 107:
		//line parser.go.y:587
		{
			yyVAL.tags = nil
		}
	case 108:
		//line parser.go.y:591
		{
			yyVAL.tags = yyS[yypt-0].tags
		}
	case 109:
		//line parser.go.y:597
		{
			yyVAL.tags = []string{yyS[yypt-0].tag}
		}
	case 110:
		//line parser.go.y:601
		{
			yyVAL.tags = append([]string{yyS[yypt-1].tag}, yyS[yypt-0].tags...)
		}
	case 111:
		//line parser.go.y:607
		{
			yyVAL.tag = yyS[yypt-0].tok.lit
		}
	case 112:
		//line parser.go.y:613
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-2].tok.pos, Statements: yyS[yypt-1].statements}}
		}
	case 113:
		//line parser.go.y:617
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}}
		}
	case 114:
		//line parser.go.y:621
		{
			yyVAL.blocks = append([]data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
