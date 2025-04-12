// receiver.go
// Название: Логика получения сообщений и вероятность потери
// Краткое описание: Реализует получение сообщений и случайную потерю с заданной вероятностью
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений: Начальная версия

package main

import (
	"fmt"
	"math/rand"
)

const lossProbability = 0.2 // Вероятность потери сообщения (20%)

func receiveMessages(messageChannel <-chan string, ackChannel chan<- bool) {
	for message := range messageChannel {
		if simulatePacketLoss() {
			// Сообщение потеряно
			fmt.Printf("Сообщение потеряно: %s\n", message)
			continue
		}

		// Сообщение получено, отправляем подтверждение
		fmt.Printf("Получено: %s\n", message)
		ackChannel <- true
	}
}

func simulatePacketLoss() bool {
	// Случайное число от 0 до 1
	return rand.Float64() < lossProbability
}
