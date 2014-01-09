package data

import (
	"testing"
)

func TestNamedTypeEquality(t *testing.T) {
	intNamedTypeA := NamedType{"int"}
	intNamedTypeB := NamedType{"int"}
	boolNamedType := NamedType{"bool"}
	boolArrayType := ArrayType{NamedType{"bool"}}
	if !intNamedTypeA.Equal(intNamedTypeB) {
		t.Error("Expect int is equal to int")
	}
	if intNamedTypeA.Equal(boolNamedType) {
		t.Error("Expect bool is not equal to int")
	}
	if intNamedTypeA.Equal(boolArrayType) {
		t.Error("Expect []bool is not equal to int")
	}
}

func TestCallaleTypeEquality(t *testing.T) {
	callableTypeA := CallableType{[]Type{NamedType{"int"}}}
	callableTypeB := CallableType{[]Type{NamedType{"int"}}}
	callableTypeC := CallableType{[]Type{NamedType{"bool"}}}
	callableTypeD := CallableType{[]Type{NamedType{"int"}, NamedType{"int"}}}
	if !callableTypeA.Equal(callableTypeB) {
		t.Error("Expect callable(int) is equal to callable(int)")
	}
	if callableTypeA.Equal(callableTypeC) {
		t.Error("Expect callable(int) is not equal to callable(bool)")
	}
	if callableTypeA.Equal(callableTypeD) {
		t.Error("Expect callable(int) is not equal to callable(int, int)")
	}
	if callableTypeA.Equal(NamedType{"int"}) {
		t.Error("Expect callable(int) is not equal to int")
	}
}

func TestArrayTypeEquality(t *testing.T) {
	arrayTypeA := ArrayType{NamedType{"int"}}
	arrayTypeB := ArrayType{NamedType{"int"}}
	arrayTypeC := ArrayType{NamedType{"bool"}}
	if !arrayTypeA.Equal(arrayTypeB) {
		t.Error("Expect []int is equal to []int")
	}
	if arrayTypeA.Equal(arrayTypeC) {
		t.Error("Expect []int is not equal to []bool")
	}
	if arrayTypeA.Equal(NamedType{"int"}) {
		t.Error("Expect []int is not equal to int")
	}
}

func TestHandshakeChannelTypeEquality(t *testing.T) {
	chTypeA := HandshakeChannelType{[]Type{NamedType{"int"}}}
	chTypeB := HandshakeChannelType{[]Type{NamedType{"int"}}}
	chTypeC := HandshakeChannelType{[]Type{NamedType{"bool"}}}
	chTypeD := HandshakeChannelType{[]Type{NamedType{"int"}, NamedType{"int"}}}
	bufCh := BufferedChannelType{nil, []Type{NamedType{"int"}}}
	if !chTypeA.Equal(chTypeB) {
		t.Error(`Expect channel {"int"} is equal to channel {"int"}`)
	}
	if chTypeA.Equal(bufCh) {
		t.Error(`Expect channel {"int"} is not equal to channel [] {"int"}`)
	}
	if chTypeA.Equal(chTypeC) {
		t.Error(`Expect channel {"int"} is not equal to channel {"bool"}`)
	}
	if chTypeA.Equal(chTypeD) {
		t.Error(`Expect channel {"int"} is not equal to channel {"int", "int"}`)
	}
}

func TestBufferedChannelTypeEquality(t *testing.T) {
	chTypeA := BufferedChannelType{nil, []Type{NamedType{"int"}}}
	chTypeB := BufferedChannelType{nil, []Type{NamedType{"int"}}}
	chTypeC := BufferedChannelType{nil, []Type{NamedType{"bool"}}}
	chTypeD := BufferedChannelType{&NumberExpression{Lit: "1"}, []Type{NamedType{"int"}}}
	chTypeE := BufferedChannelType{nil, []Type{NamedType{"int"}, NamedType{"int"}}}
	handshakeType := HandshakeChannelType{[]Type{NamedType{"int"}}}
	if !chTypeA.Equal(chTypeB) {
		t.Error(`Expect channel [] {"int"} is equal to channel [] {"int"}`)
	}
	if chTypeA.Equal(handshakeType) {
		t.Error(`Expect channel [] {"int"} is not equal to channel {"int"}`)
	}
	if chTypeA.Equal(chTypeC) {
		t.Error(`Expect channel [] {"int"} is not equal to channel [] {"bool"}`)
	}
	if !chTypeA.Equal(chTypeD) {
		t.Error(`Expect channel [] {"int"} is equal to channel [1] {"int"}`)
	}
	if chTypeA.Equal(chTypeE) {
		t.Error(`Expect channel [] {"int"} is not equal to channel [] {"int", "int"}`)
	}
}
