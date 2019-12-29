package repository

import(
	"log"
	"database/sql"
	"codixir/models"
)
 // BookRepository -> repository pattern
type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks -> get all books
func (b BookRepository) GetBooks(db *sql.DB) []models.Book{
	rows, err := db.Query("select * from books")
	logFatal(err)
	var books []models.Book
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	return books
}

// GetBook -> get book by id
func (b BookRepository) GetBook(db *sql.DB, id int) models.Book{
	row := db.QueryRow("select * from books where id=$1", id)
	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	return book
}

// AddBook -> add book
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int{
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) returning id", book.Title, book.Author, book.Year).Scan(&book.ID)
	logFatal(err)
	return book.ID
}

// UpdateBook -> update book
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64{
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id", book.Title, book.Author, book.Year, book.ID)
	logFatal(err)
	updated, err := result.RowsAffected()
	return updated
}

// RemoveBook -> delete book by id
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64{
	result, err := db.Exec("delete from books where id=$1", id)
	logFatal(err)
	rowsDeleted, err := result.RowsAffected()
	return rowsDeleted
}