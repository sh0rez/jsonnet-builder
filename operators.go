package main

import (
	"fmt"
	"strings"
)

// merge
type mergeType struct {
	value Type
}

func (t mergeType) String() string {
	return t.value.String()
}

func (t mergeType) Name() string {
	return t.value.Name()
}

func Merge(value Type) mergeType {
	if _, ok := value.(hiddenType); ok {
		panic("hiddenType cannot be a child of mergeType, it must be the other way around.")
	}
	return mergeType{value}
}

// hidden field (::)
type hiddenType struct {
	value Type
}

func Hidden(value Type) hiddenType {
	return hiddenType{value}
}

func (h hiddenType) Name() string {
	return h.value.Name()
}

func (h hiddenType) String() string {
	return h.value.String()
}

// arithmetic
type arithType struct {
	named
	operator string
	operands []Type
}

func (m arithType) String() string {
	rendered := make([]string, len(m.operands))
	for i, o := range m.operands {
		rendered[i] = o.String()
	}

	s := strings.Join(rendered, fmt.Sprintf(" %s ", m.operator))
	return s
}

func Add(name string, o ...Type) arithType {
	return arithType{named: named(name), operator: "+", operands: o}
}

func Sub(name string, o ...Type) arithType {
	return arithType{named: named(name), operator: "-", operands: o}
}

func Div(name string, o ...Type) arithType {
	return arithType{named: named(name), operator: "/", operands: o}
}

func Mul(name string, o ...Type) arithType {
	return arithType{named: named(name), operator: "*", operands: o}
}

func Mod(name string, o ...Type) arithType {
	return arithType{named: named(name), operator: "%", operands: o}
}

// string formatting
type sprintfType struct {
	named
	template string
	values   []Type
}

func Sprintf(name, format string, values ...Type) sprintfType {
	return sprintfType{
		named:    named(name),
		template: format,
		values:   values,
	}
}
