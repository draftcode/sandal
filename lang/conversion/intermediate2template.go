package conversion

import (
	"sort"
)

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
