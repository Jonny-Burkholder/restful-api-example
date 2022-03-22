package item

import (
	"errors"

	"github.com/Jonny-Burkholder/restful-api-example/internal/user"
)

var errAlreadyOut = errors.New("Item is already checked out")
var errNotCheckedOut = errors.New("Item has not been checked out")

//item is an interface that allows us to stock multiple types of items
//in the library. In our case, we will be stocking books, dvds, and
//tapes. Note that this is not how I would implement this in real life,
//but it makes a convenient example of how to implement the interface
//in an API
type Item interface {
	CheckOut(*user.User) error
	ReturnItem() error
}

type dvd struct {
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	CheckedOut  bool   `json:"checked_out"`
}

//checkOut changes the checkedout bool to true. It returns a non-nil
//error if the item is already checked out
func (d *dvd) CheckOut(user *user.User) error {
	if d.CheckedOut == true {
		return errAlreadyOut
	}
	d.CheckedOut = true
	return nil
}

//return changes the checked-out bool to false. It returns a non-nil
//error if the item is not currently checked out
func (d *dvd) ReturnItem() error {
	if d.CheckedOut == false {
		return errNotCheckedOut
	}
	d.CheckedOut = false
	return nil
}

type tape struct {
	Title   string `json:"title"`
	Runtime string `json:"runtime"`
}

type book struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	PublishingDate string `json:"publishing_date"`
}
