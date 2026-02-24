# Makefile для сборки программы на Go

# Имя исполняемого файла
BINARY_NAME=max_unique_element

# Проверка зависимостей
check-deps:
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go не установлен. Установите golang-go."; exit 1; }

# Сборка
build: check-deps
	go build -o $(BINARY_NAME) main.go

# Запуск
run: build
	./$(BINARY_NAME)

# Чистка
clean:
	rm -f $(BINARY_NAME)

# Установка зависимостей (если нужны, но в этом проекте нет)
deps:
	go mod tidy

.PHONY: build run clean check-deps deps