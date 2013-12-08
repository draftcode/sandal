// Package conversion provides a way to convert Sandal's AST to NuSMV's module.
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
//   intMainModule{
//   	Vars: []intVar{
//   		{"ch", "HandshakeChannel0(running_pid, __filled_ch, __received_ch, __value_ch)"},
//   		{"proc1", "__pid0_ProcA(running_pid, 0, __pid0_ch)"},
//   		{"__pid0_ch", "HandshakeChannel0Proxy(ch)"},
//   		{"running_pid", "{0}"},
//   	},
//   	Assigns: []intAssign{
//   		{"running_pid", "{0}"},
//   	},
//   	Defs: []intAssign{
//   		{"__filled_ch", "[__pid0_ch.next_filled]"},
//   		{"__received_ch", "[__pid0_ch.next_received]"},
//   		{"__next_value_ch", "[__pid0_ch.next_value]"},
//   	},
//   }
//   intHandshakeChannel{
//   	Name: "HandshakeChannel0",
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
//   						{"ch0.next_value", "TRUE"},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Defaults: map[string]string{
//   		"ch0.next_filled":   "ch0.filled",
//   		"ch0.next_received": "ch0.received",
//   		"ch0.next_value":    "ch0.value",
//   	},
//   	Defs: []intAssign{
//   	},
//   }
//
// Template module
//
//   tmplModule{
//   	Name: "main",
//   	Args: []string{},
//   	Vars: []tmplVar{
//   		{"running_pid", "{0}"},
//   		{"ch", "HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value)"},
//   		{"proc1", "ProcA(running_pid, 0, ch_pid0)"},
//   		{"ch_pid0", "HandshakeChannel0Proxy(ch)"},
//   	},
//   	Assigns: []tmplAssign{
//   		{"running_pid", "{0}"},
//   	},
//   	Defs: []tmplAssign{
//   		{"ch_filled", "[ch_pid0.next_filled]"}
//   		{"ch_received", "[ch_pid0.next_received]"}
//   		{"ch_next_value", "[ch_pid0.next_value]"}
//   	}
//   }
//   tmplNuSMVModule{
//   	Name: "HandshakeChannel0",
//   	Args: []string{"running_pid", "filleds", "receiveds", "values"},
//   	Vars: []tmplVar{
//   		{"filled", "boolean"},
//   		{"received", "boolean"},
//   		{"value", "boolean"},
//   	},
//   	Assigns: []tmplAssign{
//   		{"init(filled)", "FALSE"},
//   		{"next(filled)", "filleds[running_pid]"},
//   		{"init(received)", "FALSE"},
//   		{"next(received)", "receiveds[running_pid]"},
//   		{"init(value)", "FALSE"},
//   		{"next(value)", "values[running_pid]"},
//   	},
//   }
//   tmplNuSMVModule{
//   	Name: "HandshakeChannel0Proxy",
//   	Args: []string{"ch"},
//   	Vars: []tmplVar{
//   		{"next_filled", "boolean"},
//   		{"next_received", "boolean"},
//   		{"next_value", "boolean"},
//   	},
//   	Defs: []tmplAssign{
//   		{"filled", "ch.filled"},
//   		{"received", "ch.received"},
//   		{"value", "ch.value"},
//   	},
//   }
//   tmplNuSMVModule{
//   	Name: "ProcA_proc1",
//   	Args: []string{"running_pid", "pid", "ch0"},
//   	Vars: []tmplVar{
//   		{"state", "{state0, state1, state2}"},
//   	},
//   	Assigns: []tmplAssign{
//   		{"init(state)", "state0"},
//   		{"next(state)", strings.Join([]string{
//   			"case",
//   			"  running_pid = pid & state = state0 : state1;",
//   			"  running_pid = pid & state = state1 & !ch0.filled : state2;",
//   			"  TRUE : state;",
//   			"esac;",
//   		}, "\n")},
//   		{"ch0.next_filled", strings.Join([]string{
//   			"case",
//   			"  running_pid = pid & state = state1 & !ch0.filled : TRUE;"
//   			"  TRUE : ch0.filled;"
//   			"esac;",
//   		}, "\n")},
//   		{"ch0.next_received", strings.Join([]string{
//   			"case",
//   			"  running_pid = pid & state = state1 & !ch0.filled : TRUE;"
//   			"  TRUE : ch0.received;"
//   			"esac;",
//   		}, "\n")},
//   		{"ch0.next_value", strings.Join([]string{
//   			"case",
//   			"  running_pid = pid & state = state1 & !ch0.filled : TRUE;"
//   			"  TRUE : ch0.value;"
//   			"esac;",
//   		}, "\n")},
//   	},
//   }
//
// NuSMV Module
//
//   MODULE main()
//     VAR
//       running_pid : {0};
//       ch : HandshakeChannel0(running_pid, ch_filled, ch_received, ch_value);
//       proc1 : ProcA(running_pid, 0, ch_pid0);
//       ch_pid0 : HandshakeChannel0Proxy(ch);
//     ASSIGN
//       running_pid := {0};
//     DEFINE
//       ch_filled := [ch_pid0.next_filled]
//       ch_received := [ch_pid0.next_received]
//       ch_value := [ch_pid0.next_value]
//
//   MODULE HandshakeChannel0(running_pid, filleds, receiveds, values)
//     VAR
//       filled : boolean;
//       received : boolean;
//       value : boolean;
//     ASSIGN
//       init(filled) := FALSE;
//       next(filled) := filleds[running_pid];
//       init(received) := FALSE;
//       next(received) := receiveds[running_pid];
//       init(value) := FALSE;
//       next(value) := values[running_pid];
//
//   MODULE HandshakeChannel0Proxy(ch)
//     VAR
//       next_filled : boolean;
//       next_received : boolean;
//       next_value : boolean;
//     DEFINE
//       filled := ch.filled;
//       received := ch.received;
//       value := ch.value;
//
//   MODULE ProcA_proc1(running_pid, pid, ch0)
//     VAR
//       state : {state0, state1, state2};
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
//           running_pid = pid & state = state1 & !ch0.filled : TRUE;
//           TRUE : ch0.received;
//         esac;
//       ch0.next_value :=
//         case
//           running_pid = pid & state = state1 & !ch0.filled : TRUE;
//           TRUE : ch0.value;
//         esac;
package conversion
