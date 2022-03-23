package user

import (
	"fmt"
	"sync"
)

//We'll keep this nice and simple
type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	ItemsOut  sync.Map `json:"items_out"`
}

//NewUser takes a first name, last name, and email,
//and returns a pointer to a user
func NewUser(first, last, email string) *User {
	return &User{
		FirstName: first,
		LastName:  last,
		Email:     email,
		ItemsOut:  sync.Map{},
	}
}

//EmailUser emails a user to remind them that they need to return a book
func (u *User) EmailUser() {
	fmt.Println("Return your item!")
}
