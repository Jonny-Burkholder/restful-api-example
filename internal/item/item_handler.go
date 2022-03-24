package item

import (
	"encoding/json"
	"net/http"

	"github.com/Jonny-Burkholder/restful-api-example/internal/data"
)

func GetItem(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//get the item
			itemType := r.URL.Query()["type"]
			title := r.URL.Query()["title"]
			if itemType == "dvd" {
				for _, dvd := range ds.Inventory["dvd"] {
					if dvd.Title == title {
						json.NewEncoder(w).Encode(dvd)
						return
					}
				}
			} else if itemType == "tape" {
				for _, tape := range ds.Inventory["tape"] {
					if tape.Title == title {
						json.NewEncoder(w).Encode(tape)
						return
					}
				}
			} else if itemType == "book" {
				for _, book := range ds.Inventory["book"] {
					if book.Title == title {
						json.NewEncoder(w).Encode(book)
						return
					}
				}
			}
			w.Write([]byte("Item not found"))
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}

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

func HandleDVD(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//return all DVDs
			json.NewEncoder(w).Encode(ds.Inventory["dvd"])
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}

func HandleTape(ds data.DataStore) http.Handler {
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

func HandleBook(ds data.DataStore) http.Handler {
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
