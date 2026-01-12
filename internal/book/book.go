package book

// TODO: add genres and subgenres (leave it for feature development practive to work with legacy code, add feature, etc.)
type Book struct {
	ID        int      `json:"id"`
	Authors   []string `json:"authors"`
	Name      string   `json:"name"`
	ISBN      string   `json:"isbn"`
	Read      bool     `json:"read"`
	Whishlist bool     `json:"whishlist"`
	StoreLink string   `json:"store_link,omitempty"`
}

func NewBook() Book {
	//TODO: add validation to entered values
	var autor string
	var authors []string
	var name string
	var isbn string
	var read bool
	var whishlist bool
	var storeLink string

	authors = append(authors, autor)

	book := Book{
		ID:        GenerateID(),
		Authors:   authors,
		Name:      name,
		ISBN:      isbn,
		Read:      read,
		Whishlist: whishlist,
		StoreLink: storeLink,
	}

	return book
}

// TODO: in learning purposes do with two different variants, with user prompting and with cli flags
func GetBookData() {
	//prompt user for book data to fill the struct fields
	//can be passed as cli commands

}

func GenerateID() int {
	//some king of autoincrement for ID value
	//get last id in the book list and add +1 maxID+1
	// if list empty, return 1 if []Book == 0; id = 1
	var id int
	return id
}
