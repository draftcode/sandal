package conversion

import (
	. "github.com/draftcode/sandal/lang/data"
	"strings"
)

func ConvertASTToNuSMV(defs []Definition) (error, string) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		return err, ""
	}

	err, tmplMods := convertIntermediateModuleToTemplate(intMods)
	if err != nil {
		return err, ""
	}

	mods := []string{}
	for _, tmplMod := range tmplMods {
		mods = append(mods, instantiateTemplate(tmplMod))
	}

	return nil, strings.Join(mods, "")
}
