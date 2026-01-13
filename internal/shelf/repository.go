package shelf

import (
	"encoding/json"
	"os"
)

type Repository struct {
	Books []Book
}

func NewRepository() *Repository {
	return &Repository{}
}

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

func (repo *Repository) Add(book Book) {
	//receive book with filled data
	//add book to memory
	//save book to local json file
	repo.Books = append(repo.Books, book)
}

// TODO: add interactive pagination to book list (limit 25-50 books per page,  show total number and current positions listed, show page from pages 12/25)
// TODO: add cli icons in feature as app optimization step in later releases
func List() {
	//get all books
	//filter books
	//list books

}

func Search() {
	//search book by name (partial name), author (one of authors or by Author Name or Second Name), isbn
	// show is book in repository and show its status
}

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

func (repo *Repository) SetRead(id int) {
	//change read status from yes to no or backwards
	// for _, elem := range repo.Books {
	// 	if elem.ID == id {
	// 		elem.Read = true
	// 	}
	// }
}

func (repo *Repository) SetStoreLink(link string, book Book) {
	//set store link
}
