.PHONY: build up down logs restart clean

# Запуск всех сервисов
up:
	docker compose up -d

# Остановка всех сервисов
down:
	docker compose down

# Перезапуск
restart:
	docker compose restart

# Логи
logs:
	docker compose logs -f

# Сборка образов
build:
	docker compose build

# Полная очистка (включая данные БД)
clean:
	docker compose down -v
	docker system prune -f

# Обновление SSL сертификатов
renew-ssl:
	docker compose stop nginx
	sudo certbot renew --quiet
	docker compose start nginx

# Запуск в режиме разработки (с пересборкой фронта)
dev:
	docker compose up --build