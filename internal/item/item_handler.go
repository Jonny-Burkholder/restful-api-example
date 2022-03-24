package item

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Jonny-Burkholder/restful-api-example/internal/data"
)

func ReturnItem(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//return item
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}

//DonateItem is here so that there's an example of using
//a post request
func DonateItem(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//donate item
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}

func GetDVDs(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			queries := r.URL.Query()
			if len(queries) == 0 {
				//return all DVDs
				//This doesn't really account for the pattern api/dvds/0010000000X
				json.NewEncoder(w).Encode(ds.Inventory["dvd"])
				return
			} else {
				//check that all queries are valid
				for key, _ := range queries {
					if _, ok := dvdParams[strings.ToLower(strings.TrimSpace(key))]; ok != true {
						http.Error(w, "Invalid request", http.StatusBadRequest)
						return
					}
				}
				//store the results that match this query
				results := make([]*dvd, 0)
				//range over the dvd inventory
				for _, dvd := range ds.Inventory["dvd"] {
				filter:
					for key, value := range queries {
						//if any of the criteria don't match, just skip this dvd
						if dvd.key != value {
							break filter
						}
					}
					//if all the criteria matched, add this dvd to the results
					results = append(results, dvd)
				}
				//write the results to the response writer
				json.NewEncoder(w).Encode(results)
				return
			}
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func GetTapes(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//return all tapes
			json.NewEncoder(w).Encode(ds.Inventory["tape"])
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}

func GetBooks(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//return all books
			json.NewEncoder(w).Encode(ds.Inventory["book"])
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}
