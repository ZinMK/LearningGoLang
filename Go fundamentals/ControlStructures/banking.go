package main

import (
	"banking/fileOps"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const account_balance_file = "balance.txt"
 
func main() {

	account_Balance, err  := fileOps.Get_Float_from_file(account_balance_file)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err) 
		return
	}
 
	
	for {
		
		presentOptions()
		fmt.Println(randomdata.PhoneNumber(), "<- Call us if you get hacked")

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
				fileOps.Write_Float_to_file(account_Balance,account_balance_file) 
				fmt.Printf("Your new Balance is: $%f", account_Balance)	
			case 3:
				var withdraw float64
				fmt.Print("Enter the amount you want to withdraw: ")
				fmt.Scan(&withdraw)
				if withdraw>account_Balance {
					fmt.Println("You broke as hell for that")
				}
				account_Balance = account_Balance-withdraw
				fileOps.Write_Float_to_file(account_Balance,account_balance_file) 
				fmt.Printf("Your new Balance is: $%f", account_Balance)
		
			default:
				return
		}
		
	} 
	

}