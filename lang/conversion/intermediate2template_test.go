package conversion

import (
	"github.com/cookieo9/go-misc/pp"
	"strings"
	"testing"
)

func TestConvertMainModuleToTemplate(t *testing.T) {
	mod := intMainModule{
		Vars: []intVar{
			{"ch", "HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0)"},
			{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
			{"proc1", "__pid0_ProcA(running_pid, 0, __pid0_ch)"},
			{"running_pid", "{0}"},
		},
		Assigns: []intAssign{
			{"running_pid", "{0}"},
		},
		Defs: []intAssign{
			{"ch_filled", "[__pid0_ch.next_filled]"},
			{"ch_received", "[__pid0_ch.next_received]"},
			{"ch_value_0", "[__pid0_ch.next_value_0]"},
		},
	}
	expected := []tmplModule{
		{
			Name: "main",
			Args: []string{},
			Vars: []tmplVar{
				{"ch", "HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0)"},
				{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
				{"proc1", "__pid0_ProcA(running_pid, 0, __pid0_ch)"},
				{"running_pid", "{0}"},
			},
			Assigns: []tmplAssign{
				{"running_pid", "{0}"},
			},
			Defs: []tmplAssign{
				{"ch_filled", "[__pid0_ch.next_filled]"},
				{"ch_received", "[__pid0_ch.next_received]"},
				{"ch_value_0", "[__pid0_ch.next_value_0]"},
			},
		},
	}
	err, tmplMods := convertMainModuleToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}

func TestConvertHandshakeChannelToTemplate(t *testing.T) {
	mod := intHandshakeChannel{
		Name:      "HandshakeChannel0",
		ValueType: []string{"boolean"},
	}
	expected := []tmplModule{
		{
			Name: "HandshakeChannel0",
			Args: []string{"running_pid", "filleds", "receiveds", "values_0"},
			Vars: []tmplVar{
				{"filled", "boolean"},
				{"received", "boolean"},
				{"value_0", "boolean"},
			},
			Assigns: []tmplAssign{
				{"init(filled)", "FALSE"},
				{"next(filled)", "filleds[running_pid]"},
				{"init(received)", "FALSE"},
				{"next(received)", "receiveds[running_pid]"},
				{"init(value_0)", "FALSE"},
				{"next(value_0)", "values_0[running_pid]"},
			},
		},
		{
			Name: "HandshakeChannel0Proxy",
			Args: []string{"ch"},
			Vars: []tmplVar{
				{"next_filled", "boolean"},
				{"next_received", "boolean"},
				{"next_value_0", "boolean"},
			},
			Defs: []tmplAssign{
				{"filled", "ch.filled"},
				{"received", "ch.received"},
				{"value_0", "ch.value_0"},
			},
		},
	}
	err, tmplMods := convertHandshakeChannelToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}

func TestConvertProcModuleToTemplate(t *testing.T) {
	mod := intProcModule{
		Name: "__pid0_ProcA",
		Args: []string{"running_pid", "pid", "ch0"},
		Vars: []intVar{
			{"b", "0..8"},
		},
		InitState: intState("state0"),
		Trans: map[intState][]intTransition{
			"state0": []intTransition{
				{
					Condition: "",
					Actions: map[intState][]intAssign{
						"state1": nil,
					},
				},
			},
			"state1": []intTransition{
				{
					Condition: "!ch0.filled",
					Actions: map[intState][]intAssign{
						"state2": []intAssign{
							{"ch0.next_filled", "TRUE"},
							{"ch0.next_received", "FALSE"},
							{"ch0.next_value_0", "TRUE"},
						},
					},
				},
			},
		},
		Defaults: map[string]string{
			"ch0.next_filled":   "ch0.filled",
			"ch0.next_received": "ch0.received",
			"ch0.next_value_0":  "ch0.value_0",
		},
		Defs: []intAssign{},
	}
	expected := []tmplModule{
		{
			Name: "__pid0_ProcA",
			Args: []string{"running_pid", "pid", "ch0"},
			Vars: []tmplVar{
				{"state", "{state0, state1, state2}"},
				{"next_state", "{state0, state1, state2}"},
				{"b", "0..8"},
			},
			Assigns: []tmplAssign{
				{"init(state)", "state0"},
				{"next(state)", "next_state"},
				{"next_state", strings.Join([]string{
					"case",
					"  running_pid = pid & state = state0 & ((TRUE)) : {state1};",
					"  running_pid = pid & state = state1 & ((!ch0.filled)) : {state2};",
					"  TRUE : state;",
					"esac",
				}, "\n")},
				{"ch0.next_filled", strings.Join([]string{
					"case",
					"  running_pid = pid & state = state1 & next_state = state2 : TRUE;",
					"  TRUE : ch0.filled;",
					"esac",
				}, "\n")},
				{"ch0.next_received", strings.Join([]string{
					"case",
					"  running_pid = pid & state = state1 & next_state = state2 : FALSE;",
					"  TRUE : ch0.received;",
					"esac",
				}, "\n")},
				{"ch0.next_value_0", strings.Join([]string{
					"case",
					"  running_pid = pid & state = state1 & next_state = state2 : TRUE;",
					"  TRUE : ch0.value_0;",
					"esac",
				}, "\n")},
			},
		},
	}
	err, tmplMods := convertProcModuleToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}
