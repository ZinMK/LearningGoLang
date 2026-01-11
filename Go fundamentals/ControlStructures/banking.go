package main

import "fmt"


func main() {

	var account_Balance float64 = 1000

	
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
				fmt.Printf("Your new Balance is: $%f", account_Balance)
				
			case 3:
				var withdraw float64
				fmt.Print("Enter the amount you want to withdraw: ")
				fmt.Scan(&withdraw)
				if withdraw>account_Balance {
					fmt.Println("You broke as hell for that")
				}
				account_Balance = account_Balance-withdraw
				fmt.Printf("Your new Balance is: $%f", account_Balance)
		
			default:
				return
		}
		
	} 
	

}