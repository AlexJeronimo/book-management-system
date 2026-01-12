package book

// TODO: add genres and subgenres (leave it for feature development practice to work with legacy code, add features, etc.)
type Book struct {
	ID        int      `json:"id"`
	Authors   []string `json:"authors"`
	Name      string   `json:"name"`
	ISBN      string   `json:"isbn"`
	Read      bool     `json:"read"`
	Whishlist bool     `json:"whishlist"`
	StoreLink string   `json:"store_link,omitempty"`
}

// Method to create new book
func NewBook(name string, authors []string, isbn string, read bool, id int) Book {
	//create new empty book
	//fill the book with data
	//TODO: add validation to entered values
	return Book{
		ID:        id,
		Name:      name,
		Authors:   authors,
		ISBN:      isbn,
		Read:      read,
		Whishlist: false,
		StoreLink: "_",
	}

}
