package model

type ResponseDelete struct {
	Status     int                  `json:"status"`
	Message    string               `json:"message"`
	RecordLama map[string]Mahasiswa `json:"recordLama"`
}
