package item

import (
	"fmt"
	"net/http"
)

func HandleItem(ds interface{}) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//get the item
			q := r.URL.Query()
			fmt.Println(q)
		} else if r.Method == http.MethodPost {
			//return the item
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}
