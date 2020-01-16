package main

import "fmt"

type importType struct {
	named
	pkg string
	raw bool
}

func (i importType) String() string {
	op := "import"
	if i.raw {
		op += "str"
	}

	return fmt.Sprintf(`(%s "%s")`, op, i.pkg)
}

func Import(name, pkg string) importType {
	return importType{named: named(name), pkg: pkg, raw: false}
}

func ImportStr(name, pkg string) importType {
	return importType{named: named(name), pkg: pkg, raw: true}
}
