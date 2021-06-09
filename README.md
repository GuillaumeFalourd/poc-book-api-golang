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

On the repo **root directory**, run: `go run .`

_**Obs**: You should see something like this:_

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

![create book](https://user-images.githubusercontent.com/22433243/121394116-66c49980-c927-11eb-8d8b-a3a660f8c35d.png)

* * *

### Get book by author

**GET:** `http://localhost:8080/books/authors/{author}`

![get book by author](https://user-images.githubusercontent.com/22433243/121394149-6e843e00-c927-11eb-9068-6cfd75e13fe0.png)

* * *

### Get book by name

**GET:** `http://localhost:8080/books/book-name/{bookName}`

![get book by name](https://user-images.githubusercontent.com/22433243/121394171-747a1f00-c927-11eb-9ded-1e908c6329e6.png)

* * *

### Get Book by ISBN (International Standard Book Number)

**GET:** `http://localhost:8080/books/isbn/{isbn}`

![get book by isbn](https://user-images.githubusercontent.com/22433243/121394187-78a63c80-c927-11eb-970d-d5c2aa4f57ed.png)

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

![update book](https://user-images.githubusercontent.com/22433243/121394212-7e9c1d80-c927-11eb-982e-2bab2260068e.png)

* * *

### Delete Book by ISBN (International Standard Book Number)

**DELETE:** `http://localhost:8080/books/isbn/{isbn}`

![delete book](https://user-images.githubusercontent.com/22433243/121394243-83f96800-c927-11eb-9bbe-83faad3343f4.png)
