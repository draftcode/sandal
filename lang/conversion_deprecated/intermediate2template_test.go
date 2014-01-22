package conversion_deprecated

import (
	"github.com/cookieo9/go-misc/pp"
	"github.com/kylelemons/godebug/diff"
	"strings"
	"testing"
)

func TestConvertMainModuleToTemplate(t *testing.T) {
	mod := intMainModule{
		Vars: []intVar{
			{"ch", "HandshakeChannel0"},
			{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
			{"proc1", "__pid0_ProcA(__pid0_ch)"},
		},
	}
	expected := []tmplModule{
		{
			Name: "main",
			Args: []string{},
			Vars: []tmplVar{
				{"ch", "HandshakeChannel0"},
				{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
				{"proc1", "__pid0_ProcA(__pid0_ch)"},
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
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}

func TestConvertHandshakeChannelToTemplate(t *testing.T) {
	mod := intHandshakeChannel{
		Name:      "HandshakeChannel0",
		ValueType: []string{"boolean"},
		ZeroValue: []string{"FALSE"},
	}
	expected := []tmplModule{
		{
			Name: "HandshakeChannel0",
			Args: []string{},
			Vars: []tmplVar{
				{"filled", "boolean"},
				{"received", "boolean"},
				{"value_0", "boolean"},
			},
			Assigns: []tmplAssign{
				{"init(filled)", "FALSE"},
				{"init(received)", "FALSE"},
				{"init(value_0)", "FALSE"},
			},
		},
		{
			Name: "HandshakeChannel0Proxy",
			Args: []string{"ch"},
			Vars: []tmplVar{
				{"send_filled", "boolean"},
				{"send_leaving", "boolean"},
				{"recv_received", "boolean"},
				{"send_value_0", "boolean"},
			},
			Defs: []tmplAssign{
				{"ready", "ch.filled"},
				{"received", "ch.received"},
				{"value_0", "ch.value_0"},
			},
			Assigns: []tmplAssign{
				{"next(ch.filled)", strings.Join([]string{
					"case",
					"  send_filled : TRUE;",
					"  send_leaving : FALSE;",
					"  TRUE : ch.filled;",
					"esac",
				}, "\n",)},
				{"next(ch.received)", strings.Join([]string{
					"case",
					"  send_filled : FALSE;",
					"  send_leaving : FALSE;",
					"  recv_received : TRUE;",
					"  TRUE : ch.received;",
					"esac",
				}, "\n",)},
				{"next(ch.value_0)", strings.Join([]string{
					"case",
					"  send_filled : send_value_0;",
					"  TRUE : ch.value_0;",
					"esac",
				}, "\n",)},
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
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}

func TestConvertProcModuleToTemplate(t *testing.T) {
	mod := intProcModule{
		Name: "__pid0_ProcA",
		Args: []string{"ch0"},
		Vars: []intVar{
			{"b", "0..8"},
		},
		InitState: intState("state0"),
		Trans: []intTransition{
			{
				FromState: "state0",
				NextState: "state1",
				Condition: "",
			},
			{
				FromState: "state1",
				NextState: "state2",
				Condition: "!ch0.ready",
				Actions: []intAssign{
					{"ch0.send_filled", "TRUE"},
					{"ch0.send_value_0", "TRUE"},
				},
			},
		},
		Defaults: map[string]string{
			"ch0.send_filled":   "FALSE",
			"ch0.recv_received": "FALSE",
			"ch0.send_value_0":  "ch0.value_0",
		},
		Defs: []intAssign{},
	}
	expected := []tmplModule{
		{
			Name: "__pid0_ProcA",
			Args: []string{"ch0"},
			Vars: []tmplVar{
				{"state", "{state0, state1, state2}"},
				{"transition", "{notrans, trans0, trans1}"},
				{"b", "0..8"},
			},
			Trans: []string{
				"transition = trans0 -> (TRUE)",
				"transition = trans1 -> (!ch0.ready)",
			},
			Assigns: []tmplAssign{
				{"transition", strings.Join([]string{
					"case",
					"  state = state0 & ((TRUE)) : {trans0};",
					"  state = state1 & ((!ch0.ready)) : {trans1};",
					"  TRUE : notrans;",
					"esac",
				}, "\n")},
				{"init(state)", "state0"},
				{"next(state)", strings.Join([]string{
					"case",
					"  transition = trans0 : state1;",
					"  transition = trans1 : state2;",
					"  TRUE : state;",
					"esac",
				}, "\n")},
				{"ch0.send_filled", strings.Join([]string{
					"case",
					"  transition = trans1 : TRUE;",
					"  TRUE : FALSE;",
					"esac",
				}, "\n")},
				{"ch0.recv_received", strings.Join([]string{
					"case",
					"  TRUE : FALSE;",
					"esac",
				}, "\n")},
				{"ch0.send_value_0", strings.Join([]string{
					"case",
					"  transition = trans1 : TRUE;",
					"  TRUE : ch0.value_0;",
					"esac",
				}, "\n")},
			},
			Justice: "running",
		},
	}
	err, tmplMods := convertProcModuleToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}
