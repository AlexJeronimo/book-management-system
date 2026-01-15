package shelf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
	repo.Books = append(repo.Books, book)
}

// Save shelf to file
func (repo *Repository) Save(filename string) error {

	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("cannot create dir %s: %w", dir, err)
		}
	}

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
	for i := range repo.Books {
		if repo.Books[i].ID == id {
			repo.Books[i].Read = true
			return
		}
	}
}

// Move book from whishlist to shelf
func MoveFromWhishListToShelf(id int) {

}
