package models

import "time"

const (
	StatusAdded = 1 << iota
	StatusProcessed
	StatusPublishedToTelegram
	StatusPublishedToFacebook
)

type News struct {
	ID        string    `json:"id" firestore:"id"`
	Title     string    `json:"title" firestore:"title"`
	Content   string    `json:"content" firestore:"content"`
	Status    int       `json:"status" firestore:"status"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	URL       string    `json:"url" firestore:"url"`
}
