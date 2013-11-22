package parsing

import (
	"unicode"
)

const (
	EOF     = -1
	UNKNOWN = 0
)

var keywords = map[string]int{
	"data":          DATA,
	"const":         CONST,
	"module":        MODULE,
	"channel":       CHANNEL,
	"proc":          PROC,
	"var":           VAR,
	"if":            IF,
	"else":          ELSE,
	"choice":        CHOICE,
	"recv":          RECV,
	"timeout_recv":  TIMEOUT_RECV,
	"nonblock_recv": NONBLOCK_RECV,
	"peek":          PEEK,
	"timeout_peek":  TIMEOUT_PEEK,
	"nonblock_peek": NONBLOCK_PEEK,
	"send":          SEND,
	"for":           FOR,
	"break":         BREAK,
	"in":            IN,
	"range":         RANGE,
	"to":            TO,
	"init":          INIT,
	"goto":          GOTO,
	"unstable":      UNSTABLE,
	"skip":          SKIP,
}

// Position indicates a position of a first character of a token in a file.
type Position struct {
	Line   int
	Column int
}

type Mode uint

const (
	dontInsertSemis = 1 << iota
)

// Scanner is a lexer on Sandal language.
type Scanner struct {
	src       []rune
	offset    int
	lineHead  int
	line      int
	mode      Mode
	lastToken int
}

// Init initializes the scanner. mode is for internal use, expect it to be zero.
func (s *Scanner) Init(src []rune, mode Mode) {
	s.src = src
	s.mode = mode
}

// Scan returns token and literal and its position. Return UNKNOWN if unknown
// token received. Return EOF after all characters consumed.
func (s *Scanner) Scan() (tok int, lit string, pos Position) {
	if s.skipWhiteSpace() && (s.mode&dontInsertSemis) == 0 {
		switch s.lastToken {
		case IDENTIFIER, NUMBER, BREAK, ')', ']', '}':
			s.lastToken = int(';')
			tok = int(';')
			lit = "\n"
			return
		}
	}
	savedOffset := s.offset // Used in COMMENT
	pos = s.position()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit = s.scanIdentifier()
		if keyword, ok := keywords[lit]; ok {
			tok = keyword
		} else {
			tok = IDENTIFIER
		}
	case isDigit(ch):
		tok, lit = NUMBER, s.scanNumber()
	default:
		s.next()
		switch ch {
		case -1:
			tok = EOF
		case '{', '}', '(', ')', '[', ']', ',', ':', ';':
			tok = int(ch)
			lit = string(ch)
		case '+':
			switch s.peek() {
			case '=':
				s.next()
				tok = ADD_ASSIGN
				lit = "+="
			default:
				tok = ADD
				lit = "+"
			}
		case '-':
			switch s.peek() {
			case '=':
				s.next()
				tok = SUB_ASSIGN
				lit = "-="
			default:
				tok = SUB
				lit = "-"
			}
		case '*':
			switch s.peek() {
			case '=':
				s.next()
				tok = MUL_ASSIGN
				lit = "*="
			default:
				tok = MUL
				lit = "*"
			}
		case '/':
			switch s.peek() {
			case '=':
				s.next()
				tok = QUO_ASSIGN
				lit = "/="
			case '/':
				// Insert semicolon before comment.
				switch s.lastToken {
				case IDENTIFIER, NUMBER, BREAK, ')', ']', '}':
					s.offset = savedOffset
					tok = int(';')
					lit = "\n"
				default:
					s.next()
					tok = COMMENT
					lit = "//" + s.scanLineComment()
				}
			default:
				tok = QUO
				lit = "/"
			}
		case '%':
			switch s.peek() {
			case '=':
				s.next()
				tok = REM_ASSIGN
				lit = "%="
			default:
				tok = REM
				lit = "%"
			}
		case '&':
			switch s.peek() {
			case '=':
				s.next()
				tok = AND_ASSIGN
				lit = "&="
			case '&':
				s.next()
				tok = LAND
				lit = "&&"
			default:
				tok = AND
				lit = "&"
			}
		case '|':
			switch s.peek() {
			case '=':
				s.next()
				tok = OR_ASSIGN
				lit = "|="
			case '|':
				s.next()
				tok = LOR
				lit = "||"
			default:
				tok = OR
				lit = "|"
			}
		case '^':
			switch s.peek() {
			case '=':
				s.next()
				tok = XOR_ASSIGN
				lit = "^="
			default:
				tok = XOR
				lit = "^"
			}
		case '<':
			switch s.peek() {
			case '<':
				s.next()
				switch s.peek() {
				case '=':
					s.next()
					tok = SHL_ASSIGN
					lit = "<<="
				default:
					tok = SHL
					lit = "<<"
				}
			case '=':
				s.next()
				tok = LEQ
				lit = "<="
			default:
				tok = LSS
				lit = "<"
			}
		case '>':
			switch s.peek() {
			case '>':
				s.next()
				switch s.peek() {
				case '=':
					s.next()
					tok = SHR_ASSIGN
					lit = ">>="
				default:
					tok = SHR
					lit = ">>"
				}
			case '=':
				s.next()
				tok = GEQ
				lit = ">="
			default:
				tok = GTR
				lit = ">"
			}
		case '=':
			switch s.peek() {
			case '=':
				s.next()
				tok = EQL
				lit = "=="
			default:
				tok = ASSIGN
				lit = "="
			}
		case '!':
			switch s.peek() {
			case '=':
				s.next()
				tok = NEQ
				lit = "!="
			default:
				tok = NOT
				lit = "!"
			}
		}
	}
	s.lastToken = tok
	return
}

// ========================================

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) position() Position {
	return Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipWhiteSpace() (includeReturn bool) {
	for isWhiteSpace(s.peek()) {
		if s.peek() == '\n' {
			includeReturn = true
		}
		s.next()
	}
	if s.reachEOF() {
		includeReturn = true
	}
	return
}

func (s *Scanner) scanIdentifier() string {
	var ret []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

func (s *Scanner) scanNumber() string {
	var ret []rune
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

func (s *Scanner) scanLineComment() string {
	var ret []rune
	for s.peek() != '\n' && s.peek() != -1 {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}
