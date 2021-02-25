package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Account struct{

}

func (a*Account)CreteNewAccount() string  {
	return "You have access successfully"
}

var check *Account

func CheckAccount()*Account{
	if check==nil {
		once.Do(func() {
			fmt.Println("You have a new account")
			check=&Account{}
		})

	}else {
		fmt.Println("Sorry,this is exist")
	}
	return check

}
func main()  {
	CheckAccount()
	CheckAccount()
	CheckAccount()

}