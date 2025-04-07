# Web-crawler for Amazon goods
 
 Сервис предоставляющий API для получения названий товаром категории игровых мышек, спарсенных с Amazon.
 
 Используемые технологии:
 
 PostgreSQL (хранилище данных)
 Docker (запуск сервиса)
 Gin (веб-фреймворк)
 pgx (драйвер работы с PostgreSQL)
 cron (шедулер задач)
 golang-migrate (для миграций)
 chromedp(для парсинга js)
 
 Сервис написан по Clean Architecture
 
 ## Запуск
 1. Склонировать репо
 2. Создать файл .env в директории проекта и заполнить. В файле config/local.yaml можно конфигурировать http сервер
 3. Выполнить в терминале:
    ```bash
    make build
    make run 
    ``` 
 
 ## Спецификация API
 
 ### GET /goods
 
 Возвращает список названий товаров. Формат ответа такой:
 
 ```
 {
     "goods": [
         {
             "ID": 1,
             "Name": "Redragon Gaming Mouse, Wireless Mouse Gaming with 8000 DPI, PC Gaming Mice with Fire Button",
         },
      ]
 }
 ```