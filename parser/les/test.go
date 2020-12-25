package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	p := parser{
		text: "import",
		i:    1,
		word: "sadf",
		step: "asdf",
	}
	p.Parse()
}

type parser struct {
	text string
	i    int
	word string
	step string
}

func (p *parser) Parse() (string, error) {
	switch p.peek() {
	case "import":
		fmt.Println("import")
	case "action":
		fmt.Println("action")
	default:
		return "", nil
	}
	return "", nil
}

func (p *parser) peek() string {

	return "import"
}

func (p *parser) pop() string {
	return "import"
}
