package conversion

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

const expectedResult = `
MODULE HandshakeChannel0(running_pid, filleds, receiveds, values_0)
  VAR
    filled : boolean;
    received : boolean;
    value_0 : boolean;
  ASSIGN
    init(filled) := FALSE;
    next(filled) := filleds[running_pid];
    init(received) := FALSE;
    next(received) := receiveds[running_pid];
    init(value_0) := FALSE;
    next(value_0) := values_0[running_pid];

MODULE HandshakeChannel0Proxy(ch)
  VAR
    next_filled : boolean;
    next_received : boolean;
    next_value_0 : boolean;
  DEFINE
    filled := ch.filled;
    received := ch.received;
    value_0 := ch.value_0;

MODULE __pid0_ProcA(running_pid, pid, ch0)
  VAR
    state : {state0, state1, state2};
    b : 0..8;
  ASSIGN
    init(state) := state0;
    next(state) :=
      case
        running_pid = pid & state = state0 : state1;
        running_pid = pid & state = state1 & !ch0.filled : state2;
        TRUE : state;
      esac;
    ch0.next_filled :=
      case
        running_pid = pid & state = state1 & !ch0.filled : TRUE;
        TRUE : ch0.filled;
      esac;
    ch0.next_received :=
      case
        running_pid = pid & state = state1 & !ch0.filled : FALSE;
        TRUE : ch0.received;
      esac;
    ch0.next_value_0 :=
      case
        running_pid = pid & state = state1 & !ch0.filled : TRUE;
        TRUE : ch0.value_0;
      esac;

MODULE main()
  VAR
    ch : HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0);
    __pid0_ch : HandshakeChannel0Proxy(ch);
    proc1 : __pid0_ProcA(running_pid, 0, __pid0_ch);
    running_pid : {0};
  ASSIGN
    running_pid := {0};
  DEFINE
    ch_filled := [__pid0_ch.next_filled];
    ch_received := [__pid0_ch.next_received];
    ch_value_0 := [__pid0_ch.next_value_0];
`

func TestConvertASTToNuSMV(t *testing.T) {
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
	err, mod := ConvertASTToNuSMV(defs)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if mod != expectedResult {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectedResult, mod)
	}
}
