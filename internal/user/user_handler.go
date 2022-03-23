package user

import "net/http"

func HandleUser(ds interface{}) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//get user
			//render view
		} else if r.Method == http.MethodPost {
			//update user information
			//redirect
		} else {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(fn)
}
