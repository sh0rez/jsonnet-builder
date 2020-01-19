package builder

import (
	"fmt"
	"strings"
)

// function definition
type FuncType struct {
	named
	args  []Type
	value Type
}

func Func(name string, args []Type, returns Type) FuncType {
	return FuncType{
		named: named(name),
		args:  args,
		value: returns,
	}
}

func (f FuncType) Args() string {
	return argsString(f.args, len(f.args) > 3)
}

func (f FuncType) String() string {
	return f.value.String()
}

// function call
type CallType struct {
	named
	funcName string
	args     []Type
}

func (c CallType) String() string {
	args := argsString(c.args, len(c.args) > 3)
	return fmt.Sprintf("%s(%s)", c.funcName, args)
}

func Call(name, funcName string, args []Type) CallType {
	for k, v := range args {
		if v == nil {
			panic(fmt.Sprintf("argument `%v` in call to `%s` is nil", k, funcName))
		}
	}

	return CallType{
		named:    named(name),
		funcName: funcName,
		args:     args,
	}
}

// function arguments
func Args(args ...Type) []Type {
	return args
}

func argsString(m []Type, breakLine bool) string {
	s := ""
	for _, v := range m {
		if _, ok := v.(RequiredArgType); ok {
			s += fmt.Sprintf("%s, ", v.Name())
		} else {
			s += fmt.Sprintf("%s=%s, ", v.Name(), v.String())
		}

		if breakLine {
			s += "\n  "
		}
	}

	s = strings.TrimSuffix(s, "  ")
	s = strings.TrimSuffix(s, ", ")
	return s
}

type RequiredArgType struct {
	value Type
}

func (r RequiredArgType) Name() string {
	return r.value.Name()
}

func (r RequiredArgType) String() string {
	return r.value.String()
}

// mark an argument as required (no default value)
func Required(t Type) RequiredArgType {
	return RequiredArgType{t}
}

// CallChain allows to chain multiple calls
func CallChain(name string, calls ...CallType) CallType {
	if len(calls) == 1 {
		panic("callChain with a single call is redundant")
	}

	ln := ""
	if len(calls) > 1 {
		ln = "\n"
	}

	var last Type = Ref("", "")
	for i, c := range calls {
		last = Call("",
			strings.TrimPrefix(
				fmt.Sprintf("%s%s.%s", last.String(), ln, c.funcName),
				ln+".",
			),
			c.args,
		)

		if i == len(calls)-1 {
			l := last.(CallType)
			l.named = named(name)
			return l
		}
	}

	panic("loop did not return. This should never happen")
}
