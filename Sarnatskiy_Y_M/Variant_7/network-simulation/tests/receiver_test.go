// receiver_test.go
// Название: Тесты для получателя
// Краткое описание: Тесты для проверки получения сообщений и вероятности потери
// Разработчик: Ваше имя
// Лицензия: GPLv3
// История изменений: Начальная версия

package tests

import (
	"testing"
	"network-simulation"
)

func TestReceiveMessages(t *testing.T) {
	// Создаем каналы
	messageChannel := make(chan string)
	ackChannel := make(chan bool)

	// Запускаем горутины
	go network_simulation.ReceiveMessages(messageChannel, ackChannel)

	// Проверяем, что сообщения приходят и подтверждения отправляются
}

