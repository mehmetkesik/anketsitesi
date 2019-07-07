package main

import (
	"net/http"
	"runtime"
	"os/exec"
	"fmt"
	"log"
	"html/template"
	"strconv"
	"encoding/json"
	"math"
	"time"
	"strings"
	"./core"
)

func iletisim(w http.ResponseWriter, r *http.Request) {
	type Iletisim struct {
		Sayfa string
	}
	var veri Iletisim
	veri.Sayfa = "iletisim"

	html, err := template.ParseFiles("root/iletisim.html", "root/genel/navbar.html",
		"root/genel/menu.html", "root/genel/footer.html")
	hata(err)

	html.Execute(w, veri)
}

func anketler(w http.ResponseWriter, r *http.Request, link string) {
	type Anketler struct {
		Sayfa      string
		Anketler   []core.Anket
		Bolum      int
		Durum      bool
		Ileri      int
		IleriDurum bool
		Geri       int
		GeriDurum  bool
		Link       string
	}
	var veri Anketler
	veri.Sayfa = "anketler"
	var err error
	veri.Link = link

	if r.URL.Query().Get("sayfa") != "" {
		veri.Bolum, err = strconv.Atoi(r.URL.Query().Get("sayfa"))
		if err != nil || veri.Bolum < 0 {
			http.Redirect(w, r, "/notfound", 302)
			return
		}
	}
	if veri.Bolum == 0 {
		veri.Bolum = 1
	}

	veri.Geri = veri.Bolum - 1
	veri.GeriDurum = true
	veri.Ileri = veri.Bolum + 1
	veri.IleriDurum = true

	if veri.Bolum < 2 {
		veri.Geri = 1
		veri.GeriDurum = false
	}

	if veri.Bolum >= int(math.Ceil(float64(len(veri.Anketler))/12)) {
		veri.Ileri = veri.Bolum
		veri.IleriDurum = false
	}

	bit := veri.Bolum * 12
	bas := bit - 12

	switch link {
	case "oylanan":
		veri.Anketler, err = core.OylananGetir(bas, bit)
		hata(err)
		break;
	case "encok":
		veri.Anketler, err = core.EnCokGetir(bas, bit)
		hata(err)
		break;
	case "bitmis":
		veri.Anketler, err = core.BitmisGetir(bas, bit)
		hata(err)
		break;
	case "ara":
		text := r.URL.Query().Get("text")
		if text != "" {
			veri.Anketler, err = core.AnketAra(text)
			hata(err)
		}
		break
	default:
		http.Redirect(w, r, "/notfound", 302)
		return
	}

	if len(veri.Anketler) != 0 {
		veri.Durum = true
	}

	html, err := template.New("anketler.html").Funcs(template.FuncMap{
		"HesapUst": func(i int) bool {
			if i%2 == 0 {
				return true
			}
			return false
		},
		"HesapAlt": func(i int) bool {
			if i%2 == 1 || i == len(veri.Anketler)-1 {
				return true
			}
			return false
		},
	}).ParseFiles("root/anketler.html", "root/genel/navbar.html",
		"root/genel/menu.html", "root/genel/footer.html")
	hata(err)

	html.Execute(w, veri)
}

