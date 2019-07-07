package core

func OyVer(anketid, secenekid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update secenek set Oy = Oy+1 where AnketId = ? and Id = ?",
		anketid, secenekid)

	return err
}

func SecenekGetir(anketid, secenekid int) (Secenek, error) {
	var secenek Secenek
	db, err := baglan()
	if err != nil {
		return secenek, err
	}
	defer db.Close()

	err = db.QueryRow("select * from secenek where AnketId = ? and Id = ?", anketid, secenekid).
		Scan(&secenek.Id, &secenek.Ad, &secenek.Oy, &secenek.AnketId)

	return secenek, err
}

func SeceneklerGetir(anketid int) ([]Secenek, error) {
	var secenekler []Secenek
	db, err := baglan()
	if err != nil {
		return secenekler, err
	}
	defer db.Close()

	rows, err := db.Query("select * from secenek where AnketId = ? order by Oy desc", anketid)
	if err != nil {
		return secenekler, err
	}

	for rows.Next() {
		var secenek Secenek
		err = rows.Scan(&secenek.Id, &secenek.Ad, &secenek.Oy, &secenek.AnketId)
		if err != nil {
			return secenekler, err
		}
		secenekler = append(secenekler, secenek)
	}

	return secenekler, err
}

func SecenekEkle(secenek Secenek) (int, error) {
	db, err := baglan()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("insert into secenek (Ad,Oy,AnketId) values (?,?,?)", secenek.Ad, secenek.Oy, secenek.AnketId)
	if err != nil {
		return -1, err
	}
	lastid, err := result.LastInsertId()
	return int(lastid), err
}

func SecenekGuncelle(secenek Secenek) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update secenek set Ad=?,Oy=? where AnketId = ? and Id = ?",
		secenek.Ad, secenek.Oy, secenek.AnketId, secenek.Id)

	return err
}

func SecenekSil(anketid, secenekid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("delete from secenek where AnketId = ? and Id = ?", anketid, secenekid)

	return err
}

func SeceneklerSil(anketid int) error {
	db, err := baglan()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("delete from secenek where AnketId = ?", anketid)

	return err
}
