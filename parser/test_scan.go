package parser

import (
	"fmt"
)

func TestMain() {
	src := []byte(`
import asset

action create(ename issuer,easset maxAset){
	require_auth(get_self())
	esymbol sym = maxAset.symbol
	if not isSymbol(sym){
		panic("sym not valid")
	}
	if not isAsset(maxAset){
		panic("asset not valid")
	}
	if maxAset.amount<=0{
		panic("maxAset must be positive")
	}

	etable thisTable=table(myTable,get_self(),asset_to_index(maxAset))
	eindex ptr = thisTable.find(asset_to_index(maxAset))
	if ptr != nil{
		panic("token with symbol already exists")
	}

	mytable_json json = {
		"supply.symbol":maxAset.symbol,
		"max_supply":maxAset,
		"issuer":issuer
	}
	ptr.emplace(get_self(),json)
	}
action issue(ename to,easset quantity,string memo){
	esymbol sym = quantity.symbol
	if isSymbol(sym){
		panic("sym not valid")
	}
	if len(memo)>256{
		panic("memo has more than 256 bytes")
	}
	etable thisTable=table(myTable,get_self(),asset_to_index(quantity))
	eindex ptr = thisTable.find(asset_to_index(quantity))
	if ptr != nil{
		panic("token with symbol already exists")
	}
	if to != ptr.get("issuer"){
		panic("tokens can only be issued to issuer account")
	}

	require_auth(to)
	if not isAsset(quantity){
		panic("invalid quantity")
	}

	if quantity.amount <= 0{
		panic("must issue positive quantity")
	}

	if quantity.symbol != ptr.get("supply.symbol"){
		panic("symbol precision mismatch")
	}

	if quantity.amount > ptr.get("max_supply.amount") - ptr.get("supply.amount"){
		panic("quantity exceeds available supply")
	}
	ptr.modify("supply",ptr.get("supply")+quantity)
	add_balance(ptr.get("issuer"),quantity,ptr.get("issuer"))
}


table account{
	easset balance 
	index asset_to_index(balance)
}
table currency_stats{
	easset supply
	easset max_supply
	ename issuer
	index asset_to_index(balance)
}
createTable(account,"accounts")
createTable(currency_stats,"stat")
	`)
	var s Scanner
	fset := NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == EOF {
			break
		}
	}
}
