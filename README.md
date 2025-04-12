# 🎮 4 в ряд — веб-версия на Go

Это простая веб-реализация классической игры **"Четыре в ряд"** (Connect Four), написанная на языке Go с использованием стандартной библиотеки `net/http`.

## 🧩 Правила игры

Игроки по очереди бросают фишки в столбцы. Цель — первым собрать **четыре фишки подряд** по вертикали, горизонтали или диагонали.

## 💻 Возможности
- Интерфейс в браузере (HTML + CSS + JS)
- Игра на одного (человека)
- Простое управление: кликни по колонке
- Поддержка перезапуска игры без перезагрузки страницы
- Игра против бота (AI)

## 🚀 Запуск
1. Клонируй репозиторий
2. Запусти сервер:
   `go run .`
3. Открой браузер: http://localhost:8080

## 🛠️ Планы на будущее
- Онлайн режим: игрок против игрока через сеть
- Улучшения интерфейса (анимации, адаптивность)
