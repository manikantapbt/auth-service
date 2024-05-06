package repository

import (
	"database/sql"
	"log"
)

type IEventRepository interface {
	InsertEvent(event string, phoneNumber string)
}

const (
	query = "INSERT INTO user_events (phone_number, event) VALUES ($1, $2)"
)

func NewEventRepository(db *sql.DB) IEventRepository {
	return &eventRepository{db: db}
}

type eventRepository struct {
	db *sql.DB
}

func (r *eventRepository) InsertEvent(event string, phoneNumber string) {
	_, err := r.db.Exec(query, phoneNumber, event)
	if err != nil {
		// ignoring event db query errors as they are of low priority
		log.Println(err)
	}
}
