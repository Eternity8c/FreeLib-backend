package postgres

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepository struct {
	pool *pgxpool.Pool
}

func NewBookRepository(pool *pgxpool.Pool) repository.BookRepository {
	return &bookRepository{pool: pool}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	query := `SELECT * FROM books;`
	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Genre,
			&book.Content,
			&book.CoverURL,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	query := `SELECT * FROM books WHERE id = $1;`
	var book models.Book
	err := r.pool.QueryRow(context.Background(),
	query, id).Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Genre,
			&book.Content,
			&book.CoverURL,
			&book.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepository) Create(book *models.Book) error {
	query := `INSERT INTO books(
		title,
		author,
		description,
		genre,
		content,
		cover_URL,
		created_At) 
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at`
	err := r.pool.QueryRow(context.Background(), query,
		book.Title,
		book.Author,
		book.Description,
		book.Genre,
		book.Content,
		book.CoverURL,
		book.CreatedAt,
	).Scan(&book.ID, &book.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) Search(s string) ([]models.Book, error) {
	query := `SELECT * FROM books 
	WHERE title = $1 OR author = $2 OR genre = $3`
	rows, err := r.pool.Query(context.Background(), query, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Genre,
			&book.Content,
			&book.CoverURL,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) Delete(id uint) error {
	query := `DELETE FROM books WHERE id = $1`
	result, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("books with id %d not found", id)
	}

	return nil
}