package conversion

import (
	"github.com/cookieo9/go-misc/pp"
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
	expected := tmplModule{
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
	}
	err, tmplMod := convertMainModuleToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMod)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}

func TestHandshakeChannelToTemplate(t *testing.T) {
	mod := intHandshakeChannel{
		Name:      "HandshakeChannel0",
		ValueType: []string{"boolean"},
	}
	expected := tmplModule{
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
	}
	err, tmplMod := convertHandshakeChannelToTemplate(mod)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(tmplMod)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}
