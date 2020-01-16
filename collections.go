package main

import (
	"fmt"
	"strings"
)

// Objects (dicts)
type object struct {
	named
	order    []string
	children map[string]Type
}

func Object(name string, children ...Type) object {
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

	return object{named: named(name), children: c, order: order}
}

func (o object) String() string {
	s := printChildren(o.children, o.order)
	return fmt.Sprintf("{\n%s\n}", indent(s))
}

func printChildren(children map[string]Type, order []string) string {
	s := ""
	for _, key := range order {
		switch t := children[key].(type) {
		case funcType:
			s += fmt.Sprintf("%s(%s): %s,\n", t.Name(), t.Args(), t.String())
		case hiddenType:
			op := "::"

			switch h := t.value.(type) {
			case mergeType:
				op = "+::"
			case funcType:
				s += fmt.Sprintf("%s(%s)%s %s,\n", t.Name(), h.Args(), op, t.String())
				continue
			}

			s += fmt.Sprintf("%s%s %s,\n", t.Name(), op, t.String())

		case localType:
			s += fmt.Sprintf("local %s = %s,\n", t.Name(), t.String())
		case mergeType:
			s += fmt.Sprintf("%s+: %s,\n", t.Name(), t.String())
		default:
			s += fmt.Sprintf("%s: %s,\n", t.Name(), t.String())
		}
	}
	s = strings.TrimSuffix(s, "\n")
	return s
}

// Lists (arrays)
type listType struct {
	named
	items []Type
}

func List(name string, items ...Type) listType {
	return listType{named: named(name), items: items}
}

func (t listType) String() string {
	s := ""
	for _, l := range t.items {
		s += fmt.Sprintf(", %s", l.String())
	}
	s = strings.TrimPrefix(s, ", ")
	return fmt.Sprintf("[%s]", s)
}
