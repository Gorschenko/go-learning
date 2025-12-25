### ПО:

1. ValidateQueryMiddleware
2. + ValidateParamsMiddleware
3. + ValidateBodyMiddleware
4. - CorrelationIdMiddleware. Не до конца сделан. Непонятно как с этим работать в других классах. Придется всегда передавать context.
5. - LogMiddleware. Нужен хороший полноценное ПО, которое будет логировать все.
6. + HttpApi
7. Возвращать клиенту ошибку в виде объекта, а не строки.

### Скрипты:
1. Скрипт для предзаполнения БД для тестирования и т.д.