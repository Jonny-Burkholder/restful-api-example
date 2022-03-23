package item

import "net/http"

func handleItem(ds interface{}) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

	}
	return http.HandlerFunc(fn)
}
