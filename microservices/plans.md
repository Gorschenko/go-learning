1. ValidateQueryMiddleware
2. + ValidateParamsMiddleware
3. + ValidateBodyMiddleware
4. - CorrelationIdMiddleware. Не до конца сделан. Непонятно как с этим работать в других классах. Придется всегда передавать context.
5. + LogMiddleware