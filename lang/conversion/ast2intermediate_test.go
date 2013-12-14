package conversion

import (
	"github.com/cookieo9/go-misc/pp"
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

func TestConvertASTToIntModule(t *testing.T) {
	defs := []Definition{
		ProcDefinition{
			Name: "ProcA",
			Parameters: []Parameter{
				{
					Name: "ch0",
					Type: HandshakeChannelType{
						IsUnstable: false,
						Elems:      []Type{NamedType{"bool"}},
					},
				},
			},
			Statements: []Statement{
				VarDeclStatement{
					Name: "b",
					Type: NamedType{"int"},
				},
				SendStatement{
					Channel: IdentifierExpression{Pos{}, "ch0"},
					Args: []Expression{
						TrueExpression{Pos{}},
					},
				},
			},
		},
		InitBlock{
			Vars: []InitVar{
				ChannelVar{
					Name: "ch",
					Type: HandshakeChannelType{
						IsUnstable: false,
						Elems:      []Type{NamedType{"bool"}},
					},
				},
				InstanceVar{
					Name:        "proc1",
					ProcDefName: "ProcA",
					Args: []Expression{
						IdentifierExpression{Pos{}, "ch"},
					},
				},
			},
		},
	}
	expected := []intModule{
		intHandshakeChannel{
			Name:      "HandshakeChannel0",
			ValueType: []string{"boolean"},
		},
		intProcModule{
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
						Condition: "!(ch0.filled)",
						Actions: map[intState][]intAssign{
							"state2": []intAssign{
								{"ch0.next_filled", "TRUE"},
								{"ch0.next_received", "FALSE"},
								{"ch0.next_value_0", "TRUE"},
							},
						},
					},
				},
				"state2": []intTransition{
					{
						Condition: "(ch0.filled) & (ch0.received)",
						Actions: map[intState][]intAssign{
							"state3": []intAssign{
								{"ch0.next_filled", "FALSE"},
							},
						},
					},
				},
			},
			Defaults: map[string]string{
				"ch0.next_filled":   "ch0.filled",
				"ch0.next_received": "ch0.received",
				"ch0.next_value_0":  "ch0.value_0",
				"next(b)":           "b",
			},
			Defs: []intAssign{},
		},
		intMainModule{
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
		},
	}
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(intMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}
