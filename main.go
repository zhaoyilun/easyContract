/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

// import "github.com/zhaoyilun/easyContract/cmd"
import (
	"fmt"
	"github.com/zhaoyilun/easyContract/parser"
)

// import "os"
// import "fmt"

func main() {
	w := parser.Lex("./examples/eidos.ez")
	fmt.Println(w)

	// file2, error := os.OpenFile("./2.txt", os.O_RDWR|os.O_CREATE, 0766)
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// file2.Write([]byte(w))
	// file2.Close()

	// Dada()
	// cmd.Execute()
}
