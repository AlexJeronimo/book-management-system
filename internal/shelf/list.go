package shelf

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// List books in shelf
func (repo *Repository) List() {
	// TODO: add interactive pagination to book list (limit 25-50 books per page,  show total number and current positions listed, show page from pages 12/25)
	// TODO: add cli icons in feature as app optimization step in later releases
	// TODO: change to two outputs for existing books and for whishlist

	// fmt.Println("------------------------------------")
	// for _, book := range repo.Books {
	// 	fmt.Println("ID: ", book.ID)
	// 	fmt.Printf("Book Title: %s\n", book.Name)
	// 	if len(book.Authors) == 0 {
	// 		fmt.Println("Anknown Author")
	// 	} else {
	// 		fmt.Println("Author(s): ", strings.Join(book.Authors, ", "))
	// 	}

	// 	if book.ISBN == "" {
	// 		continue
	// 	} else {
	// 		fmt.Println("ISBN: ", book.ISBN)
	// 	}

	// 	if book.Whishlist == true {
	// 		fmt.Println("Book store link: ", book.StoreLink)
	// 	}
	// 	fmt.Println("------------------------------------")

	// }

}

func (repo *Repository) PrintTable(wl bool) {
	table := tablewriter.NewWriter(os.Stdout)

	if wl {
		table.Header([]string{"ID", "Title", "Authors", "Genre", "ISBN", "Read", "Wishlist"})
		for _, book := range repo.Books {
			if !book.Whishlist {
				continue
			} else {
				authors := strings.Join(book.Authors, ", ")
				read := "No"
				if book.Read {
					read = "Yes"
				}
				wishlist := "No"
				if book.Whishlist {
					wishlist = "Yes"
				}

				table.Append([]string{
					fmt.Sprint(book.ID),
					book.Name,
					authors,
					book.Genre,
					book.ISBN,
					read,
					wishlist,
				})
			}

		}
	} else {
		table.Header([]string{"ID", "Title", "Authors", "Genre", "ISBN", "Read"})
		for _, book := range repo.Books {
			authors := strings.Join(book.Authors, ", ")
			read := "No"
			if book.Read {
				read = "Yes"
			}

			table.Append([]string{
				fmt.Sprint(book.ID),
				book.Name,
				authors,
				book.Genre,
				book.ISBN,
				read,
			})
		}
	}

	table.Render()
}
