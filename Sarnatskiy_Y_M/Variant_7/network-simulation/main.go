// main.go
// Название: Программа для имитации передачи данных по ненадежной сети
// Краткое описание: Основной файл для запуска программы
// Разработчик: Ваше имя
// Лицензия: GPLv3
// История изменений: Начальная версия

package main

import (
	"fmt"
	"time"
	"network-simulation/networksimulation"
)

func main() {
	// Канал для передачи сообщений
	messageChannel := make(chan string)
	ackChannel := make(chan bool)

	// Запуск горутины для отправки сообщений
	go networksimulation.SendMessages(messageChannel, ackChannel)

	// Запуск горутины для получения сообщений
	go networksimulation.ReceiveMessages(messageChannel, ackChannel)

	// Ожидание завершения работы
	time.Sleep(10 * time.Second)
	fmt.Println("Завершение работы программы")
}
