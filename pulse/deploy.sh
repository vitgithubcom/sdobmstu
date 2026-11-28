#!/bin/bash

# ==========================================
# Автоматический деплой Pulse Dashboard (Traefik)
# ==========================================

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}   🚀 Деплой Pulse Dashboard (Traefik) ${NC}"
echo -e "${GREEN}========================================${NC}"

# ==========================================
# 1. Получение домена и email для Let's Encrypt
# ==========================================

if [ -z "$DOMAIN" ]; then
    read -p "🌐 Введите доменное имя (например, vitwebsite.ru): " DOMAIN
fi
if [ -z "$DOMAIN" ]; then
    echo -e "${RED}❌ Домен не указан!${NC}"
    exit 1
fi

if [ -z "$ACME_EMAIL" ]; then
    read -p "📧 Введите email для Let's Encrypt: " ACME_EMAIL
fi
if [ -z "$ACME_EMAIL" ]; then
    echo -e "${RED}❌ Email не указан!${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Домен: $DOMAIN${NC}"
echo -e "${GREEN}✅ Email: $ACME_EMAIL${NC}"

# ==========================================
# 2. Создание .env файла
# ==========================================

echo -e "${YELLOW}📝 Создание .env файла...${NC}"

DB_PASSWORD=$(openssl rand -base64 24 2>/dev/null || echo "StrongP@ssw0rd")
JWT_SECRET=$(openssl rand -base64 32 2>/dev/null || echo "super-secret-key")

cat > .env << EOF
# ===== ДОМЕН =====
DOMAIN=$DOMAIN
ACME_EMAIL=$ACME_EMAIL

# ===== БАЗА ДАННЫХ =====
DB_USER=pulse_user
DB_PASSWORD=$DB_PASSWORD
DB_NAME=pulse_db

# ===== JWT =====
JWT_SECRET=$JWT_SECRET

# ===== ПОРТЫ (внутренние) =====
BACKEND_PORT=8082
FRONTEND_PORT=8081
POSTGRES_PORT=5432
EOF

echo -e "${GREEN}✅ .env файл создан${NC}"

# ==========================================
# 3. Подготовка папки для сертификатов Traefik
# ==========================================

echo -e "${YELLOW}🔐 Подготовка хранилища сертификатов Traefik...${NC}"

mkdir -p traefik/letsencrypt
touch traefik/letsencrypt/acme.json
chmod 600 traefik/letsencrypt/acme.json

echo -e "${GREEN}✅ acme.json готов (права 600)${NC}"

# ==========================================
# 4. Запуск контейнеров
# ==========================================

echo -e "${YELLOW}🛠️  Сборка и запуск контейнеров...${NC}"
docker compose build --no-cache
docker compose up -d

# ==========================================
# 5. Ожидание готовности (Traefik + Let's Encrypt)
# ==========================================

echo -e "${YELLOW}⏳ Ожидание запуска сервисов и выпуска сертификата...${NC}"
sleep 30

# ==========================================
# 6. Проверка работы
# ==========================================

echo -e "${YELLOW}🔍 Проверка работы...${NC}"

RESPONSE=$(curl -s -X POST https://$DOMAIN/api/auth/login \
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
    echo "  docker compose logs traefik --tail=50"
    echo "  docker compose logs backend --tail=30"
    echo ""
    echo "Если сертификат ещё не выпущен, подождите 1-2 минуты и повторите проверку."
fi

echo ""
echo -e "${GREEN}✅ Деплой завершён!${NC}"
echo ""
echo "📋 Полезные команды:"
echo "  docker compose logs -f traefik   - логи Traefik (маршруты, сертификаты)"
echo "  docker compose logs -f backend   - логи backend"
echo "  docker compose down              - остановка"
echo "  docker compose restart           - перезапуск"
echo ""
echo -e "${GREEN}🌐 Откройте сайт: https://$DOMAIN${NC}"