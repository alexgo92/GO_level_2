### Задание 3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
-------------------------------------------------------------------------------------------
### Рассуждения
Запускаем код командой ```$ go run -race main.go``` и видим ошибку "WARNING: DATA RACE ...".
Что бы от нее избавиться достаточно добавить mutex.Lock() и mutex.Unlock(), что предотвращает возникновение одновременного чтения и/или записи несколькиими горутинами в "критической секции".