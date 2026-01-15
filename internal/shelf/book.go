package shelf

// TODO: add genres and subgenres (leave it for feature development practice to work with legacy code, add features, etc.)
type Book struct {
	ID        int      `json:"id"`
	Authors   []string `json:"authors"`
	Name      string   `json:"name"`
	ISBN      string   `json:"isbn"`
	Genre     string   `json:"genre"`
	Read      bool     `json:"read"`
	Whishlist bool     `json:"whishlist"`
	StoreLink string   `json:"store_link,omitempty"`
}

// Create new book
func NewBook(id int, name string, authors []string, isbn string, read bool, genre string) Book {

	//TODO: add validation to entered values

	// var readBook bool

	// switch read {
	// case "yes", "y", "YES", "Yes", "Y":
	// 	readBook = true
	// default:
	// 	readBook = false
	// }

	return Book{
		ID:        id,
		Name:      name,
		Authors:   authors,
		ISBN:      isbn,
		Genre:     genre,
		Read:      read,
		Whishlist: false,
		StoreLink: "",
	}

}

// Add book to whishlist and set store link
func (b *Book) AddToWhishlist(link string) {
	b.Whishlist = true
	b.StoreLink = link
}
