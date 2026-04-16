package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"net/url"
)

func StartOldEngine(query string, source string, region string) {
	fmt.Println("\n=== ЗАПУСК УНИВЕРСАЛЬНОГО ДВИЖКА ATLANTIKS ===")
	
	// Подготовка данных
	safeQuery := url.QueryEscape(query) // Кодируем запрос для URL (чтобы пробелы не ломали ссылку)

	pw, err := playwright.Run()
	if err != nil {
		fmt.Println("[ИНФО] Настройка драйверов...")
		playwright.Install()
		pw, _ = playwright.Run()
	}

	browser, _ := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	page, _ := browser.NewPage()

	// ВЫБОР ПЛОЩАДКИ
	var targetURL string
	fmt.Printf("[АНАЛИЗ] Запрос: %s | Источник: %s\n", query, source)

	switch source {
	case "avito":
		targetURL = "https://www.avito.ru/rossiya?q=" + safeQuery
	case "wb":
		targetURL = "https://www.wildberries.ru/catalog/0/search.aspx?search=" + safeQuery
	case "ozon":
		targetURL = "https://www.ozon.ru/search/?text=" + safeQuery
	case "yandex":
		targetURL = "https://market.yandex.ru/search?text=" + safeQuery
	case "saby":
		targetURL = "https://saby.ru/tenders"
	case "global":
		targetURL = "https://www.google.com/search?q=" + safeQuery
	default:
		targetURL = "https://google.com/search?q=" + safeQuery
	}

	fmt.Printf("[ПЕРЕХОД] Открываю: %s\n", targetURL)
	page.Goto(targetURL)

	// Авто-заполнение поиска для Saby (если выбран тендер)
	if source == "saby" && query != "" {
		fmt.Println("[ИНФО] Ввожу запрос в поиск СБИС...")
		// Добавляем скобки к Keyboard()
		err := page.Fill("input[type='search']", query)
		if err == nil {
			page.Keyboard().Press("Enter")
		}
	}

	fmt.Println("====================================================")
	fmt.Println("Система активна. Окно браузера под управлением Go.")
}