```
func main() {
	log.Println("Start main")
	ctx := context.Background()
	/*
		Создаем контекст, который сообщит через 2 секунлы,
		что он завершен
	*/
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	/*
		Создаем канал для сигнализации завершения
	*/
	done := make(chan struct{})
	/*
		Запускается гоурутина, которая закрывает канал после 3 секунд
	*/
	go func() {
		log.Println("Start task")
		time.Sleep(3 * time.Second)
		close(done)
	}()

	/*
		Select блокирует main, пока не придет сигнал из канала
	*/
	select {
	case <-done:
		log.Println("Done task")

	case <-ctxWithTimeout.Done():
		log.Println("Timeout")
	}

	log.Println("Finish main")
}
```