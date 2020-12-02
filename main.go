package main

import (
	"fmt"
)

type Function struct {
	Statements []fmt.Stringer
}

func main() {
	fmt.Println(&File{
		Name: "my file",
		Imports: Imports{
			&Import{
				Path: "fmt",
			},
		},
		PackageName: "main",
		Defs: Defs{
			&Def{
				Name:  "name",
				Value: "Amirreza",
			},
		},
		Decls: Decls{
			&Decl{
				Name:      "String",
				Signature: "string",
			},
		},
	})
}
