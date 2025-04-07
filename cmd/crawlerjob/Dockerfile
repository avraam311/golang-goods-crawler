# Этап сборки
FROM golang:1.24.2-alpine AS build_base

# Установка необходимых пакетов
RUN apk --no-cache add bash git make gcc gettext musl-dev chromium chromium-chromedriver nss freetype harfbuzz ttf-freefont

# Установка рабочего каталога
WORKDIR /usr/local/src

# Копирование модулей Go
COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

# Копирование исходного кода и сборка приложения
COPY . ./
RUN go build -o ./bin/job cmd/crawlerjob/main.go

# Этап запуска
FROM alpine AS runner

# Копирование собранного приложения и конфигурационных файлов
COPY --from=build_base /usr/local/src/bin/job /
COPY ./.env /
COPY config/local.yaml /config/local.yaml

# Установка необходимых пакетов для запуска приложения
RUN apk --no-cache add chromium xvfb-run

# Установка переменной окружения для Chromium
ENV CHROME_BIN=/usr/bin/chromium-browser

# Команда для запуска приложения с Xvfb
CMD ["sh", "-c", "Xvfb :99 -screen 0 1920x1080x24 & /job"]