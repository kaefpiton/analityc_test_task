<p><h1>Сервис сбора действий пользователя</h1></p>
Для поднятия сервиса:<br>
Поднять docker-compose с окружением (Postgres + Prometheus):
   <p style= "text-indent: 25px;">  docker-compose up -d </p>

Установить (если не установлен) мигратор goose:
<p style= "text-indent: 25px;"> go install github.com/pressly/goose/v3/cmd/goose@latest </p>
Накатить миграцию в контейнер postgres:
<p style= "text-indent: 25px;"> goose postgres "host=localhost port=5432 user=pguser password=pguser dbname=pgdb sslmode=disable up</p>
(Как альтернатива можно не накатывать миграцию, а подключиться к контейнеру и вставить таблицу
    из cmd/migrations/20230704115302_actions.sql вручную)

В папке configs создать файл config.json с конфигом приложения:

            {
            "HttpServer": {
              "port":"8080",
              "metrics_port":"9092"
            },
            "Postgres": {
                "host": "localhost",
                "port": "5432",
                "user": "pguser",
                "password": "pguser",
                "db_name": "pgdb",
                "ssl_mode": "disable",
                "maxOpenConns":10,
                "maxIdleConns":5
            },
                "logger": {
                    "lvl": "info"
                }
            }

Запустить сервис сбора аналитики из корня проекта:
<p style= "text-indent: 25px;">  go run cmd/main.go (Выход из приложения - ctrl+c)</p>

Для показателей rps нужно постучаться на http://localhost:9090/graph и на языке promQL ввести в поле
Expression запрос:
<p style= "text-indent: 25px;">rate(REQUESTS_TOTAL[20s])</p>

    
        