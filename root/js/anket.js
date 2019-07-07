function degerHesapla(secenekler) {
    var obj = {
        s1: null,
        s2: null,
        toplamOy: null,
        birinciYuzde: null,
        ikinciYuzde: null,
        ikinciBaslangic: null,
        text1: null,
        text2: null
    };

    obj.s1 = secenekler[0];
    obj.s2 = secenekler[1];
    obj.toplamOy = obj.s1.Oy + obj.s2.Oy;
    obj.birinciYuzde = ((obj.s1.Oy * 100) / obj.toplamOy).toFixed(2); //toFixed() virgülden sonra 2 basamak alır.
    obj.ikinciYuzde = (100 - obj.birinciYuzde).toFixed(2);
    obj.ikinciBaslangic = (2 * Math.PI * obj.birinciYuzde) / 100;
    obj.text1 = obj.s1.Ad + ": " + obj.s1.Oy + ", %" + obj.birinciYuzde;
    obj.text2 = obj.s2.Ad + ": " + obj.s2.Oy + ", %" + obj.ikinciYuzde;

    return obj;
}

function ciz(nesne) {
    var obj = degerHesapla(nesne.secenekler);
    //ilk girişteki anket durumunu çizdirecek
    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, Math.PI / 180, obj.ikinciBaslangic);
    nesne.cizim.strokeStyle = "#1ca8dd";
    nesne.cizim.lineWidth = nesne.satirGenisligi;
    nesne.cizim.stroke();

    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, obj.ikinciBaslangic + (Math.PI / 180), Math.PI * 2);
    nesne.cizim.strokeStyle = "#1bc98e";
    nesne.cizim.lineWidth = nesne.satirGenisligi;
    nesne.cizim.stroke();
}

function kalinciz1(nesne, text) {
    var obj = degerHesapla(nesne.secenekler);
    //burası ilk değeri kalın çizdirecek ve kaç oy aldığını gösterecek
    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, Math.PI / 180, obj.ikinciBaslangic);
    nesne.cizim.strokeStyle = "#1ca8dd";
    nesne.cizim.lineWidth = nesne.satirGenisligi + nesne.satirarti;
    nesne.cizim.stroke();

    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, obj.ikinciBaslangic + (Math.PI / 180), Math.PI * 2);
    nesne.cizim.strokeStyle = "#1bc98e";
    nesne.cizim.lineWidth = nesne.satirGenisligi;
    nesne.cizim.stroke();

    nesne.cizim.font = "12px Arial";
    var txtolcum = nesne.cizim.measureText(obj.text1);
    var txtgenislik = txtolcum.width;
    var txtyukseklik = parseInt(nesne.cizim.font);

    nesne.cizim.beginPath();
    var merkez = nesne.width / 2;
    var xgenislik = txtgenislik / 2 + nesne.yaricap / 10;//yaricap / 2;
    var ygenislik = txtyukseklik / 2 + nesne.yaricap / 10;//yaricap / 4;
    nesne.cizim.lineTo(merkez - xgenislik, merkez - ygenislik);
    nesne.cizim.lineTo(merkez + xgenislik, merkez - ygenislik);
    nesne.cizim.lineTo(merkez + xgenislik, (merkez - ygenislik / 4));
    nesne.cizim.lineTo(merkez + (xgenislik * 4 / 3), (merkez + ygenislik / 1.5));
    nesne.cizim.lineTo(merkez + xgenislik, (merkez + ygenislik / 4));
    nesne.cizim.lineTo(merkez + xgenislik, merkez + ygenislik);
    nesne.cizim.lineTo(merkez - xgenislik, merkez + ygenislik);
    nesne.cizim.lineTo(merkez - xgenislik, merkez - ygenislik);
    nesne.cizim.fillStyle = "rgba(28,168,221,0.5)";
    nesne.cizim.fill();
    nesne.cizim.closePath();

    nesne.cizim.beginPath();
    nesne.cizim.fillStyle = "#ffffff";
    nesne.cizim.fillText(obj.text1, merkez - xgenislik + (nesne.yaricap / 20), merkez + 4);
    nesne.cizim.closePath();
}

