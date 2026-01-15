package shelf

import "strings"

type SearchFilter struct {
	Query     string
	Author    bool
	Title     bool
	ISBN      bool
	Whishlist *bool
	Read      *bool
}

func (repo *Repository) Search(filter SearchFilter) []Book {
	var result []Book
	query := strings.ToLower(filter.Query)

	searchAll := !filter.Author && !filter.Title && !filter.ISBN

	for _, b := range repo.Books {
		if filter.Whishlist != nil && b.Whishlist != *filter.Whishlist {
			continue
		}
		if filter.Read != nil && b.Read != *filter.Read {
			continue
		}
		if query != "" {
			match := false

			if (filter.Title || searchAll) && strings.Contains(strings.ToLower(b.Name), query) {
				match = true
			}

			if (filter.ISBN || searchAll) && strings.Contains(strings.ToLower(b.ISBN), query) {
				match = true
			}

			if filter.Author || searchAll {
				for _, a := range b.Authors {
					if strings.Contains(strings.ToLower(a), query) {
						match = true
						break
					}
				}
			}

			if !match {
				continue
			}
		}

		result = append(result, b)
	}

	return result
}
