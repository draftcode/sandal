package conversion

import (
	"strings"
	"testing"
)

func TestConvertAbstractModuleToTemplate(t *testing.T) {
	mod := absNuSMVModule{
		Name:      "ProcA",
		Args:      []string{"ch0"},
		Vars:      []absVar{},
		InitState: "state0",
		Trans: map[absState][]absTransition{
			"state0": []absTransition{
				{
					Condition: "!ch0.filled",
					Actions: map[absState][]absAssign{
						"state1": []absAssign{
							{"ch0.next_filled", "TRUE"},
							{"ch0.next_received", "FALSE"},
							{"ch0.next_value", "TRUE"},
						},
					},
				},
			},
			"state1": []absTransition{
				{
					Condition: "ch0.received",
					Actions: map[absState][]absAssign{
						"state2": []absAssign{
							{"ch0.next_filled", "FALSE"},
							{"ch0.next_received", "FALSE"},
						},
					},
				},
			},
		},
		Defaults: map[string]string{
			"ch0.next_filled":   "ch0.filled",
			"ch0.next_received": "ch0.received",
			"ch0.next_value":    "ch0.value",
		},
		Defs: []absAssign{},
	}
	expected := tmplNuSMVModule{
		Name: "ProcA",
		Args: []string{"running_pid", "pid", "ch0"},
		Vars: []tmplVar{
			{"state", "{state0, state1, state2}"},
		},
		Assigns: []tmplAssign{
			{"init(state)", "state0"},
			{"next(state)", strings.Join([]string{
				"case",
				"  running_pid = pid & state = state0 & !ch0.filled : state1;",
				"  running_pid = pid & state = state1 & ch0.received : state2;",
				"  TRUE : state;",
				"esac",
			}, "\n")},
			{"ch0.next_filled", strings.Join([]string{
				"case",
				"  running_pid = pid & state = state0 & !ch0.filled : TRUE;",
				"  running_pid = pid & state = state1 & ch0.received : FALSE;",
				"  TRUE : ch0.filled;",
				"esac",
			}, "\n")},
			{"ch0.next_received", strings.Join([]string{
				"case",
				"  running_pid = pid & state = state0 & !ch0.filled : FALSE;",
				"  running_pid = pid & state = state1 & ch0.received : FALSE;",
				"  TRUE : ch0.received;",
				"esac",
			}, "\n")},
			{"ch0.next_value", strings.Join([]string{
				"case",
				"  running_pid = pid & state = state0 & !ch0.filled : TRUE;",
				"  TRUE : ch0.value;",
				"esac",
			}, "\n")},
		},
		Defs: []tmplAssign{},
	}
	result, err := ConvertAbstractModuleToTemplate(mod)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
		return
	}
	result_string := InstantiateTemplate(result)
	expected_string := InstantiateTemplate(expected)
	if result_string != expected_string {
		t.Errorf("Expect %s to be %s", result_string, expected_string)
	}
}

const expectedTemplateResult = `
MODULE A(arg1, arg2)
  VAR
    var1 : boolean;
    var2 : 0..16;
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
	mod := tmplNuSMVModule{
		Name:    "A",
		Args:    []string{"arg1", "arg2"},
		Vars:    []tmplVar{tmplVar{"var1", "boolean"}, tmplVar{"var2", "0..16"}},
		Assigns: []tmplAssign{tmplAssign{"init(var1)", "FALSE"}, tmplAssign{"next(var1)", "case\n  TRUE : FALSE;\nesac"}},
		Defs:    []tmplAssign{tmplAssign{"var3", "TRUE"}},
	}

	if result := InstantiateTemplate(mod); result != expectedTemplateResult {
		t.Errorf("Expect %s to be %s", result, expectedTemplateResult)
	}
}
