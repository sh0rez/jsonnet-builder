package main

import (
	"fmt"
	"strconv"
)

// strings
type stringType struct {
	named
	value string
}

func String(name, value string) stringType {
	return stringType{named: named(name), value: value}
}

func (s stringType) String() string {
	return fmt.Sprintf(`"%s"`, s.value)
}

// ints
type intType struct {
	named
	value int
}

func Int(name string, value int) intType {
	return intType{named: named(name), value: value}
}

func (s intType) String() string {
	return strconv.Itoa(s.value)
}
