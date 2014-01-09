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
    init(received) := FALSE;
    init(value_0) := FALSE;
    next(filled) := filleds[running_pid];
    next(received) := receiveds[running_pid];
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
    b : 0..8;
    next_state : {state0, state1, state2, state3};
    state : {state0, state1, state2, state3};
  ASSIGN
    ch0.next_filled :=
      case
        running_pid = pid & state = state1 & next_state = state2 : TRUE;
        running_pid = pid & state = state2 & next_state = state3 : FALSE;
        TRUE : ch0.filled;
      esac;
    ch0.next_received :=
      case
        running_pid = pid & state = state1 & next_state = state2 : FALSE;
        TRUE : ch0.received;
      esac;
    ch0.next_value_0 :=
      case
        running_pid = pid & state = state1 & next_state = state2 : TRUE;
        TRUE : ch0.value_0;
      esac;
    init(state) := state0;
    next(b) := b;
    next(state) := next_state;
    next_state :=
      case
        running_pid = pid & state = state0 & ((TRUE)) : {state1};
        running_pid = pid & state = state1 & ((!(ch0.filled))) : {state2};
        running_pid = pid & state = state2 & (((ch0.filled) & (ch0.received))) : {state3};
        TRUE : state;
      esac;

MODULE main()
  VAR
    __pid0_ch : HandshakeChannel0Proxy(ch);
    ch : HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0);
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
