#!/bin/bash

# ==========================================
# АВТОМАТИЧЕСКИЙ ДЕПЛОЙ PULSE DASHBOARD
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
# 1. ЗАПРОС ДОМЕНА
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
# 2. ГЕНЕРАЦИЯ .env
# ==========================================

echo -e "${YELLOW}📝 Создание .env файла...${NC}"

DB_PASSWORD=$(openssl rand -base64 24 2>/dev/null || echo "StrongP@ssw0rd")
JWT_SECRET=$(openssl rand -base64 32 2>/dev/null || echo "super-secret-key")

cat > .env << EOF
# ==========================================
# НАСТРОЙКИ ПРОЕКТА
# ==========================================

# Домен
DOMAIN=$DOMAIN

# Имя проекта (используется для контейнеров, сети, volume)
PROJECT_NAME=pulse-v-2

# Порты
BACKEND_PORT=8083
FRONTEND_PORT=8084
POSTGRES_PORT=5433

# База данных
DB_USER=pulse_user
DB_PASSWORD=$DB_PASSWORD
DB_NAME=pulse_db

# JWT
JWT_SECRET=$JWT_SECRET
EOF

echo -e "${GREEN}✅ .env файл создан${NC}"

# ==========================================
# 3. SSL СЕРТИФИКАТЫ
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
    echo "  sudo certbot certonly --standalone -d $DOMAIN -d www.$DOMAIN"
    exit 1
fi

# ==========================================
# 4. ЗАПУСК КОНТЕЙНЕРОВ
# ==========================================

echo -e "${YELLOW}🛠️  Сборка и запуск контейнеров...${NC}"

docker compose --env-file .env -p pulse-v-2 build --no-cache
docker compose --env-file .env -p pulse-v-2 up -d

# ==========================================
# 5. ПРОВЕРКА
# ==========================================

echo -e "${YELLOW}⏳ Ожидание запуска...${NC}"
sleep 15

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
    echo "  docker compose -p pulse-v-2 logs backend --tail=30"
fi

echo ""
echo -e "${GREEN}✅ Деплой завершён!${NC}"