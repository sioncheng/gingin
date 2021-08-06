package domain

import "time"

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
