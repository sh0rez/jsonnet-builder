package main

type localType struct {
	value Type
}

func Local(value Type) localType {
	return localType{value}
}

func (t localType) String() string {
	return t.value.String()
}

func (t localType) Name() string {
	return t.value.Name()
}
