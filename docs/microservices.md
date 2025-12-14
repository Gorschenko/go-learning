# Создаем корневую структуру
mkdir -p myproject/{services,shared}
cd myproject

# Инициализируем workspace
go work init

# Создаем общий модуль
cd shared
go mod init shared
cd ..

# Создаем сервис 1
mkdir -p services/service1/cmd
cd services/service1
go mod init service1

# Создаем сервис 2  
mkdir -p services/service2/cmd
cd ../service2
go mod init service2

# Возвращаемся в корень и настраиваем workspace
cd ../..
go work use ./shared
go work use ./services/service1
go work use ./services/service2

# Итоговая структура

myproject/
├── go.work
├── shared/
│   ├── go.mod
│   └── pkg/
└── services/
    ├── service1/
    │   ├── cmd/
    │   │   └── main.go
    │   └── go.mod
    └── service2/
        ├── cmd/
        │   └── main.go
        └── go.mod