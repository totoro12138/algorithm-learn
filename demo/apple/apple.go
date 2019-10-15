package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
)

const (
	sanbox = "https://sandbox.itunes.apple.com/verifyReceipt"
	buy    = "https://buy.itunes.apple.com/verifyReceipt"
)

func main() {
	type Type struct {
		ProductId          string `json:"productId"`       //com.yiwei.youliao.money.16000",
		TransactionDate    int64  `json:"transactionDate"` // ":1570799870000,
		TransactionId      string `json:"transactionId"`   //":"1000000578253498",
		TransactionReceipt string `json:"transactionReceipt"`
	}
	bytes, err := ioutil.ReadFile("apple.json")
	if err != nil {
		panic(err)
	}
	var ty Type
	err = json.Unmarshal(bytes, &ty)
	if err != nil {
		panic(err)
	}
	type Type2 struct {
		ReceiptData string `json:"receipt-data"`
	}
	var ty2 Type2
	ty2.ReceiptData = ty.TransactionReceipt
	//fmt.Println(ty)
	req, err := httplib.Post(sanbox).
		JSONBody(ty2)
	if err != nil {
		panic(err)
	}
	str, err := req.String()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
