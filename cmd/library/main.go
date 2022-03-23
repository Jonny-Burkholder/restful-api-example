package main

import "github.com/Jonny-Burkholder/restful-api-example/internal/item"

//this is an example of how to implement a rest api that handles
//keeping track of library items. I'm using an item interface

func main() {

}

//inventory will be a map with item type string as key, and as
//a value we will have a slice of items. In real life, I would
//store the items in a database
var inventory = make(map[string][]*item.Item)
