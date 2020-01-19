package builder

import (
	"fmt"
	"strings"
)

const (
	SeparatorLong    = "\n"
	SeparatorConcise = " "
)

// Objects (dicts)
type ObjectType struct {
	named
	order    []string
	children map[string]Type
	concise  bool
}

func Object(name string, children ...Type) ObjectType {
	c := make(map[string]Type)
	order := make([]string, len(children))
	i := 0
	for _, child := range children {
		if v, ok := c[child.Name()]; ok {
			panic(fmt.Sprintf("key clash: trying to add `%v` as `%s`, but `%v` already uses this key", child, child.Name(), v))
		}

		c[child.Name()] = child
		order[i] = child.Name()
		i++
	}

	return ObjectType{
		named:    named(name),
		children: c,
		order:    order,
	}
}

func ConciseObject(name string, children ...Type) ObjectType {
	o := Object(name, children...)
	o.concise = true
	return o
}

func (o ObjectType) String() string {
	if len(o.children) == 0 {
		return "{}"
	}
	if o.concise {
		return o.ConciseString()
	}

	s := printChildren(o.children, o.order, SeparatorLong)
	return fmt.Sprintf("{\n%s\n}", indent(s))
}

func (o ObjectType) ConciseString() string {
	s := printChildren(o.children, o.order, SeparatorConcise)
	return fmt.Sprintf("{ %s }", strings.TrimSuffix(s, ","))
}

func printChildren(children map[string]Type, order []string, s string) string {
	j := ""
	for _, key := range order {
		switch t := children[key].(type) {
		case FuncType:
			j += fmt.Sprintf("%s(%s): %s,"+s, t.Name(), t.Args(), t.String())
		case HiddenType:
			op := "::"

			switch h := t.value.(type) {
			case MergeType:
				op = "+::"
			case FuncType:
				j += fmt.Sprintf("%s(%s)%s %s,"+s, t.Name(), h.Args(), op, t.String())
				continue
			}

			j += fmt.Sprintf("%s%s %s,"+s, t.Name(), op, t.String())

		case LocalType:
			j += fmt.Sprintf("local %s = %s,"+s, t.Name(), t.String())
		case MergeType:
			j += fmt.Sprintf("%s+: %s,"+s, t.Name(), t.String())
		default:
			j += fmt.Sprintf("%s: %s,"+s, t.Name(), t.String())
		}
	}
	j = strings.TrimSuffix(j, s)
	return j
}