function kalinciz2(nesne, text) {
    var obj = degerHesapla(nesne.secenekler);
    //burası ikinci değeri kalın çizdirecek ve kaç oy aldığını gösterecek
    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, Math.PI / 180, obj.ikinciBaslangic);
    nesne.cizim.strokeStyle = "#1ca8dd";
    nesne.cizim.lineWidth = nesne.satirGenisligi;
    nesne.cizim.stroke();

    nesne.cizim.beginPath();
    nesne.cizim.arc(nesne.width / 2, nesne.width / 2, nesne.yaricap, obj.ikinciBaslangic + (Math.PI / 180), Math.PI * 2);
    nesne.cizim.strokeStyle = "#1bc98e";
    nesne.cizim.lineWidth = nesne.satirGenisligi + nesne.satirarti;
    nesne.cizim.stroke();


    nesne.cizim.font = "12px Arial";
    var txtolcum = nesne.cizim.measureText(obj.text2);
    var txtgenislik = txtolcum.width;
    var txtyukseklik = parseInt(nesne.cizim.font);


    nesne.cizim.beginPath();
    var merkez = nesne.width / 2;
    var xgenislik = txtgenislik / 2 + nesne.yaricap / 10;//yaricap / 2;
    var ygenislik = txtyukseklik / 2 + nesne.yaricap / 10;//yaricap / 4;
    nesne.cizim.lineTo(merkez - xgenislik, merkez - ygenislik);
    nesne.cizim.lineTo(merkez + xgenislik, merkez - ygenislik);
    nesne.cizim.lineTo(merkez + xgenislik, (merkez - ygenislik / 4));
    nesne.cizim.lineTo(merkez + (xgenislik * 4 / 3), (merkez - ygenislik / 1.5));
    nesne.cizim.lineTo(merkez + xgenislik, (merkez + ygenislik / 4));
    nesne.cizim.lineTo(merkez + xgenislik, merkez + ygenislik);
    nesne.cizim.lineTo(merkez - xgenislik, merkez + ygenislik);
    nesne.cizim.lineTo(merkez - xgenislik, merkez - ygenislik);
    nesne.cizim.fillStyle = "rgba(27,201,142,0.5)";
    nesne.cizim.fill();
    nesne.cizim.closePath();

    nesne.cizim.beginPath();
    nesne.cizim.fillStyle = "#ffffff";
    nesne.cizim.fillText(obj.text2, merkez - xgenislik + (nesne.yaricap / 20), merkez + 4);
    nesne.cizim.closePath();

}

function temizle(nesne, renk) {
    nesne.cizim.beginPath();
    nesne.cizim.fillStyle = renk;//"white";//"rgb(37, 40, 48)";
    nesne.cizim.fillRect(0, 0, nesne.tuval.width, nesne.tuval.height);
    nesne.cizim.closePath();
    nesne.cizim.fill();
}


