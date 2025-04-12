// sender.go
// Название: Логика отправки сообщений
// Краткое описание: Реализует отправку сообщений по каналу и повторную отправку в случае потери
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений: Начальная версия

package main

import (
	"fmt"
	"time"
)

const (
	maxRetries = 3        // Максимальное количество попыток повторной отправки
	messageInterval = 2 * time.Second // Интервал отправки сообщений
)

func sendMessages(messageChannel chan<- string, ackChannel <-chan bool) {
	for i := 0; i < 5; i++ { // Отправляем 5 сообщений
		message := fmt.Sprintf("Сообщение %d", i+1)
		sendMessageWithRetry(message, messageChannel, ackChannel)
		time.Sleep(messageInterval)
	}
}

func sendMessageWithRetry(message string, messageChannel chan<- string, ackChannel <-chan bool) {
	retries := 0
	for retries < maxRetries {
		// Отправка сообщения
		messageChannel <- message
		fmt.Printf("Отправлено: %s\n", message)

		// Ожидание подтверждения
		select {
		case ack := <-ackChannel:
			if ack {
				fmt.Printf("Подтверждено получение сообщения: %s\n", message)
				return
			}
		case <-time.After(3 * time.Second): // Таймаут в случае отсутствия подтверждения
			retries++
			fmt.Printf("Не получено подтверждение, повторная отправка: %s (попытка %d)\n", message, retries)
		}
	}
	fmt.Printf("Не удалось отправить сообщение после %d попыток: %s\n", maxRetries, message)
}
