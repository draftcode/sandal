package parsing

import (
	"testing"
)

func testScanner(t *testing.T, src string, expectTok int, expectLit string) {
	s := new(Scanner)
	s.Init([]rune(src), dontInsertSemis)
	tok, lit, _ := s.Scan()
	if tok != expectTok {
		t.Errorf("Scanner{Src: \"%v\"}.Scan() = %v, _ want %v", src, tok, expectTok)
	}
	if lit != expectLit {
		t.Errorf("Scanner{Src: \"%v\"}.Scan() = _, %v want %v", src, lit, expectLit)
	}
	tok, lit, _ = s.Scan()
	if tok != EOF {
		t.Errorf("Scanner{Src: \"%v\"}.Scan() = %v, _ want %v", src, tok, EOF)
	}
}

func TestScanSymbols(t *testing.T) {
	testScanner(t, "{", '{', "{")
	testScanner(t, "}", '}', "}")
	testScanner(t, "(", '(', "(")
	testScanner(t, ")", ')', ")")
	testScanner(t, "[", '[', "[")
	testScanner(t, "]", ']', "]")
	testScanner(t, ",", ',', ",")
	testScanner(t, ":", ':', ":")
	testScanner(t, ";", ';', ";")
	testScanner(t, "@", '@', "@")

	testScanner(t, "+", ADD, "+")
	testScanner(t, "-", SUB, "-")
	testScanner(t, "*", MUL, "*")
	testScanner(t, "/", QUO, "/")
	testScanner(t, "%", REM, "%")

	testScanner(t, "&", AND, "&")
	testScanner(t, "|", OR, "|")
	testScanner(t, "^", XOR, "^")
	testScanner(t, "<<", SHL, "<<")
	testScanner(t, ">>", SHR, ">>")

	testScanner(t, "+=", ADD_ASSIGN, "+=")
	testScanner(t, "-=", SUB_ASSIGN, "-=")
	testScanner(t, "*=", MUL_ASSIGN, "*=")
	testScanner(t, "/=", QUO_ASSIGN, "/=")
	testScanner(t, "%=", REM_ASSIGN, "%=")

	testScanner(t, "&=", AND_ASSIGN, "&=")
	testScanner(t, "|=", OR_ASSIGN, "|=")
	testScanner(t, "^=", XOR_ASSIGN, "^=")
	testScanner(t, "<<=", SHL_ASSIGN, "<<=")
	testScanner(t, ">>=", SHR_ASSIGN, ">>=")

	testScanner(t, "&&", LAND, "&&")
	testScanner(t, "||", LOR, "||")

	testScanner(t, "==", EQL, "==")
	testScanner(t, "<", LSS, "<")
	testScanner(t, ">", GTR, ">")
	testScanner(t, "=", ASSIGN, "=")
	testScanner(t, "!", NOT, "!")

	testScanner(t, "!=", NEQ, "!=")
	testScanner(t, "<=", LEQ, "<=")
	testScanner(t, ">=", GEQ, ">=")
	testScanner(t, "->", THEN, "->")
}

func TestScanIdentifier(t *testing.T) {
	testScanner(t, "Sample", IDENTIFIER, "Sample")
}

func TestScanNumber(t *testing.T) {
	testScanner(t, "0", NUMBER, "0")
	testScanner(t, "1", NUMBER, "1")
	testScanner(t, "119", NUMBER, "119")
}

func TestScanKeyword(t *testing.T) {
	testScanner(t, "data", DATA, "data")
	testScanner(t, "const", CONST, "const")
	testScanner(t, "module", MODULE, "module")
	testScanner(t, "channel", CHANNEL, "channel")
	testScanner(t, "proc", PROC, "proc")
	testScanner(t, "var", VAR, "var")
	testScanner(t, "if", IF, "if")
	testScanner(t, "else", ELSE, "else")
	testScanner(t, "choice", CHOICE, "choice")
	testScanner(t, "recv", RECV, "recv")
	testScanner(t, "timeout_recv", TIMEOUT_RECV, "timeout_recv")
	testScanner(t, "nonblock_recv", NONBLOCK_RECV, "nonblock_recv")
	testScanner(t, "peek", PEEK, "peek")
	testScanner(t, "timeout_peek", TIMEOUT_PEEK, "timeout_peek")
	testScanner(t, "nonblock_peek", NONBLOCK_PEEK, "nonblock_peek")
	testScanner(t, "send", SEND, "send")
	testScanner(t, "for", FOR, "for")
	testScanner(t, "break", BREAK, "break")
	testScanner(t, "in", IN, "in")
	testScanner(t, "range", RANGE, "range")
	testScanner(t, "to", TO, "to")
	testScanner(t, "init", INIT, "init")
	testScanner(t, "goto", GOTO, "goto")
	testScanner(t, "skip", SKIP, "skip")
	testScanner(t, "true", TRUE, "true")
	testScanner(t, "false", FALSE, "false")
}

func TestScanComment(t *testing.T) {
	testScanner(t, "// This is a comment", COMMENT, "// This is a comment")
}

func TestEOF(t *testing.T) {
	s := new(Scanner)
	s.Init([]rune(""), 0)
	for i := 0; i < 5; i++ {
		tok, lit, pos := s.Scan()
		if tok != EOF || lit != "" {
			t.Errorf("Expect EOF and \"\" but got %v and %v.", tok, lit)
		}
		if pos.Line != 1 || pos.Column != 1 {
			t.Errorf("Expect Pos{1, 1} but got %v.", pos)
		}
	}
}

func TestInsertSemi(t *testing.T) {
	{
		s := new(Scanner)
		s.Init([]rune("token\n"), 0)
		s.Scan()
		tok, lit, _ := s.Scan()
		if tok != ';' || lit != "\n" {
			t.Errorf("Expect ; and \"\\n\" but got %v and %v.", tok, lit)
		}
	}
	{
		s := new(Scanner)
		s.Init([]rune("token"), 0)
		s.Scan()
		tok, lit, _ := s.Scan()
		if tok != ';' || lit != "\n" {
			t.Errorf("Expect ; and \"\\n\" but got %v and %v.", tok, lit)
		}
	}
	{
		s := new(Scanner)
		s.Init([]rune("token // comment"), 0)
		var tok int
		if tok, _, _ = s.Scan(); tok != IDENTIFIER {
			t.Errorf("Expect %v but got %v", IDENTIFIER, tok)
		}
		if tok, _, _ = s.Scan(); tok != ';' {
			t.Errorf("Expect %v but got %v", ';', tok)
		}
		if tok, _, _ = s.Scan(); tok != COMMENT {
			t.Errorf("Expect %v but got %v", COMMENT, tok)
		}
		if tok, _, _ = s.Scan(); tok != EOF {
			t.Errorf("Expect %v but got %v", EOF, tok)
		}
	}
}
