package main

import (
	"fmt"
	"strings"
)

type Module struct {
	Name     string
	Packages []*Package
}

type Package struct {
	Name  string
	Files []*File
}
type Import struct {
	Name string
	Path string
}

type Imports []*Import

func (is Imports) String() string {
	var imports []string
	for _, i := range is {
		imports = append(imports, i.String())
	}
	return fmt.Sprintf("import(\n%s\n)", strings.Join(imports, "\n"))
}

func (i *Import) String() string {
	out := "\t"
	if i.Name != "" {
		out += i.Name + " "
	}
	out += fmt.Sprintf("\"%s\"", i.Path)
	return out
}

type File struct {
	Name        string
	PackageName string
	Imports     Imports
	Defs        Defs
	Decls       Decls
}

func (f *File) String() string {
	return fmt.Sprintf("package %s\n\n%s\n\n%s\n\n%s", f.PackageName, f.Imports.String(), f.Defs.String(), f.Decls.String())
}

type Decl struct {
	Name      string
	Signature string
}

func (d *Decl) String() string {
	return fmt.Sprintf("type %s %s", d.Name, d.Signature)
}

type Decls []*Decl

func (d Decls) String() string {
	var ds []string
	for _, decl := range d {
		ds = append(ds, decl.String())
	}
	return strings.Join(ds, "\n\n")
}

type Def struct {
	Name        string
	IsDefAsWell bool //:= or =
	Value       string
}
type Defs []*Def

func (d Defs) String() string {
	var ds []string
	for _, Def := range d {
		ds = append(ds, Def.String())
	}
	return strings.Join(ds, "\n\n")
}

func (d *Def) String() string {
	sign := "="
	if d.IsDefAsWell {
		sign = ":="
	}
	return fmt.Sprintf("%s %s %s", d.Name, sign, d.Value)
}
