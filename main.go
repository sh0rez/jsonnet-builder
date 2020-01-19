package builder

import (
	"fmt"
	"strings"
)

type named string

func (n named) Name() string {
	return string(n)
}

type Doc struct {
	Locals []LocalType
	Root   Type
}

func (d Doc) String() string {
	s := ""
	for _, l := range d.Locals {
		s += fmt.Sprintf("local %s = %s;\n", l.Name(), l.String())
	}

	s += d.Root.String()
	return s
}

type Type interface {
	String() string
	Name() string
}

// func main() {
// 	o := Object("",
// 		Hidden(Object("deployment",
// 			Func("new",
// 				Args(Required(String("name", "")), Int("retain", 3)),
// 				Object("", Ref("name", "name")),
// 			),
// 		)),
// 		Call("deploy", "$.deployment.new", Args(String("name", "hi"))),
// 	)

// 	d := Doc{
// 		Root: o,
// 		Locals: []localType{
// 			Local(String("nice", "ho")),
// 		},
// 	}

// 	fmt.Println(d.String())
// }

func indent(s string) string {
	split := strings.Split(s, "\n")
	for i := range split {
		split[i] = "  " + split[i]
	}
	return strings.Join(split, "\n")
}

func dedent(s string) string {
	split := strings.Split(s, "\n")
	for i := range split {
		split[i] = strings.TrimPrefix(split[i], "  ")
	}
	return strings.Join(split, "\n")
}
