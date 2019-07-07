function cizx(nesne) {
    //ilk girişteki anket durumunu çizdirecek
    let maxOy = nesne.secenekler[0].Oy;
    for (let i = 0; i < nesne.secenekler.length; i++) {
        let oy = nesne.secenekler[i].Oy;
        let yuzde = (100 * oy) / maxOy;
        nesne.cizim.beginPath();
        if (i == nesne.kalinCiz) {
            nesne.cizim.lineWidth = nesne.satirGenisligi * 3 / 2;
            if (nesne.aktif) {
                nesne.cizim.strokeStyle = "#1bc98e";
            } else {
                nesne.cizim.strokeStyle = "#1ca8dd";
            }
        } else {
            nesne.cizim.lineWidth = nesne.satirGenisligi;
            nesne.cizim.strokeStyle = "#1ca8dd";
        }
        nesne.cizim.lineCap = "round";
        nesne.cizim.moveTo(nesne.aralik / 2, nesne.aralik * (i + 1));
        nesne.cizim.lineTo((nesne.aralik / 2) + ((nesne.genislik * yuzde) / 100), nesne.aralik * (i + 1));
        nesne.cizim.stroke();

    }
    //---------------------------//yazılar

    for (let i = 0; i < nesne.secenekler.length; i++) {
        nesne.cizim.beginPath();
        nesne.cizim.fillStyle = "#ffffff";
        if (i == nesne.kalinCiz) {
            nesne.cizim.font = "14px Arial";
        } else {
            nesne.cizim.font = "13px Arial";
        }
        let yuzde = Math.round((100 * nesne.secenekler[i].Oy) / nesne.toplamOy);
        nesne.cizim.fillText(nesne.secenekler[i].Ad + ": " + nesne.secenekler[i].Oy + ", %" + yuzde, nesne.aralik / 2, (nesne.aralik * (i + 1)) + 3.5);
        nesne.cizim.fill();
    }
}

function temizlex(nesne) {
    nesne.cizim.beginPath();
    nesne.cizim.fillStyle = nesne.arkaplan;//"white";//"rgb(37, 40, 48)";
    nesne.cizim.fillRect(0, 0, nesne.tuval.width, nesne.tuval.height);
    nesne.cizim.closePath();
    nesne.cizim.fill();
}


function Runx(id, secenekler, arkaplanRenk, aktif) {
    idsaf = id;
    id = "tuval" + id;

    if (arkaplanRenk == null) {
        arkaplanRenk = "white";
    }
    if (aktif == null) {
        aktif = false;
    }

    var obj = JSON.parse(secenekler);
    secenekler = obj.secenekler;

    var nesne = {
        tuval: null,
        cizim: null,
        width: null,
        genislik: null,
        uzunluk: null,
        aralik: null,
        satirGenisligi: null,
        arkaplan: null,
        oylar: null,
        toplamOy: null,
        yazilar: null,
        kalinCiz: null,
        aktif: null,
        secenekler: null
    };

    nesne.tuval = document.getElementById(id);
    nesne.cizim = nesne.tuval.getContext("2d");
    nesne.tuval.setAttribute("width", 250);
    nesne.width = nesne.tuval.getAttribute("width");
    nesne.satirGenisligi = 20;//lineWidth yani oy çizgisinin genişliği
    nesne.aralik = nesne.satirGenisligi * 2;//köşe ara oy çizgileri aralığı
    nesne.genislik = nesne.width - nesne.aralik;
    nesne.arkaplan = arkaplanRenk;
    nesne.secenekler = secenekler;
    nesne.tuval.setAttribute("height", (nesne.secenekler.length + 1) * nesne.aralik);
    nesne.height = nesne.tuval.getAttribute("height");
    nesne.aktif = aktif;
    nesne.kalinCiz = -1;

    nesne.toplamOy = 0;
    for (var i = 0; i < nesne.secenekler.length; i++) {
        nesne.toplamOy += nesne.secenekler[i].Oy;
    }

    temizlex(nesne);
    cizx(nesne);

    nesne.tuval.addEventListener("mousemove", function (e) {
        var konumx = parseFloat(nesne.tuval.getBoundingClientRect().left);
        var konumy = parseFloat(nesne.tuval.getBoundingClientRect().top);
        var posy = e.y;

        for (let i = 0; i < nesne.secenekler.length; i++) {
            if (posy > konumy + parseInt(nesne.aralik * (i + 1)) - (nesne.satirGenisligi / 2) &&
                posy < konumy + (nesne.aralik * (i + 1)) + (nesne.satirGenisligi / 2)) {
                nesne.kalinCiz = i;
                temizlex(nesne);
                cizx(nesne);
                if (localStorage.getItem(id) == null && nesne.aktif) {
                    document.getElementById("body").style.cursor = "pointer";
                }
                break;
            } else {
                nesne.kalinCiz = -1;
                temizlex(nesne);
                cizx(nesne);
                document.getElementById("body").style.cursor = "context-menu";
            }
        }

    });

    if (nesne.aktif) {
        nesne.tuval.addEventListener("click", function (e) {
            document.getElementById("deneme").innerHTML = nesne.kalinCiz + " e tıklandı";
            if (typeof(Storage) !== "undefined") {
                if (localStorage.getItem(id) == null) {
                    localStorage.setItem(id, nesne.kalinCiz + ""); //Burada hangi ankete hangi oyu verdiğini kaydediyoruz.
                    let o = {
                        "oykullan": "tiktak",
                        "anketid": idsaf,
                        "secenekid": nesne.secenekler[nesne.kalinCiz].Id
                    };
                    $.post("/post", o, function (result, status) {
                        if (status == "success") {
                            nesne.secenekler[nesne.kalinCiz].Oy += 1;
                            temizlex(nesne);
                            cizx(nesne);
                        } else {
                            alert("hata var");
                        }
                    }, "json");
                    document.getElementById("deneme").innerHTML += " oy kullanıldı";
                } else {
                    document.getElementById("deneme").innerHTML += " daha önce oy kullanılmış";
                }
            } else {
                alert("Tarayıcınız çok eski bu yüzden oy kullanamazsınız :(");
            }
        });
    }
}

//Runx("tuval",[500,150,18,5],["Galatasaray","Fenerbahçe","Beşiktaş","Trabzonspor"],"rgb(37, 40, 48)");