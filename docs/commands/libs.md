## Загрузить библиотеку:
```
<!-- Старый -->
go get github.com/joho/godotenv

<!-- Новый -->
go install github.com/joho/godotenv
```

## Удалить библиотеку и актуализировать зависимости:

```
go clean -i github.com/joho/godotenv
go mod tidy
<!-- Удаляем кэш на всякий случай-->
go clean -modcache
```

## Удалить глобальную библиотеку

```
Get-Command air
del "path from terminal"
```