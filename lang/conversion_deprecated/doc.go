// package conversion_deprecated provides a way to convert Sandal's AST to NuSMV's module.
//
// Sandal
//
//   proc ProcA(ch0 channel {bool}) {
//     var b int
//     send(ch0, true)
//   }
//
//   init {
//     ch:    channel {bool},
//     proc1: ProcA(ch),
//   }
//
// Intermediate module
//
//   intHandshakeChannel{
//   	Name:      "HandshakeChannel0",
//   	ValueType: []string{"boolean"},
//   }
//   intProcModule{
//   	Name: "__pid0_ProcA",
//   	Args: []string{"running_pid", "pid", "ch0"},
//   	Vars: []intVar{
//   		{"b", "0..8"},
//   	},
//   	InitState: intState("state0"),
//   	Trans: map[intState][]intTransition{
//   		"state0": []intTransition{
//   			{
//   				Condition: "",
//   				Actions: map[intState][]intAssign{
//   					"state1": nil,
//   				},
//   			},
//   		},
//   		"state1": []intTransition{
//   			{
//   				Condition: "!ch0.filled",
//   				Actions: map[intState][]intAssign{
//   					"state2": []intAssign{
//   						{"ch0.next_filled", "TRUE"},
//   						{"ch0.next_received", "FALSE"},
//   						{"ch0.next_value_0", "TRUE"},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Defaults: map[string]string{
//   		"ch0.next_filled":   "ch0.filled",
//   		"ch0.next_received": "ch0.received",
//   		"ch0.next_value_0":  "ch0.value_0",
//   	},
//   	Defs: []intAssign{},
//   }
//   intMainModule{
//   	Vars: []intVar{
//   		{"ch", "HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0)"},
//   		{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
//   		{"proc1", "__pid0_ProcA(running_pid, 0, __pid0_ch)"},
//   		{"running_pid", "{0}"},
//   	},
//   	Assigns: []intAssign{
//   		{"running_pid", "{0}"},
//   	},
//   	Defs: []intAssign{
//   		{"ch_filled", "[__pid0_ch.next_filled]"},
//   		{"ch_received", "[__pid0_ch.next_received]"},
//   		{"ch_value_0", "[__pid0_ch.next_value_0]"},
//   	},
//   }
//
// Template module
//
//   tmplModule{
//   	Name: "HandshakeChannel0",
//   	Args: []string{"running_pid", "filleds", "receiveds", "values_0"},
//   	Vars: []tmplVar{
//   		{"filled", "boolean"},
//   		{"received", "boolean"},
//   		{"value_0", "boolean"},
//   	},
//   	Assigns: []tmplAssign{
//   		{"init(filled)", "FALSE"},
//   		{"next(filled)", "filleds[running_pid]"},
//   		{"init(received)", "FALSE"},
//   		{"next(received)", "receiveds[running_pid]"},
//   		{"init(value_0)", "FALSE"},
//   		{"next(value_0)", "values_0[running_pid]"},
//   	},
//   }
//   tmplModule{
//   	Name: "HandshakeChannel0Proxy",
//   	Args: []string{"ch"},
//   	Vars: []tmplVar{
//   		{"next_filled", "boolean"},
//   		{"next_received", "boolean"},
//   		{"next_value_0", "boolean"},
//   	},
//   	Defs: []tmplAssign{
//   		{"filled", "ch.filled"},
//   		{"received", "ch.received"},
//   		{"value_0", "ch.value_0"},
//   	},
//   }
//   intProcModule{
//   	Name: "__pid0_ProcA",
//   	Args: []string{"running_pid", "pid", "ch0"},
//   	Vars: []intVar{
//   		{"b", "0..8"},
//   	},
//   	InitState: intState("state0"),
//   	Trans: map[intState][]intTransition{
//   		"state0": []intTransition{
//   			{
//   				Condition: "",
//   				Actions: map[intState][]intAssign{
//   					"state1": nil,
//   				},
//   			},
//   		},
//   		"state1": []intTransition{
//   			{
//   				Condition: "!ch0.filled",
//   				Actions: map[intState][]intAssign{
//   					"state2": []intAssign{
//   						{"ch0.next_filled", "TRUE"},
//   						{"ch0.next_received", "FALSE"},
//   						{"ch0.next_value_0", "TRUE"},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Defaults: map[string]string{
//   		"ch0.next_filled":   "ch0.filled",
//   		"ch0.next_received": "ch0.received",
//   		"ch0.next_value_0":  "ch0.value_0",
//   	},
//   	Defs: []intAssign{},
//   }
//   tmplModule{
//   	Name: "main",
//   	Args: []string{},
//   	Vars: []tmplVar{
//   		{"ch", "HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0)"},
//   		{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
//   		{"proc1", "__pid0_ProcA(running_pid, 0, __pid0_ch)"},
//   		{"running_pid", "{0}"},
//   	},
//   	Assigns: []tmplAssign{
//   		{"running_pid", "{0}"},
//   	},
//   	Defs: []tmplAssign{
//   		{"ch_filled", "[__pid0_ch.next_filled]"},
//   		{"ch_received", "[__pid0_ch.next_received]"},
//   		{"ch_value_0", "[__pid0_ch.next_value_0]"},
//   	},
//   }
//
// NuSMV Module
//
//   MODULE HandshakeChannel0(running_pid, filleds, receiveds, values_0)
//     VAR
//       filled : boolean;
//       received : boolean;
//       value_0 : boolean;
//     ASSIGN
//       init(filled) := FALSE;
//       next(filled) := filleds[running_pid];
//       init(received) := FALSE;
//       next(received) := receiveds[running_pid];
//       init(value_0) := FALSE;
//       next(value_0) := values_0[running_pid];
//
//   MODULE HandshakeChannel0Proxy(ch)
//     VAR
//       next_filled : boolean;
//       next_received : boolean;
//       next_value_0 : boolean;
//     DEFINE
//       filled := ch.filled;
//       received := ch.received;
//       value_0 := ch.value_0;
//
//   MODULE __pid0_ProcA(running_pid, pid, ch0)
//     VAR
//       state : {state0, state1, state2};
//       b : 0..8;
//     ASSIGN
//       init(state) := state0;
//       next(state) :=
//         case
//           running_pid = pid & state = state0 : state1;
//           running_pid = pid & state = state1 & !ch0.filled : state2;
//           TRUE : state;
//         esac;
//       ch0.next_filled :=
//         case
//           running_pid = pid & state = state1 & !ch0.filled : TRUE;
//           TRUE : ch0.filled;
//         esac;
//       ch0.next_received :=
//         case
//           running_pid = pid & state = state1 & !ch0.filled : FALSE;
//           TRUE : ch0.received;
//         esac;
//       ch0.next_value_0 :=
//         case
//           running_pid = pid & state = state1 & !ch0.filled : TRUE;
//           TRUE : ch0.value_0;
//         esac;
//
//   MODULE main()
//     VAR
//       ch : HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value_0);
//       __pid0_ch : HandshakeChannel0Proxy(ch);
//       proc1 : __pid0_ProcA(running_pid, 0, __pid0_ch);
//       running_pid : {0};
//     ASSIGN
//       running_pid := {0};
//     DEFINE
//       ch_filled := [__pid0_ch.next_filled];
//       ch_received := [__pid0_ch.next_received];
//       ch_value_0 := [__pid0_ch.next_value_0];
package conversion_deprecated
