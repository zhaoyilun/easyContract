package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Lex 用来进行分词
func Lex(filepath string) string {
	if !isExists(filepath) {
		panic("file does not exist")
	}

	src := readFile(filepath)
	var s Scanner
	fset := NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	var x string
	for {
		pos, tok, lit := s.Scan()
		// fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)
		tk := fmt.Sprintf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)
		if tok == EOF {
			break
		}
		x = x + tk
	}
	fmt.Println(x)
	return x
}

//IsExists  用于判断文件是否存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	return false
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
