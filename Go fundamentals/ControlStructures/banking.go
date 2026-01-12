package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const account_balance_file = "balance.txt"

func write_balance_to_file(balance float64){
	balanceText:= fmt.Sprint(balance)
	os.WriteFile(account_balance_file, []byte(balanceText), 0644)

}

func get_balance_from_file() (balance float64, err error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("Could not read from File")
	}
	balance_text := string(data)
	balance, err = strconv.ParseFloat(balance_text, 64) //convert from float64 to string
	if err != nil {
		return 1000, errors.New("Could not read from File")
	}
	err = nil
	return 
}

func main() {

	account_Balance, err  := get_balance_from_file()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err) 
		return
	}
 
	
	for {
		
		fmt.Printf(`
		Welcome to Zin's Bank In Go
	    1.Check Balance
		2.Deposit Money
		3.Withdraw Money
		4.Exit
		`)

		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Printf("Your Balance is: $%v", account_Balance)
				
			case 2:
				var deposit float64
				fmt.Print("Enter the amount you want to deposit: ")
				fmt.Scan(&deposit)
				account_Balance = account_Balance+deposit
				write_balance_to_file(account_Balance) 
				fmt.Printf("Your new Balance is: $%f", account_Balance)	
			case 3:
				var withdraw float64
				fmt.Print("Enter the amount you want to withdraw: ")
				fmt.Scan(&withdraw)
				if withdraw>account_Balance {
					fmt.Println("You broke as hell for that")
				}
				account_Balance = account_Balance-withdraw
				write_balance_to_file(account_Balance)
				fmt.Printf("Your new Balance is: $%f", account_Balance)
		
			default:
				return
		}
		
	} 
	

}