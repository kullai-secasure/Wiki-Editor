package storage

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"wikiflex/internal/models"
)

type PostgresDB struct {
	conn *sql.DB
}

func NewPostgresDB() *PostgresDB {
	dsn := os.Getenv("DATABASE_URL")
	conn, _ := sql.Open("postgres", dsn)
	return &PostgresDB{conn: conn}
}

func (db *PostgresDB) GetPageBySlug(slug string) (*models.WikiPage, error) {
	page := &models.WikiPage{}
	err := db.conn.QueryRow("SELECT id, slug, title, content, author, updated_at FROM pages WHERE slug = $1", slug).Scan(&page.ID, &page.Slug, &page.Title, &page.Content, &page.Author, &page.UpdatedAt)
	return page, err
}

func (db *PostgresDB) Close() { db.conn.Close() }
