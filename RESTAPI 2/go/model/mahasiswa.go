package model

import "time"

type Mahasiswa struct {
	ID        int       `json:"id"`
	NIM       int       `json:"nim"`
	Name      string    `name:"name"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseUpdate struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Mahasiswa
}

type ResponseDelete struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Mahasiswa
}
