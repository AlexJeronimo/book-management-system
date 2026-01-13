package shelf

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

// Create new book
func NewBook(id int, name string, authors []string, isbn string, read bool) Book {
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
		StoreLink: "",
	}

}

// Add book to whishlist and set store link
func (b *Book) AddToWhishlist(link string) {
	//add book to whishlist
	//set storelink
	b.Whishlist = true
	b.StoreLink = link
}
