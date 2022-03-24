package item

import "fmt"

//generator handles creating a simple ID system
//for library items
type generator struct {
	dvd  uint8
	tape uint8
	book uint8
}

//NewGenerator returns a new generator object.
//Note that because this is just an example, and
//not a commercial product, this is just going to
//start with ones no matter what
func NewGenerator() *generator {
	return &generator{
		dvd:  1,
		tape: 1,
		book: 1,
	}
}

//getID returns a stringified, zero-padded version
//of the current number for the given category, then
//iterates that number
func (g *generator) GetID(which string) string {
	var s string
	switch which {
	case "dvd":
		s = fmt.Sprintf("%08d", g.dvd)
	case "tape":
		s = fmt.Sprintf("%08d", g.tape)
	case "book":
		s = fmt.Sprintf("%08d", g.book)
	default:
		return fmt.Sprintf("%q is not a valid type. Please enter \"dvd\", \"tape\", or \"book\"", which)
	}
	return s
}