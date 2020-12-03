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
				Name: "haha",
				Args: Args{
					&Arg{"x", "int"},
					&Arg{"y", "int"},
				},
				Output: []string{"int"},
				Body: Statements{
					&IfElse{
						If: &If{
							Condition: "x<y",
							Body: Statements{
								&Return{Code("x")},
							},
						},
						Else: &Else{
							Body: Statements{
								&Return{Code("y")},
							},
						},
					},
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
