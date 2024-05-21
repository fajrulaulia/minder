package model

import "time"

type Matcher struct {
	ID        int
	User1ID   int
	User2ID   int
	Action    string
	CreatetAt time.Time
}
