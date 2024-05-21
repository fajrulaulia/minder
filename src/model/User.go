package model

import (
	"time"
)

type User struct {
	ID               int
	Username         string
	Email            string
	Password         string
	SubscribedEndate *time.Time
	CreatetAt        time.Time
	UpdatedAt        time.Time
}
