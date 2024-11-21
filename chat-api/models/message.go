package models

import "time"

// Message adalah model untuk menyimpan pesan
type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `json:"content"`
	Sender    string    `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
}