func anket(w http.ResponseWriter, r *http.Request) {
	type Anket struct {
		Sayfa      string
		Link       string
		Anket      core.Anket
		Bolum      int
		Ileri      int
		IleriDurum bool
		Geri       int
		GeriDurum  bool
	}
	var veri Anket
	veri.Sayfa = "anket"
	veri.Link = r.URL.Query().Get("anket")
	anketid, err := strconv.Atoi(veri.Link)
	hata(err)
	veri.Anket, err = core.AnketGetir(anketid)
	hata(err)

	if r.URL.Query().Get("sayfa") != "" {
		veri.Bolum, err = strconv.Atoi(r.URL.Query().Get("sayfa"))
		if err != nil || veri.Bolum < 0 {
			http.Redirect(w, r, "/notfound", 302)
			return
		}
	}
	if veri.Bolum == 0 {
		veri.Bolum = 1
	}

	veri.Geri = veri.Bolum - 1
	veri.GeriDurum = true
	veri.Ileri = veri.Bolum + 1
	veri.IleriDurum = true

	if veri.Bolum < 2 {
		veri.Geri = 1
		veri.GeriDurum = false
	}

	if veri.Bolum >= int(math.Ceil(float64(len(veri.Anket.Yorumlar))/12)) {
		veri.Ileri = veri.Bolum
		veri.IleriDurum = false
	}

	bit := veri.Bolum * 12
	bas := bit - 12

	if len(veri.Anket.Yorumlar) > bit {
		veri.Anket.Yorumlar = veri.Anket.Yorumlar[bas:bit]
	} else {
		veri.Anket.Yorumlar = veri.Anket.Yorumlar[bas:]
	}

	html, err := template.ParseFiles("root/anket.html", "root/genel/navbar.html",
		"root/genel/menu.html", "root/genel/footer.html")
	hata(err)

	html.Execute(w, veri)
}

func anasayfa(w http.ResponseWriter, r *http.Request) {
	type Anasayfa struct {
		Sayfa    string
		Anketler []core.Anket
	}
	var err error
	var veri Anasayfa
	hata(err)
	veri.Sayfa = "anasayfa"
	veri.Anketler, err = core.AktifSon5Getir()
	hata(err)

	html, err := template.ParseFiles("root/index.html", "root/genel/navbar.html",
		"root/genel/menu.html", "root/genel/slider.html", "root/genel/footer.html")
	hata(err)

	html.Execute(w, veri)
}

func web(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("anket") != "" {
		anket(w, r)
	} else if r.URL.Query().Get("anketler") != "" {
		link := r.URL.Query().Get("anketler")
		anketler(w, r, link)
	} else if r.URL.Query().Get("iletisim") != "" {
		iletisim(w, r)
	} else {
		anasayfa(w, r)
	}
}