function Run(id, secenekler, arkaplanRenk, aktif) {
    var idsaf = id;
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
        yaricap: null,
        satirarti: null,
        satirGenisligi: null,
        merkezx: null,
        merkezy: null,
        birinciYuzde: null,
        ikinciBaslangic: null,
        secenekler: null,
        aktif: null
    };

    /*var s1 = secenekler[0];
    var s2 = secenekler[1];
    var toplamOy = s1.Oy + s2.Oy;
    var ilkYuzde = ((s1.Oy * 100) / toplamOy).toFixed(2); //toFixed() virgülden sonra 2 basamak alır.
    var ikinciYuzde = (100 - ilkYuzde).toFixed(2);
    var text1 = s1.Ad + ": " + s1.Oy + ", %" + ilkYuzde;
    var text2 = s2.Ad + ": " + s2.Oy + ", %" + ikinciYuzde;*/

    nesne.tuval = document.getElementById(id);
    nesne.cizim = nesne.tuval.getContext("2d");
    nesne.width = nesne.tuval.getAttribute("width");
    nesne.yaricap = (nesne.width / 2) - (nesne.width / 8);
    nesne.satirarti = nesne.width / 10;
    nesne.satirGenisligi = (nesne.satirarti / 6) * 5;
    nesne.satirarti = nesne.satirarti - nesne.satirGenisligi;
    nesne.merkezx = (nesne.width / 2) + parseFloat(nesne.tuval.getBoundingClientRect().left);
    nesne.merkezy = (nesne.width / 2) + parseFloat(nesne.tuval.getBoundingClientRect().top);
    nesne.secenekler = secenekler;
    nesne.aktif = aktif;
    nesne.kalinCiz = -1;

    temizle(nesne, arkaplanRenk);
    ciz(nesne);

    nesne.tuval.addEventListener("mousemove", function (e) {
        nesne.merkezx = (nesne.width / 2) + parseFloat(nesne.tuval.getBoundingClientRect().left);
        nesne.merkezy = (nesne.width / 2) + parseFloat(nesne.tuval.getBoundingClientRect().top);
        var posx = e.x;
        var posy = e.y;
        var uzaklık = Math.sqrt(Math.pow(nesne.merkezx - posx, 2) + Math.pow(nesne.merkezy - posy, 2));
        var egim = (-1 * (posy - nesne.merkezy)) / (posx - nesne.merkezx);

        var radyandegeri = Math.atan(egim);
        var gercekradyandegeri = 0;

        if (posy <= nesne.merkezy && posx >= nesne.merkezx) {
            //1. bölge
            gercekradyandegeri = (Math.PI * 2) - Math.abs(radyandegeri);
        } else if (posy <= nesne.merkezy && posx <= nesne.merkezx) {
            // 2. bölge
            gercekradyandegeri = (Math.PI * 2) - ((Math.PI / 2) + radyandegeri + (Math.PI / 2));
        } else if (posy >= nesne.merkezy && posx <= nesne.merkezx) {
            // 3. bölge
            gercekradyandegeri = (Math.PI * 2) - (Math.abs(radyandegeri) + Math.PI);
        } else {
            // 4. bölge
            gercekradyandegeri = (Math.PI * 2) - ((Math.PI / 2) + radyandegeri + (3 * (Math.PI / 2)));
        }

        // 1 radyan 57,2958 derecedir

        var obj = degerHesapla(nesne.secenekler);
        if (uzaklık > nesne.yaricap - (nesne.satirGenisligi / 2) && uzaklık < nesne.yaricap + (nesne.satirGenisligi / 2)) {
            if (gercekradyandegeri >= 0 && gercekradyandegeri < obj.ikinciBaslangic) {
                temizle(nesne, arkaplanRenk);
                kalinciz1(nesne);
                nesne.kalinCiz = nesne.secenekler[0].Id;
            } else {
                temizle(nesne, arkaplanRenk);
                kalinciz2(nesne);
                nesne.kalinCiz = nesne.secenekler[1].Id;
            }
            if (localStorage.getItem(id) == null && nesne.aktif) {
                document.getElementById("body").style.cursor = "pointer";
            }
        } else {
            temizle(nesne, arkaplanRenk);
            ciz(nesne);
            nesne.kalinCiz = -1;
            document.getElementById("body").style.cursor = "context-menu";
        }
    });

    if (nesne.aktif) {
        nesne.tuval.addEventListener("click", function (e) {
            document.getElementById("deneme").innerHTML = nesne.kalinCiz + " e tıklandı";
            if (typeof(Storage) !== "undefined") {
                if (localStorage.getItem(id) == null) {
                    localStorage.setItem(id, nesne.kalinCiz + ""); //Burada hangi ankete hangi oyu verdiğini kaydediyoruz.
                    var obj = {
                        "oykullan": "tiktak",
                        "anketid": idsaf,
                        "secenekid": nesne.kalinCiz
                    };
                    $.post("/post", obj, function (result, status) {
                        if (status == "success") {
                            temizle(nesne, arkaplanRenk);
                            switch (nesne.kalinCiz) {
                                case nesne.secenekler[0].Id:
                                    nesne.secenekler[0].Oy += 1;
                                    kalinciz1(nesne);
                                    break;
                                case nesne.secenekler[1].Id:
                                    nesne.secenekler[1].Oy += 1;
                                    kalinciz2(nesne);
                                    break;
                                default:
                                    ciz(nesne);
                                    break;
                            }
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