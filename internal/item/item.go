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
	ReturnItem(*user.User) error
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
	//first, let's check to see if the item is already checked out
	if d.CheckedOut {
		//if so, return an error
		return errAlreadyOut
	}
	//if not, add it to the user's checked out items
	user.ItemsOut.Store(d, true)
	//mark it as checked out
	d.CheckedOut = true
	//no error, return nil
	return nil
}

//return changes the checked-out bool to false. It returns a non-nil
//error if the item is not currently checked out
func (d *dvd) ReturnItem(user *user.User) error {
	if !d.CheckedOut {
		return errNotCheckedOut
	}
	user.ItemsOut.Delete(d)
	d.CheckedOut = false
	return nil
}

//tapes are boring, we'll keep them to a minimum
type tape struct {
	Title      string `json:"title"`
	Runtime    string `json:"runtime"`
	CheckedOut bool   `json:"checked_out"`
}

func (t *tape) CheckOut(user *user.User) error {
	if t.CheckedOut {
		return errAlreadyOut
	}
	user.ItemsOut.Store(t, true)
	t.CheckedOut = true
	return nil
}

func (t *tape) ReturnItem(user *user.User) error {
	if !t.CheckedOut {
		return errNotCheckedOut
	}
	user.ItemsOut.Delete(t)
	t.CheckedOut = false
	return nil
}

type book struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	PublishingDate string `json:"publishing_date"`
	CheckedOut     bool   `json:"checked_out"`
}

func (b *book) CheckOut(user *user.User) error {
	if b.CheckedOut {
		return errAlreadyOut
	}
	user.ItemsOut.Store(b, true)
	b.CheckedOut = true
	return nil
}

func (b *book) ReturnItem(user *user.User) error {
	if !b.CheckedOut {
		return errAlreadyOut
	}
	user.ItemsOut.Delete(b)
	b.CheckedOut = false
	return nil
}
