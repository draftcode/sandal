package conversion_deprecated

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

func convertIntermediateModuleToTemplate(mods []intModule) (error, []tmplModule) {
	retTmplMods := []tmplModule{}
	for _, mod := range mods {
		var tmplMods []tmplModule
		var err error
		switch mod := mod.(type) {
		case intMainModule:
			err, tmplMods = convertMainModuleToTemplate(mod)
		case intHandshakeChannel:
			err, tmplMods = convertHandshakeChannelToTemplate(mod)
		case intBufferedChannel:
			err, tmplMods = convertBufferedChannelToTemplate(mod)
		case intProcModule:
			err, tmplMods = convertProcModuleToTemplate(mod)
		}
		if err != nil {
			return err, nil
		}
		retTmplMods = append(retTmplMods, tmplMods...)
	}
	return nil, retTmplMods
}

func convertMainModuleToTemplate(mod intMainModule) (error, []tmplModule) {
	vars := []tmplVar{}
	for _, intvar := range mod.Vars {
		vars = append(vars, tmplVar{intvar.Name, intvar.Type})
	}
	return nil, []tmplModule{
		{
			Name:     "main",
			Vars:     vars,
			LtlSpecs: mod.LtlSpecs,
		},
	}
}
func convertHandshakeChannelToTemplate(mod intHandshakeChannel) (error, []tmplModule) {
	channelModule := tmplModule{}
	{
		channelModule.Name = mod.Name
		channelModule.Vars = []tmplVar{
			{"filled", "boolean"},
			{"received", "boolean"},
		}
		for i, elem := range mod.ValueType {
			channelModule.Vars = append(channelModule.Vars,
				tmplVar{fmt.Sprintf("value_%d", i), elem},
			)
		}
		channelModule.Assigns = []tmplAssign{
			{"init(filled)", "FALSE"},
			{"init(received)", "FALSE"},
		}
		for i, _ := range mod.ValueType {
			channelModule.Assigns = append(channelModule.Assigns, tmplAssign{
				fmt.Sprintf("init(value_%d)", i),
				mod.ZeroValue[i],
			})
		}
	}
	proxyModule := tmplModule{}
	{
		proxyModule.Name = mod.Name + "Proxy"
		proxyModule.Args = []string{"ch"}
		proxyModule.Vars = []tmplVar{
			{"send_filled", "boolean"},
			{"send_leaving", "boolean"},
			{"recv_received", "boolean"},
		}
		for i, elem := range mod.ValueType {
			proxyModule.Vars = append(proxyModule.Vars, tmplVar{
				fmt.Sprintf("send_value_%d", i),
				elem,
			})
		}
		proxyModule.Defs = []tmplAssign{
			{"ready", "ch.filled"},
			{"received", "ch.received"},
		}
		for i, _ := range mod.ValueType {
			proxyModule.Defs = append(proxyModule.Defs, tmplAssign{
				fmt.Sprintf("value_%d", i),
				fmt.Sprintf("ch.value_%d", i),
			})
		}
		proxyModule.Assigns = []tmplAssign{
			{"next(ch.filled)", strings.Join([]string{
				"case",
				"  send_filled : TRUE;",
				"  send_leaving : FALSE;",
				"  TRUE : ch.filled;",
				"esac",
			}, "\n")},
			{"next(ch.received)", strings.Join([]string{
				"case",
				"  send_filled : FALSE;",
				"  send_leaving : FALSE;",
				"  recv_received : TRUE;",
				"  TRUE : ch.received;",
				"esac",
			}, "\n")},
		}
		for i, _ := range mod.ValueType {
			proxyModule.Assigns = append(proxyModule.Assigns, tmplAssign{
				fmt.Sprintf("next(ch.value_%d)", i),
				strings.Join([]string{
					"case",
					fmt.Sprintf("  send_filled : send_value_%d;", i),
					fmt.Sprintf("  TRUE : ch.value_%d;", i),
					"esac",
				}, "\n"),
			})
		}
	}
	return nil, []tmplModule{channelModule, proxyModule}
}
func convertBufferedChannelToTemplate(mod intBufferedChannel) (error, []tmplModule) {
	channelModule := tmplModule{}
	{
		channelModule.Name = mod.Name
		channelModule.Vars = []tmplVar{
			{"filled", fmt.Sprintf("array 0..%d of boolean", mod.Length-1)},
			{"next_idx", fmt.Sprintf("0..%d", mod.Length)},
		}
		for i, elem := range mod.ValueType {
			channelModule.Vars = append(channelModule.Vars, tmplVar{
				fmt.Sprintf("value_%d", i),
				fmt.Sprintf("array 0..%d of %s", mod.Length-1, elem),
			})
		}
		channelModule.Assigns = []tmplAssign{
			{"init(next_idx)", "0"},
		}
		for i, _ := range mod.ValueType {
			for bufIdx := 0; bufIdx < mod.Length; bufIdx++ {
				channelModule.Assigns = append(channelModule.Assigns, tmplAssign{
					fmt.Sprintf("init(filled[%d])", bufIdx),
					"FALSE",
				})
				channelModule.Assigns = append(channelModule.Assigns, tmplAssign{
					fmt.Sprintf("init(value_%d[%d])", i, bufIdx),
					mod.ZeroValue[i],
				})
			}
		}
	}
	proxyModule := tmplModule{}
	{
		proxyModule.Name = mod.Name + "Proxy"
		proxyModule.Args = []string{"ch"}
		proxyModule.Vars = []tmplVar{
			{"send_filled", "boolean"},
			{"recv_received", "boolean"},
		}
		for i, elem := range mod.ValueType {
			proxyModule.Vars = append(proxyModule.Vars, tmplVar{
				fmt.Sprintf("send_value_%d", i),
				elem,
			})
		}
		proxyModule.Defs = []tmplAssign{
			{"full", fmt.Sprintf("ch.next_idx = %d", mod.Length)},
			{"ready", "ch.filled[0]"},
		}
		for i, _ := range mod.ValueType {
			proxyModule.Defs = append(proxyModule.Defs, tmplAssign{
				fmt.Sprintf("value_%d", i),
				fmt.Sprintf("ch.value_%d[0]", i),
			})
		}
		proxyModule.Assigns = []tmplAssign{
			{"next(ch.next_idx)", strings.Join([]string{
				"case",
				fmt.Sprintf("  send_filled & ch.next_idx < %d : ch.next_idx + 1;", mod.Length),
				"  recv_received & ch.next_idx > 0 : ch.next_idx - 1;",
				"  TRUE : ch.next_idx;",
				"esac",
			}, "\n")},
		}
		for i, _ := range mod.ValueType {
			for bufIdx := 0; bufIdx < mod.Length; bufIdx++ {
				proxyModule.Assigns = append(proxyModule.Assigns, tmplAssign{
					fmt.Sprintf("next(ch.filled[%d])", bufIdx),
					createFilledCase(bufIdx, mod.Length),
				})
				proxyModule.Assigns = append(proxyModule.Assigns, tmplAssign{
					fmt.Sprintf("next(ch.value_%d[%d])", i, bufIdx),
					createValueCase(i, bufIdx, mod.Length),
				})
			}
		}
	}
	return nil, []tmplModule{channelModule, proxyModule}
}
func createFilledCase(bufIdx, length int) string {
	buf := []string{"case"}
	buf = append(buf, fmt.Sprintf("  send_filled & ch.next_idx = %d : TRUE;", bufIdx))
	if bufIdx+1 < length {
		buf = append(buf, fmt.Sprintf("  recv_received : ch.filled[%d];", bufIdx+1))
	} else {
		buf = append(buf, "  recv_received : FALSE;")
	}
	buf = append(buf, fmt.Sprintf("  TRUE : ch.filled[%d];", bufIdx))
	buf = append(buf, "esac")
	return strings.Join(buf, "\n")
}
func createValueCase(elemIdx, bufIdx, length int) string {
	buf := []string{"case"}
	buf = append(buf, fmt.Sprintf("  send_filled & ch.next_idx = %d : send_value_%d;", bufIdx, elemIdx))
	if bufIdx+1 < length {
		buf = append(buf, fmt.Sprintf("  recv_received : ch.value_%d[%d];", elemIdx, bufIdx+1))
	}
	buf = append(buf, fmt.Sprintf("  TRUE : ch.value_%d[%d];", elemIdx, bufIdx))
	buf = append(buf, "esac")
	return strings.Join(buf, "\n")
}
func convertProcModuleToTemplate(mod intProcModule) (error, []tmplModule) {
	transitions := nameTransitions(mod.Trans)
	vars := []tmplVar{
		{"state", "{" + argJoin(collectStates(mod)) + "}"},
		{"transition", "{notrans, " + argJoin(collectTransitions(transitions)) + "}"},
	}
	for _, intvar := range mod.Vars {
		vars = append(vars, tmplVar{intvar.Name, intvar.Type})
	}
	trans, assigns := convertTransition(mod.InitState, mod.Defaults, transitions)
	return nil, []tmplModule{
		{
			Name:    mod.Name,
			Args:    mod.Args,
			Vars:    vars,
			Trans:   trans,
			Assigns: assigns,
			Justice: "running",
		},
	}
}

