package main

import (
	"codesejin/golearn/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
