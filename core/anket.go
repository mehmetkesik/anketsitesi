package core

import (
	"html/template"
	"encoding/json"
	"database/sql"
)

func YorumlarGetir(anketid int) ([]Yorum, error) {
	var yorumlar []Yorum
	db, err := baglan()
	if err != nil {
		return yorumlar, err
	}
	defer db.Close()

	rows, err := db.Query("select * from yorum where AnketId = ? order by Id desc", anketid)

	for rows.Next() {
		var yorum Yorum
		rows.Scan(&yorum.Id, &yorum.AnketId, &yorum.Nick, &yorum.Yorum, &yorum.Zaman)
		yorumlar = append(yorumlar, yorum)
	}

	return yorumlar, err
}

func YorumArttir(anketid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update anket set YorumSayisi = YorumSayisi+1 where Id=?", anketid)

	return err
}

func YorumEkle(yorum Yorum) (int, error) {
	db, err := baglan()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("insert into yorum (AnketId,Nick,Yorum,Zaman) values"+
		" (?,?,?,?)", yorum.AnketId, yorum.Nick, yorum.Yorum, yorum.Zaman)
	if err != nil {
		return -1, err
	}
	lastid, err := result.LastInsertId()
	return int(lastid), err
}

func OyArttir(anketid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update anket set OySayisi = OySayisi+1 where Id=?", anketid)

	return err
}

func BitmisGetir(bas, bit int) ([]Anket, error) {
	var anketler []Anket
	db, err := baglan()
	if err != nil {
		return anketler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket where Aktif = 0 order by Id desc limit ?,? ", bas, bit)
	if err != nil {
		return anketler, err
	}

	return anketGetir(rows)
}

func EnCokGetir(bas, bit int) ([]Anket, error) {
	var anketler []Anket
	db, err := baglan()
	if err != nil {
		return anketler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket order by OySayisi desc limit ?,? ", bas, bit)
	if err != nil {
		return anketler, err
	}
	return anketGetir(rows)
}

func OylananGetir(bas, bit int) ([]Anket, error) {
	var anketler []Anket
	db, err := baglan()
	if err != nil {
		return anketler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket where Aktif = 1 order by Id desc limit ?,? ", bas, bit)
	if err != nil {
		return anketler, err
	}

	return anketGetir(rows)
}

func AktifSon5Getir() ([]Anket, error) {
	var anketler []Anket
	db, err := baglan()
	if err != nil {
		return anketler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket where Aktif = 1 order by Id desc limit 5 ")
	if err != nil {
		return anketler, err
	}

	return anketGetir(rows)
}

func AnketGetir(id int) (Anket, error) {
	var anket Anket
	db, err := baglan()
	if err != nil {
		return anket, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket where Id = ?", id)
	anketler, err := anketGetir(rows)

	return anketler[0], err
}

func AnketEkle(anket Anket) (int, error) {
	db, err := baglan()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("insert into anket (Ad,Baslik,Tip,OySayisi,YorumSayisi,Aktif,Zaman) values"+
		" (?,?,?,?,?,?,?)", anket.Ad, anket.Baslik, anket.Tip, anket.OySayisi, anket.YorumSayisi, anket.Aktif, anket.Zaman)
	if err != nil {
		return -1, err
	}
	lastid, err := result.LastInsertId()

	if len(anket.Secenekler) > 0 {
		for _, v := range anket.Secenekler {
			v.AnketId = int(lastid)
			_, err = SecenekEkle(v)
			if err != nil {
				return -1, err
			}
		}
	}

	if len(anket.Yorumlar) > 0 {
		for _, v := range anket.Yorumlar {
			_, err = YorumEkle(v)
			if err != nil {
				return -1, err
			}
		}
	}

	return int(lastid), err
}

func AnketGuncelle(anket Anket) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update anket set Ad=?,Baslik=?,Tip=?,OySayisi=?,YorumSayisi=?,Aktif=?,Zaman=? where Id=?",
		anket.Ad, anket.Baslik, anket.Tip, anket.OySayisi, anket.YorumSayisi, anket.Aktif, anket.Zaman, anket.Id)
	if err != nil {
		return err
	}

	SeceneklerSil(anket.Id)
	if len(anket.Secenekler) > 0 {
		for _, v := range anket.Secenekler {
			_, err = SecenekEkle(v)
			if err != nil {
				return err
			}
		}
	}

	return err

}

func AnketSil(anketid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("delete from anket where Id=?)", anketid)

	return err
}

func MesajYolla(baslik, mesaj, zaman string) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("insert into mesaj (Baslik,Mesaj,Zaman) values"+
		" (?,?,?)", baslik, mesaj, zaman)
	if err != nil {
		return err
	}

	return err
}

func AnketAra(text string) ([]Anket, error) {
	text = "%" + text + "%"
	var anketler []Anket
	db, err := baglan()
	if err != nil {
		return anketler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from anket where Ad like ? or Baslik like ? order by Id desc ", text, text)
	if err != nil {
		return anketler, err
	}

	return anketGetir(rows)
}

func YoneticiGetir() (string, string, error) {
	db, err := baglan()
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	var id int
	var nick, sifre string

	err = db.QueryRow("select * from yonetici where Id = ?", 1).Scan(&id, &nick, &sifre)

	return nick, sifre, err
}

func YoneticiGuncelle(nick, sifre string) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update yonetici set nick=?,sifre=? where Id = ?", nick, sifre, 1)

	return err
}

func anketGetir(rows *sql.Rows) ([]Anket, error) {
	var anketler []Anket
	var err error
	for rows.Next() {
		var anket Anket
		err = rows.Scan(&anket.Id, &anket.Ad, &anket.Baslik,
			&anket.Tip, &anket.OySayisi, &anket.YorumSayisi, &anket.Aktif, &anket.Zaman)
		if err != nil {
			return anketler, err
		}
		secenekler, err := SeceneklerGetir(anket.Id)
		if err != nil {
			return anketler, err
		}
		anket.Secenekler = secenekler
		m := make(map[string]interface{})
		m["secenekler"] = secenekler
		b, err := json.Marshal(m)
		if err != nil {
			return anketler, err
		}
		anket.SecenekJson = template.JS(string(b))
		anket.Yorumlar, err = YorumlarGetir(anket.Id)
		if err != nil {
			return anketler, err
		}
		anketler = append(anketler, anket)
	}

	return anketler, err
}