// ========================================

func nameTransitions(trans []intTransition) map[string]intTransition {
	ret := make(map[string]intTransition)
	for num, transition := range trans {
		ret[fmt.Sprintf("trans%d", num)] = transition
	}
	return ret
}

func collectTransitions(trans map[string]intTransition) (ret []string) {
	for transName, _ := range trans {
		ret = append(ret, transName)
	}
	sort.Strings(ret)
	return
}

func collectStates(mod intProcModule) []string {
	m := make(map[string]bool)
	m[string(mod.InitState)] = true
	for _, intTrans := range mod.Trans {
		m[string(intTrans.FromState)] = true
		m[string(intTrans.NextState)] = true
	}
	states := []string{}
	for state, _ := range m {
		states = append(states, state)
	}
	sort.Strings(states)
	return states
}

func convertTransition(initState intState, defaults map[string]string, trans map[string]intTransition) ([]string, []tmplAssign) {
	assigns := []tmplAssign{
		{"transition", instantiateCaseTemplate(caseTmplValue{
			Cases:   buildTransitionAssignment(trans),
			Default: "notrans;",
		})},
		{"init(state)", string(initState)},
		{"next(state)", instantiateCaseTemplate(caseTmplValue{
			Cases:   buildNextStateAssignment(trans),
			Default: "state;",
		})},
	}
	assigns = append(assigns, buildVariableAssignments(defaults, trans)...)
	return buildTransitionTransitions(trans), assigns
}

