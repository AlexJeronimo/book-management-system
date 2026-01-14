package shelf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Repository struct {
	Books []Book
}

// Create new bookshelf
func NewRepository() *Repository {
	return &Repository{}
}

// Generate book ID
func (repo *Repository) GenerateID() int {
	//some king of autoincrement for ID value
	//get last id in the book list and add +1 maxID+1
	//1. Load book list from file if exists
	//2. if list empty, return 1 if []Book == 0; id = 1
	maxID := 0
	for _, book := range repo.Books {
		if book.ID > maxID {
			maxID = book.ID
		}
	}

	return maxID + 1
}

// Add book to shelf
func (repo *Repository) Add(book Book) {
	//receive book with filled data
	//add book to memory
	//save book to local json file
	repo.Books = append(repo.Books, book)
}

// List books in shelf
func (repo *Repository) List() {
	//list books
	// TODO: add interactive pagination to book list (limit 25-50 books per page,  show total number and current positions listed, show page from pages 12/25)
	// TODO: add cli icons in feature as app optimization step in later releases
	// TODO: change to two outputs for existing books and for whishlist

	fmt.Println("------------------------------------")
	for _, book := range repo.Books {
		fmt.Println("ID: ", book.ID)
		fmt.Printf("Book Title: %s\n", book.Name)
		if len(book.Authors) == 0 {
			fmt.Println("Anknown Author")
		} else {
			fmt.Println("Author(s): ", strings.Join(book.Authors, ", "))
		}

		if book.ISBN == "" {
			continue
		} else {
			fmt.Println("ISBN: ", book.ISBN)
		}

		if book.Whishlist == true {
			fmt.Println("Book store link: ", book.StoreLink)
		}
		fmt.Println("------------------------------------")

	}

}

func Search(whishlist bool, search string) {
	//search book by name (partial name), author (one of authors or by Author Name or Second Name), isbn
	// show is book in repository and show its status
	if whishlist == true {
		//do search in whichlist books
	} else {
		//do search in already existed books
	}
}

// Save shelf to file
func (repo *Repository) Save(filename string) error {
	//save books to json file
	data, err := json.MarshalIndent(repo, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Load books from file to shelf
func (repo *Repository) Load(filename string) error {
	//get books list from json file
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, repo)
	if err != nil {
		return err
	}
	return nil

}

// Mark book as red
func (repo *Repository) SetRead(id int) {
	//repo.List()
	//fmt.Println("Call list to identify does method receive books")
	//fmt.Printf("ID: %d, type: %T", id, id)

	for i := range repo.Books {
		if repo.Books[i].ID == id {
			//fmt.Printf("The book where I whant to change ID. %v+\n", repo.Books[i])
			repo.Books[i].Read = true
			//fmt.Printf("The read status after changing %v+\n", repo.Books[i].Read)
		}
	}
}

// Move book from whishlist to shelf
func MoveFromWhishListToShelf(id int) {

}
