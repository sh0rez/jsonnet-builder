package main

import (
	"fmt"
	"strings"
)

// function definition
type funcType struct {
	named
	args  []Type
	value Type
}

func Func(name string, args []Type, returns Type) funcType {
	return funcType{
		named: named(name),
		args:  args,
		value: returns,
	}
}

func (f funcType) Args() string {
	return argsString(f.args)
}

func (f funcType) String() string {
	return f.value.String()
}

// function call
type callType struct {
	named
	funcName string
	args     []Type
}

func (c callType) String() string {
	return fmt.Sprintf("%s(%s)", c.funcName, argsString(c.args))
}

func Call(name, funcName string, args []Type) callType {
	for k, v := range args {
		if v == nil {
			panic(fmt.Sprintf("argument `%s` in call to `%s` is nil", k, funcName))
		}
	}

	return callType{
		named:    named(name),
		funcName: funcName,
		args:     args,
	}
}

// function arguments
func Args(args ...Type) []Type {
	return args
}

func argsString(m []Type) string {
	s := ""
	for _, v := range m {
		if _, ok := v.(requiredArgType); ok {
			s += fmt.Sprintf(", %s", v.Name())
			continue
		}

		s += fmt.Sprintf(", %s=%s", v.Name(), v.String())
	}

	s = strings.TrimPrefix(s, ", ")
	return s
}

type requiredArgType struct {
	value Type
}

func (r requiredArgType) Name() string {
	return r.value.Name()
}

func (r requiredArgType) String() string {
	return r.value.String()
}

// mark an argument as required (no default value)
func Required(t Type) requiredArgType {
	return requiredArgType{t}
}

// CallChain allows to chain multiple calls
func CallChain(name string, calls ...callType) callType {
	if len(calls) == 1 {
		panic("callChain with a single call is redundant")
	}

	var last Type = Ref("", "")
	for i, c := range calls {
		last = Call("",
			strings.TrimPrefix(Field("", last, c.funcName).String(), "."),
			c.args,
		)

		if i == len(calls)-1 {
			l := last.(callType)
			l.named = named(name)
			return l
		}
	}

	panic("loop did not return. This should never happen")
}
