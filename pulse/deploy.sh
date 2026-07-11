#!/bin/bash

# ==========================================
# Автоматический деплой Pulse Dashboard
# ==========================================

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}   🚀 Деплой Pulse Dashboard          ${NC}"
echo -e "${GREEN}========================================${NC}"

# ==========================================
# 1. Получение домена
# ==========================================

if [ -z "$DOMAIN" ]; then
    echo ""
    read -p "🌐 Введите доменное имя (например, vitwebsite.ru): " DOMAIN
fi

if [ -z "$DOMAIN" ]; then
    echo -e "${RED}❌ Домен не указан!${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Домен: $DOMAIN${NC}"

# ==========================================
# 2. Создание .env файла
# ==========================================

echo -e "${YELLOW}📝 Создание .env файла...${NC}"

# Генерируем случайные пароли
DB_PASSWORD=$(openssl rand -base64 24 2>/dev/null || echo "StrongP@ssw0rd")
JWT_SECRET=$(openssl rand -base64 32 2>/dev/null || echo "super-secret-key")

cat > .env << EOF
# ===== ДОМЕН =====
DOMAIN=$DOMAIN

# ===== БАЗА ДАННЫХ =====
DB_USER=pulse_user
DB_PASSWORD=$DB_PASSWORD
DB_NAME=pulse_db

# ===== JWT =====
JWT_SECRET=$JWT_SECRET

# ===== ПОРТЫ =====
BACKEND_PORT=8082
FRONTEND_PORT=8081
POSTGRES_PORT=5432
EOF

echo -e "${GREEN}✅ .env файл создан${NC}"

# ==========================================
# 3. Проверка SSL сертификатов
# ==========================================

echo -e "${YELLOW}🔐 Проверка SSL сертификатов...${NC}"

if [ -f "ssl/fullchain.pem" ] && [ -f "ssl/privkey.pem" ]; then
    echo -e "${GREEN}✅ SSL сертификаты найдены в папке ssl/${NC}"
else
    echo -e "${RED}❌ SSL сертификаты не найдены!${NC}"
    echo "Пожалуйста, скопируйте сертификаты в папку ssl/:"
    echo "  ssl/fullchain.pem"
    echo "  ssl/privkey.pem"
    echo ""
    echo "Или получите их через certbot:"
    echo "  certbot certonly --standalone -d $DOMAIN -d www.$DOMAIN"
    exit 1
fi

# ==========================================
# 4. Запуск контейнеров
# ==========================================

echo -e "${YELLOW}🛠️  Сборка и запуск контейнеров...${NC}"
docker compose build --no-cache
docker compose up -d

# ==========================================
# 5. Ожидание готовности
# ==========================================

echo -e "${YELLOW}⏳ Ожидание запуска сервисов...${NC}"
sleep 15

# ==========================================
# 6. Проверка работы
# ==========================================

echo -e "${YELLOW}🔍 Проверка работы...${NC}"

RESPONSE=$(curl -s -k -X POST https://$DOMAIN/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"login":"admin","password":"admin123"}' 2>/dev/null)

if echo "$RESPONSE" | grep -q "token"; then
    echo -e "${GREEN}✅ Система успешно запущена!${NC}"
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}  🌐 Сайт: https://$DOMAIN${NC}"
    echo -e "${GREEN}  👤 Логин: admin${NC}"
    echo -e "${GREEN}  🔑 Пароль: admin123${NC}"
    echo -e "${GREEN}========================================${NC}"
else
    echo -e "${YELLOW}⚠️  API не отвечает. Проверьте логи:${NC}"
    echo "  docker compose logs backend --tail=30"
fi

echo ""
echo -e "${GREEN}✅ Деплой завершён!${NC}"
echo ""
echo "📋 Полезные команды:"
echo "  docker compose logs -f   - просмотр логов"
echo "  docker compose down      - остановка"
echo "  docker compose restart   - перезапуск"
echo ""
echo -e "${GREEN}🌐 Откройте сайт: https://$DOMAIN${NC}"