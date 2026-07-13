package main

import (
	"fmt"
	"os"
)

type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type BankAccount interface {
	WithdrawMoney (amount float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println(obj.Name)
	fmt.Println(obj.Age)
	fmt.Println(obj.Money)
}

func (obj user) ReceiveMoney(amount float64) float64 {
	obj.Money = obj.Money + amount
	return obj.Money
}

func (obj user) WithdrawMoney (amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

func main() {
	var usr1 People
	usr1 = user{
		Name:  "Abdullah Al Towsif",
		Age:   24,
		Money: 10.08,
	}

	usr2 := user{
		Name:  "Habibur Rahman",
		Age:   30,
		Money: 10.08,
	}

	var usr3 BankAccount
	usr3 = user{
		Name: "Hello",
		Age: 34,
		Money: 8464,
	}
	usr3.WithdrawMoney(125)

	obj, ok := usr3.(user)
	if !ok {
		fmt.Println("Sorry usr3 is not type of user struct")
		os.Exit(1)
	}
	obj.PrintDetails()
	obj.ReceiveMoney(54)

	usr1.PrintDetails()
	usr2.PrintDetails()
}
