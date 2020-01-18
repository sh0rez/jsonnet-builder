package main

import "fmt"

type refType struct {
	named
	to string
}

func Ref(name, to string) refType {
	return refType{named(name), to}
}

func (r refType) String() string {
	return r.to
}

func Field(name string, parent Type, fieldName string) refType {
	return refType{
		named: named(name),
		to:    fmt.Sprintf("%s.%s", parent.String(), fieldName),
	}
}
