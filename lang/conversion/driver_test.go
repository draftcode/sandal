package conversion

import (
	. "github.com/draftcode/sandal/lang/data"
	"testing"
)

const expectedResult1 = `
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
    state : {state0, state1, state2, state3};
    transition : {notrans, trans0, trans1, trans2};
  TRANS transition = trans0 -> (TRUE);
  TRANS transition = trans1 -> (!(ch0.filled));
  TRANS transition = trans2 -> ((ch0.filled) & (ch0.received));
  ASSIGN
    ch0.next_filled :=
      case
        transition = trans1 : TRUE;
        transition = trans2 : FALSE;
        TRUE : ch0.filled;
      esac;
    ch0.next_received :=
      case
        transition = trans1 : FALSE;
        TRUE : ch0.received;
      esac;
    ch0.next_value_0 :=
      case
        transition = trans1 : TRUE;
        TRUE : ch0.value_0;
      esac;
    init(state) := state0;
    next(b) :=
      case
        TRUE : b;
      esac;
    next(state) :=
      case
        transition = trans0 : state1;
        transition = trans1 : state2;
        transition = trans2 : state3;
        TRUE : state;
      esac;
    transition :=
      case
        running_pid = pid & state = state0 & ((TRUE)) : {trans0};
        running_pid = pid & state = state1 & ((!(ch0.filled))) : {trans1};
        running_pid = pid & state = state2 & (((ch0.filled) & (ch0.received))) : {trans2};
        TRUE : notrans;
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

func TestConvertASTToNuSMV1(t *testing.T) {
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
	err, mod := ConvertASTToNuSMV(defs)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if mod != expectedResult1 {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectedResult1, mod)
	}
}

const expectedResult2 = `
MODULE BufferedChannel0(running_pid, filleds, receiveds, values_0)
  VAR
    filled : array 0..2 of boolean;
    next_idx : 0..3;
    value_0 : array 0..2 of boolean;
  ASSIGN
    init(filled[0]) := FALSE;
    init(filled[1]) := FALSE;
    init(filled[2]) := FALSE;
    init(next_idx) := 0;
    next(filled[0]) :=
      case
        filleds[running_pid] & next_idx = 0 : TRUE;
        receiveds[running_pid] : filled[1];
        TRUE : filled[0];
      esac;
    next(filled[1]) :=
      case
        filleds[running_pid] & next_idx = 1 : TRUE;
        receiveds[running_pid] : filled[2];
        TRUE : filled[1];
      esac;
    next(filled[2]) :=
      case
        filleds[running_pid] & next_idx = 2 : TRUE;
        receiveds[running_pid] : FALSE;
        TRUE : filled[2];
      esac;
    next(next_idx) :=
      case
        filleds[running_pid] : (next_idx + 1) mod 4;
        TRUE : next_idx;
      esac;
    next(value_0[0]) :=
      case
        filleds[running_pid] & next_idx = 0 : values_0[running_pid];
        receiveds[running_pid] : value_0[1];
        TRUE : value_0[0];
      esac;
    next(value_0[1]) :=
      case
        filleds[running_pid] & next_idx = 1 : values_0[running_pid];
        receiveds[running_pid] : value_0[2];
        TRUE : value_0[1];
      esac;
    next(value_0[2]) :=
      case
        filleds[running_pid] & next_idx = 2 : values_0[running_pid];
        TRUE : value_0[2];
      esac;

MODULE BufferedChannel0Proxy(ch)
  VAR
    filled : boolean;
    next_value_0 : boolean;
    received : boolean;
  DEFINE
    full := ch.next_idx = 3;
    ready := ch.filled[0];
    value_0 := ch.value_0[0];

MODULE __pid0_ProcA(running_pid, pid, ch0)
  VAR
    b : 0..8;
    state : {state0, state1, state2};
    transition : {notrans, trans0, trans1};
  TRANS transition = trans0 -> (TRUE);
  TRANS transition = trans1 -> (!(ch0.full));
  ASSIGN
    ch0.filled :=
      case
        transition = trans1 : TRUE;
        TRUE : FALSE;
      esac;
    ch0.next_value_0 :=
      case
        transition = trans1 : TRUE;
        TRUE : ch0.value_0;
      esac;
    ch0.received :=
      case
        TRUE : FALSE;
      esac;
    init(state) := state0;
    next(b) :=
      case
        TRUE : b;
      esac;
    next(state) :=
      case
        transition = trans0 : state1;
        transition = trans1 : state2;
        TRUE : state;
      esac;
    transition :=
      case
        running_pid = pid & state = state0 & ((TRUE)) : {trans0};
        running_pid = pid & state = state1 & ((!(ch0.full))) : {trans1};
        TRUE : notrans;
      esac;

MODULE main()
  VAR
    __pid0_ch : BufferedChannel0Proxy(ch);
    ch : BufferedChannel0(running_pid, ch_filled, ch_received, ch_value_0);
    proc1 : __pid0_ProcA(running_pid, 0, __pid0_ch);
    running_pid : {0};
  ASSIGN
    running_pid := {0};
  DEFINE
    ch_filled := [__pid0_ch.filled];
    ch_received := [__pid0_ch.received];
    ch_value_0 := [__pid0_ch.next_value_0];
`

func TestConvertASTToNuSMV2(t *testing.T) {
	defs := []Definition{
		ProcDefinition{
			Name: "ProcA",
			Parameters: []Parameter{
				{
					Name: "ch0",
					Type: BufferedChannelType{
						BufferSize: NumberExpression{Pos{}, "3"},
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
					Type: BufferedChannelType{
						BufferSize: NumberExpression{Pos{}, "3"},
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
	if mod != expectedResult2 {
		t.Errorf("Unmatched\nExpected %s\nGot      %s", expectedResult2, mod)
	}
}
