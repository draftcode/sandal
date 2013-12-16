package conversion

import (
	"testing"
)

const expectedTemplateResult = `
MODULE A(arg1, arg2)
  VAR
    var1 : boolean;
    var2 : 0..16;
  TRANS arg1 = FALSE;
  TRANS arg2 = TRUE;
  ASSIGN
    init(var1) := FALSE;
    next(var1) :=
      case
        TRUE : FALSE;
      esac;
  DEFINE
    var3 := TRUE;
`

func TestInstantiateTemplate(t *testing.T) {
	mod := tmplModule{
		Name:    "A",
		Args:    []string{"arg1", "arg2"},
		Vars:    []tmplVar{tmplVar{"var1", "boolean"}, tmplVar{"var2", "0..16"}},
		Trans:   []string{"arg1 = FALSE", "arg2 = TRUE"},
		Assigns: []tmplAssign{tmplAssign{"init(var1)", "FALSE"}, tmplAssign{"next(var1)", "case\n  TRUE : FALSE;\nesac"}},
		Defs:    []tmplAssign{tmplAssign{"var3", "TRUE"}},
	}

	if result := instantiateTemplate(mod); result != expectedTemplateResult {
		t.Errorf("Expect %s to be %s", result, expectedTemplateResult)
	}
}
