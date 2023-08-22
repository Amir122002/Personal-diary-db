package models

import "time"

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DiaryEntry struct {
	Id       int       `json:"id"`
	Text     string    `json:"text"`
	Active   bool      `json:"active"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
}
