// sender_test.go
// Название: Тесты для отправителя
// Краткое описание: Тесты для проверки функциональности отправки сообщений
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений: Начальная версия

package tests

import (
	"testing"
	"network-simulation" // Исправленный импорт
)

func TestSendMessages(t *testing.T) {
	// Создаем каналы
	messageChannel := make(chan string)
	ackChannel := make(chan bool)

	// Запускаем горутины
	go network_simulation.SendMessages(messageChannel, ackChannel)
	go network_simulation.ReceiveMessages(messageChannel, ackChannel)

	// Проверка на количество отправленных сообщений
	// Это можно проверить через таймеры или дополнительные механизмы
	// (например, счётчики сообщений)
}
