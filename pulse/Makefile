.PHONY: up down logs build clean

# Запуск (сборка и старт)
up:
	docker compose up -d --build

# Остановка
down:
	docker compose down

# Логи
logs:
	docker compose logs -f

# Очистка (всё, включая данные БД)
clean:
	docker compose down -v
	docker system prune -f

# Перезапуск
restart:
	docker compose restart