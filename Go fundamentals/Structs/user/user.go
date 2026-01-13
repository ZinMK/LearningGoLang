package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u User) ShowUserDetails() {
fmt.Println(u.firstName,u.lastName,u.birthDate)
}

func (u *User) ChangeUserFirstName(newName string) {
(*u).firstName = newName
}


func (u *User) ChangeUserLastName(newName string) {
(*u).firstName = newName
}

//Inheritence
type Admin struct {
	email string
	password string
	User // This value will ensure that it is inherited
}

func NewAdmin(new_email, new_password string) Admin {
	return Admin{	
		email: new_email,
		password: new_password,
		User: User{
			firstName: "joe",
			lastName: "biden",
			birthDate: "09/89/1992",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstname, lastname, birthdate string) (*User, error) {

	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("The values are empty")
	}

	return &User{
		firstName: firstname,
		lastName: lastname,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

