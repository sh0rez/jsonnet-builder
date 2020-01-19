package builder

import (
	"fmt"
	"strings"
)

// Objects (dicts)
type ObjectType struct {
	named
	order    []string
	children map[string]Type
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

	return ObjectType{named: named(name), children: c, order: order}
}

func (o ObjectType) String() string {
	if len(o.children) == 0 {
		return "{}"
	}
	s := printChildren(o.children, o.order)
	return fmt.Sprintf("{\n%s\n}", indent(s))
}

func printChildren(children map[string]Type, order []string) string {
	s := ""
	for _, key := range order {
		switch t := children[key].(type) {
		case FuncType:
			s += fmt.Sprintf("%s(%s): %s,\n", t.Name(), t.Args(), t.String())
		case HiddenType:
			op := "::"

			switch h := t.value.(type) {
			case MergeType:
				op = "+::"
			case FuncType:
				s += fmt.Sprintf("%s(%s)%s %s,\n", t.Name(), h.Args(), op, t.String())
				continue
			}

			s += fmt.Sprintf("%s%s %s,\n", t.Name(), op, t.String())

		case LocalType:
			s += fmt.Sprintf("local %s = %s,\n", t.Name(), t.String())
		case MergeType:
			s += fmt.Sprintf("%s+: %s,\n", t.Name(), t.String())
		default:
			s += fmt.Sprintf("%s: %s,\n", t.Name(), t.String())
		}
	}
	s = strings.TrimSuffix(s, "\n")
	return s
}

// Lists (arrays)
type ListType struct {
	named
	items []Type
}

func List(name string, items ...Type) ListType {
	return ListType{named: named(name), items: items}
}

func (t ListType) String() string {
	s := ""
	for _, l := range t.items {
		s += fmt.Sprintf(", %s", l.String())
	}
	s = strings.TrimPrefix(s, ", ")
	return fmt.Sprintf("[%s]", s)
}
