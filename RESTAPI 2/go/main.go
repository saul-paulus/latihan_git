package main

import (
	"be3gomy/config"
	"be3gomy/mahasiswa"
	"be3gomy/model"
	"be3gomy/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const addr = `:7001`

type Server struct {
	db      *sql.DB
	ViewDir string
}
type handler func(w http.ResponseWriter, r *http.Request)
type H map[string]interface{}

func InitServer() *Server {
	db, err := config.Mysql()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		db:      db,
		ViewDir: `views/`,
	}
}
func (s *Server) Listen() {
	log.Println(`listen at ` + addr)
	http.HandleFunc(`/mahasiswa`, s.Mahasiswa())
	//http.HandleFunc(`/mahasiswa/show`, s.MahasiswaGetFromID())
	http.HandleFunc(`/mahasiswa/create`, s.MahasiswaCreate())
	http.HandleFunc(`/mahasiswa/update`, s.MahasiswaUpdate())
	http.HandleFunc(`/mahasiswa/delete`, s.MahasiswaDelete())
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *Server) Mahasiswa() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		mahasiswas, err := mahasiswa.SelectAll(s.db)
		if utils.IsError(w, err) {
			return
		}
		utils.ResponseJson(w, mahasiswas)
	}
}

func (s *Server) MahasiswaCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == `GET` {
			http.ServeFile(w, r, s.ViewDir+`mahasiswa_create.html`)
			return
		}
		m := model.Mahasiswa{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if utils.IsError(w, err) {
			return
		}
		err = mahasiswa.Insert(s.db, &m)
		if utils.IsError(w, err) {
			return
		}
		utils.ResponseJson(w, m)
	}
}

func (s *Server) MahasiswaUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == `GET` {
			http.ServeFile(w, r, s.ViewDir+`mahasiswa_update.html`)
			return
		}
		m := model.Mahasiswa{}
		ru := model.ResponseUpdate{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if utils.IsError(w, err) {
			return
		}
		err = mahasiswa.Update(s.db, &m, &ru)
		if utils.IsError(w, err) {
			return
		}
		utils.ResponseJson(w, ru)
	}
}

func (s *Server) MahasiswaDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == `GET` {
			http.ServeFile(w, r, s.ViewDir+`mahasiswa_delete.html`)
			return
		}
		m := model.Mahasiswa{}
		rd := model.ResponseDelete{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if utils.IsError(w, err) {
			return
		}
		err = mahasiswa.Delete(s.db, &m, &rd)
		if utils.IsError(w, err) {
			return
		}
		utils.ResponseJson(w, rd)
	}
}
func main() {
	server := InitServer()
	server.Listen()
}
