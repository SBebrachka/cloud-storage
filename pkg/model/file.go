package model

import "time"

type File struct {
	ID        int       `json:"id" db:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	Filename  string    `json:"filename" db:"filename"`
	Size      int64     `json:"size" db:"size"`
	Path      string    `json:"path,omitempty" db:"path"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
