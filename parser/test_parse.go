// package main
package parser

// import (
// 	"go/ast"
// 	"go/parser"
// 	"go/token"
// 	"log"
// )

// func main() {
// 	src := []byte(`
// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, world!")
// }
// `)

// 	fset := token.NewFileSet()

// 	file, err := parser.ParseFile(fset, "", src, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ast.Print(fset, file)
// }
