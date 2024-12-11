package models

import "time"

type News struct {
	ID        string    `json:"id" firestore:"id"`
	Title     string    `json:"title" firestore:"title"`
	Content   string    `json:"content" firestore:"content"`
	Status    string    `json:"status" firestore:"status"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}
