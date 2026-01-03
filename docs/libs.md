1. Swagger
    -   https://github.com/swaggo/swag
    -   go get -tool github.com/swaggo/swag/cmd/swag
    -   go tool swag init -g cmd/server/main.go -o ./docs - локально

2. Envs
    -   https://github.com/joho/godotenv
    -   go get github.com/joho/godotenv

3. Hot reloader
    -   https://github.com/air-verse/air
    -   go install github.com/air-verse/air@latest - глобально
    -   go get -tool github.com/air-verse/air@latest - локально
    -   Создать минимальный конфиг в корне проекта.

4. Validator
    -   https://github.com/go-playground/validator
    -   go get github.com/go-playground/validator/v10

5. Выбор ORM:
    -   https://github.com/d-tsuji/awesome-go-orms

6. ORM:
    -   https://github.com/go-gorm/gorm?tab=readme-ov-file
    -   go get -u gorm.io/gorm
    -   go get -u gorm.io/driver/postgres

7. Crypto:
    -   https://github.com/golang/crypto
    -   Встроена в стандартную библиотеку. Нужно просто указать импорт, например, import "golang.org/x/crypto/bcrypt".

8. JWT:
    -   https://github.com/golang-jwt/jwt
    -   go get -u github.com/golang-jwt/jwt/v5

9. SQL Mock:
    -   https://github.com/DATA-DOG/go-sqlmock
    -   go get github.com/DATA-DOG/go-sqlmock

10. Создание фейковых данных:
    -   https://github.com/brianvoe/gofakeit
    -   go get github.com/brianvoe/gofakeit/v7

11. Тестирование. Проверки. Хорошие текстовки:
    -   https://github.com/stretchr/testify
    -   go get github.com/stretchr/testify/assert

12. UUID:
    - https://github.com/google/uuid
    - go get github.com/google/uuid

13. HTTP Requests:
    -   https://github.com/go-resty/resty
    -   go get -u github.com/go-resty/resty/v2

14. Преобразование класса в URL Query:
    - https://github.com/google/go-querystring
    - go get github.com/google/go-querystring/query