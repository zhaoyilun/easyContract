package eos

import (
	"errors"
)

const (
	include    string = "#include <eosio/eosio.hpp>"
	namespace  string = "using namespace eosio"
	classText1 string = "class [[eosio::contract(\""
	classText2 string = "\")]] "
	classText3 string = " : public eosio::contract"
	funcText0  string = "[[eosio::action]] void "
	account    string = "eosio.token"
)

const (
	typeAction string = "[[eosio::action]]"
	typeNotify string = "[[eosio::on_notify(\"eosio.token::transfer\")]]"
	typeTable  string = "[[eosio::table]]"
)

func newClass(name string) string {
	return classText1 + name + classText2 + name + classText3
}

func newNotify(account string) string {
	return "[[eosio::on_notify(\"" + account + "::transfer\")]]"
}

func newFunctionTitle(types, returnDefine string, variables []variable) (function, error) {
	var _function = function{}
	if types != typeAction && types != typeNotify && types != typeTable {
		return _function, errors.New("Function type error")
	}
	if returnDefine != "int" && returnDefine != "void" {
		return _function, errors.New("returnDefine  error")
	}
	_function = function{
		types:        types,
		returnDefine: returnDefine,
		variables:    variables,
		text:         "{}",
	}
	return _function, nil
}

//newVariable 创建新变量
func newVariable(name, types string) variable {
	var _variable = variable{
		name:  name,
		types: types,
	}
	return _variable
}

func (fc *function) toString() string {
	inside := "("
	for i, v := range fc.variables {
		if i != 0 {
			inside = inside + ","
		}
		inside = inside + v.types + " " + v.name
	}
	inside = inside + ")"
	return fc.types + fc.returnDefine + inside + fc.text
}

type function struct {
	types        string
	returnDefine string
	variables    []variable
	text         string
}

type variable struct {
	name  string
	types string
}

type table struct {
}

type cppStruct struct {
	include   string
	namespace string
	name      string

	functions []function
	tables    []table
}

//目前仅为测试
func (c *cppStruct) newCpp() string {
	var result string = ""
	result = c.include + "\n" + c.namespace + "\n"
	classDefine := newClass(c.name)
	result = result + classDefine + "{\n"
	for _, f := range c.functions {
		result = result + f.toString() + "\n"
	}
	result = result + "};"
	return result
}
