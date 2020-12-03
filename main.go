package main

import (
	"fmt"
)

func main() {
	fmt.Println(&File{
		Name: "my file",
		Imports: Imports{
			&Import{
				Path: "fmt",
			},
		},
		PackageName: "main",
		Functions: Functions{
			&Function{
				Name: "sum",
				Args: Args{
					&Arg{"x", "int"},
					&Arg{"y", "int"},
				},
				Output: []string{"int"},
				Body: Statements{
					&Return{Code("x+y")},
				},
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
