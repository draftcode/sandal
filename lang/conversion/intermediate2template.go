package conversion

import (
	"fmt"
	"sort"
)

func convertIntermediateModuleToTemplate(mods []intModule) (error, []tmplModule) {
	tmplMods := []tmplModule{}
	for _, mod := range mods {
		var tmplMod tmplModule
		var err error
		switch mod := mod.(type) {
		case intMainModule:
			err, tmplMod = convertMainModuleToTemplate(mod)
		case intHandshakeChannel:
			err, tmplMod = convertHandshakeChannelToTemplate(mod)
		case intBufferedChannel:
			err, tmplMod = convertBufferedChannelToTemplate(mod)
		case intProcModule:
			err, tmplMod = convertProcModuleToTemplate(mod)
		}
		if err != nil {
			return err, nil
		}
		tmplMods = append(tmplMods, tmplMod)
	}
	// tmpl.Name = module.Name
	// tmpl.Args = append([]string{"running_pid", "pid"}, module.Args...)
	// tmpl.Vars = []tmplVar{
	// 	{"state", "{" + strings.Join(extractStates(module), ", ") + "}"},
	// }
	// for _, absvar := range module.Vars {
	// 	tmpl.Vars = append(tmpl.Vars, tmplVar{absvar.Name, absvar.Type})
	// }
	// assignCond := make(map[string]map[string]string)
	// for state, transes := range module.Trans {
	// 	for _, trans := range transes {
	// 		extractAssignCondition(state, trans, assignCond)
	// 	}
	// }
	// for variable, cases := range assignCond {
	// 	var defaultValue string
	// 	if variable == "next(state)" {
	// 		defaultValue = "state"
	// 	} else if defaultValue, hasValue := module.Defaults[variable]; !hasValue {
	// 		return tmplNuSMVModule{}, fmt.Errorf("There is no default value for %s", variable)
	// 	}
	// }
	return nil, tmplMods
}

func convertMainModuleToTemplate(mod intMainModule) (error, tmplModule) {
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
	tmplMod := tmplModule{
		Name:    "main",
		Vars:    vars,
		Assigns: assigns,
		Defs:    defs,
	}
	return nil, tmplMod
}
func convertHandshakeChannelToTemplate(mod intHandshakeChannel) (error, tmplModule) {
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
		assigns = append(assigns, tmplAssign{
			fmt.Sprintf("init(value_%d)", i),
			zeroValueInNuSMV(elem),
		})
		assigns = append(assigns, tmplAssign{
			fmt.Sprintf("next(value_%d)", i),
			fmt.Sprintf("values_%d[running_pid]", i),
		})
	}
	return nil, tmplModule{
		Name:    mod.Name,
		Args:    args,
		Vars:    vars,
		Assigns: assigns,
	}
}
func convertBufferedChannelToTemplate(mod intModule) (error, tmplModule) {
	panic("Not implemented")
}
func convertProcModuleToTemplate(mod intModule) (error, tmplModule) {
	panic("Not implemented")
}

// ========================================

type assignCond struct {
	state     string
	condition string
	value     string
}

func extractStates(module intProcModule) (states []string) {
	states_map := make(map[intState]bool)
	states_map[module.InitState] = true
	for s, transes := range module.Trans {
		states_map[s] = true
		for _, trans := range transes {
			for t, _ := range trans.Actions {
				states_map[t] = true
			}
		}
	}

	for state, _ := range states_map {
		states = append(states, string(state))
	}
	sort.StringSlice(states).Sort()
	return
}

// ========================================

func zeroValueInNuSMV(ty string) string {
	switch ty {
	case "boolean":
		return "FALSE"
	default:
		panic("Not implemented")
	}
}
