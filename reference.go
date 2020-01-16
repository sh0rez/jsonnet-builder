package main

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
