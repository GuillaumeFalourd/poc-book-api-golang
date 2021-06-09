# poc-book-api-golang

## Description

POC of a _Book REST API_ using [gorilla/mux](https://github.com/gorilla/mux) (a powerful HTTP router and URL matcher for building Go web servers).

## Prerequisites

### Golang

- Go installed
- GOPATH configured

### Dependency

- `go get -u github.com/gorilla/mux`

## How to use?

On the repo **root directory**, run: `go run .`, you should see something like this:

```sh
YYYY/MM/DD HH:mm:ss file load took 11.617198ms
YYYY/MM/DD HH:mm:ss Start Book Rest API
```

Then, consume the services below on your Rest Client tool.

## Services

### Create new book

**POST:** `http://localhost:8080/books`

```json
{
    "book_id": "50000",
    "title": "Book Title",
    "authors": "Author Name",
    "average_rating": "4.99",
    "isbn":"1234567890",
    "isbn_13":"1234567890123",
    "language_code":"eng",
    "num_pages":"500",
    "ratings":"100",
    "reviews":"50"
}
```

* * *

### Get book by author

**GET:** `http://localhost:8080/books/authors/{author}`

* * *

### Get book by name

**GET:** `http://localhost:8080/books/book-name/{bookName}`

* * *

### Get Book by ISBN (International Standard Book Number)

**GET:** `http://localhost:8080/books/isbn/{isbn}`

* * *

### Update Book by ISBN (International Standard Book Number)

**PUT:** `http://localhost:8080/books/isbn/{isbn}`

```json
{
    "book_id": "8429",
    "title": "Insects & Spiders",
    "authors": "David Burnie",
    "average_rating": "4.33",
    "isbn":"0783548818",
    "isbn_13":"9780783548814",
    "language_code":"eng",
    "num_pages":"64",
    "ratings":"3",
    "reviews":"0"
}
```

* * *

### Delete Book by ISBN (International Standard Book Number)

**DELETE:** `http://localhost:8080/books/isbn/{isbn}`
