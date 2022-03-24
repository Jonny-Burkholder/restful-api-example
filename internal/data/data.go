package data

import (
	"github.com/Jonny-Burkholder/restful-api-example/internal/user"
)

type DataStore struct {
	Inventory map[string][]interface{}
	Users     map[string]*user.User
}