func buildTransitionAssignment(trans map[string]intTransition) (ret []caseTmplCase) {
	m := make(map[intState][]string) // Transitions keyed with FromState
	for transName, transition := range trans {
		m[transition.FromState] = append(m[transition.FromState], transName)
	}
	for state, transNames := range m {
		sort.Strings(transNames)
		conds := []string{}
		for _, transName := range transNames {
			if trans[transName].Condition == "" {
				conds = append(conds, "(TRUE)")
			} else {
				conds = append(conds, "("+trans[transName].Condition+")")
			}
		}
		cond := fmt.Sprintf("state = %s & (%s)",
			state, strings.Join(uniqAndSort(conds), " | "))
		ret = append(ret, caseTmplCase{
			Condition: cond,
			Value:     "{" + argJoin(transNames) + "};",
		})
	}
	return
}

func buildTransitionTransitions(trans map[string]intTransition) (ret []string) {
	for transName, transition := range trans {
		cond := "TRUE"
		if transition.Condition != "" {
			cond = transition.Condition
		}
		ret = append(ret, fmt.Sprintf("transition = %s -> (%s)", transName, cond))
	}
	return
}

func buildNextStateAssignment(trans map[string]intTransition) (ret []caseTmplCase) {
	for transName, transition := range trans {
		ret = append(ret, caseTmplCase{
			Condition: "transition = " + transName,
			Value:     string(transition.NextState) + ";",
		})
	}
	return
}

func buildVariableAssignments(defaults map[string]string, trans map[string]intTransition) (ret []tmplAssign) {
	m := make(map[string][]caseTmplCase) // Assignments keyed with variables
	for transName, transition := range trans {
		for _, action := range transition.Actions {
			m[action.LHS] = append(m[action.LHS], caseTmplCase{
				Condition: "transition = " + transName,
				Value:     action.RHS + ";",
			})
		}
	}

	for variable, defaultValue := range defaults {
		ret = append(ret, tmplAssign{
			variable, instantiateCaseTemplate(caseTmplValue{
				Cases:   m[variable],
				Default: defaultValue + ";",
			}),
		})
	}
	return
}

// ========================================

func zeroValueInNuSMV(ty string) string {
	switch ty {
	case "boolean":
		return "FALSE"
	default:
		return ""
	}
}

const caseTemplate = `case{{range .Cases}}
  {{.Condition}} : {{.Value}}{{end}}
  TRUE : {{.Default}}
esac`

type (
	caseTmplCase struct {
		Condition string
		Value     string
	}

	caseTmplValue struct {
		Cases   []caseTmplCase
		Default string
	}

	caseTmplCases []caseTmplCase
)

func (l caseTmplCases) Len() int           { return len(l) }
func (l caseTmplCases) Less(i, j int) bool { return l[i].Condition < l[j].Condition }
func (l caseTmplCases) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func instantiateCaseTemplate(val caseTmplValue) string {
	tmpl, err := template.New("NuSMVCase").Parse(caseTemplate)
	if err != nil {
		panic(err)
	}

	sort.Sort(caseTmplCases(val.Cases))

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, val)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

// ========================================

func uniqAndSort(strs []string) []string {
	m := make(map[string]bool)
	for _, s := range strs {
		m[s] = true
	}
	r := []string{}
	for s, _ := range m {
		r = append(r, s)
	}
	sort.Strings(r)
	return r
}
