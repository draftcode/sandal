package sandal

import (
	"testing"
)

func TestNamedTypeEquality(t *testing.T) {
	intNamedTypeA := NamedType{"int"}
	intNamedTypeB := NamedType{"int"}
	boolNamedType := NamedType{"bool"}
	boolArrayType := ArrayType{NamedType{"bool"}}
	if !intNamedTypeA.equal(intNamedTypeB) {
		t.Error("Expect int is equal to int")
	}
	if intNamedTypeA.equal(boolNamedType) {
		t.Error("Expect bool is not equal to int")
	}
	if intNamedTypeA.equal(boolArrayType) {
		t.Error("Expect []bool is not equal to int")
	}
}

func TestCallaleTypeEquality(t *testing.T) {
	callableTypeA := CallableType{[]Type{NamedType{"int"}}}
	callableTypeB := CallableType{[]Type{NamedType{"int"}}}
	callableTypeC := CallableType{[]Type{NamedType{"bool"}}}
	callableTypeD := CallableType{[]Type{NamedType{"int"}, NamedType{"int"}}}
	if !callableTypeA.equal(callableTypeB) {
		t.Error("Expect callable(int) is equal to callable(int)")
	}
	if callableTypeA.equal(callableTypeC) {
		t.Error("Expect callable(int) is not equal to callable(bool)")
	}
	if callableTypeA.equal(callableTypeD) {
		t.Error("Expect callable(int) is not equal to callable(int, int)")
	}
	if callableTypeA.equal(NamedType{"int"}) {
		t.Error("Expect callable(int) is not equal to int")
	}
}

func TestArrayTypeEquality(t *testing.T) {
	arrayTypeA := ArrayType{NamedType{"int"}}
	arrayTypeB := ArrayType{NamedType{"int"}}
	arrayTypeC := ArrayType{NamedType{"bool"}}
	if !arrayTypeA.equal(arrayTypeB) {
		t.Error("Expect []int is equal to []int")
	}
	if arrayTypeA.equal(arrayTypeC) {
		t.Error("Expect []int is not equal to []bool")
	}
	if arrayTypeA.equal(NamedType{"int"}) {
		t.Error("Expect []int is not equal to int")
	}
}

func TestHandshakeChannelTypeEquality(t *testing.T) {
	chTypeA := HandshakeChannelType{false, []Type{NamedType{"int"}}}
	chTypeB := HandshakeChannelType{false, []Type{NamedType{"int"}}}
	chTypeC := HandshakeChannelType{false, []Type{NamedType{"bool"}}}
	chTypeD := HandshakeChannelType{true, []Type{NamedType{"int"}}}
	chTypeE := HandshakeChannelType{false, []Type{NamedType{"int"}, NamedType{"int"}}}
	bufCh := BufferedChannelType{false, nil, []Type{NamedType{"int"}}}
	if !chTypeA.equal(chTypeB) {
		t.Error("Expect channel {\"int\"} is equal to channel {\"int\"}")
	}
	if chTypeA.equal(bufCh) {
		t.Error("Expect channel {\"int\"} is not equal to channel [] {\"int\"}")
	}
	if chTypeA.equal(chTypeC) {
		t.Error("Expect channel {\"int\"} is not equal to channel {\"bool\"}")
	}
	if !chTypeA.equal(chTypeD) {
		t.Error("Expect channel {\"int\"} is equal to unstable channel {\"int\"}")
	}
	if chTypeA.equal(chTypeE) {
		t.Error("Expect channel {\"int\"} is not equal to channel {\"int\", \"int\"}")
	}
}

func TestBufferedChannelTypeEquality(t *testing.T) {
	chTypeA := BufferedChannelType{false, nil, []Type{NamedType{"int"}}}
	chTypeB := BufferedChannelType{false, nil, []Type{NamedType{"int"}}}
	chTypeC := BufferedChannelType{false, nil, []Type{NamedType{"bool"}}}
	chTypeD := BufferedChannelType{true, nil, []Type{NamedType{"int"}}}
	chTypeE := BufferedChannelType{false, &NumberExpression{"1"}, []Type{NamedType{"int"}}}
	chTypeF := BufferedChannelType{false, nil, []Type{NamedType{"int"}, NamedType{"int"}}}
	handshakeType := HandshakeChannelType{false, []Type{NamedType{"int"}}}
	if !chTypeA.equal(chTypeB) {
		t.Error("Expect channel [] {\"int\"} is equal to channel [] {\"int\"}")
	}
	if chTypeA.equal(handshakeType) {
		t.Error("Expect channel [] {\"int\"} is not equal to channel {\"int\"}")
	}
	if chTypeA.equal(chTypeC) {
		t.Error("Expect channel [] {\"int\"} is not equal to channel [] {\"bool\"}")
	}
	if !chTypeA.equal(chTypeD) {
		t.Error("Expect channel [] {\"int\"} is equal to unstable channel [] {\"int\"}")
	}
	if !chTypeA.equal(chTypeE) {
		t.Error("Expect channel [] {\"int\"} is equal to channel [1] {\"int\"}")
	}
	if chTypeA.equal(chTypeF) {
		t.Error("Expect channel [] {\"int\"} is not equal to channel [] {\"int\", \"int\"}")
	}
}
