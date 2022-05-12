// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package models

import "time"

type User struct {
	ID         int64     `json:"id"`
	Login      string    `json:"-"`
	Password   string    `json:"-"`
	Photo      Photo     `json:"photo"`
	ScreenName string    `json:"screen_name"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Blocked    string    `json:"-"`
	Deleted    string    `json:"-"`
	CreatedAt  time.Time `json:"-"`
}
