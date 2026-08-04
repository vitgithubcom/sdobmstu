#!/bin/bash

# Цветной вывод
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Проверка работы сервиса (порты: 84, 8084, 5434) ===${NC}"

# 1. Проверяем здоровье
echo -e "\n${YELLOW}1. Health check:${NC}"
if curl -s http://localhost:84/health | jq . 2>/dev/null; then
    echo -e "${GREEN}✓ Health check OK${NC}"
else
    echo -e "${RED}✗ Health check FAILED${NC}"
    exit 1
fi

# 2. Одиночный запрос к пользователю
echo -e "\n${YELLOW}2. Запрос пользователя с id=1:${NC}"
if curl -s http://localhost:84/users/1 | jq . 2>/dev/null; then
    echo -e "${GREEN}✓ Запрос пользователя OK${NC}"
else
    echo -e "${RED}✗ Запрос пользователя FAILED${NC}"
fi

# 3. Метрики пула соединений
echo -e "\n${YELLOW}3. Метрики БД:${NC}"
if curl -s http://localhost:84/debug/metrics | jq . 2>/dev/null; then
    echo -e "${GREEN}✓ Метрики получены${NC}"
else
    echo -e "${RED}✗ Метрики не доступны${NC}"
fi

# 4. Стресс-тест (если установлен wrk)
echo -e "\n${YELLOW}4. Запуск стресс-теста:${NC}"
if command -v wrk &> /dev/null; then
    echo -e "Запуск wrk на 10 секунд с 10 потоками и 100 соединениями..."
    wrk -t 10 -c 100 -d 10s http://localhost:84/users/1
else
    echo -e "${YELLOW}wrk не установлен. Установите:${NC}"
    echo "  Ubuntu/Debian: sudo apt-get install wrk"
    echo "  MacOS: brew install wrk"
    echo -e "\nИспользуем ab (Apache Benchmark):"
    if command -v ab &> /dev/null; then
        ab -n 1000 -c 50 http://localhost:84/users/1
    else
        echo -e "${RED}Ни wrk, ни ab не установлены${NC}"
    fi
fi

# 5. Проверка статуса Nginx
echo -e "\n${YELLOW}5. Статус Nginx:${NC}"
curl -s http://localhost:84/nginx-status 2>/dev/null || echo "Статус недоступен"

# 6. Просмотр логов Nginx
echo -e "\n${YELLOW}6. Последние 10 строк логов Nginx:${NC}"
docker logs --tail 10 benchmark-nginx 2>&1 | grep -E "(access|error)" || echo "Логов нет"

# 7. Проверка использования портов
echo -e "\n${YELLOW}7. Использование портов на хосте:${NC}"
echo "Порт 84 (Nginx):"
ss -tlnp | grep ":84" || echo "  Не используется"
echo "Порт 8084 (Go):"
ss -tlnp | grep ":8084" || echo "  Не используется"
echo "Порт 5434 (Postgres):"
ss -tlnp | grep ":5434" || echo "  Не используется"

echo -e "\n${GREEN}=== Диагностика завершена ===${NC}"