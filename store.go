package main

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

// CreateHightLights implements Storage.
func (s *Store) CreateHightLights([]Highlight) error {
	panic("unimplemented")
}

type Storage interface {
	CreateBook(Book) error
	CreateHightLights([]Highlight) error
	GetBookByISBN(string) (*Book, error)
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateBook(b Book) error {
	_, err := s.db.Exec(`
	INSERT INTO books (isbn, title ,authors)
	    VALUES(?,?,?)
		`, b.ISBN, b.Title, b.Authors)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateHightlights(hs []Highlight) error {
	values := []interface{}{}

	query := "INSERT INTO highlights (text, location, note, userId, bookId) VALUES "
	for _, h := range hs {
		query += "(?, ?, ?, ?, ?),"
		values = append(values, h.Text, h.Location, h.Note, h.UserID, h.BookID)
	}

	query = query[:len(query)-1]

	_, err := s.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetBookByISBN(isbn string) (*Book, error) {
	rows, err := s.db.Query(`
	SELECT * FROM books WHERE isbn = ?`, isbn)
	if err != nil {
		return nil, err
	}
	book := new(Book)
	for rows.Next() {
		if err := rows.Scan(&book.ISBN, &book.Title, &book.Authors); err != nil {
			return nil, err
		}
	}
	return book, nil
}
