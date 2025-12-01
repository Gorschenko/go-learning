- Value

```
func main() {
	log.Println("Start main")
	type contextKey struct{}
	var EmailKey = contextKey{}
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, 2, "a@a.ru")

	userEmail, ok := ctxWithValue.Value(EmailKey).(string)

	if ok {
		fmt.Println(userEmail)
	} else {
		fmt.Println("Not value")
	}

	log.Println("Finish main")
}
```

- Cancel

```
func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
```