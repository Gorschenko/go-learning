- Swagger
    -   https://github.com/swaggo/swag
    -   go get -tool github.com/swaggo/swag/cmd/swag
    -   go tool swag init -g cmd/server/main.go -o ./docs - локально

- Envs
    -   https://github.com/joho/godotenv
    -   go get github.com/joho/godotenv

- Hot reloader
    -   https://github.com/air-verse/air
    -   go install github.com/air-verse/air@latest - глобально
    -   go get -tool github.com/air-verse/air@latest - локально
    -   Создать минимальный конфиг в корне проекта.

- Выбор ORM:
    -   https://github.com/d-tsuji/awesome-go-orms

- ORM:
    -   https://github.com/go-gorm/gorm
    -   go get -u gorm.io/gorm
    -   go get -u gorm.io/driver/postgres

- Cron:
    -   https://github.com/robfig/cron
    -   go get github.com/robfig/cron/v3@v3.0.0

- Redis:
    -  https://github.com/redis/go-redis
    -  go get github.com/redis/go-redis/v9

- MQTT:
    -  https://github.com/eclipse-paho/paho.mqtt.golang
    -  go get github.com/eclipse/paho.mqtt.golang

- Validator
    -   https://github.com/go-playground/validator
    -   go get github.com/go-playground/validator/v10

- Crypto:
    -   https://github.com/golang/crypto
    -   Встроена в стандартную библиотеку. Нужно просто указать импорт, например, import "golang.org/x/crypto/bcrypt".

- JWT:
    -   https://github.com/golang-jwt/jwt
    -   go get -u github.com/golang-jwt/jwt/v5

- SQL Mock:
    -   https://github.com/DATA-DOG/go-sqlmock
    -   go get github.com/DATA-DOG/go-sqlmock

- Создание фейковых данных:
    -   https://github.com/brianvoe/gofakeit
    -   go get github.com/brianvoe/gofakeit/v7

- Тестирование. Проверки. Хорошие текстовки:
    -   https://github.com/stretchr/testify
    -   go get github.com/stretchr/testify/assert

- UUID:
    - https://github.com/google/uuid
    - go get github.com/google/uuid

- HTTP Requests:
    -   https://github.com/go-resty/resty
    -   go get -u github.com/go-resty/resty/v2

- Преобразование класса в URL Query:
    - https://github.com/google/go-querystring
    - go get github.com/google/go-querystring/query