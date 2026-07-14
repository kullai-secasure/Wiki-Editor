package models

import "time"

type WikiPage struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PageRevision struct {
	ID        int64     `json:"id"`
	PageID    int64     `json:"page_id"`
	Content   string    `json:"content"`
	Editor    string    `json:"editor"`
	CreatedAt time.Time `json:"created_at"`
}
