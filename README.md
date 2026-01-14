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

