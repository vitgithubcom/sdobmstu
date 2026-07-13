# Pulse Dashboard

Корпоративная система мониторинга KPI «Пульс»

## 📦 Требования

- Docker
- Docker Compose
- Доменное имя (для HTTPS)
- SSL сертификаты (или возможность получить через Let's Encrypt)

## 🚀 Быстрый старт

```bash
# 1. Клонируем репозиторий
git clone <your-repo> /opt/pulse
cd /opt/pulse

# 2. Делаем скрипт деплоя исполняемым
chmod +x deploy.sh

# 3. Создаём папку для SSL сертификатов
mkdir -p ssl

# 4. Копируем SSL сертификаты (если есть)
cp /etc/letsencrypt/live/ваш-домен.ru/*.pem ssl/ 2>/dev/null || echo "Сертификаты не найдены"

# 5. Запускаем деплой
./deploy.sh