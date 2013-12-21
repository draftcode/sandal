package conversion

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
		{"next_state", "{" + argJoin(collectStates(mod)) + "}"},
	}
	for _, intvar := range mod.Vars {
		vars = append(vars, tmplVar{intvar.Name, intvar.Type})
	}
	trans, cases := buildStateTransition(mod)
	assigns := []tmplAssign{
		{"init(state)", string(mod.InitState)},
		{"next(state)", "next_state"},
		{"next_state", instantiateCaseTemplate(caseTmplValue{
			Cases:   cases,
			Default: "state;",
		})},
	}
	assigns = append(assigns, buildAssignments(mod)...)
	return nil, []tmplModule{
		{
			Name:    mod.Name,
			Args:    mod.Args,
			Vars:    vars,
			Trans:   trans,
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
			if tr.NextState != "" {
				m[string(tr.NextState)] = true
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

func buildStateTransition(mod intProcModule) ([]string, []caseTmplCase) {
	trans := []string{}
	cases := []caseTmplCase{}
	for state, intTrans := range mod.Trans {
		nextStates := []string{}
		conds := []string{}
		nextStateAndCond := make(map[string][]string)
		for _, tr := range intTrans {
			if tr.NextState != "" {
				cond := tr.Condition
				if cond == "" {
					cond = "TRUE"
				}
				nextStateAndCond[string(tr.NextState)] = append(
					nextStateAndCond[string(tr.NextState)],
					cond,
				)
				nextStates = append(nextStates, string(tr.NextState))
				conds = append(conds, fmt.Sprintf("(%s)", cond))
			}
		}

		cond := ""
		if len(nextStateAndCond) > 1 {
			cond = fmt.Sprintf("running_pid = pid & state = %s", state)
			for nextState, conds := range nextStateAndCond {
				cond := strings.Join(uniqAndSort(conds), " | ")
				trans = append(
					trans,
					fmt.Sprintf("state = %s & next_state = %s -> %s", state, nextState, cond),
				)
			}
		} else {
			cond = strings.Join(uniqAndSort(conds), " | ")
			cond = fmt.Sprintf("running_pid = pid & state = %s & (%s)", state, cond)
		}
		cases = append(cases, caseTmplCase{cond, "{" + argJoin(uniqAndSort(nextStates)) + "};"})
	}
	return trans, cases
}

func buildAssignments(mod intProcModule) []tmplAssign {
	assignss := make(map[string][]caseTmplCase)
	for state, intTrans := range mod.Trans {
		for _, tr := range intTrans {
			cond := ""
			if tr.NextState == "" {
				cond = fmt.Sprintf("running_pid = pid & state = %s", state)
			} else {
				cond = fmt.Sprintf("running_pid = pid & state = %s & next_state = %s", state, tr.NextState)
			}
			for _, assign := range tr.Actions {
				assignss[assign.LHS] = append(
					assignss[assign.LHS],
					caseTmplCase{cond, assign.RHS + ";"},
				)
			}
		}
	}

	retAssigns := []tmplAssign{}
	defaultAssigned := make(map[string]bool)
	for lhs, assigns := range assignss {
		defaultValue := mod.Defaults[lhs]
		if defaultValue == "" {
			panic("No default value")
		}
		defaultAssigned[lhs] = true
		retAssigns = append(retAssigns, tmplAssign{
			LHS: lhs,
			RHS: instantiateCaseTemplate(caseTmplValue{
				Cases:   assigns,
				Default: defaultValue + ";",
			}),
		})
	}
	for lhs, defaultValue := range mod.Defaults {
		if !defaultAssigned[lhs] {
			retAssigns = append(retAssigns, tmplAssign{
				LHS: lhs, RHS: defaultValue,
			})
		}
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

func (l caseTmplCases) Len() int { return len(l) }
func (l caseTmplCases) Less(i, j int) bool { return l[i].Condition < l[j].Condition }
func (l caseTmplCases) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

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
