package main

import (
	"fmt"
	"github.com/zhaoyilun/easyContract/ast"
	"github.com/zhaoyilun/easyContract/parser"
	"github.com/zhaoyilun/easyContract/token"
	"io/ioutil"
	"os"
)

type astNode struct {
}

// func (node *astNode) Pos() token.Pos {}
// func (node *astNode) End() token.Pos {}

var a = astNode{}

func Dada() {
	// src is the input for which we want to print the AST.
	// src := `
	// package main
	// func main() {
	// 	println("Hello, World!")
	// }
	// `
	// w := parser.Lex("./examples/eidos.ez")
	// w := parser.Lex("./examples/eidos.ez")
	src := readFile("./examples/eidos.ez")
	// fmt.Println(w)
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	// f, err := parser.ParseFile(fset, "", []byte(w), 0)
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
func readFile(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("read file fail")
		panic(err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail")
		panic(err)
	}
	return fd
}
