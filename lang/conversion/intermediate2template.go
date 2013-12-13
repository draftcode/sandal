package conversion

import (
	"bytes"
	"fmt"
	"sort"
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
	assigns := []tmplAssign{}
	for _, intassign := range mod.Assigns {
		assigns = append(assigns, tmplAssign{intassign.LHS, intassign.RHS})
	}
	defs := []tmplAssign{}
	for _, intassign := range mod.Defs {
		defs = append(defs, tmplAssign{intassign.LHS, intassign.RHS})
	}
	return nil, []tmplModule{
		{
			Name:    "main",
			Vars:    vars,
			Assigns: assigns,
			Defs:    defs,
		},
	}
}
func convertHandshakeChannelToTemplate(mod intHandshakeChannel) (error, []tmplModule) {
	// Fields for the channel module.
	args := []string{"running_pid", "filleds", "receiveds"}
	for i, _ := range mod.ValueType {
		args = append(args, fmt.Sprintf("values_%d", i))
	}
	vars := []tmplVar{
		{"filled", "boolean"},
		{"received", "boolean"},
	}
	for i, elem := range mod.ValueType {
		vars = append(vars, tmplVar{fmt.Sprintf("value_%d", i), elem})
	}
	assigns := []tmplAssign{
		{"init(filled)", "FALSE"},
		{"next(filled)", "filleds[running_pid]"},
		{"init(received)", "FALSE"},
		{"next(received)", "receiveds[running_pid]"},
	}
	for i, elem := range mod.ValueType {
		zeroValue := zeroValueInNuSMV(elem)
		if zeroValue != "" {
			assigns = append(assigns, tmplAssign{
				fmt.Sprintf("init(value_%d)", i),
				zeroValue,
			})
		}
		assigns = append(assigns, tmplAssign{
			fmt.Sprintf("next(value_%d)", i),
			fmt.Sprintf("values_%d[running_pid]", i),
		})
	}
	// Fields for the channel proxy module
	proxyVars := []tmplVar{
		{"next_filled", "boolean"},
		{"next_received", "boolean"},
	}
	for i, elem := range mod.ValueType {
		proxyVars = append(proxyVars, tmplVar{
			fmt.Sprintf("next_value_%d", i),
			elem,
		})
	}
	proxyDefs := []tmplAssign{
		{"filled", "ch.filled"},
		{"received", "ch.received"},
	}
	for i, _ := range mod.ValueType {
		proxyDefs = append(proxyDefs, tmplAssign{
			fmt.Sprintf("value_%d", i),
			fmt.Sprintf("ch.value_%d", i),
		})
	}
	return nil, []tmplModule{
		{
			Name:    mod.Name,
			Args:    args,
			Vars:    vars,
			Assigns: assigns,
		},
		{
			Name: mod.Name + "Proxy",
			Args: []string{"ch"},
			Vars: proxyVars,
			Defs: proxyDefs,
		},
	}
}
func convertBufferedChannelToTemplate(mod intModule) (error, []tmplModule) {
	panic("Not implemented")
}
func convertProcModuleToTemplate(mod intProcModule) (error, []tmplModule) {
	vars := []tmplVar{
		{"state", "{" + argJoin(collectStates(mod)) + "}"},
	}
	for _, intvar := range mod.Vars {
		vars = append(vars, tmplVar{intvar.Name, intvar.Type})
	}
	assigns := []tmplAssign{
		{"init(state)", string(mod.InitState)},
		{"next(state)", instantiateCaseTemplate(caseTmplValue{
			Cases:   buildStateTransition(mod),
			Default: "state;",
		})},
	}
	assigns = append(assigns, buildAssignments(mod)...)
	return nil, []tmplModule{
		{
			Name:    mod.Name,
			Args:    mod.Args,
			Vars:    vars,
			Assigns: assigns,
		},
	}
}

// ========================================

func collectStates(mod intProcModule) []string {
	m := make(map[string]bool)
	for state, intTrans := range mod.Trans {
		m[string(state)] = true
		for _, tr := range intTrans {
			if len(tr.Actions) != 1 {
				// TODO
				panic("multiple actions not supported")
			}
			for nextState, _ := range tr.Actions {
				m[string(nextState)] = true
			}
		}
	}
	states := []string{}
	for state, _ := range m {
		states = append(states, state)
	}
	sort.Strings(states)
	return states
}

type condMap map[string][]string

func buildStateTransition(mod intProcModule) []caseTmplCase {
	transs := make(map[intState]condMap)
	for state, intTrans := range mod.Trans {
		tmplTrans := make(map[string][]string)
		for _, tr := range intTrans {
			cond := tr.Condition
			if cond == "" {
				cond = fmt.Sprintf("running_pid = pid & state = %s", state)
			} else {
				cond = fmt.Sprintf("running_pid = pid & state = %s & %s", state, cond)
			}
			for nextState, _ := range tr.Actions {
				tmplTrans[cond] = append(tmplTrans[cond], string(nextState))
			}
		}
		transs[state] = tmplTrans
	}

	cases := []caseTmplCase{}
	for _, condmap := range transs {
		for cond, nextStates := range condmap {
			if len(nextStates) == 1 {
				cases = append(cases, caseTmplCase{cond, nextStates[0] + ";"})
			} else {
				cases = append(cases, caseTmplCase{
					cond, "{" + argJoin(nextStates) + "};",
				})
			}
		}
	}
	return cases
}

func buildAssignments(mod intProcModule) []tmplAssign {
	assignss := make(map[string][]caseTmplCase)
	for state, intTrans := range mod.Trans {
		for _, tr := range intTrans {
			cond := tr.Condition
			if cond == "" {
				cond = fmt.Sprintf("running_pid = pid & state = %s", state)
			} else {
				cond = fmt.Sprintf("running_pid = pid & state = %s & %s", state, cond)
			}

			for _, assigns := range tr.Actions {
				for _, assign := range assigns {
					assignss[assign.LHS] = append(
						assignss[assign.LHS],
						caseTmplCase{cond, assign.RHS + ";"},
					)
				}
			}
		}
	}

	retAssigns := []tmplAssign{}
	for lhs, assigns := range assignss {
		retAssigns = append(retAssigns, tmplAssign{
			LHS: lhs,
			RHS: instantiateCaseTemplate(caseTmplValue{
				Cases:   assigns,
				Default: mod.Defaults[lhs] + ";",
			}),
		})
	}
	return retAssigns
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

type caseTmplCase struct {
	Condition string
	Value     string
}

type caseTmplValue struct {
	Cases   []caseTmplCase
	Default string
}

func instantiateCaseTemplate(val caseTmplValue) string {
	tmpl, err := template.New("NuSMVCase").Parse(caseTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, val)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
