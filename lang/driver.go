package lang

import (
	"github.com/draftcode/sandal/lang/parsing"
	"github.com/draftcode/sandal/lang/typecheck"
	"github.com/draftcode/sandal/lang/conversion"
)

func CompileFile(body string) (error, string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)
	if err := typecheck.TypeCheck(defs); err != nil {
		return err, ""
	}
	return conversion.ConvertASTToNuSMV(defs)
}
