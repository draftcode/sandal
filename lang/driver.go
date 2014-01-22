package lang

import (
	"github.com/draftcode/sandal/lang/parsing"
	"github.com/draftcode/sandal/lang/typecheck"
	"github.com/draftcode/sandal/lang/conversion"
	"github.com/draftcode/sandal/lang/conversion_deprecated"
)

func CompileFile(body string, deprecated bool) (error, string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)
	if err := typecheck.TypeCheck(defs); err != nil {
		return err, ""
	}
	if deprecated {
		return conversion_deprecated.ConvertASTToNuSMV(defs)
	} else {
		return conversion.ConvertASTToNuSMV(defs)
	}
}
