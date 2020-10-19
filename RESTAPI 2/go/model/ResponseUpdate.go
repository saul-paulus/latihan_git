package model

type ResponseUpdate struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Mahasiswa
}
