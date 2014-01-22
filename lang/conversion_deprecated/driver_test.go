package conversion_deprecated

import (
	. "github.com/draftcode/sandal/lang/data"
	"github.com/kylelemons/godebug/diff"
	"testing"
)

const expectedResult1 = `
MODULE HandshakeChannel0()
  VAR
    filled : boolean;
    received : boolean;
    value_0 : boolean;
  ASSIGN
    init(filled) := FALSE;
    init(received) := FALSE;
    init(value_0) := FALSE;

MODULE HandshakeChannel0Proxy(ch)
  VAR
    recv_received : boolean;
    send_filled : boolean;
    send_leaving : boolean;
    send_value_0 : boolean;
  ASSIGN
    next(ch.filled) :=
      case
        send_filled : TRUE;
        send_leaving : FALSE;
        TRUE : ch.filled;
      esac;
    next(ch.received) :=
      case
        send_filled : FALSE;
        send_leaving : FALSE;
        recv_received : TRUE;
        TRUE : ch.received;
      esac;
    next(ch.value_0) :=
      case
        send_filled : send_value_0;
        TRUE : ch.value_0;
      esac;
  DEFINE
    ready := ch.filled;
    received := ch.received;
    value_0 := ch.value_0;

MODULE __pid0_ProcA(__orig_ch0)
  JUSTICE
    running
  VAR
    b : 0..8;
    ch0 : HandshakeChannel0Proxy(__orig_ch0);
    state : {state0, state1, state2, state3};
    transition : {notrans, trans0, trans1, trans2};
  TRANS transition = trans0 -> (TRUE);
  TRANS transition = trans1 -> (!(ch0.ready));
  TRANS transition = trans2 -> ((ch0.ready) & (ch0.received));
  ASSIGN
    ch0.recv_received :=
      case
        TRUE : FALSE;
      esac;
    ch0.send_filled :=
      case
        transition = trans1 : TRUE;
        TRUE : FALSE;
      esac;
    ch0.send_leaving :=
      case
        transition = trans2 : TRUE;
        TRUE : FALSE;
      esac;
    ch0.send_value_0 :=
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
        state = state0 & ((TRUE)) : {trans0};
        state = state1 & ((!(ch0.ready))) : {trans1};
        state = state2 & (((ch0.ready) & (ch0.received))) : {trans2};
        TRUE : notrans;
      esac;

MODULE main()
  VAR
    ch : HandshakeChannel0;
    proc1 : process __pid0_ProcA(ch);
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
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectedResult1, mod))
	}
}

const expectedResult2 = `
MODULE BufferedChannel0()
  VAR
    filled : array 0..2 of boolean;
    next_idx : 0..3;
    value_0 : array 0..2 of boolean;
  ASSIGN
    init(filled[0]) := FALSE;
    init(filled[1]) := FALSE;
    init(filled[2]) := FALSE;
    init(next_idx) := 0;
    init(value_0[0]) := FALSE;
    init(value_0[1]) := FALSE;
    init(value_0[2]) := FALSE;

MODULE BufferedChannel0Proxy(ch)
  VAR
    recv_received : boolean;
    send_filled : boolean;
    send_value_0 : boolean;
  ASSIGN
    next(ch.filled[0]) :=
      case
        send_filled & ch.next_idx = 0 : TRUE;
        recv_received : ch.filled[1];
        TRUE : ch.filled[0];
      esac;
    next(ch.filled[1]) :=
      case
        send_filled & ch.next_idx = 1 : TRUE;
        recv_received : ch.filled[2];
        TRUE : ch.filled[1];
      esac;
    next(ch.filled[2]) :=
      case
        send_filled & ch.next_idx = 2 : TRUE;
        recv_received : FALSE;
        TRUE : ch.filled[2];
      esac;
    next(ch.next_idx) :=
      case
        send_filled & ch.next_idx < 3 : ch.next_idx + 1;
        recv_received & ch.next_idx > 0 : ch.next_idx - 1;
        TRUE : ch.next_idx;
      esac;
    next(ch.value_0[0]) :=
      case
        send_filled & ch.next_idx = 0 : send_value_0;
        recv_received : ch.value_0[1];
        TRUE : ch.value_0[0];
      esac;
    next(ch.value_0[1]) :=
      case
        send_filled & ch.next_idx = 1 : send_value_0;
        recv_received : ch.value_0[2];
        TRUE : ch.value_0[1];
      esac;
    next(ch.value_0[2]) :=
      case
        send_filled & ch.next_idx = 2 : send_value_0;
        TRUE : ch.value_0[2];
      esac;
  DEFINE
    full := ch.next_idx = 3;
    ready := ch.filled[0];
    value_0 := ch.value_0[0];

MODULE __pid0_ProcA(__orig_ch0)
  JUSTICE
    running
  VAR
    b : 0..8;
    ch0 : BufferedChannel0Proxy(__orig_ch0);
    state : {state0, state1, state2};
    transition : {notrans, trans0, trans1};
  TRANS transition = trans0 -> (TRUE);
  TRANS transition = trans1 -> (!(ch0.full));
  ASSIGN
    ch0.recv_received :=
      case
        TRUE : FALSE;
      esac;
    ch0.send_filled :=
      case
        transition = trans1 : TRUE;
        TRUE : FALSE;
      esac;
    ch0.send_value_0 :=
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
        TRUE : state;
      esac;
    transition :=
      case
        state = state0 & ((TRUE)) : {trans0};
        state = state1 & ((!(ch0.full))) : {trans1};
        TRUE : notrans;
      esac;

MODULE main()
  VAR
    ch : BufferedChannel0;
    proc1 : process __pid0_ProcA(ch);
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
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectedResult2, mod))
	}
}
