package main

import (
	"fmt"
	"strings"
)

type Code string

func (c Code) String() string {
	return string(c)
}

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
	Functions   Functions
}

func (f *File) String() string {
	return fmt.Sprintf("package %s\n\n%s\n\n%s\n\n%s\n\n%s",
		f.PackageName, f.Imports.String(), f.Defs.String(), f.Decls.String(), f.Functions.String())
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

type Statement fmt.Stringer
type Statements []Statement

func (s Statements) Join(sep string) string {
	var ss []string
	for _, statement := range s {
		ss = append(ss, statement.String())
	}
	return strings.Join(ss, sep)
}

type Arg struct {
	Name string
	Type string
}

func (a *Arg) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

type Args []*Arg

func (a Args) String() string {
	var as []string
	for _, arg := range a {
		as = append(as, arg.String())
	}
	return strings.Join(as, ", ")
}

type Function struct {
	Name     string
	Args     Args
	Receiver string
	Output   []string
	Body     Statements
}

func (f *Function) String() string {
	out := "func "
	if f.Receiver != "" {
		out += fmt.Sprintf("(%s) ", f.Receiver)
	}
	out += fmt.Sprintf("%s(%s) ", f.Name, f.Args.String())
	if f.Output != nil {
		out += fmt.Sprintf("(%s) ", strings.Join(f.Output, ","))
	}
	out += "{\n"
	out += f.Body.Join("\n\t")
	out += "\n}"
	return out
}

type Functions []*Function

func (f Functions) String() string {
	var fs []string
	for _, f := range f {
		fs = append(fs, f.String())
	}
	return strings.Join(fs, "\n")
}

type If struct {
	isElse    bool
	Condition string
	Body      Statements
}

func (i *If) String() string {
	return fmt.Sprintf("if %s {\n%s\n}", i.Condition, i.Body.Join("\n"))
}

type Else struct {
	Body Statements
}

func (e *Else) String() string {
	return fmt.Sprintf("else {\n%s\n}", e.Body.Join("\n"))
}

type For struct {
	Condition string
	Body      Statements
}

func (f *For) String() string {
	return fmt.Sprintf("for %s {\n%s\n}", f.Condition, f.Body.Join("\n"))
}

type Call struct {
	Name string
	Args Statements
}

func (c *Call) String() string {
	return fmt.Sprintf("%s(%s)", c.Name, c.Args.Join(", "))
}

type Return struct {
	Value Statement
}

func (r *Return) String() string {
	return fmt.Sprintf("return %s", r.Value)
}

type Goroutine struct {
	Call Call
}

func (g *Goroutine) String() string {
	return fmt.Sprintf("go %s", g.Call)
}
