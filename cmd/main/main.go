package main

import (
	"log"
	"net/http"

	"github.com/Jonny-Burkholder/restful-api-example/internal/item"
	"github.com/Jonny-Burkholder/restful-api-example/internal/user"
)

//this is an example of how to implement a rest api that handles
//keeping track of library items. I'm using an item interface

func main() {

	http.Handle("/item", item.HandleItem(inventory))
	http.Handle("/user", user.HandleUser(inventory))

	log.Println("Serving on port 8080")
	http.ListenAndServe(":8080", nil)

}

//inventory will be a map with item type string as key, and as
//a value we will have a slice of items. In real life, I would
//store the items in a database. I'm using a map for simplicity's
//sake, rather than something thread safe, because honestly this
//is just an example and it will only have one person (me) using
//it at a time. Eventually I'll convert this to MongoDB, for fun
var inventory = make(map[string]*item.Item) //this will bite us in the butt if a book and movie have the same name. Need tables
var users = make(map[string]*user.User)

type dataStore struct {
	inventory map[string]*item.Item
	users     map[string]*user.User
}
