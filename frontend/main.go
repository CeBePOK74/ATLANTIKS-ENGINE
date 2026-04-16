package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"time"
)

const htmlInterface = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>ATLANTIKS ENGINE | Terminal</title>
    <style>
        body { background-color: #050505; color: #00ff41; font-family: 'Courier New', monospace; display: flex; justify-content: center; padding-top: 20px; margin: 0; }
        .main-frame { width: 800px; background: #0f0f0f; border: 1px solid #1a1a1a; padding: 25px; border-radius: 5px; box-shadow: 0 0 25px rgba(0,255,65,0.1); }
        h1 { color: #fff; text-align: center; letter-spacing: 5px; font-size: 24px; margin-bottom: 20px; border-bottom: 1px solid #1a1a1a; padding-bottom: 10px; }
        .query-input { width: 100%; padding: 18px; background: #000; border: 2px solid #00ff41; color: #fff; font-size: 20px; margin-bottom: 20px; box-sizing: border-box; outline: none; }
        .cat-group { margin-bottom: 8px; }
        .main-btn { width: 100%; padding: 14px; background: #111; border: 1px solid #333; color: #58a6ff; font-weight: bold; cursor: pointer; text-align: left; font-size: 14px; text-transform: uppercase; }
        .main-btn:hover { background: #1a1a1a; border-color: #58a6ff; }
        .menu-content { display: none; background: #080808; padding: 20px; border: 1px solid #1a1a1a; border-top: none; }
        .grid { display: grid; grid-template-columns: 1fr 1fr; gap: 10px; }
        label { font-size: 11px; color: #aaa; cursor: pointer; display: flex; align-items: center; padding: 5px; }
        label:hover { background: #111; color: #fff; }
        input[type="checkbox"] { margin-right: 10px; accent-color: #00ff41; }
        .start-btn { width: 100%; padding: 22px; background: #00ff41; color: #000; border: none; font-size: 20px; font-weight: bold; cursor: pointer; margin-top: 25px; }
        .start-btn:hover { background: #fff; box-shadow: 0 0 30px #00ff41; }
    </style>
</head>
<body>
    <div class="main-frame">
        <h1>ATLANTIKS ENGINE</h1>
        
        <input type="text" id="query" class="query-input" placeholder="ВВЕДИТЕ ЗАПРОС..." required autofocus>

        <div class="cat-group">
            <button type="button" class="main-btn" onclick="openMenu('tenders')">[+] ТЕНДЕРНЫЕ ПЛОЩАДКИ (ПОЛНЫЙ СПИСОК)</button>
            <div id="tenders" class="menu-content">
                <div class="grid">
                    <label><input type="checkbox" class="site-check" value="zakupki_gov"> ЕИС ЗАКУПКИ (44/223)</label>
                    <label><input type="checkbox" class="site-check" value="sber_ast"> СБЕРБАНК-АСТ</label>
                    <label><input type="checkbox" class="site-check" value="rts_tender"> РТС-ТЕНДЕР</label>
                    <label><input type="checkbox" class="site-check" value="roseltorg"> РОСЭЛТОРГ</label>
                    <label><input type="checkbox" class="site-check" value="b2b_center"> B2B-CENTER</label>
                    <label><input type="checkbox" class="site-check" value="fabrikant"> ФАБРИКАНТ</label>
                    <label><input type="checkbox" class="site-check" value="tektorg"> ТЭК-ТОРГ</label>
                    <label><input type="checkbox" class="site-check" value="etp_gpb"> ЭТП ГПБ</label>
                </div>
            </div>
        </div>

        <div class="cat-group">
            <button type="button" class="main-btn" onclick="openMenu('market')" style="color:#00ff41;">[+] МАРКЕТПЛЕЙСЫ И УСЛУГИ</button>
            <div id="market" class="menu-content">
                <div class="grid">
                    <label><input type="checkbox" class="site-check" value="avito"> AVITO</label>
                    <label><input type="checkbox" class="site-check" value="youla"> ЮЛА</label>
                    <label><input type="checkbox" class="site-check" value="yandex_market"> ЯНДЕКС.МАРКЕТ</label>
                    <label><input type="checkbox" class="site-check" value="ozon"> OZON</label>
                    <label><input type="checkbox" class="site-check" value="wildberries"> WILDBERRIES</label>
                    <label><input type="checkbox" class="site-check" value="vseinstrumenti"> ВСЕ ИНСТРУМЕНТЫ</label>
                    <label><input type="checkbox" class="site-check" value="profi"> ПРОФИ.РУ</label>
                    <label><input type="checkbox" class="site-check" value="uslugi_yandex"> ЯНДЕКС УСЛУГИ</label>
                </div>
            </div>
        </div>

        <div class="cat-group">
            <button type="button" class="main-btn" onclick="openMenu('global')" style="color:#fff;">[+] GLOBAL SEARCH</button>
            <div id="global" class="menu-content">
                <div class="grid">
                    <label><input type="checkbox" class="site-check" value="google"> GOOGLE</label>
                    <label><input type="checkbox" class="site-check" value="yandex"> YANDEX</label>
                </div>
            </div>
        </div>

        <button type="button" class="start-btn" onclick="startSearch()">ЗАПУСТИТЬ ПРОТОКОЛ</button>
        
    </div>

    <script>
        function openMenu(id) {
            var menu = document.getElementById(id);
            menu.style.display = (menu.style.display === "block") ? "none" : "block";
        }

        function startSearch() {
            const query = document.getElementById('query').value;
            if (!query) { alert("Введите запрос!"); return; }
            
            const encodedQuery = encodeURIComponent(query);
            const selected = document.querySelectorAll('.site-check:checked');
            
            const urls = {
                "google": "https://www.google.com/search?q=" + encodedQuery,
                "yandex": "https://yandex.ru/search/?text=" + encodedQuery,
                "avito": "https://www.avito.ru/chelyabinsk?q=" + encodedQuery,
                "youla": "https://youla.ru/all?q=" + encodedQuery,
                "ozon": "https://www.ozon.ru/search/?text=" + encodedQuery,
                "wildberries": "https://www.wildberries.ru/catalog/0/search.aspx?search=" + encodedQuery,
                "yandex_market": "https://market.yandex.ru/search?text=" + encodedQuery,
                "vseinstrumenti": "https://www.vseinstrumenti.ru/search/main/query=" + encodedQuery,
                "profi": "https://profi.ru/poisk/?name=" + encodedQuery,
                "uslugi_yandex": "https://uslugi.yandex.ru/search?text=" + encodedQuery,
                "zakupki_gov": "https://zakupki.gov.ru/epz/order/extendedsearch/results.html?searchString=" + encodedQuery,
                "sber_ast": "https://www.sberbank-ast.ru/purchaseList.aspx?search=" + encodedQuery,
                "rts_tender": "https://www.rts-tender.ru/poisk?search=" + encodedQuery,
                "roseltorg": "https://www.roseltorg.ru/search?q=" + encodedQuery,
                "b2b_center": "https://www.b2b-center.ru/market/?keyword=" + encodedQuery,
                "fabrikant": "https://www.fabrikant.ru/trades/procedure/search/?query=" + encodedQuery,
                "tektorg": "https://www.tektorg.ru/search?q=" + encodedQuery,
                "etp_gpb": "https://etp.gpb.ru/procedure/search?q=" + encodedQuery
            };

            selected.forEach(box => {
                const url = urls[box.value];
                if (url) {
                    window.open(url, '_blank');
                }
            });
        }
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.New("index").Parse(htmlInterface)
		tmpl.Execute(w, nil)
	})
	fmt.Println("ATLANTIKS ENGINE ONLINE: http://localhost:8080")
	go func() {
		time.Sleep(1 * time.Second)
		exec.Command("cmd", "/c", "start", "http://localhost:8080").Run()
	}()
	http.ListenAndServe(":8080", nil)
}