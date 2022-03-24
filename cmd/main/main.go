package main

import (
	"log"
	"net/http"

	"github.com/Jonny-Burkholder/restful-api-example/internal/data"
	"github.com/Jonny-Burkholder/restful-api-example/internal/item"
	"github.com/Jonny-Burkholder/restful-api-example/internal/user"
)

//this is an example of how to implement a rest api that handles
//keeping track of library items. I'm using an item interface

func main() {

	makeItems()
	ds.Users["admin"] = user.NewUser("Bob", "Jones", "therealbobjones@bju.edu")

	http.Handle("/get-item", item.GetItem(ds))
	http.Handle("/return-item", item.ReturnItem(ds))
	http.Handle("/donate-item", item.DonateItem(ds))
	http.Handle("/dvd", item.HandleDVD(ds))
	http.Handle("/tape", item.HandleTape(ds))
	http.Handle("/book", item.HandleBook(ds))
	http.Handle("/user", user.HandleUser(ds))

	log.Println("Serving on port 8080")
	http.ListenAndServe(":8080", nil)

}

//inventory will be a map with item type string as key, and as
//a value we will have a slice of items. In real life, I would
//store the items in a database. I'm using a map for simplicity's
//sake, rather than something thread safe, because honestly this
//is just an example and it will only have one person (me) using
//it at a time. Eventually I'll convert this to MongoDB, for fun
var inventory = make(map[string][]interface{}) //this is not ideal. We really need tables
var users = make(map[string]*user.User)

var ds = data.DataStore{
	Inventory: inventory,
	Users:     users,
}

//just until we get an actual database in this business
func makeItems() {
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook("Moby Dick", "Hermin Melville", "Boring", "1851"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook("Gardens of the Moon", "Steven Erikson", "Fantasy", "1999"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook("Words of Radiance", "Brandon Sanderson", "Science Fantasy", "2014"))
	ds.Inventory["tape"] = append(ds.Inventory["tape"], item.NewTape("Adventures in Odyssey", "109 minutes"))
	ds.Inventory["tape"] = append(ds.Inventory["tape"], item.NewTape("How to Make Friends and Influence People", "12 minutes"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD("The Matrix", "Action", "Unclear", "1999", "2-ish hours"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD("Infinity War", "Marvel", "PG-13", "2018", "149 minutes"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD("Milltown Pride", "Baseball", "G", "2011", "Literally forever"))
}
