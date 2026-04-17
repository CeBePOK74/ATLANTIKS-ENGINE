FROM golang:1.22-bookworm

# Установка необходимых библиотек для браузеров
RUN apt-get update && apt-get install -y \
    libnss3 libatk-bridge2.0-0 libxcomposite1 libxdamage1 libxrandr2 libgbm1 libasound2 libpangocairo-1.0-0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Сборка
# Установка ВСЕХ необходимых библиотек (включая те, что в рамке на скриншоте)
RUN apt-get update && apt-get install -y \
    libnss3 \
    libatk-bridge2.0-0 \
    libxcomposite1 \
    libxdamage1 \
    libxrandr2 \
    libgbm1 \
    libasound2 \
    libpangocairo-1.0-0 \
    libxcursor1 \
    libgtk-3-0 \
    libgdk-pixbuf-2.0-0 \
    libcairo-gobject2 \
    && rm -rf /var/lib/apt/lists/*

# Прямая команда на установку браузера (без поиска путей)
RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps chromium

CMD ["./main"]