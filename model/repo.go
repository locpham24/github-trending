package model

import "time"

type Repo struct {
	Name        string    `json:"name" db:"name,omitempty"`
	Description string    `json:"description" db:"description,omitempty"`
	Url         string    `json:"url" db:"url, omitempty"`
	Color       string    `json:"color" db:"color, omitempty"`
	Lang        string    `json:"lang" db:"lang, omitempty"`
	Fork        string    `json:"fork" db:"fork, omitempty"`
	Stars       string    `json:"stars" db:"stars, omitempty"`
	StarsToday  string    `json:"stars_today" db:"stars_today, omitempty"`
	Author      string    `json:"author" db:"author, omitempty"`
	Bookmarked  bool      `json:"bookmarked"`
	CreatedAt   time.Time `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at, omitempty"`
}
