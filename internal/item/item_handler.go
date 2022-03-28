package item

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Jonny-Burkholder/restful-api-example/internal/data"
)

func ReturnItem(ds *data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//return item
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func CheckOutItem(ds *data.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check item out
	})
}

//DonateItem is here so that there's an example of using
//a post request
func DonateItem(ds *data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//donate item
			//I'll probably never actually implement this, since it's kind of bad practice,
			//although I think it would be really fun
		} else {
			http.Error(w, "This method is not supported for this endpoint", http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

func GetDVDs(ds *data.DataStore, g *generator) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint reached: dvds\nIP:%q\ntime:%v", r.RemoteAddr, time.Now().Unix())
		if r.Method == http.MethodGet {
			queries := r.URL.Query()
			if len(queries) == 0 {
				//return all DVDs
				//This doesn't really account for the pattern api/dvds/0010000000X
				json.NewEncoder(w).Encode(ds.Inventory["dvd"])
				return
			} else {
				//check that all queries are valid, and that each key only has exactly one value
				log.Println(queries) //you know, for debugging
				for key, value := range queries {
					if _, ok := dvdParams[strings.ToLower(strings.TrimSpace(key))]; ok != true || len(value) != 1 {
						http.Error(w, "Invalid request", http.StatusBadRequest)
						return
					}
				}
				//store the results that match this query
				results := make([]*dvd, 0)
				//range over the dvd inventory
				for _, v := range ds.Inventory["dvd"] {
					item := v.(*dvd)
					//if the field meets the criteria, add will be true
					add := true
					for key, value := range queries {
						//if any of the criteria don't match, just skip this dvd
						var field string
						switch strings.ToLower(strings.TrimSpace(key)) {
						case "id":
							field = item.ID
						case "title":
							field = item.Title
						case "genre":
							field = item.Genre
						case "rating":
							field = item.Rating
						case "releasedate":
							field = item.ReleaseDate
						case "runtime":
							field = item.Runtime
						case "checkedout":
							if item.CheckedOut == true {
								field = "true"
							} else {
								field = "false"
							}
						case "checkedoutby":
							field = item.CheckedOutBy
						}
						if strings.ToLower(strings.TrimSpace(field)) != value[0] {
							log.Printf("Field: %v, Value: %v", strings.ToLower(strings.TrimSpace(field)), value[0])
							add = false
						}
					}
					//if all the criteria matched, add this dvd to the results
					if !add {
						//do nothing
					} else {
						results = append(results, item)
					}
				}
				//if there are no results, send a 204 error
				if len(results) < 1 {
					http.Error(w, "No results found", http.StatusNoContent)
					return
				}
				//write the results to the response writer
				json.NewEncoder(w).Encode(results)
				return
			}
		} else if r.Method == http.MethodPost {
			//handle post request
			var disk *dvd
			err := json.NewDecoder(r.Body).Decode(disk)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}
			disk.ID = g.GetID("dvd")
			//Should probably check for duplicates first, but I don't think it matters much
			ds.Inventory["dvd"] = append(ds.Inventory["dvd"], disk)
			log.Println("Post request handled successfully")
			log.Println(ds.Inventory["dvd"]...)
		} else {
			http.Error(w, "This method is not supported", http.StatusNotImplemented)
			return
		}
	}
	return http.HandlerFunc(fn)
}

//GetDVDByID does what it says on the tin. This handles the /api/dvds/ url
func GetDVDByID(ds *data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint reached: DVD by ID")
		if r.Method != http.MethodGet {
			http.Error(w, "This method is not supported for this endpoint", http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Path[len("/api/dvds/"):]
		if len(id) < 1 {
			json.NewEncoder(w).Encode(ds.Inventory["dvd"])
			return
		}
		log.Printf("DVD by ID requested: %s\n", id)
		res := &dvd{}
		for _, v := range ds.Inventory["dvd"] {
			disk := v.(*dvd)
			if disk.ID == id {
				res = disk
				break
			}
		}
		json.NewEncoder(w).Encode(res)
	}
	return http.HandlerFunc(fn)
}

func GetTapes(ds *data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint reached: tapes\nIP:%q\ntime:%v", r.RemoteAddr, time.Now().Unix())
		if r.Method == http.MethodGet {
			queries := r.URL.Query()
			if len(queries) == 0 {
				//if there are no specific queries, return all tapes
				json.NewEncoder(w).Encode(ds.Inventory["tape"])
				return
			} else {
				//check that all queries are valid, and that each key only has exactly one value
				for key, value := range queries {
					if _, ok := tapeParams[strings.ToLower(strings.TrimSpace(key))]; ok != true || len(value) != 1 {
						http.Error(w, "Invalid request", http.StatusBadRequest)
						return
					}
				}
				//store the results that match this query
				results := make([]*tape, 0)
				//range over the dvd inventory
				for _, v := range ds.Inventory["tape"] {
					item := v.(*tape)
				filter:
					for key, value := range queries {
						//if any of the criteria don't match, just skip this dvd
						var field string
						switch strings.ToLower(strings.TrimSpace(key)) {
						case "id":
							field = item.ID
						case "title":
							field = item.Title
						case "runtime":
							field = item.Runtime
						case "checkedout":
							if item.CheckedOut == true {
								field = "true"
							} else {
								field = "false"
							}
						case "checkedoutby":
							field = item.CheckedOutBy
						}
						if field != value[0] {
							break filter
						}
					}
					//if all the criteria matched, add this dvd to the results
					results = append(results, item)
				}
				//write the results to the response writer
				json.NewEncoder(w).Encode(results)
				return
			}
		} else {
			http.Error(w, "This method is not supported", http.StatusNotImplemented)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func GetBooks(ds *data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Printf("Endpoint reached: books\nIP:%q\ntime:%v", r.RemoteAddr, time.Now().Unix())
			queries := r.URL.Query()
			if len(queries) == 0 {
				//return all books if there is no specific request.
				//heh, this would be massive amounts of data for a
				//real city library. Probably would not have this,
				//or would throttle the connection
				json.NewEncoder(w).Encode(ds.Inventory["book"])
				return
			} else {
				//check that all queries are valid, and that each key only has exactly one value
				for key, value := range queries {
					if _, ok := bookParams[strings.ToLower(strings.TrimSpace(key))]; ok != true || len(value) != 1 {
						http.Error(w, "Invalid request", http.StatusBadRequest)
						return
					}
				}
				//store the results that match this query
				results := make([]*book, 0)
				//range over the dvd inventory
				for _, v := range ds.Inventory["book"] {
					item := v.(*book)
				filter:
					for key, value := range queries {
						//if any of the criteria don't match, just skip this dvd
						var field string
						switch strings.ToLower(strings.TrimSpace(key)) {
						case "id":
							field = item.ID
						case "title":
							field = item.Title
						case "author":
							field = item.Author
						case "genre":
							field = item.Genre
						case "publishingdate":
							field = item.PublishingDate
						case "checkedout":
							if item.CheckedOut == true {
								field = "true"
							} else {
								field = "false"
							}
						case "checkedoutby":
							field = item.CheckedOutBy
						}
						if field != value[0] {
							break filter
						}
					}
					//if all the criteria matched, add this dvd to the results
					results = append(results, item)
				}
				//write the results to the response writer
				json.NewEncoder(w).Encode(results)
				return
			}
		} else {
			http.Error(w, "This method is not supported", http.StatusNotImplemented)
			return
		}
	}
	return http.HandlerFunc(fn)
}
