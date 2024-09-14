package models

type Note struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

var Notes []Note
var NoteID = 0
