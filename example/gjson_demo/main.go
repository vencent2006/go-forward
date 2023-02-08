package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	getNotify()
}

func getNotify() {
	json := `{
		"tokenAddress": "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		"address": "TQceSn8A1EAvACYE8wBSWZnuMkz91NrxFn",
		"tokenSymbol": "USDT",
		"txid": "19ade82ad55d6ddb643f3643474b281772744f67325a0c4e3cededba20e583b0",
		"time": 1675682847,
		"confirmations": 1,
		"tokenValue": "-22921.636103",
		"value": "0",
		"coin": "TRX",
		"height": 48359865
	}`

	addr := gjson.Get(json, "tokenAddress")
	fmt.Println("addr:", addr.String())
	time := gjson.Get(json, "time")
	fmt.Println("time:", time.Int())
	tokenValue := gjson.Get(json, "tokenValue")
	fmt.Println("tokenValue:", tokenValue.Int())
	fmt.Println("tokenValue:", tokenValue.Value())
	kkk := gjson.Get(json, "kkk")
	if kkk.String() == "" {
		fmt.Println("kkk is empty string")
	}
}

func demo() {
	json := `{"name":{"first":"han","last":"tuo"},"age":22}`
	lastName := gjson.Get(json, "name.last")
	fmt.Println("last name:", lastName.String())

	age := gjson.Get(json, "age")
	fmt.Println("age:", age.Int())
	fmt.Println("age:", age.String())
	last := gjson.Get(json, "name.last")
	fmt.Println("age:", last.String())
	fmt.Println("age:", last.Int())
	kkk := gjson.Get(json, "name.last.kkk")
	fmt.Println("kkk:", kkk.String())
	fmt.Println("kkk:", kkk.Int())
	fmt.Printf("age:%+v", kkk)
}
