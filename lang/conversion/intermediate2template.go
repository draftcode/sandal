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
				conds = append(conds, "(" + trans[transName].Condition + ")")
			}
		}
		cond := fmt.Sprintf("running_pid = pid & state = %s & (%s)",
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
