package eos

import (
	"fmt"
	"testing"
)

func TestNewVariable(t *testing.T) {
	va := newVariable("asdf", "asdf")
	fmt.Println(va)
}

func TestFunction(t *testing.T) {
	va := newVariable("abcd", "int")
	fmt.Println(va)
	vb := newVariable("efgh", "int")
	fmt.Println(vb)
	d := []variable{va, vb}

	c, err := newFunctionTitle(typeAction, "int", d)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(c)
	fmt.Println(c.toString())
}

func TestNewCpp(t *testing.T) {
	va := newVariable("abcd", "int")

	vb := newVariable("efgh", "int")

	d := []variable{va, vb}
	c, _ := newFunctionTitle(typeAction, "int", d)

	ve := newVariable("a2323", "int")

	vf := newVariable("e2323h", "int")

	g := []variable{ve, vf}

	h, _ := newFunctionTitle(typeNotify, "void", g)

	i := []function{c, h}
	cpp := cppStruct{
		include:   include,
		namespace: namespace,
		name:      "testing",
		functions: i,
	}
	j := cpp.newCpp()
	fmt.Println(j)
}
