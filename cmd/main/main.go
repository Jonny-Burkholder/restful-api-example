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

	mux := http.NewServeMux()

	mux.Handle("/api/return-book/", item.ReturnItem(ds))
	mux.Handle("/api/return-dvd/", item.ReturnItem(ds))
	mux.Handle("/api/return-tape/", item.ReturnItem(ds))
	mux.Handle("/api/donate-item/", item.DonateItem(ds)) //should be able to tell what this is via reflection, and I don't feel like writing all the handlers out
	mux.Handle("/api/dvds", item.GetDVDs(ds))
	mux.Handle("/api/tapes", item.GetTapes(ds))
	mux.Handle("/api/books", item.GetBooks(ds))
	mux.Handle("/api/users", user.HandleUser(ds))

	log.Println("Serving on port 8080")
	http.ListenAndServe(":8080", mux)

}

var ds = data.NewDataStore()

var generator = item.NewGenerator()

//just until we get an actual database in this business
func makeItems() {
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Moby Dick", "Hermin Melville", "Boring", "1851"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Gardens of the Moon", "Steven Erikson", "Fantasy", "1999"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Deadhouse Gates", "Steven Erikson", "Fantasy", "2000"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Memories of Ice", "Steven Erikson", "Mind-Melter", "2001"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Words of Radiance", "Brandon Sanderson", "Science Fantasy", "2014"))
	ds.Inventory["book"] = append(ds.Inventory["book"], item.NewBook(generator.GetID("book"), "Bands of Mourning", "Brandon Sanderson", "Western Science Fantasy", "2016"))
	ds.Inventory["tape"] = append(ds.Inventory["tape"], item.NewTape(generator.GetID("tape"), "Adventures in Odyssey", "109 minutes"))
	ds.Inventory["tape"] = append(ds.Inventory["tape"], item.NewTape(generator.GetID("tape"), "How to Make Friends and Influence People", "12 minutes"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD(generator.GetID("dvd"), "The Matrix", "Action", "Unclear", "1999", "2-ish hours"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD(generator.GetID("dvd"), "Infinity War", "Marvel", "PG-13", "2018", "149 minutes"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD(generator.GetID("dvd"), "Milltown Pride", "Baseball", "G", "2011", "Literally forever"))
	ds.Inventory["dvd"] = append(ds.Inventory["dvd"], item.NewDVD(generator.GetID("dvd"), "Mission: Impossible", "Action", "PG-13", "1996", "110 minutes"))
}
