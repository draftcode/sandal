package lang

import (
	"testing"
)

func TestConvertNuSMV(t *testing.T) {
	defs := []Definition{&ProcDefinition{"A",
		[]Parameter{Parameter{"ch", HandshakeChannelType{false, []Type{NamedType{"bool"}}}},
			Parameter{"chs", ArrayType{HandshakeChannelType{false, []Type{NamedType{"bit"}}}}}},
		[]Statement{&NullStatement{}}}}
	ConvertNuSMV(defs)
}
