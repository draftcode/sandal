package conversion_deprecated

import (
	"github.com/cookieo9/go-misc/pp"
	. "github.com/draftcode/sandal/lang/data"
	"github.com/kylelemons/godebug/diff"
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
						Elems: []Type{NamedType{"bool"}},
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
						Elems: []Type{NamedType{"bool"}},
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
			ZeroValue: []string{"FALSE"},
		},
		intProcModule{
			Name: "__pid0_ProcA",
			Args: []string{"__orig_ch0"},
			Vars: []intVar{
				{"ch0", "HandshakeChannel0Proxy(__orig_ch0)"},
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
					Condition: "!(ch0.ready)",
					Actions: []intAssign{
						{"ch0.send_filled", "TRUE"},
						{"ch0.send_value_0", "TRUE"},
					},
				},
				{
					FromState: "state2",
					NextState: "state3",
					Condition: "(ch0.ready) & (ch0.received)",
					Actions: []intAssign{
						{"ch0.send_leaving", "TRUE"},
					},
				},
			},
			Defaults: map[string]string{
				"ch0.send_leaving":  "FALSE",
				"ch0.send_filled":   "FALSE",
				"ch0.recv_received": "FALSE",
				"ch0.send_value_0":  "ch0.value_0",
				"next(b)":           "b",
			},
			Defs: []intAssign{},
		},
		intMainModule{
			Vars: []intVar{
				{"ch", "HandshakeChannel0"},
				{"proc1", "process __pid0_ProcA(ch)"},
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
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}
