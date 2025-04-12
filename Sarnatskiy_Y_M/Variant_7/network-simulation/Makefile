# Makefile

all: build

build:
	@echo "Сборка программы"
	go build -o network-simulation .

run:
	@echo "Запуск программы"
	go run main.go

test:
	@echo "Запуск тестов"
	go test ./tests -v

clear:
	@echo "Очистка"
	rm -f network-simulation

deploy:
	@echo "Деплой не настроен"

