# Book management system

CLI application written in Go for managing a personal book library.

The project is planned as a GoLang learning project, and in the future will grow from a simple CLI tool to a web-based tool.
Also, it is planned to be used as my book management system (I have plenty of books, and sometimes when ordering a new book, you realize that you have one already).



## Features

- Add book with title, authors, ISBN
- Mark books as read / unread
- Manage whishlist (books to buy later)
- Search books by:
    - title
    - author
    - ISBN
- List books with filters
- Local storage JSON 
- Pagination for large collections

## Project Status

This project is under active development and is used as learning project for Go and CLI applications.

## Requirements

- Go 1.25+

## Instalaltions

Clone the repository

```bash
git clone https://github.com/AlexJeronimo/book-management-system.git
cd book-management-system
go build -o bms ./cmd
```

## Usage

`bms [command]`

Available Commands:
    add         Add the book to the shelf
    list        List all books in the shelf
    read        Set book as read
    help        Help about any command

### help
    -h, --help   help for bms

`bms help`
`bms -h`

### add
Usage:
  bms add [flags]

Flags:
  -a, --author strings     Book author
  -g, --genre string       add book genre
  -h, --help               help for add
  -s, --isbn string        ISBN
  -r, --read               is book read
  -l, --storelink string   add storelink
  -t, --title string       Book name
  -w, --whishlist          add book to whishlist

  `bms add -a "Author" --author "Another Author" --title "Book title" -g "Book genre" -r -w -l "https://link_where_you_can_order_the_book" -s "Book ISBN"`

### list
Usage:
  bms list

Flags:
  -h, --help        help for list
  -w, --whishlist   show books from whishlist

`bms list`
`bms list -w`

  ### read
  Usage:
  bms read [flags]

Flags:
  -h, --help        help for read
  -i, --id string   set book id

`bms read --id 15`