func yonet(w http.ResponseWriter, r *http.Request) {
	type Yonet struct {
		AraDurum      bool
		Anketler      []core.Anket
		Anket         core.Anket
		AnketDurum    bool
		AnketEkle     bool
		AnasayfaDurum bool
		Nick          string
		Sifre         string
	}
	var veri Yonet

	if r.Method == "POST" {
		if r.FormValue("giris") != "" {
			nick := r.FormValue("nick")
			sifre := r.FormValue("sifre")
			http.SetCookie(w, &http.Cookie{
				Name:  "anketnick",
				Value: nick,
			})
			http.SetCookie(w, &http.Cookie{
				Name:  "anketsifre",
				Value: sifre,
			})
		} else if r.FormValue("anketguncelle") != "" {
			anketid, err := strconv.Atoi(r.FormValue("anketid"))
			hata(err)
			silinecekidler := r.FormValue("silinecekid")
			ad := r.FormValue("ad")
			baslik := r.FormValue("baslik")

			var toplamSecenek int
			var sxAd []string
			for name, value := range r.Form {
				if len(name) > 7 {
					if name[:7] == "secenek" {
						toplamSecenek += 1
					} else if name[:7] == "secenex" {
						toplamSecenek += 1
						sxAd = append(sxAd, value[0])
					}
				}
			}

			tip := 0
			if toplamSecenek > 2 {
				tip = 1
			}

			anket, err := core.AnketGetir(anketid)
			hata(err)
			anket.Ad = ad
			anket.Baslik = baslik
			anket.Tip = tip

			var secenekler []core.Secenek
			sidler := strings.Split(strings.Trim(silinecekidler, " "), ",")
			for _, secenek := range anket.Secenekler {
				var ekleme = true
				for i := 0; i < len(sidler); i++ {
					sidint, err := strconv.Atoi(sidler[i])
					if err != nil {
						break
					}
					if secenek.Id == sidint {
						ekleme = false
					}
				}
				if ekleme {
					secenekler = append(secenekler, secenek)
				}
			}

			for _, k := range sxAd {
				var secenek core.Secenek
				secenek.AnketId = anketid
				secenek.Ad = k
				secenek.Oy = 0
				secenekler = append(secenekler, secenek)
			}

			if len(secenekler) > 1 {
				anket.Secenekler = secenekler
			}

			core.AnketGuncelle(anket)
		} else if r.FormValue("anketyolla") != "" {
			ad := r.FormValue("ad")
			baslik := r.FormValue("baslik")

			var anket core.Anket
			anket.Ad = ad
			anket.Baslik = baslik

			var toplamSecenek int
			for name, value := range r.Form {
				if len(name) > 7 {
					if name[:7] == "secenek" {
						toplamSecenek += 1
						var secenek core.Secenek
						secenek.Ad = value[0]
						secenek.Oy = 1
						anket.Secenekler = append(anket.Secenekler, secenek)
					}
				}
			}

			anket.Tip = 0
			if toplamSecenek > 2 {
				anket.Tip = 1
			}

			anket.OySayisi = toplamSecenek
			anket.Aktif = 1
			t := time.Now()
			anket.Zaman = zamankontrol(t.Year()) + "-" + zamankontrol(int(t.Month())) + "-" + zamankontrol(t.Day()) + " " +
				zamankontrol(t.Hour()) + ":" + zamankontrol(t.Minute()) + ":" + zamankontrol(t.Second())
			core.AnketEkle(anket)

		} else if r.FormValue("profilguncelle") != "" {
			nick := r.FormValue("nick")
			sifre := r.FormValue("sifre")
			if nick != "" || sifre != "" {
				err := core.YoneticiGuncelle(nick, sifre)
				hata(err)

				http.SetCookie(w, &http.Cookie{
					Name:  "anketnick",
					Value: nick,
				})
				http.SetCookie(w, &http.Cookie{
					Name:  "anketsifre",
					Value: sifre,
				})
			}
		}
		http.Redirect(w, r, r.URL.String(), 302)
	}

	var oturum = true
	cookienick, err := r.Cookie("anketnick")
	if err != nil {
		oturum = false
	}
	cookiesifre, err := r.Cookie("anketsifre")
	if err != nil {
		oturum = false
	}
	nick, sifre, err := core.YoneticiGetir()
	hata(err)
	var html *template.Template
	if oturum && cookienick.Value == nick && cookiesifre.Value == sifre {
		veri.Nick = nick
		veri.Sifre = sifre
		if r.URL.Query().Get("ara") != "" {
			veri.AraDurum = true
			ara := r.URL.Query().Get("ara")
			veri.Anketler, err = core.AnketAra(ara)
			hata(err)
		} else if r.URL.Query().Get("anket") != "" {
			anketid, err := strconv.Atoi(r.URL.Query().Get("anket"))
			if err != nil {
				http.Redirect(w, r, "/notfound", 302)
			}
			veri.Anket, err = core.AnketGetir(anketid)
			veri.AnketDurum = true
		} else if r.URL.Query().Get("ekle") != "" {
			veri.AnketEkle = true
		} else if r.URL.Query().Get("cikis") != "" {
			http.SetCookie(w, &http.Cookie{
				Name:  "anketnick",
				Value: "",
			})
			http.SetCookie(w, &http.Cookie{
				Name:  "anketsifre",
				Value: "",
			})
			http.Redirect(w, r, "/yonet", 302)
		} else {
			veri.AnasayfaDurum = true
		}
		html, err = template.New("yonet.html").Funcs(template.FuncMap{
			"HesapUst": func(i int) bool {
				if i%2 == 0 {
					return true
				}
				return false
			},
			"HesapAlt": func(i int) bool {
				if i%2 == 1 || i == len(veri.Anketler)-1 {
					return true
				}
				return false
			},
		}).ParseFiles("root/yonet.html")
		hata(err)
	} else {
		html, err = template.ParseFiles("root/yonetgiris.html")
		hata(err)
	}

	html.Execute(w, veri)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("root/notfound.html")
	hata(err)
	html.Execute(w, nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.FormValue("oykullan") != "" {
			m := make(map[string]interface{})
			anketid, err := strconv.Atoi(r.FormValue("anketid"))
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}
			secenekid, err := strconv.Atoi(r.FormValue("secenekid"))
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			err = core.OyVer(anketid, secenekid)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}
			err = core.OyArttir(anketid)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			m["durum"] = "ok"
			m["sonuc"] = ""
			b, _ := json.Marshal(m)
			fmt.Fprint(w, string(b))
			return
		}

		if r.FormValue("yorumyolla") != "" {
			m := make(map[string]interface{})
			anketidstr := r.FormValue("anketid")
			nick := r.FormValue("nick")
			yorumstr := r.FormValue("yorum")
			if anketidstr == "" || nick == "" || yorumstr == "" {
				m["durum"] = "hata"
				m["sonuc"] = "Verilerden herhangi biri boş olamaz!"
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}
			anketid, err := strconv.Atoi(anketidstr)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			if len(nick) > 20 || len(yorumstr) > 200 {
				m["durum"] = "hata"
				m["sonuc"] = "Veri çok uzun!"
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			var yorum core.Yorum
			yorum.AnketId = anketid
			yorum.Nick = nick
			yorum.Yorum = yorumstr
			t := time.Now()
			yorum.Zaman = zamankontrol(t.Year()) + "-" + zamankontrol(int(t.Month())) + "-" + zamankontrol(t.Day()) + " " +
				zamankontrol(t.Hour()) + ":" + zamankontrol(t.Minute()) + ":" + zamankontrol(t.Second())
			lastid, err := core.YorumEkle(yorum)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}
			err = core.YorumArttir(anketid)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			m["durum"] = "ok"
			m["sonuc"] = lastid
			b, _ := json.Marshal(m)
			fmt.Fprint(w, string(b))
			return
		} //yorum yolla son

		if r.FormValue("mesajyolla") != "" {
			m := make(map[string]interface{})
			baslik := r.FormValue("baslik")
			mesaj := r.FormValue("mesaj")
			if baslik == "" || mesaj == "" {
				m["durum"] = "hata"
				m["sonuc"] = "Verilerden herhangi biri boş olamaz!"
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			if len(baslik) > 50 || len(mesaj) > 200 {
				m["durum"] = "hata"
				m["sonuc"] = "Veri çok uzun!"
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			t := time.Now()
			zaman := zamankontrol(t.Year()) + "-" + zamankontrol(int(t.Month())) + "-" + zamankontrol(t.Day()) + " " +
				zamankontrol(t.Hour()) + ":" + zamankontrol(t.Minute()) + ":" + zamankontrol(t.Second())
			err := core.MesajYolla(baslik, mesaj, zaman)
			if err != nil {
				m["durum"] = "hata"
				m["sonuc"] = err.Error()
				b, _ := json.Marshal(m)
				fmt.Fprint(w, string(b))
				return
			}

			m["durum"] = "ok"
			b, _ := json.Marshal(m)
			fmt.Fprint(w, string(b))
			return
		}

	} //post son
}

func main() {
	openbrowser("http://localhost:4182")
	http.Handle("/root/", http.StripPrefix("/root/", http.FileServer(http.Dir("root"))))
	http.HandleFunc("/", web)
	http.HandleFunc("/post", post)
	http.HandleFunc("/notfound", notfound)
	http.HandleFunc("/yonet", yonet)
	http.ListenAndServe(":4182", nil)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func hata(err error) {
	if err != nil {
		panic(err)
	}
}

func zamankontrol(i int) string {
	str := strconv.Itoa(i)
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}
