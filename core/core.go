package core

import (
	"html/template"
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

type Mesaj struct {
	Id     int
	Baslik string
	Mesaj  string
	Zaman  string
}

type Yorum struct {
	Id      int
	AnketId int
	Nick    string
	Yorum   string
	Zaman   string
}

type Secenek struct {
	Id      int
	Ad      string
	Oy      int
	AnketId int
}

type Anket struct {
	Id          int
	Ad          string
	Baslik      string
	Tip         int
	OySayisi    int
	YorumSayisi int
	Aktif       int
	Zaman       string
	Secenekler  []Secenek
	SecenekJson template.JS
	Yorumlar    []Yorum
}

func baglan() (*sql.DB, error) {
	var dbName = "anket.db"
	if _, err := os.Stat(dbName); os.IsNotExist(err) { //veri tabanı dosyası yoksa
		return nil, err
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
