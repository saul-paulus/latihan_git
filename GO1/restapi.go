package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Bangundatar
type BangunDatar struct {
	Panjang int `json:"panjang"`
	Lebar   int `json:"lebar"`
	Tinggi  int `json:"tinggi"`
}
type Output struct {
	Tipe   string `json:"tipe"`
	Volume int    `json:"volume"`
}

// TerimaData
func TerimaData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Bng BangunDatar
	var Hsl Output
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Bng); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getP := r.PostFormValue("panjang")
			panjang, _ := strconv.Atoi(getP)
			getL := r.PostFormValue("lebar")
			lebar, _ := strconv.Atoi(getL)
			getT := r.PostFormValue("lebar")
			tinggi, _ := strconv.Atoi(getT)
			Bng = BangunDatar{
				Panjang: panjang,
				Lebar:   lebar,
				Tinggi:  tinggi,
			}
			if panjang == lebar {
				volume := panjang * lebar * tinggi
				tipe := "Kubus"
				Hsl = Output{
					Tipe:   tipe,
					Volume: volume,
				}
			} else {
				volume := (panjang * lebar) * tinggi
				tipe := "Balok"
				Hsl = Output{
					Tipe:   tipe,
					Volume: volume,
				}
			}
		}

		// dataBangundatar, _ := json.Marshal(Bng) // to byte
		dataOutput, _ := json.Marshal(Hsl)
		// w.Write(dataBangundatar)                // cetak di browser
		w.Write(dataOutput) // cetak di browser
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/index", TerimaData)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
