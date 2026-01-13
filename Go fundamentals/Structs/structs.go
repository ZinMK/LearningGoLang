package main

import (
	"LearningStructs/user"
	"fmt"
)

func main() {
	firstname := getUserInput("Enter Your First Name")
	lastname  := getUserInput("Enter Your Last Name")
	birthdate := getUserInput("Enter your birthday (DD/MM/YYYY)")

	var appUser *user.User

	appUser, err := user.NewUser(firstname,lastname,birthdate)
	
	if err != nil {
		return
	}

	appUser.ShowUserDetails()

	admin := user.NewAdmin("joe@gmail.com", "123455")
	admin.ShowUserDetails()
	admin.ChangeUserFirstName("newadmin")
	admin.ShowUserDetails()
}



func getUserInput(textinput string) (userinput string){
	fmt.Println(textinput)
	fmt.Scanln(&userinput)
	return userinput
}
