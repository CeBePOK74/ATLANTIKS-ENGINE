package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/playwright-community/playwright-go"
)

func main() {
	StartDidEngine("Плитка", "avito", "74")
}

func StartDidEngine(query string, source string, region string) {
	fmt.Println("\n=== ЗАПУСК УНИВЕРСАЛЬНОГО ДВИЖКА ATLANTIKS-EVOLUTION ===")

	// Автоматическая установка драйвера, если его нет
	err := playwright.Install()
	if err != nil {
		log.Printf("[ИНФО] Установка драйверов: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("[ОШИБКА] Не удалось запустить Playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args: []string{
			"--no-sandbox",
			"--disable-setuid-sandbox",
			"--disable-dev-shm-usage",
		},
	})
	if err != nil {
		log.Fatalf("[ОШИБКА] Не удалось запустить Chromium: %v", err)
	}

	page, _ := browser.NewPage()
	targetURL := "https://www.avito.ru/rossiya?q=" + url.QueryEscape(query)

	fmt.Printf("[ПЕРЕХОД] Открываю: %s\n", targetURL)
	_, err = page.Goto(targetURL)
	if err != nil {
		fmt.Printf("[ОШИБКА] Не удалось загрузить страницу: %v\n", err)
	} else {
		fmt.Println("====================================================")
		fmt.Println("Система ATLANTIKS активна. Робот успешно зашел на сайт.")
	}

	browser.Close()
	pw.Stop()
}