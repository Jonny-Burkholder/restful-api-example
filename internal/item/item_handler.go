package item

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jonny-Burkholder/restful-api-example/internal/data"
)

func GetItem(ds data.DataStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//get the item
			q := r.URL.Query()
			fmt.Println(q)
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
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}
